/*
 * Copyright 2018 - Praneet Bachheti
 *
 * Master Election Implementation
 *
 */

// TODO
// - callback func to caller if we r master

package election

import (
  "fmt"
  "strings"
  "time"
  "github.com/satori/go.uuid"
  etcl "github.com/Juniper/contrail/pkg/intent/compilation/etcdclient"
)


const (
  ELECTION_STATE_NONE     = iota
  ELECTION_STATE_MASTER   = iota
  ELECTION_STATE_BACKUP   = iota
)

type MemberType struct {
  Uuid     string // My UUID
  State    int    // My State
  Muuid    string // Master's UUID
  Client   *etcl.IntentEtcdClient // etcd client
  Timer    *time.Timer // Timer to check master
}

// Create a New Member
func NewMember(client *etcl.IntentEtcdClient) (*MemberType) {
  member := new(MemberType)
  member.State = ELECTION_STATE_NONE
  member.Client = client

  // Create UUID for the master
  uid := uuid.NewV4()
  fmt.Printf("Created UUID: %s\n", uid)

  member.Uuid = uid.String()

  for ; ; {
    err := member.TryBeMaster()
    if err != nil {
      time.Sleep(time.Second *5)
      continue
    }
    break
  }

  fmt.Printf(" Master is %s\n", member.Muuid)
  fmt.Printf(" Watch Master election\n")
  go member.WatchMasterElection()

  // Start a Timer for Master Refresh
  fmt.Printf(" Start Master Refresh Timer\n")
  go member.MasterRefreshTimer()

  return member
}

// Refresh TTL if we are the Master
func (m *MemberType) MasterRefreshTimer() {
  for ; ; {
    m.Timer = time.NewTimer(5 * time.Second)
    <-m.Timer.C
    if m.State == ELECTION_STATE_MASTER {
      m.Client.UpdateKeyWithTTL("master", 20 * time.Second)
    }
  }
  return
}

// Check Master Key
func (m *MemberType) TryBeMaster() error {
  m.State = ELECTION_STATE_BACKUP
  master, err := m.Client.Get("master")
  if err != nil {
    fmt.Print("Error: ", err)
    return err
  }
  if master == "" {
    err = m.Client.SetWithTTL("master", m.Uuid, 20 * time.Second)
    if err != nil {
      fmt.Print("Error: ", err)
      return err
    }
    master = m.Uuid
    m.State = ELECTION_STATE_MASTER
    fmt.Printf("I am the mstr: %s\n", m.Uuid)
  }
  m.Muuid = master
  return nil
}

// Callback function for Master Watch
func (m *MemberType) MasterChange(client *etcl.IntentEtcdClient,
                                  index uint64, key, newValue string) {
  fmt.Printf("Master Changed %s: %s\n", key, newValue)
  m.Muuid = newValue

  m.TryBeMaster()

  if strings.Compare(m.Uuid, m.Muuid) == 0 {
    fmt.Printf("I am Master: %s\n", m.Muuid)
    m.State = ELECTION_STATE_MASTER
  } else {
    fmt.Printf("Master is: %s\n", m.Muuid)
    m.State = ELECTION_STATE_BACKUP
  }
  return
}

// Watch for Master Change
func (m *MemberType) WatchMasterElection() {
  err := m.Client.WatchRecursive("master", m.MasterChange)
  if err != nil {
    fmt.Print("Error: ", err)
    return
  }
  return
}
