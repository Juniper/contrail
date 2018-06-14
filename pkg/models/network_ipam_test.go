package models

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twinj/uuid"
)

func TestValidateIpInSubnet(t *testing.T) {
	name := "testnetwork"
	subnet := net.IPNet{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(24, 32)}

	t.Run("Validate ip from provided subnet", func(t *testing.T) {
		ip := "10.0.0.1"
		err := validateIPinSubnet(&subnet, name, ip)
		assert.NoError(t, err)
	})

	t.Run("Validate ip which is not from provided subnet", func(t *testing.T) {
		ip := "11.0.0.1"
		err := validateIPinSubnet(&subnet, name, ip)
		assert.Error(t, err)
	})
}

func TestValidateAllocationPools(t *testing.T) {
	subnet := net.IPNet{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(24, 32)}

	t.Run("Validate allocation pools which belong to provided subnet", func(t *testing.T) {
		allocPool := AllocationPoolType{Start: "10.0.0.5", End: "10.0.0.6"}
		err := allocPool.Validate(&subnet)
		assert.NoError(t, err)
	})

	t.Run("Validate allocation pools which don't belong to provided subnet", func(t *testing.T) {
		allocPool := AllocationPoolType{Start: "10.1.0.5", End: "10.1.0.6"}
		err := allocPool.Validate(&subnet)
		assert.Error(t, err)
	})
}

func TestValidateSubnetParams(t *testing.T) {
	subnet := SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}
	allocPools := []*AllocationPoolType{&AllocationPoolType{Start: "10.0.0.5", End: "10.0.0.6"}}
	wrongAllocPools := []*AllocationPoolType{&AllocationPoolType{Start: "10.0.0.5", End: "10.0.0.6"}, &AllocationPoolType{Start: "11.0.0.5", End: "11.0.0.6"}}

	t.Run("Validate subnet which allocation pools belong to the subnet provided", func(t *testing.T) {
		ipamSubnet := IpamSubnetType{Subnet: &subnet, AllocationPools: allocPools}
		err := ipamSubnet.ValidateSubnetParams()
		assert.NoError(t, err)
	})

	t.Run("Validate subnet which allocation pools where one don't belong to the subnet provided", func(t *testing.T) {
		ipamSubnet := IpamSubnetType{Subnet: &subnet, AllocationPools: wrongAllocPools}
		err := ipamSubnet.ValidateSubnetParams()
		assert.Error(t, err)
	})

	t.Run("Validate subnet with only ip and mask provided", func(t *testing.T) {
		ipamSubnet := IpamSubnetType{Subnet: &subnet}
		err := ipamSubnet.ValidateSubnetParams()
		assert.NoError(t, err)
	})

	t.Run("Validate subnet with gateway which belongs to provided subnet", func(t *testing.T) {
		ipamSubnet := IpamSubnetType{Subnet: &subnet, DefaultGateway: "10.0.0.1"}
		err := ipamSubnet.ValidateSubnetParams()
		assert.NoError(t, err)
	})

	t.Run("Validate subnet with gateway which doesn't belong to provided subnet", func(t *testing.T) {
		ipamSubnet := IpamSubnetType{Subnet: &subnet, DefaultGateway: "11.0.0.1"}
		err := ipamSubnet.ValidateSubnetParams()
		assert.Error(t, err)
	})

	t.Run("Validate subnet with DNS server in provided subnet", func(t *testing.T) {
		ipamSubnet := IpamSubnetType{Subnet: &subnet, DNSServerAddress: "10.0.0.2"}
		err := ipamSubnet.ValidateSubnetParams()
		assert.NoError(t, err)
	})

	t.Run("Validate succeeds with UUID provided", func(t *testing.T) {
		ipamSubnet := IpamSubnetType{Subnet: &subnet, SubnetUUID: uuid.NewV4().String()}
		err := ipamSubnet.ValidateSubnetParams()
		assert.NoError(t, err)
	})
}

func TestValidateIpamSubnet(t *testing.T) {
	subnet := SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}
	allocPools := []*AllocationPoolType{&AllocationPoolType{Start: "10.0.0.3", End: "10.0.0.25"}}

	t.Run("Validate ipam subnet with correct UUID", func(t *testing.T) {
		ipamSubnet := IpamSubnetType{Subnet: &subnet, SubnetUUID: uuid.NewV4().String(), DNSServerAddress: "10.0.0.2", DefaultGateway: "10.0.0.1", AllocationPools: allocPools}
		err := ipamSubnet.Validate()
		assert.NoError(t, err)
	})

	t.Run("Validate ipam subnet with wrong UUID", func(t *testing.T) {
		ipamSubnet := IpamSubnetType{Subnet: &subnet, SubnetUUID: "uuid", DNSServerAddress: "10.0.0.2", DefaultGateway: "10.0.0.1", AllocationPools: allocPools}
		err := ipamSubnet.Validate()
		assert.Error(t, err)
	})
}
