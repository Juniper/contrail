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

	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/compilation/config"
	"github.com/Juniper/contrail/pkg/compilation/watch"
	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/db/etcd"
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

//IntentCompilationService represents Intent Compilation Service.
type IntentCompilationService struct {
	Etcd    *etcd.Client // TODO (Michal): Use interface.
	Cfg     *config.Config
	Service *compilationif.CompilationService
	locker  locker
}

// NewIntentCompilationService makes a new Intent Compilation Service
func NewIntentCompilationService() (*IntentCompilationService, error) {
	ics := &IntentCompilationService{}

	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	watch.WatcherInit(conf.DefaultCfg.MaxJobQueueLen)
	watch.InitDispatcher(conf.DefaultCfg.NumberOfWorkers, ics.Service.HandleEtcdMessages)

	ics.Cfg = conf

	ics.Etcd, err = etcd.DialByConfig()
	if err != nil {
		return nil, err
	}

	ics.locker, err = etcd.NewDistributedLocker()
	if err != nil {
		return nil, err
	}

	ics.Service = SetupService()

	log.Debug("Created Intent Compilation Service")
	return ics, nil
}

// HandleMessage handles message received from etcd pubsub.
func (ics *IntentCompilationService) HandleMessage(
	ctx context.Context, index int64, oper int32, key string, newValue []byte,
) {

	log.Debugf("Index: %d, oper: %d, Got Message %s: %s",
		index, oper, key, newValue)

	messageIndexKey := ics.Cfg.EtcdNotifierCfg.MsgIndexString
	lockTTL := time.Second * time.Duration(ics.Cfg.EtcdNotifierCfg.MsgQueueLockTime) // TODO(Michal): Change field type to time.Duration

	err := ics.locker.DoWithLock(ctx, messageIndexKey, lockTTL, func(ctx context.Context) error {
		storedIndex, err := ics.getStoredIndex(ctx)
		if err != nil {
			log.WithError(err).Debug("Error getting stored message index, skipping the message")
			return nil
		}

		if index <= storedIndex {
			log.Debugf("index %d <= storedIndex %d\n", index, storedIndex)
			return nil
		}
		log.Debugf("index %d > storedIndex %d!", index, storedIndex)

		if err := ics.putStoredIndex(ctx, index); err != nil {
			log.WithError(err).Println("Cannot Set MessageIndex")
		}

		watch.AddJob(ctx, index, oper, key, string(newValue))
		time.Sleep(5 * time.Second) // TODO(Michal): Use some kind of synchronization primitive instead of sleep.
		log.Debugf("#goroutines: %d", runtime.NumGoroutine())
		return nil
	})

	if err != nil {
		log.WithError(err).Error("etcd transaction failed")
	}
}

func (ics *IntentCompilationService) getStoredIndex(ctx context.Context) (int64, error) {
	messageIndexKey := ics.Cfg.EtcdNotifierCfg.MsgIndexString

	storedIndexData, err := ics.Etcd.Get(ctx, messageIndexKey)
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
	return ics.Etcd.Put(ctx, messageIndexKey, []byte(newIndexStr))
}

//Run runs the IntentCompilationService.
func (ics *IntentCompilationService) Run() error {
	log.Debug("Running Service")
	ctx := context.Background()

	watch.RunDispatcher()

	err := ics.Etcd.Create(ctx, ics.Cfg.EtcdNotifierCfg.MsgIndexString, []byte("0"))
	if err != nil {
		log.Println("Cannot Set MessageIndex")
		return err
	}

	// Watch the Configured etcd directory for messages
	ics.Etcd.WatchRecursive(ctx, "/"+ics.Cfg.EtcdNotifierCfg.WatchPath, ics.HandleMessage)

	return nil
}

//Close closes IntentCompilationService
func (ics *IntentCompilationService) Close() error {
	return nil
}
