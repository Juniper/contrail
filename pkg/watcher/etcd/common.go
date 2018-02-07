package etcd

import (
	"time"

	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
)

const (
	kvClientRequestTimeout = 60 * time.Second
)

type noopKVClient struct{}

func (kv *noopKVClient) Put(ctx context.Context, key, val string, opts ...clientv3.OpOption) (
	*clientv3.PutResponse, error) {
	return nil, nil
}

func (kv *noopKVClient) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (
	*clientv3.GetResponse, error) {
	return nil, nil
}

func (kv *noopKVClient) Delete(ctx context.Context, key string, opts ...clientv3.OpOption) (
	*clientv3.DeleteResponse, error) {
	return nil, nil
}

func (kv *noopKVClient) Compact(ctx context.Context, rev int64, opts ...clientv3.CompactOption) (
	*clientv3.CompactResponse, error) {
	return nil, nil
}

func (kv *noopKVClient) Do(ctx context.Context, op clientv3.Op) (clientv3.OpResponse, error) {
	return clientv3.OpResponse{}, nil
}

func (kv *noopKVClient) Txn(ctx context.Context) clientv3.Txn { return nil }
