package logic_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/neutron/logic"
	neutronmock "github.com/Juniper/contrail/pkg/neutron/mock"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
				{Start: "192.168.1.2", End: "192.168.1.255"},
			},
		},
		{
			name:       "Allocation Pools when ipam has subnet and alloc pools don't exist",
			allocPools: []*models.AllocationPoolType{},
			hasSubnet:  true,
			sr:         &logic.SubnetResponse{GatewayIP: "192.168.1.1", Cidr: "192.168.1.0/24"},
			expected: []*logic.AllocationPool{
				{Start: "192.168.1.2", End: "192.168.1.254"},
			},
		},
		{
			name:       "Allocation Pools when ipam has subnet, but gateway and alloc pools don't exist",
			allocPools: []*models.AllocationPoolType{},
			hasSubnet:  true,
			sr:         &logic.SubnetResponse{GatewayIP: "0.0.0.0", Cidr: "192.168.1.0/24"},
			expected: []*logic.AllocationPool{
				{Start: "192.168.1.1", End: "192.168.1.254"},
			},
		},
		{
			name:       "Allocation Pools without ipam subnet",
			allocPools: []*models.AllocationPoolType{},
			hasSubnet:  false,
			sr:         &logic.SubnetResponse{},
			expected: []*logic.AllocationPool{
				{Start: "0.0.0.0", End: "255.255.255.255"},
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

// func TestSubnetResponse_DNSNameServersFromVnc(t *testing.T) {
// 	tests := []struct {
// 		name        string
// 		dhcpOptions *models.DhcpOptionsListType
// 		sr          *logic.SubnetResponse
// 		expected    []*logic.DnsNameserver
// 	}{
// 		{
// 			name:        "DHCP options does not exist",
// 			dhcpOptions: nil,
// 			sr:          &logic.SubnetResponse{},
// 			expected:    []*logic.DnsNameserver{},
// 		},
// 		{
// 			name: "DHCP with option 1",
// 			dhcpOptions: &models.DhcpOptionsListType{
// 				DHCPOption: []*models.DhcpOptionType{
// 					{
// 						DHCPOptionName:  "1",
// 						DHCPOptionValue: "0.0.255.255",
// 					},
// 				},
// 			},
// 			sr:       &logic.SubnetResponse{},
// 			expected: []*logic.DnsNameserver{},
// 		},
// 		{
// 			name: "DHCP with option 6",
// 			dhcpOptions: &models.DhcpOptionsListType{
// 				DHCPOption: []*models.DhcpOptionType{
// 					{
// 						DHCPOptionName:  "6",
// 						DHCPOptionValue: "10.0.2.1",
// 					},
// 				},
// 			},
// 			sr: &logic.SubnetResponse{ID: "fake-subnet-id"},
// 			expected: []*logic.DnsNameserver{
// 				{
// 					Address:  "10.0.2.1",
// 					SubnetID: "fake-subnet-id",
// 				},
// 			},
// 		},
// 		{
// 			name: "DHCP with option 6 and multiple values",
// 			dhcpOptions: &models.DhcpOptionsListType{
// 				DHCPOption: []*models.DhcpOptionType{
// 					{
// 						DHCPOptionName: "6",
// 						DHCPOptionValue: "10.0.2.2 10.0.3.12	10.0.4.12  10.0.5.5",
// 					},
// 				},
// 			},
// 			sr: &logic.SubnetResponse{ID: "fake-subnet-id"},
// 			expected: []*logic.DnsNameserver{
// 				{
// 					Address:  "10.0.2.2",
// 					SubnetID: "fake-subnet-id",
// 				},
// 				{
// 					Address:  "10.0.3.12",
// 					SubnetID: "fake-subnet-id",
// 				},
// 				{
// 					Address:  "10.0.4.12",
// 					SubnetID: "fake-subnet-id",
// 				},
// 				{
// 					Address:  "10.0.5.5",
// 					SubnetID: "fake-subnet-id",
// 				},
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.sr.DNSNameServersFromVnc(tt.dhcpOptions)
// 			assert.Equal(t, tt.expected, tt.sr.DNSNameservers)
// 		})
// 	}
// }

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
			expected:   []*logic.RouteTableType{},
		},
		{
			name: "Route table is empty",
			routeTable: &models.RouteTableType{
				Route: []*models.RouteType{},
			},
			sr:       &logic.SubnetResponse{},
			expected: []*logic.RouteTableType{},
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

// func TestSubnet_ReadAll(t *testing.T) {
// 	type mockVN struct {
// 		VirtualNetworks *services.ListVirtualNetworkResponse
// 		Error           error
// 	}

// 	type mockKVs struct {
// 		Response *services.RetrieveValuesResponse
// 		Error    error
// 	}

// 	tests := []struct {
// 		name    string
// 		filters logic.Filters
// 		fields  logic.Fields

// 		expected []*logic.SubnetResponse

// 		mockVN  mockVN
// 		mockKVs mockKVs
// 	}{
// 		{
// 			name: "No virtual networks",
// 			mockVN: mockVN{
// 				VirtualNetworks: &services.ListVirtualNetworkResponse{},
// 			},
// 			expected: []*logic.SubnetResponse{},
// 		},
// 		{
// 			name:    "Without filters",
// 			filters: logic.Filters{},
// 			mockVN: mockVN{
// 				VirtualNetworks: &services.ListVirtualNetworkResponse{
// 					VirtualNetworks: []*models.VirtualNetwork{
// 						fakeVirtualNetwork("blue", 2, false),
// 						fakeVirtualNetwork("red", 1, false),
// 						fakeVirtualNetwork("green", 0, false),
// 					},
// 				},
// 			},
// 			expected: []*logic.SubnetResponse{
// 				{
// 					NetworkID:       "virtual_network_blue",
// 					ID:              "subnet_blue_1_uuid",
// 					Cidr:            "10.0.100.0/24",
// 					GatewayIP:       "10.0.100.1",
// 					AllocationPools: []*logic.AllocationPool{{Start: "10.0.100.2", End: "10.0.100.254"}},
// 					HostRoutes:      []*logic.RouteTableType{},
// 					DNSNameservers:  []*logic.DnsNameserver{},
// 					IPVersion:       4,
// 				},
// 				{
// 					NetworkID:       "virtual_network_blue",
// 					ID:              "subnet_blue_2_uuid",
// 					Cidr:            "10.0.101.0/24",
// 					GatewayIP:       "10.0.101.1",
// 					AllocationPools: []*logic.AllocationPool{{Start: "10.0.101.2", End: "10.0.101.254"}},
// 					HostRoutes:      []*logic.RouteTableType{},
// 					DNSNameservers:  []*logic.DnsNameserver{},
// 					IPVersion:       4,
// 				},
// 				{
// 					NetworkID:       "virtual_network_red",
// 					ID:              "subnet_red_1_uuid",
// 					Cidr:            "10.0.100.0/24",
// 					GatewayIP:       "10.0.100.1",
// 					AllocationPools: []*logic.AllocationPool{{Start: "10.0.100.2", End: "10.0.100.254"}},
// 					HostRoutes:      []*logic.RouteTableType{},
// 					DNSNameservers:  []*logic.DnsNameserver{},
// 					IPVersion:       4,
// 				},
// 			},
// 		},
// 		{
// 			name:    "With ID filters",
// 			filters: logic.Filters{"id": []string{"subnet_blue_1_uuid"}},
// 			mockVN: mockVN{
// 				VirtualNetworks: &services.ListVirtualNetworkResponse{
// 					VirtualNetworks: []*models.VirtualNetwork{
// 						fakeVirtualNetwork("blue", 1, false),
// 					},
// 				},
// 			},
// 			expected: []*logic.SubnetResponse{
// 				{
// 					NetworkID:       "virtual_network_blue",
// 					ID:              "subnet_blue_1_uuid",
// 					Cidr:            "10.0.100.0/24",
// 					GatewayIP:       "10.0.100.1",
// 					AllocationPools: []*logic.AllocationPool{{Start: "10.0.100.2", End: "10.0.100.254"}},
// 					HostRoutes:      []*logic.RouteTableType{},
// 					DNSNameservers:  []*logic.DnsNameserver{},
// 					IPVersion:       4,
// 				},
// 			},
// 		},
// 		{
// 			name: "With shared and router:external filters",
// 			filters: logic.Filters{
// 				"router:external": []string{"router_blue"},
// 				"shared":          []string{"true"},
// 			},
// 			mockVN: mockVN{
// 				VirtualNetworks: &services.ListVirtualNetworkResponse{
// 					VirtualNetworks: []*models.VirtualNetwork{
// 						fakeVirtualNetwork("blue", 1, true),
// 					},
// 				},
// 			},
// 			expected: []*logic.SubnetResponse{
// 				{
// 					Shared:          true,
// 					NetworkID:       "virtual_network_blue",
// 					ID:              "subnet_blue_1_uuid",
// 					Cidr:            "10.0.100.0/24",
// 					GatewayIP:       "10.0.100.1",
// 					AllocationPools: []*logic.AllocationPool{{Start: "10.0.100.2", End: "10.0.100.254"}},
// 					HostRoutes:      []*logic.RouteTableType{},
// 					DNSNameservers:  []*logic.DnsNameserver{},
// 					IPVersion:       4,
// 				},
// 			},
// 		},
// 		{
// 			name:    "Duplicated virtual network should be skipped",
// 			filters: logic.Filters{},
// 			mockVN: mockVN{
// 				VirtualNetworks: &services.ListVirtualNetworkResponse{
// 					VirtualNetworks: []*models.VirtualNetwork{
// 						fakeVirtualNetwork("blue", 1, false),
// 						fakeVirtualNetwork("blue", 1, false),
// 						fakeVirtualNetwork("blue", 1, false),
// 					},
// 				},
// 			},
// 			expected: []*logic.SubnetResponse{
// 				{
// 					NetworkID:       "virtual_network_blue",
// 					ID:              "subnet_blue_1_uuid",
// 					Cidr:            "10.0.100.0/24",
// 					GatewayIP:       "10.0.100.1",
// 					AllocationPools: []*logic.AllocationPool{{Start: "10.0.100.2", End: "10.0.100.254"}},
// 					HostRoutes:      []*logic.RouteTableType{},
// 					DNSNameservers:  []*logic.DnsNameserver{},
// 					IPVersion:       4,
// 				},
// 			},
// 		},
// 	}

// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			rp := logic.RequestParameters{
// 				ReadService: mockReadService(mockCtrl, tt.mockVN.VirtualNetworks, tt.mockVN.Error),
// 				UserAgentKV: mockUserAgentService(mockCtrl, tt.mockKVs.Response, tt.mockKVs.Error),
// 			}

// 			subnet := &logic.Subnet{}
// 			result, err := subnet.ReadAll(context.Background(), rp, tt.filters, tt.fields)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tt.expected, result)
// 		})
// 	}
// }

// func TestSubnet_Read(t *testing.T) {
// 	type mockVN struct {
// 		VirtualNetworks *services.ListVirtualNetworkResponse
// 		Error           error
// 	}

// 	type mockKVs struct {
// 		Response *services.RetrieveValuesResponse
// 		Error    error
// 	}

// 	tests := []struct {
// 		name string
// 		id   string

// 		expected interface{}
// 		fails    bool

// 		mockVN  mockVN
// 		mockKVs mockKVs
// 	}{
// 		{
// 			name: "No virtual networks",
// 			mockVN: mockVN{
// 				VirtualNetworks: &services.ListVirtualNetworkResponse{},
// 			},
// 			fails: true,
// 		},
// 		{
// 			name: "With correct id",
// 			id:   "subnet_green_1_uuid",
// 			mockVN: mockVN{
// 				VirtualNetworks: &services.ListVirtualNetworkResponse{
// 					VirtualNetworks: []*models.VirtualNetwork{
// 						fakeVirtualNetwork("green", 1, false),
// 					},
// 				},
// 			},
// 			expected: &logic.SubnetResponse{
// 				NetworkID:       "virtual_network_green",
// 				ID:              "subnet_green_1_uuid",
// 				Cidr:            "10.0.100.0/24",
// 				GatewayIP:       "10.0.100.1",
// 				AllocationPools: []*logic.AllocationPool{{Start: "10.0.100.2", End: "10.0.100.254"}},
// 				HostRoutes:      []*logic.RouteTableType{},
// 				DNSNameservers:  []*logic.DnsNameserver{},
// 				IPVersion:       4,
// 			},
// 		},
// 		{
// 			name: "With incorrect id",
// 			id:   "does_not_exist",
// 			mockVN: mockVN{
// 				VirtualNetworks: &services.ListVirtualNetworkResponse{
// 					VirtualNetworks: []*models.VirtualNetwork{
// 						fakeVirtualNetwork("green", 1, false),
// 					},
// 				},
// 			},
// 			fails: true,
// 		},
// 	}

// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			rp := logic.RequestParameters{
// 				ReadService: mockReadService(mockCtrl, tt.mockVN.VirtualNetworks, tt.mockVN.Error),
// 				UserAgentKV: mockUserAgentService(mockCtrl, tt.mockKVs.Response, tt.mockKVs.Error),
// 			}

// 			subnet := &logic.Subnet{}
// 			result, err := subnet.Read(context.Background(), rp, tt.id)
// 			if tt.fails {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 			assert.Equal(t, tt.expected, result)
// 		})
// 	}
// }

func mockReadService(
	mockCtrl *gomock.Controller,
	vns *services.ListVirtualNetworkResponse,
	err error,
) *servicesmock.MockService {
	mock := servicesmock.NewMockService(mockCtrl)
	mock.EXPECT().ListVirtualNetwork(gomock.Any(), gomock.Any()).Return(vns, err).AnyTimes()
	return mock
}

func mockUserAgentService(
	mockCtrl *gomock.Controller,
	res *services.RetrieveValuesResponse,
	err error,
) *neutronmock.MockuserAgentKVServer {
	mock := neutronmock.NewMockuserAgentKVServer(mockCtrl)
	mock.EXPECT().RetrieveValues(gomock.Any(), gomock.Any()).Return(res, err).AnyTimes()
	return mock
}

func fakeVirtualNetwork(name string, subnets int, shared bool) *models.VirtualNetwork {
	var ipamSubnets []*models.IpamSubnetType
	for i := 1; i <= subnets; i++ {
		ipamSubnets = append(ipamSubnets, &models.IpamSubnetType{
			DefaultGateway: fmt.Sprintf("10.0.%d.1", i+99),
			SubnetUUID:     fmt.Sprintf("subnet_%s_%d_uuid", name, i),
			Subnet: &models.SubnetType{
				IPPrefix:    fmt.Sprintf("10.0.%d.0", i+99),
				IPPrefixLen: 24,
			},
		})
	}

	return &models.VirtualNetwork{
		IsShared: shared,
		UUID:     fmt.Sprintf("virtual_network_%s", name),
		NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
			{Attr: &models.VnSubnetsType{IpamSubnets: ipamSubnets}},
		},
	}
}

func Test_GetHostPrefixes(t *testing.T) {
	tests := []struct {
		name       string
		hostRoutes []*logic.RouteTableType
		subnetCIDR string
		want       map[string][]string
	}{
		{
			name:       "Simple",
			subnetCIDR: "12.5.3.0/24",
			hostRoutes: []*logic.RouteTableType{
				{
					Destination: "10.0.0.0/24",
					Nexthop:     "12.5.3.2",
				},
				{
					Destination: "12.0.0.0/24",
					Nexthop:     "12.5.3.4",
				},
				{
					Destination: "14.0.0.0/24",
					Nexthop:     "12.5.3.23",
				},
			},
			want: map[string][]string{
				"12.5.3.2":  []string{"10.0.0.0/24"},
				"12.5.3.4":  []string{"12.0.0.0/24"},
				"12.5.3.23": []string{"14.0.0.0/24"},
			},
		},
		{
			name:       "not Simple",
			subnetCIDR: "8.0.0.0/24",
			hostRoutes: []*logic.RouteTableType{
				{
					Destination: "10.0.0.0/24",
					Nexthop:     "8.0.0.2",
				},
				{
					Destination: "12.0.0.0/24",
					Nexthop:     "10.0.0.4",
				},
				{
					Destination: "14.0.0.0/24",
					Nexthop:     "12.0.0.23",
				},
				{
					Destination: "16.0.0.0/24",
					Nexthop:     "8.0.0.4",
				},
				{
					Destination: "15.0.0.0/24",
					Nexthop:     "16.0.0.2",
				},
				{
					Destination: "20.0.0.0/24",
					Nexthop:     "8.0.0.12",
				},
			},
			want: map[string][]string{
				"8.0.0.2":  []string{"10.0.0.0/24", "12.0.0.0/24", "14.0.0.0/24"},
				"8.0.0.4":  []string{"16.0.0.0/24", "15.0.0.0/24"},
				"8.0.0.12": []string{"20.0.0.0/24"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := logic.GetHostPrefixes(tt.hostRoutes, tt.subnetCIDR); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHostPrefixes() = %v, want %v", got, tt.want)
			}
		})
	}
}
