package db

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/twinj/uuid"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

func TestAddressManagerAllocations(t *testing.T) {
	tests := []struct {
		name               string
		ipamSubnet         *models.IpamSubnetType
		allocationMode     string
		expectedValidIPs   []string
		expectedInvalidIPs []string
	}{
		{
			name: "Test subnet without allocations pools",
			ipamSubnet: &models.IpamSubnetType{
				SubnetUUID: "uuid-1",
				Subnet: &models.SubnetType{
					IPPrefix:    "10.0.0.0",
					IPPrefixLen: 24,
				},
			},
			allocationMode: models.UserDefinedSubnetOnly,
			expectedValidIPs: []string{
				"10.0.0.1",
				"10.0.0.127",
				"10.0.0.254",
			},
			expectedInvalidIPs: []string{
				"10.1.0.0",
				"127.0.0.1",
			},
		},
		{
			name: "Test subnet with any subnetUUID",
			ipamSubnet: &models.IpamSubnetType{
				Subnet: &models.SubnetType{
					IPPrefix:    "10.0.0.0",
					IPPrefixLen: 24,
				},
				AllocationPools: []*models.AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.255",
					},
				},
			},
			allocationMode: models.UserDefinedSubnetOnly,
			expectedValidIPs: []string{
				"10.0.0.0",
				"10.0.0.254",
			},
			expectedInvalidIPs: []string{
				"10.1.0.0",
				"127.0.0.1",
			},
		},
		{
			name: "Test subnet with provided subnetUUID",
			ipamSubnet: &models.IpamSubnetType{
				SubnetUUID: "uuid-1",
				Subnet: &models.SubnetType{
					IPPrefix:    "10.0.0.0",
					IPPrefixLen: 24,
				},
				AllocationPools: []*models.AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.255",
					},
				},
			},
			allocationMode: models.UserDefinedSubnetOnly,
			expectedValidIPs: []string{
				"10.0.0.0",
				"10.0.0.254",
			},
		},
		{
			name: "Test subnet with multiple allocation pools",
			ipamSubnet: &models.IpamSubnetType{
				Subnet: &models.SubnetType{
					IPPrefix:    "10.0.0.0",
					IPPrefixLen: 16,
				},
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
			allocationMode: models.UserDefinedSubnetOnly,
			expectedValidIPs: []string{
				"10.0.0.0",
				"10.0.0.254",
				"10.0.3.0",
				"10.0.3.254",
			},
			expectedInvalidIPs: []string{
				"10.0.2.1",
				"10.0.4.1",
				"127.0.0.1",
			},
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

					for _, invalidIP := range tt.expectedInvalidIPs {
						_, inErr := db.IsIPAllocated(ctx, &ipam.IsIPAllocatedRequest{
							IPAddress:      invalidIP,
							VirtualNetwork: virtualNetwork,
						})
						assert.Error(t, inErr)

						_, _, inErr = db.AllocateIP(ctx, &ipam.AllocateIPRequest{
							IPAddress:      invalidIP,
							VirtualNetwork: virtualNetwork,
						})
						assert.Error(t, inErr)

						inErr = db.DeallocateIP(ctx, &ipam.DeallocateIPRequest{
							IPAddress:      invalidIP,
							VirtualNetwork: virtualNetwork,
						})
						assert.Error(t, inErr)
					}

					for _, validIP := range tt.expectedValidIPs {
						isAllocated, inErr := db.IsIPAllocated(ctx, &ipam.IsIPAllocatedRequest{
							IPAddress:      validIP,
							VirtualNetwork: virtualNetwork,
						})
						assert.NoError(t, inErr)
						assert.False(t, isAllocated, "IP %v shouldn't be allocated", validIP)

						allocatedIP, subnetUUID, inErr := db.AllocateIP(ctx, &ipam.AllocateIPRequest{
							IPAddress:      validIP,
							VirtualNetwork: virtualNetwork,
						})
						assert.Equal(t, validIP, allocatedIP, "Unexpected ip value")
						assert.NoError(t, inErr)
						assert.Equal(t, gotSubnetUUID, subnetUUID)

						isAllocated, inErr = db.IsIPAllocated(ctx, &ipam.IsIPAllocatedRequest{
							IPAddress:      validIP,
							VirtualNetwork: virtualNetwork,
						})
						assert.NoError(t, inErr)
						assert.True(t, isAllocated, "IP %v should be already allocated", validIP)

						inErr = db.DeallocateIP(ctx, &ipam.DeallocateIPRequest{
							IPAddress:      validIP,
							VirtualNetwork: virtualNetwork,
						})
						assert.NoError(t, inErr, "Couldn't deallocate ip %v", validIP)

						isAllocated, inErr = db.IsIPAllocated(ctx, &ipam.IsIPAllocatedRequest{
							IPAddress:      validIP,
							VirtualNetwork: virtualNetwork,
						})
						assert.NoError(t, inErr)
						assert.False(t, isAllocated, "IP %v is still allocated, but it shouldn't", validIP)
					}

					err = db.DeleteIpamSubnet(ctx, &ipam.DeleteIpamSubnetRequest{
						SubnetUUID: request.IpamSubnet.SubnetUUID,
					})
					assert.NoError(t, err)

					err = clearIPAddressPool(ctx)
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
				Subnet: &models.SubnetType{
					IPPrefix:    "10.0.0.0",
					IPPrefixLen: 24,
				},
				AllocationPools: []*models.AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.255",
					},
				},
			},
			ipsToAllocate: []string{
				"10.0.0.0",
				"10.0.0.127",
				"10.0.0.254",
			},
			allocationMode: models.UserDefinedSubnetOnly,
		},
		{
			name: "Test allocation with provided ip addresses (end of pool)",
			ipamSubnet: &models.IpamSubnetType{
				SubnetUUID: "uuid-1",
				Subnet: &models.SubnetType{
					IPPrefix:    "10.0.0.0",
					IPPrefixLen: 24,
				},
				AllocationPools: []*models.AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.255",
					},
				},
			},
			ipsToAllocate: []string{
				"10.0.0.255",
			},
			allocationMode: models.UserDefinedSubnetOnly,
			// TODO: This test SHOULDN'T be failing
			fails: true,
		},
		{
			name: "Test allocation without provided ip address",
			ipamSubnet: &models.IpamSubnetType{
				SubnetUUID: "uuid-1",
				Subnet: &models.SubnetType{
					IPPrefix:    "10.0.0.0",
					IPPrefixLen: 24,
				},
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
			allocationMode: models.UserDefinedSubnetOnly,
		},
		{
			name: "Test subnet exhaust",
			ipamSubnet: &models.IpamSubnetType{
				SubnetUUID: "uuid-1",
				Subnet: &models.SubnetType{
					IPPrefix:    "10.0.0.0",
					IPPrefixLen: 24,
				},
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
				"",
			},
			allocationMode: models.UserDefinedSubnetOnly,
			fails:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			err := db.DoInTransaction(ctx,
				func(ctx context.Context) error {
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

					err = clearIPAddressPool(ctx)
					assert.NoError(t, err)
					return nil
				})
			assert.NoError(t, err)
		})
	}
}

