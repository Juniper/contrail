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
	"github.com/Juniper/contrail/pkg/serviceif"
)

// SetupService setups all required services and chains them.
func SetupService() *compilationif.CompilationService {
	// create services
	compilationService := compilationif.NewCompilationService()

	// chain them
	serviceif.Chain(
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
	Store     store
	Cfg       *config.Config
	Service   *compilationif.CompilationService
	locker    locker
	eventChan chan etcd.Event

	log logrus.FieldLogger
}

// NewIntentCompilationService makes a new Intent Compilation Service
func NewIntentCompilationService() (*IntentCompilationService, error) {
	ics := &IntentCompilationService{log: log.NewLogger("intent-compilation")}

	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	ics.Cfg = conf

	watch.WatcherInit(ics.Cfg.DefaultCfg.MaxJobQueueLen)
	watch.InitDispatcher(ics.Cfg.DefaultCfg.NumberOfWorkers, ics.Service.HandleEtcdMessages)

	e, err := etcd.DialByConfig()
	if err != nil {
		return nil, err
	}

	ics.Store = etcd.NewClient(e)

	ics.locker, err = etcd.NewDistributedLocker()
	if err != nil {
		return nil, err
	}

	ics.Service = SetupService()

	ics.log.Debug("Created Intent Compilation Service")
	return ics, nil
}

// HandleMessage handles message received from etcd pubsub.
func (ics *IntentCompilationService) HandleMessage(
	ctx context.Context, index int64, oper int32, key string, newValue []byte,
) {

	ics.log.Printf("Index: %d, oper: %d, Got Message %s: %s\n",
		index, oper, key, newValue)

	messageIndexKey := ics.Cfg.EtcdNotifierCfg.MsgIndexString
	lockTTL := time.Second * time.Duration(ics.Cfg.EtcdNotifierCfg.MsgQueueLockTime) // TODO(Michal): Change field type to time.Duration

	err := ics.locker.DoWithLock(ctx, messageIndexKey, lockTTL, func(ctx context.Context) error {
		storedIndex, err := ics.getStoredIndex(ctx)
		if err != nil {
			ics.log.WithError(err).Printf("Error getting stored message index\n")
			return nil
		}

		if index <= storedIndex {
			ics.log.Printf("index %d <= storedIndex %d\n", index, storedIndex)
			return nil
		}
		ics.log.Printf("index %d > storedIndex %d!\n", index, storedIndex)

		if err := ics.putStoredIndex(ctx, index); err != nil {
			ics.log.WithError(err).Println("Cannot Set MessageIndex")
		}

		watch.AddJob(ctx, index, oper, key, string(newValue))
		time.Sleep(5 * time.Second) // TODO(Michal): Use some kind of synchronization primitive instead of sleep.
		ics.log.Printf("#goroutines: %d\n", runtime.NumGoroutine())
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

// Init sets vars in store and initializes store watch.
func (ics *IntentCompilationService) Init(ctx context.Context) error {
	ics.log.Debug("Setting MessageIndex to 0 (if not exists)")
	err := ics.Store.Create(ctx, ics.Cfg.EtcdNotifierCfg.MsgIndexString, []byte("0"))
	if err != nil {
		ics.log.Println("Cannot Set MessageIndex")
		return err
	}

	// Init watching channel
	watchPath := ics.Cfg.EtcdNotifierCfg.WatchPath
	ics.log.WithField("watchPath", watchPath).Debug("Starting recursive watch")
	ics.eventChan = ics.Store.WatchRecursive(ctx, "/"+watchPath, int64(0))

	return nil
}

// Run runs the IntentCompilationService.
func (ics *IntentCompilationService) Run(ctx context.Context) error {
	ics.log.Debug("Running Service")

	watch.RunDispatcher()

	ics.log.Debug("Starting handle loop")
	for e := range ics.eventChan {
		ics.HandleMessage(ctx, e.Revision, e.Type, e.Key, e.Value)
	}

	return nil
}
