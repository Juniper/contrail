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
	"runtime"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/compilation/config"
	"github.com/Juniper/contrail/pkg/compilation/watch"
	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/services"
)

// SetupService setups all required services and chains them.
func SetupService() *compilationif.CompilationService {
	// create services
	compilationService := compilationif.NewCompilationService()

	// chain them
	services.Chain(
		compilationService,
	)

	// return entry service
	return compilationService
}

type locker interface {
	DoWithLock(context.Context, string, time.Duration, func(ctx context.Context) error) error
}

type store interface {
	Create(context.Context, string, []byte) error
	Put(context.Context, string, []byte) error
	Get(context.Context, string) ([]byte, error)
	WatchRecursive(context.Context, string, int64) chan etcd.Event
}

//IntentCompilationService represents Intent Compilation Service.
type IntentCompilationService struct {
	Store   store
	Cfg     *config.Config
	Service *compilationif.CompilationService
	locker  locker

	log logrus.FieldLogger
}

// NewIntentCompilationService makes a new Intent Compilation Service
func NewIntentCompilationService() (*IntentCompilationService, error) {
	conf := config.ReadConfig()

	e, err := etcd.DialByConfig()
	if err != nil {
		return nil, err
	}

	l, err := etcd.NewDistributedLocker()
	if err != nil {
		return nil, err
	}

	return &IntentCompilationService{
		Service: SetupService(),
		Store:   etcd.NewClient(e),
		locker:  l,
		Cfg:     &conf,
		log:     log.NewLogger("intent-compilation"),
	}, nil
}

// HandleMessage handles message received from etcd pubsub.
func (ics *IntentCompilationService) HandleMessage(
	ctx context.Context, index int64, oper int32, key string, newValue []byte,
) {

	ics.log.Debugf("Index: %d, oper: %d, Got Message %s: %s\n",
		index, oper, key, newValue)

	messageIndexKey := ics.Cfg.EtcdNotifierCfg.MsgIndexString
	lockTTL := time.Second * time.Duration(ics.Cfg.EtcdNotifierCfg.MsgQueueLockTime) // TODO(Michal): Change field type to time.Duration

	err := ics.locker.DoWithLock(ctx, messageIndexKey, lockTTL, func(ctx context.Context) error {
		storedIndex, err := ics.getStoredIndex(ctx)
		if err != nil {
			ics.log.WithError(err).Debug("Error getting stored message index, skipping the message")
			return nil
		}

		if index <= storedIndex {
			ics.log.Debugf("index %d <= storedIndex %d", index, storedIndex)
			return nil
		}
		ics.log.Debugf("index %d > storedIndex %d!", index, storedIndex)

		if err := ics.putStoredIndex(ctx, index); err != nil {
			ics.log.WithError(err).Println("Cannot Set MessageIndex")
		}

		watch.AddJob(ctx, index, oper, key, string(newValue))
		time.Sleep(5 * time.Second) // TODO(Michal): Use some kind of synchronization primitive instead of sleep.
		ics.log.Debugf("#goroutines: %d", runtime.NumGoroutine())
		return nil
	})

	if err != nil {
		ics.log.WithError(err).Error("etcd transaction failed")
	}
}

func (ics *IntentCompilationService) getStoredIndex(ctx context.Context) (int64, error) {
	messageIndexKey := ics.Cfg.EtcdNotifierCfg.MsgIndexString

	storedIndexData, err := ics.Store.Get(ctx, messageIndexKey)
	if err != nil {
		return 0, err
	}
	storedIndex, err := strconv.ParseInt(string(storedIndexData), 10, 64)
	if err != nil {
		return 0, err
	}

	return storedIndex, nil
}

func (ics *IntentCompilationService) putStoredIndex(ctx context.Context, index int64) error {
	messageIndexKey := ics.Cfg.EtcdNotifierCfg.MsgIndexString

	newIndexStr := strconv.FormatInt(index, 10)
	return ics.Store.Put(ctx, messageIndexKey, []byte(newIndexStr))
}

// Run runs the IntentCompilationService.
func (ics *IntentCompilationService) Run(ctx context.Context) error {
	ics.log.Debug("Running Service")

	watch.WatcherInit(ics.Cfg.DefaultCfg.MaxJobQueueLen)
	watch.InitDispatcher(ics.Cfg.DefaultCfg.NumberOfWorkers, ics.Service.HandleEtcdMessages)

	ics.log.Debug("Setting MessageIndex to 0 (if not exists)")
	err := ics.Store.Create(ctx, ics.Cfg.EtcdNotifierCfg.MsgIndexString, []byte("0"))
	if err != nil {
		ics.log.Println("Cannot Set MessageIndex")
		return err
	}

	// Init watching channel
	watchPath := ics.Cfg.EtcdNotifierCfg.WatchPath
	ics.log.WithField("watchPath", watchPath).Debug("Starting recursive watch")
	eventChan := ics.Store.WatchRecursive(ctx, "/"+watchPath, int64(0))

	watch.RunDispatcher()

	ics.log.Debug("Starting handle loop")
	for {
		select {
		case <-ctx.Done():
			return nil
		case e := <-eventChan:
			ics.HandleMessage(ctx, e.Revision, e.Type, e.Key, e.Value)
		}
	}
}
