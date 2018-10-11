package etcd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/log"
)

func TestClient_InTransaction(t *testing.T) {
	tests := []struct {
		name    string
		ctx     context.Context
		do      func(context.Context) error
		wantErr bool
	}{
		{
			name:    "transaction already in context function returns no error",
			ctx:     WithTxn(context.Background(), &stmTxn{}),
			do:      func(context.Context) error { return nil },
			wantErr: false,
		},
		{
			name:    "transaction already in context function returns error",
			ctx:     WithTxn(context.Background(), &stmTxn{}),
			do:      func(context.Context) error { return assert.AnError },
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				log: log.NewLogger("testclient-log"),
			}
			err := c.InTransaction(tt.ctx, tt.do)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestClient_InTransactionTxnPassing(t *testing.T) {
	c := &Client{
		log: log.NewLogger("testclient-log"),
	}
	txn := &stmTxn{}
	ctx := WithTxn(context.Background(), txn)

	err := c.InTransaction(ctx, func(ctx context.Context) error {
		if GetTxn(ctx) == nil {
			return assert.AnError
		}
		return nil
	})

	assert.NoError(t, err)
}
