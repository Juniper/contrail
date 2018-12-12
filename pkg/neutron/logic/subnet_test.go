package logic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/neutron/logic"
)

func TestSubnetResponse_CIDRFromVnc(t *testing.T) {
	tests := []struct {
		name              string
		subnet            *models.SubnetType
		sr                *logic.SubnetResponse
		expectedCIDR      string
		expectedIPVersion int64
	}{
		{
			name:              "CIDR from network with IPv4",
			subnet:            &models.SubnetType{IPPrefix: "10.0.1.2", IPPrefixLen: 32},
			sr:                &logic.SubnetResponse{},
			expectedCIDR:      "10.0.1.2/32",
			expectedIPVersion: 4,
		},
		{
			name:              "CIDR from network with IPv6",
			subnet:            &models.SubnetType{IPPrefix: "FC00::", IPPrefixLen: 8},
			sr:                &logic.SubnetResponse{},
			expectedCIDR:      "FC00::/8",
			expectedIPVersion: 6,
		},
		{
			name:              "Default CIDR if subnet is nil",
			subnet:            nil,
			sr:                &logic.SubnetResponse{},
			expectedCIDR:      "0.0.0.0/0",
			expectedIPVersion: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sr.CIDRFromVnc(tt.subnet)
			assert.Equal(t, tt.expectedCIDR, tt.sr.Cidr)
			assert.Equal(t, tt.expectedIPVersion, tt.sr.IPVersion)
		})
	}
}

func TestSubnetResponse_GatewayFromVnc(t *testing.T) {
	tests := []struct {
		name, gateway, expected string
		sr                      *logic.SubnetResponse
	}{
		{
			name:     "Gateway is not defined",
			gateway:  "",
			expected: "",
			sr:       &logic.SubnetResponse{},
		},
		{
			name:     "Gateway is set to 0.0.0.0",
			gateway:  "0.0.0.0",
			expected: "",
			sr:       &logic.SubnetResponse{},
		},
		{
			name:     "Gateway is defined",
			gateway:  "10.0.5.1",
			expected: "10.0.5.1",
			sr:       &logic.SubnetResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sr.GatewayFromVnc(tt.gateway)
			assert.Equal(t, tt.expected, tt.sr.GatewayIP)
		})
	}
}

