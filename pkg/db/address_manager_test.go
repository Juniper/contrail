package db

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/twinj/uuid"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/types/ipam"
	"github.com/stretchr/testify/assert"
)

func makeVirtualNetworkWithSubnets(ipamSubnets []*models.IpamSubnetType) *models.VirtualNetwork {
	virtualNetwork := models.MakeVirtualNetwork()
	vnSubnet := models.MakeVnSubnetsType()
	vnSubnet.IpamSubnets = append(vnSubnet.IpamSubnets, ipamSubnets...)

	networkIpamRefs := &models.VirtualNetworkNetworkIpamRef{Attr: vnSubnet}
	virtualNetwork.NetworkIpamRefs = append(virtualNetwork.NetworkIpamRefs, networkIpamRefs)

	return virtualNetwork
}

func testIpamSubnetWithIPs(ctx context.Context,
	virtualNetwork *models.VirtualNetwork, validIPs []string, testFun func(bool, string, error)) {
	for _, ip := range validIPs {
		isAllocated, err := db.IsIPAllocated(ctx, &ipam.IsIPAllocatedRequest{
			IPAddress:      ip,
			VirtualNetwork: virtualNetwork,
		})

		testFun(isAllocated, ip, err)
	}
}

func testIPAllocationWithIPs(ctx context.Context,
	virtualNetwork *models.VirtualNetwork, validIPs []string, testFun func(string, string, error)) {
	for _, ip := range validIPs {
		testFun(db.AllocateIP(ctx, &ipam.AllocateIPRequest{
			IPAddress:      ip,
			VirtualNetwork: virtualNetwork,
		}))
	}
}

func testIPDeallocationWithIPs(ctx context.Context,
	virtualNetwork *models.VirtualNetwork, validIPs []string, testFun func(string, error)) {
	for _, ip := range validIPs {
		testFun(ip, db.DeallocateIP(ctx, &ipam.DeallocateIPRequest{
			IPAddress:      ip,
			VirtualNetwork: virtualNetwork,
		}))
	}
}

func validateSubnetUUID(t *testing.T, ipamSubnet *models.IpamSubnetType, subnetUUID string) {
	if len(ipamSubnet.SubnetUUID) > 0 {
		assert.Equal(t, ipamSubnet.SubnetUUID, subnetUUID)
		return
	}
	_, err := uuid.Parse(subnetUUID)
	assert.NoError(t, err)
}

func TestAddressManagerSubnet(t *testing.T) {

	tests := []struct {
		name           string
		ipamSubnet     *models.IpamSubnetType
		allocationMode string
		validIPs       []string
		inValidIPs     []string
	}{
		{
			name: "Test subnet with any subnetUUID",
			ipamSubnet: &models.IpamSubnetType{
				AllocationPools: []*models.AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.255",
					},
				},
			},
			validIPs: []string{
				"10.0.0.0",
				"10.0.0.254",
			},
			inValidIPs: []string{
				"10.1.0.0",
				"127.0.0.1",
			},
			allocationMode: "user-defined-subnet-only",
		},
		{
			name: "Test subnet with provided subnetUUID",
			ipamSubnet: &models.IpamSubnetType{
				SubnetUUID: "uuid-1",
				AllocationPools: []*models.AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.255",
					},
				},
			},
			validIPs: []string{
				"10.0.0.0",
				"10.0.0.254",
			},
			allocationMode: "user-defined-subnet-only",
		},
		{
			name: "Test subnet with multiple allocation pools",
			ipamSubnet: &models.IpamSubnetType{
				AllocationPools: []*models.AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.255",
					},
					{
						Start: "10.0.3.0",
						End:   "10.0.3.255",
					},
				},
			},
			validIPs: []string{
				"10.0.0.0",
				"10.0.0.254",
				"10.0.3.0",
				"10.0.3.254",
			},
			inValidIPs: []string{
				"10.0.2.1",
				"10.0.4.1",
				"127.0.0.1",
			},
			allocationMode: "user-defined-subnet-only",
		},

		// TODO: Add test cases:
		// TODO: check allocation pool
		// TODO: check gw
		// TODO: check service addr
		// TODO: check dns nameservers
		// TODO: check allocation units
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			err := db.DoInTransaction(ctx,
				func(ctx context.Context) error {
					defer GetTransaction(ctx).ExecContext(ctx, "delete from ipaddress_pool") // nolint: errcheck
					request := &ipam.CreateIpamSubnetRequest{
						IpamSubnet: tt.ipamSubnet,
					}

					gotSubnetUUID, err := db.CreateIpamSubnet(ctx, request)
					assert.NoError(t, err)

					validateSubnetUUID(t, request.IpamSubnet, gotSubnetUUID)
					request.IpamSubnet.SubnetUUID = gotSubnetUUID

					virtualNetwork := makeVirtualNetworkWithSubnets([]*models.IpamSubnetType{
						request.IpamSubnet,
					})
					virtualNetwork.AddressAllocationMode = tt.allocationMode

					testIpamSubnetWithIPs(ctx, virtualNetwork, tt.inValidIPs,
						func(isAllocated bool, ip string, err error) {
							assert.Error(t, err)
						})

					testIpamSubnetWithIPs(ctx, virtualNetwork, tt.validIPs,
						func(isAllocated bool, ip string, err error) {
							assert.NoError(t, err)
							assert.False(t, isAllocated, "IP %v shouldn't be allocated", ip)
						})

					testIPAllocationWithIPs(ctx, virtualNetwork, tt.validIPs,
						func(address string, subnetUUID string, err error) {
							assert.NotNil(t, net.ParseIP(address), "Unexpected IP address format: %v", address)
							assert.NoError(t, err)
							assert.Equal(t, gotSubnetUUID, subnetUUID)
						})

					testIpamSubnetWithIPs(ctx, virtualNetwork, tt.validIPs,
						func(isAllocated bool, ip string, err error) {
							assert.NoError(t, err)
							assert.True(t, isAllocated, "IP %v should be already allocated", ip)
						})

					testIPDeallocationWithIPs(ctx, virtualNetwork, tt.validIPs,
						func(ip string, err error) {
							assert.NoError(t, err, "Couldn't deallocate ip %v", ip)
						})

					testIpamSubnetWithIPs(ctx, virtualNetwork, tt.validIPs,
						func(isAllocated bool, ip string, err error) {
							assert.NoError(t, err)
							assert.False(t, isAllocated, "IP %v is still allocated, but it should", ip)
						})

					err = db.DeleteIpamSubnet(ctx, &ipam.DeleteIpamSubnetRequest{
						SubnetUUID: request.IpamSubnet.SubnetUUID,
					})

					assert.NoError(t, err)
					return nil
				})
			assert.NoError(t, err)
		})
	}
}

