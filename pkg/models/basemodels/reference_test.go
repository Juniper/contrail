package basemodels_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

func TestReferenceFieldName(t *testing.T) {
	tests := []struct {
		name string
		r    basemodels.Reference
		want string
	}{
		{"vn_vn ref", &models.VirtualNetworkVirtualNetworkRef{}, "virtual_network_refs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, basemodels.ReferenceFieldName(tt.r))
		})
	}
}