func TestSubnetResponse_AllocationPoolsFromVnc(t *testing.T) {
	tests := []struct {
		name       string
		allocPools []*models.AllocationPoolType
		hasSubnet  bool
		sr         *logic.SubnetResponse
		expected   []*logic.AllocationPool
	}{
		{
			name:       "Allocation Pools when ipam has subnet and alloc pools are defined",
			allocPools: []*models.AllocationPoolType{{Start: "192.168.1.2", End: "192.168.1.255"}},
			hasSubnet:  true,
			sr:         &logic.SubnetResponse{GatewayIP: "192.168.1.1"},
			expected: []*logic.AllocationPool{
				{FirstIP: "192.168.1.2", LastIP: "192.168.1.255"},
			},
		},
		{
			name:       "Allocation Pools when ipam has subnet and alloc pools don't exist",
			allocPools: []*models.AllocationPoolType{},
			hasSubnet:  true,
			sr:         &logic.SubnetResponse{GatewayIP: "192.168.1.1", Cidr: "192.168.1.0/24"},
			expected: []*logic.AllocationPool{
				{FirstIP: "192.168.1.2", LastIP: "192.168.1.254"},
			},
		},
		{
			name:       "Allocation Pools when ipam has subnet, but gateway and alloc pools don't exist",
			allocPools: []*models.AllocationPoolType{},
			hasSubnet:  true,
			sr:         &logic.SubnetResponse{GatewayIP: "0.0.0.0", Cidr: "192.168.1.0/24"},
			expected: []*logic.AllocationPool{
				{FirstIP: "192.168.1.1", LastIP: "192.168.1.254"},
			},
		},
		{
			name:       "Allocation Pools without ipam subnet",
			allocPools: []*models.AllocationPoolType{},
			hasSubnet:  false,
			sr:         &logic.SubnetResponse{},
			expected: []*logic.AllocationPool{
				{FirstIP: "0.0.0.0", LastIP: "255.255.255.255"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sr.AllocationPoolsFromVnc(tt.allocPools, tt.hasSubnet)
			assert.Equal(t, tt.expected, tt.sr.AllocationPools)
		})
	}
}

func TestSubnetResponse_DNSNameServersFromVnc(t *testing.T) {
	tests := []struct {
		name        string
		dhcpOptions *models.DhcpOptionsListType
		sr          *logic.SubnetResponse
		expected    []*logic.DnsNameserver
	}{
		{
			name:        "DHCP options does not exist",
			dhcpOptions: nil,
			sr:          &logic.SubnetResponse{},
			expected:    nil,
		},
		{
			name: "DHCP with option 1",
			dhcpOptions: &models.DhcpOptionsListType{
				DHCPOption: []*models.DhcpOptionType{
					{
						DHCPOptionName:  "1",
						DHCPOptionValue: "0.0.255.255",
					},
				},
			},
			sr:       &logic.SubnetResponse{},
			expected: nil,
		},
		{
			name: "DHCP with option 6",
			dhcpOptions: &models.DhcpOptionsListType{
				DHCPOption: []*models.DhcpOptionType{
					{
						DHCPOptionName:  "6",
						DHCPOptionValue: "10.0.2.1",
					},
				},
			},
			sr: &logic.SubnetResponse{ID: "fake-subnet-id"},
			expected: []*logic.DnsNameserver{
				{
					Address:  "10.0.2.1",
					SubnetID: "fake-subnet-id",
				},
			},
		},
		{
			name: "DHCP with option 6 and multiple values",
			dhcpOptions: &models.DhcpOptionsListType{
				DHCPOption: []*models.DhcpOptionType{
					{
						DHCPOptionName:  "6",
						DHCPOptionValue: "10.0.2.2 10.0.3.12	10.0.4.12  10.0.5.5",
					},
				},
			},
			sr: &logic.SubnetResponse{ID: "fake-subnet-id"},
			expected: []*logic.DnsNameserver{
				{
					Address:  "10.0.2.2",
					SubnetID: "fake-subnet-id",
				},
				{
					Address:  "10.0.3.12",
					SubnetID: "fake-subnet-id",
				},
				{
					Address:  "10.0.4.12",
					SubnetID: "fake-subnet-id",
				},
				{
					Address:  "10.0.5.5",
					SubnetID: "fake-subnet-id",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sr.DNSNameServersFromVnc(tt.dhcpOptions)
			assert.Equal(t, tt.expected, tt.sr.DNSNameservers)
		})
	}
}

func TestSubnetResponse_HostRoutesFromVnc(t *testing.T) {
	tests := []struct {
		name       string
		routeTable *models.RouteTableType
		sr         *logic.SubnetResponse
		expected   []*logic.RouteTableType
	}{
		{
			name:       "Route table is nil",
			routeTable: nil,
			sr:         &logic.SubnetResponse{},
			expected:   nil,
		},
		{
			name: "Route table is empty",
			routeTable: &models.RouteTableType{
				Route: []*models.RouteType{},
			},
			sr:       &logic.SubnetResponse{},
			expected: nil,
		},
		{
			name: "Route table is defined",
			routeTable: &models.RouteTableType{
				Route: []*models.RouteType{
					{
						Prefix:  "10.0.3.10",
						NextHop: "10.0.3.1",
					},
				},
			},
			sr: &logic.SubnetResponse{ID: "fake-subnet-id"},
			expected: []*logic.RouteTableType{
				{
					Destination: "10.0.3.10",
					Nexthop:     "10.0.3.1",
					SubnetID:    "fake-subnet-id",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sr.HostRoutesFromVnc(tt.routeTable)
			assert.Equal(t, tt.expected, tt.sr.HostRoutes)
		})
	}
}
