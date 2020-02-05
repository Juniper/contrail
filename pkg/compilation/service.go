/*
 * Copyright 2018 - Juniper Networks
 * Author: Praneet Bachheti
 *
 * main Implementation
 *
 */

package compilation

import (
	"context"
	"fmt"
	"runtime"
	"strconv"

	"github.com/Juniper/asf/pkg/keystone"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/compilation/config"
	"github.com/Juniper/contrail/pkg/compilation/dependencies"
	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/compilation/logic"
	"github.com/Juniper/contrail/pkg/compilation/watch"
	"github.com/Juniper/contrail/pkg/etcd"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/sirupsen/logrus"

	asfclient "github.com/Juniper/asf/pkg/client"
)

const serviceName = "intent-compiler"

// SetupService setups all required services and chains them.
func SetupService(
	WriteService services.WriteService,
	ReadService services.ReadService,
	allocator services.IntPoolAllocator,
) (services.Service, error) {
	reactions, err := dependencies.ParseReactions(
		[]byte(logic.ReactionsYAML),
		"intent-compiler",
	)
	if err != nil {
		return nil, err
	}

	logicService := logic.NewService(
		WriteService,
		ReadService,
		allocator,
		intent.NewCache(),
		dependencies.NewDependencyProcessor(reactions),
	)

	return logicService, nil
}

// Store represents data store that is source of events.
type Store interface {
	Create(context.Context, string, []byte) error
	Put(context.Context, string, []byte) error
	Get(context.Context, string) ([]byte, error)
	WatchRecursive(context.Context, string, int64) chan etcd.Message
	DoInTransaction(ctx context.Context, do func(context.Context) error) error
	Close() error
}

//IntentCompilationService represents Intent Compilation Service.
type IntentCompilationService struct {
	config    *config.Config
	Store     Store
	service   services.Service
	apiClient *client.HTTP

	log logrus.FieldLogger
}

// NewIntentCompilationService makes a new Intent Compilation Service.
func NewIntentCompilationService() (*IntentCompilationService, error) {
	c := config.ReadConfig()

	ec, err := etcd.NewClientByViper(serviceName)
	if err != nil {
		return nil, err
	}

	apiClient := newAPIClient(&c.APIClientConfig)

	logicService, err := SetupService(apiClient, apiClient, apiClient)
	if err != nil {
		return nil, err
	}

	return &IntentCompilationService{
		service:   logicService,
		apiClient: apiClient,
		Store:     ec,
		config:    &c,
		log:       logutil.NewLogger(serviceName),
	}, nil
}

func newAPIClient(c *config.APIClientConfig) *client.HTTP {
	client := client.NewHTTP(&asfclient.HTTPConfig{
		ID:       c.ID,
		Password: c.Password,
		Endpoint: c.URL,
		AuthURL:  c.AuthURL,
		Scope: keystone.NewScope(
			c.DomainID,
			c.DomainName,
			c.ProjectID,
			c.ProjectName,
		),
		Insecure: c.Insecure,
	})

	return client
}

// handleMessage handles message received from etcd pubsub.
// TODO(Michal): [FIXME] Handle etcd transaction events.
// When sync does a etcd transaction all events are sent at once,
// with equal revision. This is handled in etcd.Client.Watch by splitting the list
// into separate events. However handleMessage relies on revision value and
// events having same revision are discarded. That causes that only one event
// from single transaction could be handled.
// Jira-Bug: JBE-993
func (ics *IntentCompilationService) handleMessage(
	ctx context.Context, index int64, oper int32, key string, newValue []byte,
) {

	ics.log.Debugf("Index: %d, oper: %d, Got Message %s: %s\n",
		index, oper, key, newValue)

	var skipMessage bool
	if err := ics.Store.DoInTransaction(ctx, func(ctx context.Context) error {
		skipMessage = true
		storedIndex, err := ics.getStoredIndex(ctx)
		if err != nil {
			ics.log.WithError(err).Debug("Error getting stored message index, skipping the message")
			return nil
		}

		if index <= storedIndex {
			ics.log.Debugf("index %d <= storedIndex %d", index, storedIndex)
			return nil
		}
		ics.log.Debugf("index %d > storedIndex %d", index, storedIndex)

		ics.putStoredIndex(ctx, index)

		skipMessage = false
		return nil
	}); err != nil {
		ics.log.WithError(err).Error("etcd transaction failed")
	}

	if !skipMessage {
		watch.AddJob(ctx, index, oper, key, string(newValue))
		ics.log.Debugf("#goroutines: %d", runtime.NumGoroutine())
	}
}

