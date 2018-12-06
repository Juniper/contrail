package etcd

import (
	"context"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	integrationetcd "github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name   string
		config *Config
		fails  bool
	}{
		{
			name: "succeeds when TLS disabled and correct credentials given",
			config: &Config{
				Config: *etcdConfig(),
			},
			fails: false,
		},
		{
			name: "fails when TLS enabled no certificates given",
			config: &Config{
				Config: *etcdConfig(),
				TLSConfig: TLSConfig{
					Enabled: true,
				},
			},
			fails: true,
		},
		{
			name: "fails when TLS enabled invalid certificate paths given",
			config: &Config{
				Config: *etcdConfig(),
				TLSConfig: TLSConfig{
					Enabled:         true,
					CertificatePath: "invalid-path",
					KeyPath:         "invalid-path",
					TrustedCAPath:   "invalid-path",
				},
			},
			fails: true,
		},
		//{ // TODO: generate test certs
		//	name: "succeeds when correct TLS configuration given",
		//	config: &Config{
		//		Config: *etcdConfig(),
		//		TLSConfig: TLSConfig{
		//			Enabled: true,
		//			CertificatePath: "testdata/cert.pem",
		//			KeyPath: "testdata/key.pem",
		//			TrustedCAPath: "testdata/trusted-ca.pem",
		//		},
		//	},
		//	fails: false,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient(tt.config)

			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				defer closeClient(t, c)

				ctx, cancel := context.WithTimeout(context.Background(), integrationetcd.ETCDRequestTimeout)
				defer cancel()

				_, err = c.ETCD.Maintenance.Status(ctx, integrationetcd.Endpoint)
				assert.NoError(t, err)
			}
		})
	}
}

func etcdConfig() *clientv3.Config {
	return &clientv3.Config{
		Endpoints:   []string{integrationetcd.Endpoint},
		DialTimeout: 10 * time.Millisecond, // TODO: reduce
	}
}

func closeClient(t *testing.T, c *Client) {
	assert.NoError(t, c.Close())
}

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
			c, err := NewClient(&Config{
				Client:      &clientv3.Client{},
				ServiceName: t.Name(),
			})
			require.NoError(t, err)

			err = c.InTransaction(tt.ctx, tt.do)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
