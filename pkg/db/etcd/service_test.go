package etcd_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/gogo/protobuf/types"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

func TestEtcdNotifierService(t *testing.T) {
	tests := []struct {
		name     string
		ops      func(*testing.T, context.Context, *etcd.NotifierService)
		watchers []watcher
	}{
		{
			name: "create and update virtual network",
			ops: func(t *testing.T, ctx context.Context, sv *etcd.NotifierService) {
				_, err := sv.CreateVirtualNetwork(ctx, &services.CreateVirtualNetworkRequest{
					VirtualNetwork: &models.VirtualNetwork{
						UUID: "vn-blue",
						Name: "vn_blue",
					},
				})
				assert.NoError(t, err, "create virtual network failed")

				_, err = sv.UpdateVirtualNetwork(ctx, &services.UpdateVirtualNetworkRequest{
					VirtualNetwork: &models.VirtualNetwork{
						UUID: "vn-blue",
						Name: "vn_bluuee",
					},
					FieldMask: types.FieldMask{Paths: []string{"name"}},
				})
				assert.NoError(t, err, "update virtual network failed")
			},
			watchers: []watcher{
				{
					key: "/json/virtual_network/vn-blue",
					values: []map[string]interface{}{
						{
							"name": "vn_blue",
						},
						{
							"name": "vn_bluuee",
						},
					},
				},
			},
		},
		{
			name: "create and delete reference from virtual network to logical router",
			ops: func(t *testing.T, ctx context.Context, sv *etcd.NotifierService) {
				_, err := sv.CreateVirtualNetwork(ctx, &services.CreateVirtualNetworkRequest{
					VirtualNetwork: &models.VirtualNetwork{
						UUID: "vn-blue",
						Name: "vn_blue",
					},
				})
				assert.NoError(t, err, "create virtual network failed")

				_, err = sv.CreateLogicalRouter(ctx, &services.CreateLogicalRouterRequest{
					LogicalRouter: &models.LogicalRouter{
						UUID: "lr-blue",
						Name: "lr_blue",
					},
				})
				assert.NoError(t, err, "create logical router failed")

				_, err = sv.CreateVirtualNetworkLogicalRouterRef(ctx,
					&services.CreateVirtualNetworkLogicalRouterRefRequest{
						ID: "vn-blue",
						VirtualNetworkLogicalRouterRef: &models.VirtualNetworkLogicalRouterRef{
							UUID: "lr-blue",
						}})
				assert.NoError(t, err, "create vn-lr reference failed")

				_, err = sv.DeleteVirtualNetworkLogicalRouterRef(ctx,
					&services.DeleteVirtualNetworkLogicalRouterRefRequest{
						ID: "vn-blue",
						VirtualNetworkLogicalRouterRef: &models.VirtualNetworkLogicalRouterRef{
							UUID: "lr-blue",
						}})
				assert.NoError(t, err, "delete vn-lr reference failed")
			},
			watchers: []watcher{
				{
					key: "/json/virtual_network/vn-blue",
					values: []map[string]interface{}{
						{
							"name": "vn_blue",
						},
						{
							"name": "vn_blue",
							"logical_router_refs": []interface{}{
								map[string]interface{}{
									"uuid": "lr-blue",
								},
							},
						},
						{
							"name": "vn_blue",
						},
					},
				},
				{
					key: "/json/logical_router/lr-blue",
					values: []map[string]interface{}{
						{
							"name": "lr_blue",
						},
						{
							"name": "lr_blue",
							"virtual_network_backrefs": []interface{}{
								map[string]interface{}{
									"uuid": "vn-blue",
								},
							},
						},
						{
							"name": "lr_blue",
						},
					},
				},
			},
		},
	}

	viper.Set("etcd.endpoints", integrationetcd.Endpoint)
	viper.Set("etcd.path", "json")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := integrationetcd.NewEtcdClient(t)
			defer ec.Close(t)

			// Clean the database
			ec.DeleteKey(t, "json", clientv3.WithPrefix())

			check := startWatchers(t, tt.watchers)
			sv, err := etcd.NewNotifierService("json", models.JSONCodec)
			require.NoError(t, err)

			tt.ops(t, context.Background(), sv)

			check(t)
		})
	}
}

type watcher struct {
	key    string
	values []map[string]interface{}
}

func startWatchers(t *testing.T, watchers []watcher) func(t *testing.T) {
	checks := []func(t *testing.T){}

	ec := integrationetcd.NewEtcdClient(t)
	for _, w := range watchers {
		w := w

		collect := ec.WatchKeyN(w.key, len(w.values), 2*time.Second)
		checks = append(checks, func(t *testing.T) {
			collected := collect()
			assert.Equal(t, len(w.values), len(collected))
			for i, e := range w.values[:len(collected)] {
				var data interface{}
				err := json.Unmarshal([]byte(collected[i]), &data)
				assert.NoError(t, err)
				testutil.AssertEqual(t, e, data, "different data under key %s\n", w.key)
			}
		})
	}

	return func(t *testing.T) {
		for _, c := range checks {
			c(t)
		}
	}
}