func TestAddressManagerAllocateIP(t *testing.T) {

	tests := []struct {
		name           string
		ipamSubnet     *models.IpamSubnetType
		allocationMode string
		ipsToAllocate  []string
		fails          bool
	}{
		{
			name: "Test allocation with provided ip addresses",
			ipamSubnet: &models.IpamSubnetType{
				SubnetUUID: "uuid-1",
				AllocationPools: []*models.AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.255",
					},
				},
			},
			ipsToAllocate: []string{
				"10.0.0.0",
				"10.0.0.254",
			},
			allocationMode: "user-defined-subnet-only",
		},
		{
			name: "Test allocation without provided ip address",
			ipamSubnet: &models.IpamSubnetType{
				SubnetUUID: "uuid-1",
				AllocationPools: []*models.AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.255",
					},
				},
			},
			ipsToAllocate: []string{
				"",
				"",
			},
			allocationMode: "user-defined-subnet-only",
		},
		{
			name: "Test subnet exhaust",
			ipamSubnet: &models.IpamSubnetType{
				SubnetUUID: "uuid-1",
				AllocationPools: []*models.AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.2",
					},
				},
			},
			ipsToAllocate: []string{
				"",
				"",
				"",
			},
			allocationMode: "user-defined-subnet-only",
			// TODO: Check if it is actually correct
			fails: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			err := db.DoInTransaction(ctx,
				func(ctx context.Context) error {
					defer GetTransaction(ctx).ExecContext(ctx, "delete from ipaddress_pool") // nolint: errcheck
					request := &ipam.CreateIpamSubnetRequest{
						IpamSubnet: tt.ipamSubnet,
					}
					_, err := db.CreateIpamSubnet(ctx, request)
					assert.NoError(t, err)
					virtualNetwork := makeVirtualNetworkWithSubnets([]*models.IpamSubnetType{
						request.IpamSubnet,
					})
					virtualNetwork.AddressAllocationMode = tt.allocationMode

					for _, ipToAllocate := range tt.ipsToAllocate {
						var allocatedIP string
						allocatedIP, _, err = db.AllocateIP(ctx, &ipam.AllocateIPRequest{
							IPAddress:      ipToAllocate,
							VirtualNetwork: virtualNetwork,
						})
						if err != nil {
							break
						}
						if len(ipToAllocate) > 0 {
							assert.Equal(t, ipToAllocate, allocatedIP)
						} else {
							assert.NotNil(t, net.ParseIP(allocatedIP), "Unexpected IP address format: %v", allocatedIP)
						}
					}

					if tt.fails {
						// TODO: Check error type
						assert.Error(t, err)
					} else {
						assert.NoError(t, err)
					}
					return nil
				})
			assert.NoError(t, err)
		})
	}
}