func TestAddressManagerCheckIfIpamSubnetExists(t *testing.T) {

	tests := []struct {
		name       string
		ipamSubnet *models.IpamSubnetType
		subnetUUID string
		expects    bool
	}{
		{
			name: "Check existing subnet",
			ipamSubnet: &models.IpamSubnetType{
				SubnetUUID: "uuid-1",
				Subnet: &models.SubnetType{
					IPPrefix:    "10.0.0.0",
					IPPrefixLen: 24,
				},
			},
			subnetUUID: "uuid-1",
			expects:    true,
		},
		{
			name: "Check non-existing subnet",
			ipamSubnet: &models.IpamSubnetType{
				SubnetUUID: "uuid-1",
				Subnet: &models.SubnetType{
					IPPrefix:    "10.0.0.0",
					IPPrefixLen: 24,
				},
			},
			subnetUUID: "non-existing-uuid",
			expects:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			err := db.DoInTransaction(ctx,
				func(ctx context.Context) error {
					_, err := db.CreateIpamSubnet(ctx, &ipam.CreateIpamSubnetRequest{
						IpamSubnet: tt.ipamSubnet,
					})
					assert.NoError(t, err)
					res, err := db.CheckIfIpamSubnetExists(ctx, tt.subnetUUID)
					assert.NoError(t, err)
					assert.Equal(t, tt.expects, res)

					err = clearIPAddressPool(ctx)
					assert.NoError(t, err)
					return nil
				})
			assert.NoError(t, err)
		})
	}
}

func makeVirtualNetworkWithSubnets(ipamSubnets []*models.IpamSubnetType) *models.VirtualNetwork {
	virtualNetwork := models.MakeVirtualNetwork()
	vnSubnet := models.MakeVnSubnetsType()
	vnSubnet.IpamSubnets = append(vnSubnet.IpamSubnets, ipamSubnets...)

	networkIpamRefs := &models.VirtualNetworkNetworkIpamRef{Attr: vnSubnet}
	virtualNetwork.NetworkIpamRefs = append(virtualNetwork.NetworkIpamRefs, networkIpamRefs)

	return virtualNetwork
}

func validateSubnetUUID(t *testing.T, ipamSubnet *models.IpamSubnetType, subnetUUID string) {
	if len(ipamSubnet.SubnetUUID) > 0 {
		assert.Equal(t, ipamSubnet.SubnetUUID, subnetUUID)
		return
	}
	_, err := uuid.Parse(subnetUUID)
	assert.NoError(t, err)
}
