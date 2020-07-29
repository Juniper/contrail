package services

import (
	"context"

	asfservices "github.com/Juniper/asf/pkg/services"
)

// UserAgentKVService is a service which manages operations on key-value store
type UserAgentKVService interface {
	StoreKeyValue(ctx context.Context, key string, value string) error
	RetrieveValues(ctx context.Context, keys []string) (vals []string, err error)
	DeleteKey(ctx context.Context, key string) error
	RetrieveKVPs(ctx context.Context) (kvps []*asfservices.KeyValuePair, err error)
}
