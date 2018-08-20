package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

// Message contains message data reveived from WatchRecursive.
type Message struct {
	Revision int64
	Type     int32
	Key      string
	Value    []byte
}

// Message type values.
const (
	MessageCreate = iota
	MessageModify
	MessageDelete
	MessageUnknown
)

func NewMessage(e *clientv3.Event) Message {
	return Message{
		Revision: e.Kv.ModRevision,
		Type:     messageTypeFromEvent(e),
		Key:      string(e.Kv.Key),
		Value:    e.Kv.Value,
	}
}

func messageTypeFromEvent(e *clientv3.Event) int32 {
	switch {
	case e.IsCreate():
		return MessageCreate
	case e.IsModify():
		return MessageModify
	case e.Type == mvccpb.DELETE:
		return MessageDelete
	}
	return MessageUnknown
}
