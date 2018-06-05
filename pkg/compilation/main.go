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

	"github.com/Juniper/contrail/pkg/compilation/config"
	"github.com/Juniper/contrail/pkg/compilation/watch"
	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/serviceif"

	"github.com/Juniper/contrail/pkg/db/etcd"
	log "github.com/sirupsen/logrus"
)

// isMessageIndexBehind checks current index value
//   if currIndex <= storedIndex ignore message, return false
//   if currIndex > storedIndex return true
func isMessageIndexBehind(storedIndexStr string, currIndex int64) bool {
	storedIndex, err := strconv.ParseInt(storedIndexStr, 10, 64)
	if err != nil {
		log.Printf("storedIndex malformed %d\n", storedIndex)
		// ignore message
		return false
	}

	// compare with passed index
	if currIndex <= storedIndex {
		// ignore message
		log.Printf("storedIndex %d >= currIndex %d\n", storedIndex, currIndex)
		return false
	}

	log.Printf("storedIndex %d < currIndex %d!\n", storedIndex, currIndex)
	return true
}

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
	Etcdcl  *etcd.Client // TODO (Michal): Use interface.
	Cfg     *config.Config
	Service *compilationif.CompilationService
}

// NewIntentCompilationService makes a new Intent Compilation Service
func NewIntentCompilationService() (*IntentCompilationService, error) {
	ics := &IntentCompilationService{}
	log.Debug("Created New IntentCompilation Service")
	return ics, nil
}

// Init setup the IntentCompilationService.
func (ics *IntentCompilationService) Init(configFile string) error {
	conf, err := config.NewConfig(configFile)
	if err != nil {
		log.Print("Error: ", err)
		return err
	}
	ics.Cfg = conf

	watch.WatcherInit(conf.DefaultCfg.MaxJobQueueLen)
	watch.InitDispatcher(conf.DefaultCfg.NumberOfWorkers, ics.Service.HandleEtcdMessages)

	ics.Etcdcl, err = etcd.Dial()
	if err != nil {
		log.Print("Error: ", err)
		return err
	}

	ics.Service = SetupService()

	log.Debug("Initialized Intent Compilation Service")
	return nil
}

// HandleMessage : Callback function
//
// Try Locking
// if Lock Acquired
//  - check index > current_index
//    - if true
//      - set current index to Index
//      - process Message
//    - if false
//      - ignore msg
//  - Unlock
// if Lock not Acquired
// - Wait on lock
func (ics *IntentCompilationService) HandleMessage(
	ctx context.Context, index int64, oper int32, key string, newValue []byte,
) {

	log.Printf("Index: %d, oper: %d, Got Message %s: %s\n",
		index, oper, key, newValue)

	messageIndexKey := ics.Cfg.EtcdNotifierCfg.MsgIndexString

	err := ics.Etcdcl.InTransaction(ctx, func(txn etcd.Txn) error {
		storedIndex := txn.Get(messageIndexKey)

		shouldHandle := isMessageIndexBehind(string(storedIndex), index)
		if !shouldHandle {
			return nil
		}

		newMessageIndex := strconv.FormatInt(index, 10)
		txn.Put(messageIndexKey, []byte(newMessageIndex))

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

	err := ics.Etcdcl.Put(ctx, ics.Cfg.EtcdNotifierCfg.MsgIndexString, []byte("0"))
	if err != nil {
		log.Println("Cannot Set MessageIndex")
		return err
	}

	// Watch the Configured etcd directory for messages
	ics.Etcdcl.WatchRecursive(ctx, ics.Cfg.EtcdNotifierCfg.WatchPath, ics.HandleMessage)

	return nil
}

//Close closes IntentCompilationService
func (ics *IntentCompilationService) Close() error {
	return nil
}
