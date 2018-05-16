package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidMultiPolicyServiceChainConfig(t *testing.T) {
	var tests = []struct {
		name           string
		virtualNetwork *VirtualNetwork
		fails          bool
	}{
		{
			"check for rt",
			&VirtualNetwork{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList: &RouteTargetList{
					RouteTarget: []string{"100:101"},
				},
				ExportRouteTargetList: &RouteTargetList{
					RouteTarget: []string{"100:102"},
				},
			},
			false,
		},
		{
			"check for rt",
			&VirtualNetwork{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList: &RouteTargetList{
					RouteTarget: []string{"100:101"},
				},
				ExportRouteTargetList: &RouteTargetList{
					RouteTarget: []string{"100:101"},
				},
			},
			true,
		},
		{
			"check for MultiPolicyServiceChainsEnabled",
			&VirtualNetwork{
				MultiPolicyServiceChainsEnabled: false,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.virtualNetwork.IsValidMultiPolicyServiceChainConfig()
			if tt.fails {
				assert.False(t, res, tt.name)
			} else {
				assert.True(t, res, tt.name)
			}
		})
	}
}

func TestMakeNeutronCompatible(t *testing.T) {
	var tests = []struct {
		name           string
		virtualNetwork *VirtualNetwork
		expectedVN     *VirtualNetwork
	}{
		{
			"check for is shared",
			&VirtualNetwork{
				IsShared: true,
				Perms2:   &PermType2{},
			},
			&VirtualNetwork{
				IsShared: true,
				Perms2: &PermType2{
					GlobalAccess: PermsRWX,
				},
			},
		},
		{
			"check for RWX global access",
			&VirtualNetwork{
				IsShared: false,
				Perms2: &PermType2{
					GlobalAccess: PermsRWX,
				},
			},
			&VirtualNetwork{
				IsShared: true,
				Perms2: &PermType2{
					GlobalAccess: PermsRWX,
				},
			},
		},
		{
			"check for is shared and RWX global access ",
			&VirtualNetwork{
				IsShared: true,
				Perms2: &PermType2{
					GlobalAccess: PermsRWX,
				},
			},
			&VirtualNetwork{
				IsShared: true,
				Perms2: &PermType2{
					GlobalAccess: PermsRWX,
				},
			},
		},
		{
			"check for is shared and not RWX global access ",
			&VirtualNetwork{
				IsShared: true,
				Perms2: &PermType2{
					GlobalAccess: PermsW,
				},
			},
			&VirtualNetwork{
				IsShared: true,
				Perms2: &PermType2{
					GlobalAccess: PermsRWX,
				},
			},
		},
		{
			"check for PermsW global access ",
			&VirtualNetwork{
				IsShared: false,
				Perms2: &PermType2{
					GlobalAccess: PermsW,
				},
			},
			&VirtualNetwork{
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
			assert.Equal(t, tt.virtualNetwork, tt.expectedVN)
		})
	}
}
