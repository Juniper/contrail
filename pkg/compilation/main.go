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
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/compilation/config"
	"github.com/Juniper/contrail/pkg/compilation/watch"
	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/db/etcd"
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

//IntentCompilationService represents Intent Compilation Service.
type IntentCompilationService struct {
	Etcd    *etcd.Client // TODO (Michal): Use interface.
	Cfg     *config.Config
	Service *compilationif.CompilationService
}

// NewIntentCompilationService makes a new Intent Compilation Service
func NewIntentCompilationService(configFile string) (*IntentCompilationService, error) {
	ics := &IntentCompilationService{}

	conf, err := config.NewConfig(configFile)
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

	ics.Service = SetupService()

	log.Debug("Created Intent Compilation Service")
	return ics, nil
}

// HandleMessage handles message received from etcd pubsub.
func (ics *IntentCompilationService) HandleMessage(
	ctx context.Context, index int64, oper int32, key string, newValue []byte,
) {

	log.Printf("Index: %d, oper: %d, Got Message %s: %s\n",
		index, oper, key, newValue)

	messageIndexKey := ics.Cfg.EtcdNotifierCfg.MsgIndexString

	err := ics.Etcd.InTransaction(ctx, func(ctx context.Context) error {
		txn := etcd.GetTxn(ctx)
		storedIndex, err := etcd.GetInt64InTxn(txn, messageIndexKey)
		if err != nil {
			log.Printf("Error getting stored message index %s\n", err)
			return nil
		}

		if index <= storedIndex {
			log.Printf("index %d <= storedIndex %d\n", index, storedIndex)
			return nil
		}
		log.Printf("index %d > storedIndex %d!\n", index, storedIndex)

		etcd.PutInt64InTxn(txn, messageIndexKey, index)

		watch.AddJob(ctx, index, oper, key, string(newValue))
		time.Sleep(5 * time.Second)
		log.Printf("#goroutines: %d\n", runtime.NumGoroutine())
		return nil
	})

	if err != nil {
		log.Errorf("etcd transaction failed: %v", err)
	}
}

//Run runs the IntentCompilationService.
func (ics *IntentCompilationService) Run() error {
	log.Debug("Running Service")
	ctx := context.Background()

	watch.RunDispatcher()

	err := ics.Etcd.Put(ctx, ics.Cfg.EtcdNotifierCfg.MsgIndexString, []byte("0"))
	if err != nil {
		log.Println("Cannot Set MessageIndex")
		return err
	}

	// Watch the Configured etcd directory for messages
	ics.Etcd.WatchRecursive(ctx, ics.Cfg.EtcdNotifierCfg.WatchPath, ics.HandleMessage)

	return nil
}

//Close closes IntentCompilationService
func (ics *IntentCompilationService) Close() error {
	return nil
}
