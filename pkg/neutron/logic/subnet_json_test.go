package logic

import (
	"github.com/Juniper/contrail/pkg/format"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubnet_ApplyMap(t *testing.T) {
	tests := []struct {
		name     string
		m        map[string]interface{}
		expected Subnet
		fails    bool
	}{
		{
			name: "integer_instead_of_ipam_fq_name",
			m: map[string]interface{}{
				"ipam_fq_name": 1,
			},
			fails: true,
		},
		{
			name: "array_of_strings",
			m: map[string]interface{}{
				"ipam_fq_name": []interface{}{
					"default-domain",
					"default-project",
					"default-network-ipam",
				},
			},
			expected: Subnet{
				IpamFQName: []string{
					"default-domain",
					"default-project",
					"default-network-ipam",
				},
			},
		},
		{
			name: "string",
			m: map[string]interface{}{
				"ipam_fq_name": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Subnet{}
			err := format.ApplyMap(tt.m, &s)
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expected, s)
		})
	}
}
