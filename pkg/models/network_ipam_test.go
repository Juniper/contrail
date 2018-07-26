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
		err := isIPInSubnet(&subnet, ip)
		assert.NoError(t, err)
	})

	t.Run("Validate ip which is not from provided subnet", func(t *testing.T) {
		ip := "11.0.0.1"
		err := isIPInSubnet(&subnet, ip)
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
			err := ipamSubnet.ValidateSubnetParams()
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

func TestIpamSubnetTypeContains(t *testing.T) {
	type fields struct {
		AllocationPools []*AllocationPoolType
		Subnet          *SubnetType
	}
	type args struct {
		ip net.IP
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "IP in subnet",
			fields: fields{
				Subnet: &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
			},
			args: args{
				ip: net.ParseIP("10.0.0.1"),
			},
			want: true,
		},
		{
			name: "IP in subnet (begin)",
			fields: fields{
				Subnet: &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
			},
			args: args{
				ip: net.ParseIP("10.0.0.0"),
			},
			want: true,
		},
		{
			name: "IP in subnet (end)",
			fields: fields{
				Subnet: &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
			},
			args: args{
				ip: net.ParseIP("10.0.0.255"),
			},
			want: true,
		},
		{
			name: "IP in allocation pool",
			fields: fields{
				Subnet: &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				AllocationPools: []*AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.255",
					},
				},
			},
			args: args{
				ip: net.ParseIP("10.0.0.1"),
			},
			want: true,
		},
		{
			name: "IP in allocation pool (begin)",
			fields: fields{
				Subnet: &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				AllocationPools: []*AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.255",
					},
				},
			},
			args: args{
				ip: net.ParseIP("10.0.0.0"),
			},
			want: true,
		},
		{
			name: "IP in allocation pool (end)",
			fields: fields{
				Subnet: &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				AllocationPools: []*AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.255",
					},
				},
			},
			args: args{
				ip: net.ParseIP("10.0.0.255"),
			},
			want: true,
		},
		{
			name: "IP outside of allocation pool",
			fields: fields{
				Subnet: &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				AllocationPools: []*AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.255",
					},
				},
			},
			args: args{
				ip: net.ParseIP("127.0.0.1"),
			},
		},
		{
			name: "IP outside of subnet",
			fields: fields{
				Subnet: &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
			},
			args: args{
				ip: net.ParseIP("127.0.0.1"),
			},
		},
		{
			name: "IP in second allocation pool",
			fields: fields{
				Subnet: &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				AllocationPools: []*AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.125",
					},
					{
						Start: "10.0.0.130",
						End:   "10.0.0.255",
					},
				},
			},
			args: args{
				ip: net.ParseIP("10.0.0.131"),
			},
			want: true,
		},
		{
			name: "IP between allocation pools",
			fields: fields{
				Subnet: &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				AllocationPools: []*AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "10.0.0.125",
					},
					{
						Start: "10.0.0.130",
						End:   "10.0.0.255",
					},
				},
			},
			args: args{
				ip: net.ParseIP("10.0.0.127"),
			},
		},
		{
			name: "Invalid data",
			fields: fields{
				Subnet: &SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				AllocationPools: []*AllocationPoolType{
					{
						Start: "10.0.0.0",
						End:   "dead-beaf",
					},
				},
			},
			args: args{
				ip: net.ParseIP("10.0.0.1"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &IpamSubnetType{
				Subnet:          tt.fields.Subnet,
				AllocationPools: tt.fields.AllocationPools,
			}
			got, err := m.Contains(tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("IpamSubnetType.Contains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IpamSubnetType.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIpamSubnetsSubstract(t *testing.T) {
	type args struct {
		left  []*IpamSubnetType
		right []*IpamSubnetType
	}
	tests := []struct {
		name string
		args args
		want []*IpamSubnetType
	}{
		{
			name: "Simple substraction",
			args: args{
				left: []*IpamSubnetType{
					{
						SubnetUUID: "uuid-1",
					},
					{
						SubnetUUID: "uuid-2",
					},
				},
				right: []*IpamSubnetType{
					{
						SubnetUUID: "uuid-1",
					},
				},
			},
			want: []*IpamSubnetType{
				{
					SubnetUUID: "uuid-2",
				},
			},
		},
		{
			name: "Equal sets substraction",
			args: args{
				left: []*IpamSubnetType{
					{
						SubnetUUID: "uuid-1",
					},
				},
				right: []*IpamSubnetType{
					{
						SubnetUUID: "uuid-1",
					},
				},
			},
			want: nil,
		},
		{
			name: "Bigger set substraction",
			args: args{
				left: []*IpamSubnetType{
					{
						SubnetUUID: "uuid-1",
					},
				},
				right: []*IpamSubnetType{
					{
						SubnetUUID: "uuid-1",
					},
					{
						SubnetUUID: "uuid-2",
					},
				},
			},
			want: nil,
		},
		{
			name: "Substract empty set",
			args: args{
				left: []*IpamSubnetType{
					{
						SubnetUUID: "uuid-1",
					},
					{
						SubnetUUID: "uuid-2",
					},
				},
				right: []*IpamSubnetType{},
			},
			want: []*IpamSubnetType{
				{
					SubnetUUID: "uuid-1",
				},
				{
					SubnetUUID: "uuid-2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IpamSubnetsSubstract(tt.args.left, tt.args.right)
			assert.Equal(t, tt.want, got)
		})
	}
}
