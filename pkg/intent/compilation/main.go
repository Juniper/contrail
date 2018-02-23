package main

import (
  "fmt"
  "log"
  "runtime"
  "time"
  "github.com/Juniper/contrail/pkg/intent/compilation/config"
  elec "github.com/Juniper/contrail/pkg/intent/compilation/election"
  etcl "github.com/Juniper/contrail/pkg/intent/compilation/etcdclient"
  "github.com/Juniper/contrail/pkg/intent/compilation/watch"
)

type PluginConf struct {
    PluginDirectory   string
    WatchDirectory    string
    Functions         []string
    NumberOfWorkers   int
    MaxJobQueueLen    int
}

type Plugin struct {
  Name        string
  Conf        *PluginConf
  PluginNames []string
}

var counter int

// Callback function
func HandleMsg(etcdClient *etcl.IntentEtcdClient,
               index uint64, key, newValue string) {
  fmt.Printf("Index: %d, Got Msg %s: %s\n", index, key, newValue)

  // Try Locking
  // if Lock Acquired
  //  - check index > current_index
  //    - if true
  //      - set current index to Index
  //      - process Msg
  //    - if false
  //      - ignore msg
  //  - Unlock
  // if Lock not Acquired
  // - Wait on lock

  etcdClient.MsgLock()
  log.Printf("Acquired the lock!")

  ret := etcdClient.CheckSetIndex(index)
  if ret == 0 {
    m := make(map[string]string)
    m[key] = newValue
    counter += 1
    watch.AddJob(counter, m)
    time.Sleep(8*time.Second)
    fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
  }

  etcdClient.MsgUnLock()
  log.Printf("Released the lock!")
}

func main() {
  fmt.Print("Hello\n\n")

  conf := config.NewConfig("main.ini")
  //conf.ReadPlugin()
  //conf.CheckPlugin()

  watch.WatcherInit(conf.PluginConfig.MaxJobQueueLen)
  watch.InitDispatcher(conf.PluginConfig.NumberOfWorkers)
  watch.RunDispatcher()

  etcdcl, err := etcl.Dial(conf.EtcdServersList)
  if err != nil {
    fmt.Print("Error: ", err)
    return
  }

  // Master Election
  elec.NewMember(etcdcl)

  // Create Work Queues
  // etcl.NewEtcdQueues(etcdcl)

  lock := etcdcl.CreateMsgLock(conf.EtcdServersList,
    conf.DefaultConfig.MsgIndexString,
    conf.DefaultConfig.QueueLockTime)
	if lock == nil {
		log.Printf("Lock failed")
		return
	}

  //  Watch the Configured etcd directory for messages
  err = etcdcl.WatchRecursive(conf.PluginConfig.WatchDirectory, HandleMsg)
  if err != nil {
    fmt.Print("Error: ", err)
    return
  }

  for ; ; {
    fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
    time.Sleep(100 * time.Second)
  }
}
