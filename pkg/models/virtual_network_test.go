package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidMultiPolicyServiceChainConfig(t *testing.T) {
	var tests = []struct {
		name           string
		virtualNetwork *VirtualNetwork
		expected       bool
	}{
		{
			name:     "check for rt",
			expected: true,
			virtualNetwork: &VirtualNetwork{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList: &RouteTargetList{
					RouteTarget: []string{"100:101"},
				},
				ExportRouteTargetList: &RouteTargetList{
					RouteTarget: []string{"100:102"},
				},
			},
		},
		{
			name:     "check for rt",
			expected: false,
			virtualNetwork: &VirtualNetwork{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList: &RouteTargetList{
					RouteTarget: []string{"100:101"},
				},
				ExportRouteTargetList: &RouteTargetList{
					RouteTarget: []string{"100:101"},
				},
			},
		},
		{
			name:     "check for multi-policy service chains disabled",
			expected: true,
			virtualNetwork: &VirtualNetwork{
				MultiPolicyServiceChainsEnabled: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.virtualNetwork.IsValidMultiPolicyServiceChainConfig()
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestMakeNeutronCompatible(t *testing.T) {
	var tests = []struct {
		name           string
		virtualNetwork *VirtualNetwork
		expected       *VirtualNetwork
	}{
		{
			name: "check for is shared",
			virtualNetwork: &VirtualNetwork{
				IsShared: true,
				Perms2:   &PermType2{},
			},
			expected: &VirtualNetwork{
				IsShared: true,
				Perms2: &PermType2{
					GlobalAccess: PermsRWX,
				},
			},
		},
		{
			name: "check for RWX global access",
			virtualNetwork: &VirtualNetwork{
				IsShared: false,
				Perms2: &PermType2{
					GlobalAccess: PermsRWX,
				},
			},
			expected: &VirtualNetwork{
				IsShared: true,
				Perms2: &PermType2{
					GlobalAccess: PermsRWX,
				},
			},
		},
		{
			name: "check for is shared and RWX global access ",
			virtualNetwork: &VirtualNetwork{
				IsShared: true,
				Perms2: &PermType2{
					GlobalAccess: PermsRWX,
				},
			},
			expected: &VirtualNetwork{
				IsShared: true,
				Perms2: &PermType2{
					GlobalAccess: PermsRWX,
				},
			},
		},
		{
			name: "check for is shared and not RWX global access ",
			virtualNetwork: &VirtualNetwork{
				IsShared: true,
				Perms2: &PermType2{
					GlobalAccess: PermsW,
				},
			},
			expected: &VirtualNetwork{
				IsShared: true,
				Perms2: &PermType2{
					GlobalAccess: PermsRWX,
				},
			},
		},
		{
			name: "check for PermsW global access ",
			virtualNetwork: &VirtualNetwork{
				IsShared: false,
				Perms2: &PermType2{
					GlobalAccess: PermsW,
				},
			},
			expected: &VirtualNetwork{
				IsShared: false,
				Perms2: &PermType2{
					GlobalAccess: PermsW,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.virtualNetwork.MakeNeutronCompatible()
			assert.Equal(t, tt.virtualNetwork, tt.expected)
		})
	}
}