func (ics *IntentCompilationService) getStoredIndex(ctx context.Context) (int64, error) {
	txn := etcd.GetTxn(ctx)
	messageIndexKey := ics.config.EtcdNotifierCfg.MsgIndexString

	storedIndexData := txn.Get(messageIndexKey)

	storedIndex, err := strconv.ParseInt(string(storedIndexData), 10, 64)
	if err != nil {
		return 0, err
	}

	return storedIndex, nil
}

func (ics *IntentCompilationService) putStoredIndex(ctx context.Context, index int64) {
	txn := etcd.GetTxn(ctx)
	messageIndexKey := ics.config.EtcdNotifierCfg.MsgIndexString

	newIndexStr := strconv.FormatInt(index, 10)
	txn.Put(messageIndexKey, []byte(newIndexStr))
}

// Run runs the IntentCompilationService.
func (ics *IntentCompilationService) Run(ctx context.Context) error {
	ics.log.Debug("Running Service")

	if err := ics.apiClient.Login(ctx); err != nil {
		return err
	}

	watch.WatcherInit(ics.config.DefaultCfg.MaxJobQueueLen)
	watch.InitDispatcher(ics.config.DefaultCfg.NumberOfWorkers, ics.handleEtcdMessage)

	ics.log.Debug("Setting MessageIndex to 0 (if not exists)")
	if err := ics.Store.Create(ctx, ics.config.EtcdNotifierCfg.MsgIndexString, []byte("0")); err != nil {
		ics.log.Println("Cannot Set MessageIndex")
		return err
	}

	// Init watching channel
	watchPath := ics.config.EtcdNotifierCfg.WatchPath
	ics.log.WithField("watchPath", watchPath).Debug("Starting recursive watch")
	eventChan := ics.Store.WatchRecursive(ctx, "/"+watchPath, int64(0))

	watch.RunDispatcher()

	ics.log.Debug("Starting handle loop")
	for {
		select {
		case <-ctx.Done():
			return nil
		case e, ok := <-eventChan:
			if !ok {
				ics.log.Info("event channel unsuspectingly closed, restarting etcd watch")
				eventChan = ics.Store.WatchRecursive(ctx, "/"+watchPath, int64(0))
			} else {
				ics.handleMessage(ctx, e.Revision, e.Type, e.Key, e.Value)
			}
		}
	}
}

// HandleEtcdMessage handles messages received from etcd.
func (ics *IntentCompilationService) handleEtcdMessage(ctx context.Context, oper int32, key, value string) {
	messageFields := logrus.Fields{"operation": oper, "key": key, "value": value}
	ics.log.WithFields(messageFields).Print("HandleEtcdMessages: Got a message")
	event, err := etcd.ParseEvent(oper, key, []byte(value))
	if err != nil {
		logrus.WithFields(messageFields).WithField(
			logrus.ErrorKey, fmt.Sprintf("%+v", err)).Error("failed to parse etcd event")
	}
	processor := services.ServiceEventProcessor{
		Service: ics.service,
	}
	_, err = processor.Process(ctx, event)
	if err != nil {
		ics.log.WithFields(messageFields).WithField(
			logrus.ErrorKey, fmt.Sprintf("%+v", err)).Error("Failed to handle etcd message")
	}
}
