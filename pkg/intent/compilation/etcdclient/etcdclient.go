/*
 * Copyright 2018 - Praneet Bachheti
 *
 * ETCD Client interface
 *
 */

package etcdclient

import (
  "fmt"
  "strconv"
  "time"
  "github.com/coreos/etcd/client"
  "github.com/zieckey/etcdsync"
  "golang.org/x/net/context"
)


type Callback func(cl *IntentEtcdClient, index uint64, key, newValue string)

type EtcdClient interface {
  // Get gets a value in Etcd
  Get(key string) (string, error)

  // Set sets a value in Etcd
  Set(key, value string) error

  // Set sets a value in Etcd
  SetWithTTL(key, value string, ttl time.Duration) error

  // UpdateKeyWithTTL updates a key with a ttl value
  UpdateKeyWithTTL(key string, ttl time.Duration) error

  // MkDir creates an empty directory in etcd
  MkDir(directory string) error

  // Recursively Watches a Directory for changes
  WatchRecursive(directory string, callback Callback) error

  // Recursively Watches a Directory for changes After an Index
  WatchRecursiveAfterIndex(directory string, callback Callback) error
}

// IntentEtcdClient implements EtcdClient
type IntentEtcdClient struct {
    CurrIndex  uint64
    etcd       client.Client
    Lock       *etcdsync.Mutex
}


// Dial constructs a new EtcdClient
func Dial(etcdURI []string) (*IntentEtcdClient, error) {
  cfg := client.Config{
    Endpoints: etcdURI,
    //Transport: DefaultTransport,
  }

  etcd, err := client.New(cfg)
  if err != nil {
    fmt.Printf("Error connecting to ETCD: %s\n", etcdURI)
    return nil, err
  }
  return &IntentEtcdClient{0, etcd, nil}, nil
}

// Get gets a value in Etcd
func (etcdClient *IntentEtcdClient) Get(key string) (string, error) {
  api := client.NewKeysAPI(etcdClient.etcd)
  response, err := api.Get(context.Background(), key, nil)
  if err != nil {
    if client.IsKeyNotFound(err) {
      return "", nil
    }
    return "", err
  }
  return response.Node.Value, nil
}

// Set sets a value in Etcd
func (etcdClient *IntentEtcdClient) Set(key, value string) error {
  api := client.NewKeysAPI(etcdClient.etcd)
  _, err := api.Set(context.Background(), key, value, nil)
  return err
}

// Set sets a value in Etcd with TTL
func (etcdClient *IntentEtcdClient) SetWithTTL(key, value string,
      ttl time.Duration) error {
  api := client.NewKeysAPI(etcdClient.etcd)
  opts := &client.SetOptions{TTL: ttl}
  _, err := api.Set(context.Background(), key, value, opts)
  return err
}

// Updatekey updates a key with a ttl value
func (etcdClient *IntentEtcdClient) UpdateKeyWithTTL(key string,
      ttl time.Duration) error {
    api := client.NewKeysAPI(etcdClient.etcd)
    refreshopts := &client.SetOptions{Refresh: true,
       PrevExist: client.PrevExist, TTL: ttl}
    _, err := api.Set(context.Background(), key, "", refreshopts)
    return err
}

func (etcdClient *IntentEtcdClient) MkDir(directory string) error {
  api := client.NewKeysAPI(etcdClient.etcd)

  // Check if Directory exists
  res, err := api.Get(context.Background(), directory, nil)
  if err != nil && !client.IsKeyNotFound(err) {
    // Directory exists, ignore error
    return nil
  }

  if err != nil && client.IsKeyNotFound(err) {
    // Directory doesn't exist, create it
    opts := &client.SetOptions{Dir: true, PrevExist: client.PrevIgnore}
    _, err = api.Set(context.Background(), directory, "", opts)
    return err
  }

  // directory exists as a keyname
  if !res.Node.Dir {
    return fmt.Errorf("Cannot overwrite key/value with a directory: %v",
                       directory)
  }

  return nil
}

func (etcdClient *IntentEtcdClient) WatchRecursive(directory string,
      callback Callback) error {
  api := client.NewKeysAPI(etcdClient.etcd)
  afterIndex := uint64(0)

  for {
    watcher := api.Watcher(directory, &client.WatcherOptions{Recursive: true,
                           AfterIndex: afterIndex})
    resp, err := watcher.Next(context.Background())
    if err != nil {
      if shouldIgnoreError(err) {
        continue
      }
      return err
    }

    afterIndex = resp.Index
    callback(etcdClient, resp.Index, resp.Node.Key, resp.Node.Value)
  }
}

func (etcdClient *IntentEtcdClient) WatchRecursiveAfterIndex(directory string,
      callback Callback) error {
  api := client.NewKeysAPI(etcdClient.etcd)

  for {
    watcher := api.Watcher(directory, &client.WatcherOptions{Recursive: true,
                           AfterIndex: etcdClient.CurrIndex})
    resp, err := watcher.Next(context.Background())
    if err != nil {
      if shouldIgnoreError(err) {
        continue
      }
      return err
    }

    etcdClient.CurrIndex = resp.Index
    callback(etcdClient, resp.Index, resp.Node.Key, resp.Node.Value)
  }
}

func(etcdClient *IntentEtcdClient) CreateMsgLock(Uri []string, lockKey string,
    lockTimeout int) *etcdsync.Mutex {
  lk, err := etcdsync.New(lockKey, lockTimeout, Uri)
  if lk == nil || err != nil {
    fmt.Println("etcdsync.New failed ", err)
    return nil
  }
  etcdClient.Lock = lk
  return lk
}

func(etcdClient *IntentEtcdClient) MsgLock() {
  err := etcdClient.Lock.Lock()
	if err != nil {
		fmt.Printf("etcdsync.Lock failed")
	} else {
		fmt.Printf("etcdsync.Lock OK")
	}
}

func(etcdClient *IntentEtcdClient) MsgUnLock() {
	err := etcdClient.Lock.Unlock()
	if err != nil {
		fmt.Printf("etcdsync.Unlock failed")
	} else {
		fmt.Printf("etcdsync.Unlock OK")
	}
}

func(etcdClient *IntentEtcdClient) CheckSetIndex(index uint64) int {
  // Read Stored index
  storedIndexStr, err := etcdClient.Get("/MsgIndex")
  if err == nil {
    // compare with passed index
    storedIndex, _ := strconv.ParseUint(storedIndexStr, 10, 64)
    if index <= storedIndex {
      fmt.Printf("StoredIndex %d >= CurrIndex %d\n", storedIndex, index)
      return -1
    }
    fmt.Printf("StoredIndex %d < CurrIndex %d\n", storedIndex, index)
    newIndexStr := strconv.FormatUint(index, 10)
    etcdClient.Set("/MsgIndex", newIndexStr)
    if err != nil {
      fmt.Println("Cannot Set MsgIndex")
    }
    return 0
  }
  // Set passed index to Stored Index if greater
  return 0
}

func shouldIgnoreError(err error) bool {
  switch err := err.(type) {
  default:
    return false
  case *client.Error:
    return err.Code == client.ErrorCodeEventIndexCleared
  }
}
