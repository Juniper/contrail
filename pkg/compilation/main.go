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

	"github.com/DavidCai1993/etcd-lock"
	"github.com/Juniper/contrail/pkg/compilation/config"
	"github.com/Juniper/contrail/pkg/compilation/watch"
	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/labstack/echo"

	etcl "github.com/Juniper/contrail/pkg/db/etcd"
	log "github.com/sirupsen/logrus"
)

// setMessageIndex(ctx context.Context, etcdcl *etcl.IntentEtcdClient,
func setMessageIndex(ctx context.Context, etcdcl *etcl.IntentEtcdClient,
	key string, index int64) {
	newIndexStr := strconv.FormatInt(index, 10)
	err := etcdcl.Update(ctx, key, newIndexStr)
	if err != nil {
		log.Println("Cannot Set MessageIndex")
	}
}

// checkMessageIndex checks current index value
//   if currIndex <= storedIndex ignore message, return -1
//   if currIndex > storedIndex return currIndex
func checkMessageIndex(ctx context.Context, etcdcl *etcl.IntentEtcdClient,
	key string, currIndex int64) bool {
	// Read Stored index
	storedIndexStr, err := etcdcl.Get(ctx, key)
	if err == nil {
		// compare with passed index
		storedIndex, error := strconv.ParseInt(storedIndexStr, 10, 64)
		if error != nil {
			// ignore message
			return false
		}
		if currIndex <= storedIndex {
			// ignore message
			log.Printf("storedIndex %d >= currIndex %d\n", storedIndex,
				currIndex)
			return false
		}
		log.Printf("storedIndex %d < currIndex %d!\n", storedIndex, currIndex)
		return true
	}
	return false
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
func HandleMessage(ctx context.Context, etcdcl *etcl.IntentEtcdClient,
	index int64, oper int32, key, newValue string) {

	log.Printf("Index: %d, oper: %d, Got Message %s: %s\n",
		index, oper, key, newValue)

	// Get Intent Compilation Handle from the context
	val := ctx.Value(IntentCompilationHandle)
	if val == nil {
		log.Println("IntentCompilationHandle not found:")
		return
	}
	ics := val.(*IntentCompilationService)

	lock, err := etcdcl.AcquireLock(ctx, ics.ELock,
		ics.Cfg.EtcdNotifierCfg.MsgIndexString,
		ics.Cfg.EtcdNotifierCfg.MsgQueueLockTime)
	if err != nil {
		log.Printf("Acquire Lock failed")
		return
	}

	log.Printf("Acquired the lock!")
	defer etcdcl.ReleaseLock(ctx, lock)

	ret := checkMessageIndex(ctx, etcdcl,
		ics.Cfg.EtcdNotifierCfg.MsgIndexString, index)
	if ret {
		setMessageIndex(ctx, etcdcl,
			ics.Cfg.EtcdNotifierCfg.MsgIndexString, index)
		watch.AddJob(ctx, index, oper, key, newValue)
		time.Sleep(5 * time.Second)
		log.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}

	log.Printf("Released the lock!")
}

// HandleEtcdMessages handles callbacks from workers
func HandleEtcdMessages(ctx context.Context, oper int32, key, value string) {
	// Get Intent Compilation Handle from the context
	val := ctx.Value(IntentCompilationHandle)
	if val == nil {
		log.Println("IntentCompilationHandle not found:")
		return
	}
	ics := val.(*IntentCompilationService)
	ics.Service.HandleEtcdMessages(ctx, oper, key, value)
	return
}

//IntentCompilationService represents Intent Compilation Service.
type IntentCompilationService struct {
	Echo    *echo.Echo
	Etcdcl  *etcl.IntentEtcdClient
	ELock   *etcdlock.Locker
	Cfg     *config.Config
	Service *compilationif.CompilationService
}

// NewIntentCompilationService makes a new Intent Compilation Service
func NewIntentCompilationService() (*IntentCompilationService, error) {
	ics := &IntentCompilationService{
		Echo: echo.New(),
	}
	log.Debug("Created New IntentCompilation Service")
	return ics, nil
}

//AddServices adds services to compilation chain
func (ics *IntentCompilationService) AddServices() *compilationif.CompilationService {
	var serviceChain []serviceif.Service

	compilationService := compilationif.NewCompilationService()
	serviceChain = append(serviceChain, compilationService)

	serviceif.Chain(serviceChain)

	return compilationService
}

//Init setup the IntentCompilationService.
func (ics *IntentCompilationService) Init(configFile string) error {
	conf, err := config.NewConfig(configFile)
	if err != nil {
		log.Print("Error: ", err)
		return err
	}
	ics.Cfg = conf

	watch.WatcherInit(conf.DefaultCfg.MaxJobQueueLen)
	watch.InitDispatcher(conf.DefaultCfg.NumberOfWorkers, HandleEtcdMessages)

	etcdcl, err := etcl.Dial(conf.EtcdServersUrls)
	if err != nil {
		log.Print("Error: ", err)
		return err
	}
	ics.Etcdcl = etcdcl

	// Create Lock
	locker, err := etcdcl.CreateLock(conf.EtcdServers[0])
	if err != nil {
		log.Fatal("Cannot Acquire Lock")
		return err
	}
	ics.ELock = locker

	ics.Service = ics.AddServices()

	log.Debug("Initialized Intent Compilation Service")
	return nil
}

//Run runs the IntentCompilationService.
func (ics *IntentCompilationService) Run() error {
	log.Debug("Running Service")
	ctx := context.Background()

	watch.RunDispatcher()

	err := ics.Etcdcl.Set(ctx, ics.Cfg.EtcdNotifierCfg.MsgIndexString, "0")
	if err != nil {
		log.Println("Cannot Set MessageIndex")
		return err
	}

	// Store IntentCompilationService Handle in context
	newCtx := context.WithValue(context.Background(), IntentCompilationHandle, ics)

	// Watch the Configured etcd directory for messages
	ics.Etcdcl.WatchRecursive(newCtx, ics.Cfg.EtcdNotifierCfg.WatchPath, HandleMessage)

	return nil
}

//Close closes IntentCompilationService
func (ics *IntentCompilationService) Close() error {
	return nil
}
