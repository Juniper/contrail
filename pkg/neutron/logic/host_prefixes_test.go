package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getMapFromHostPrefixes(hp HostPrefixes) map[string][]string {
	result := make(map[string][]string)
	for _, addr := range hp.GetIPAddresses() {
		for _, destination := range hp.GetDestinationsForIP(addr) {
			result[addr] = append(result[addr], destination)
		}
	}
	return result
}

func Test_GetHostPrefixes(t *testing.T) {
	tests := []struct {
		name       string
		hostRoutes []*RouteTableType
		subnetCIDR string
		want       map[string][]string
	}{
		{
			name:       "Simple",
			subnetCIDR: "12.5.3.0/24",
			hostRoutes: []*RouteTableType{
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
				"12.5.3.2":  {"10.0.0.0/24"},
				"12.5.3.4":  {"12.0.0.0/24"},
				"12.5.3.23": {"14.0.0.0/24"},
			},
		},
		{
			name:       "not Simple",
			subnetCIDR: "8.0.0.0/24",
			hostRoutes: []*RouteTableType{
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
				"8.0.0.2":  {"10.0.0.0/24", "12.0.0.0/24", "14.0.0.0/24"},
				"8.0.0.4":  {"16.0.0.0/24", "15.0.0.0/24"},
				"8.0.0.12": {"20.0.0.0/24"},
			},
		},
		{
			name:       "not Simple in a different order",
			subnetCIDR: "8.0.0.0/24",
			hostRoutes: []*RouteTableType{
				{
					Destination: "20.0.0.0/24",
					Nexthop:     "8.0.0.12",
				},
				{
					Destination: "12.0.0.0/24",
					Nexthop:     "10.0.0.4",
				},
				{
					Destination: "10.0.0.0/24",
					Nexthop:     "8.0.0.2",
				},
				{
					Destination: "14.0.0.0/24",
					Nexthop:     "12.0.0.23",
				},
				{
					Destination: "15.0.0.0/24",
					Nexthop:     "16.0.0.2",
				},
				{
					Destination: "16.0.0.0/24",
					Nexthop:     "8.0.0.4",
				},
			},
			want: map[string][]string{
				"8.0.0.2":  {"10.0.0.0/24", "12.0.0.0/24", "14.0.0.0/24"},
				"8.0.0.4":  {"16.0.0.0/24", "15.0.0.0/24"},
				"8.0.0.12": {"20.0.0.0/24"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHostPrefixes(tt.hostRoutes, tt.subnetCIDR)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, getMapFromHostPrefixes(got))
		})
	}
}
