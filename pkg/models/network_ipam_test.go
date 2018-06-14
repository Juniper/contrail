package models

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twinj/uuid"
)

type testNetworkIpamParams struct {
	allocPools     []*AllocationPoolType
	subnet         *SubnetType
	defaultGateway string
	dnsServer      string
	subnetUUID     string
}

func createTestIpamSubnet(testParams *testNetworkIpamParams) *IpamSubnetType {
	ipamSubnet := MakeIpamSubnetType()
	if len(testParams.allocPools) > 0 {
		ipamSubnet.AllocationPools = testParams.allocPools
	}
	ipamSubnet.Subnet = testParams.subnet
	ipamSubnet.DefaultGateway = testParams.defaultGateway
	ipamSubnet.DNSServerAddress = testParams.dnsServer
	ipamSubnet.SubnetUUID = testParams.subnetUUID
	return ipamSubnet
}

func TestIsIpInSubnet(t *testing.T) {
	subnet := net.IPNet{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(24, 32)}

	t.Run("Validate ip from provided subnet", func(t *testing.T) {
		ip := "10.0.0.1"
		err := isIpInSubnet(&subnet, ip)
		assert.NoError(t, err)
	})

	t.Run("Validate ip which is not from provided subnet", func(t *testing.T) {
		ip := "11.0.0.1"
		err := isIpInSubnet(&subnet, ip)
		assert.Error(t, err)
	})
}

func TestAllocPoolIsInSubnet(t *testing.T) {
	subnet := net.IPNet{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(24, 32)}

	t.Run("Validate allocation pools which belong to provided subnet", func(t *testing.T) {
		allocPool := AllocationPoolType{Start: "10.0.0.5", End: "10.0.0.6"}
		err := allocPool.IsInSubnet(&subnet)
		assert.NoError(t, err)
	})

	t.Run("Validate allocation pools which don't belong to provided subnet", func(t *testing.T) {
		allocPool := AllocationPoolType{Start: "10.1.0.5", End: "10.1.0.6"}
		err := allocPool.IsInSubnet(&subnet)
		assert.Error(t, err)
	})
}

func TestCheckIfSubnetParamsAreValid(t *testing.T) {
	tests := []struct {
		name       string
		testParams *testNetworkIpamParams
		fails      bool
	}{
		{
			name: "Validate subnet which allocation pools belong to the subnet provided",
			testParams: &testNetworkIpamParams{
				subnet:     &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				allocPools: []*AllocationPoolType{{Start: "10.0.0.5", End: "10.0.0.6"}},
			},
			fails: false,
		},
		{
			name: "Validate subnet which allocation pools where one don't belong to the subnet provided",
			testParams: &testNetworkIpamParams{
				subnet:     &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				allocPools: []*AllocationPoolType{{Start: "10.0.0.5", End: "10.0.0.6"}, {Start: "11.0.0.5", End: "11.0.0.6"}},
			},
			fails: true,
		},
		{
			name: "Validate subnet with only ip and mask provided",
			testParams: &testNetworkIpamParams{
				subnet: &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
			},
			fails: false,
		},
		{
			name: "Validate subnet with gateway which belongs to provided subnet",
			testParams: &testNetworkIpamParams{
				subnet:         &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				defaultGateway: "10.0.0.1",
			},
			fails: false,
		},
		{
			name: "Validate subnet with gateway which doesn't belong to provided subnet",
			testParams: &testNetworkIpamParams{
				subnet:         &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				defaultGateway: "11.0.0.1",
			},
			fails: true,
		},
		{
			name: "Validate subnet with DNS server in provided subnet",
			testParams: &testNetworkIpamParams{
				subnet:    &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				dnsServer: "10.0.0.2",
			},
			fails: false,
		},
		{
			name: "Validate subnet with DNS server not in provided subnet",
			testParams: &testNetworkIpamParams{
				subnet:    &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				dnsServer: "11.0.0.2",
			},
			fails: true,
		},
		{
			name: "Validate succeeds with UUID provided",
			testParams: &testNetworkIpamParams{
				subnet:     &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				subnetUUID: uuid.NewV4().String(),
			},
			fails: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ipamSubnet := createTestIpamSubnet(tt.testParams)
			err := ipamSubnet.CheckIfSubnetParamsAreValid()
			if tt.fails {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestIpamSubnetValidate(t *testing.T) {
	t.Run("Validate ipam subnet with correct UUID", func(t *testing.T) {
		testParams := &testNetworkIpamParams{
			subnet:         &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
			allocPools:     []*AllocationPoolType{{Start: "10.0.0.3", End: "10.0.0.25"}},
			subnetUUID:     uuid.NewV4().String(),
			dnsServer:      "10.0.0.2",
			defaultGateway: "10.0.0.1",
		}
		ipamSubnet := createTestIpamSubnet(testParams)
		err := ipamSubnet.Validate()
		assert.NoError(t, err)
	})

	t.Run("Validate ipam subnet with wrong UUID", func(t *testing.T) {
		testParams := &testNetworkIpamParams{
			subnet:         &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
			allocPools:     []*AllocationPoolType{{Start: "10.0.0.3", End: "10.0.0.25"}},
			subnetUUID:     "uuid",
			dnsServer:      "10.0.0.2",
			defaultGateway: "10.0.0.1",
		}
		ipamSubnet := createTestIpamSubnet(testParams)
		err := ipamSubnet.Validate()
		assert.Error(t, err)
	})
}
