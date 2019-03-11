package replication

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
)

func TestReplicator_portToVNCPortJSON(t *testing.T) {
	tests := []struct {
		name     string
		port     *models.Port
		wantJSON string
	}{
		{name: "nil", wantJSON: `{"port": null}`},
		{name: "empty", port: &models.Port{}, wantJSON: `{"port": {"parent_type": "end-system"}}`},
		{
			name:     "with bms info",
			port:     &models.Port{BMSPortInfo: &models.BaremetalPortInfo{}},
			wantJSON: `{"port": {"parent_type": "end-system", "port_bms_port_info": {}, "bms_port_info": {}}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Replicator{}
			b, err := json.Marshal(r.portToVNCPort(tt.port))
			assert.NoError(t, err)
			assert.JSONEq(t, tt.wantJSON, string(b))
		})
	}
}

func TestReplicator_nodeToVNCEndSystemJSON(t *testing.T) {
	tests := []struct {
		name     string
		node     *models.Node
		wantJSON string
	}{
		{name: "nil", wantJSON: `{"end-system": null}`},
		{name: "empty", node: &models.Node{}, wantJSON: `{"end-system": {}}`},
		{
			name:     "with hostname",
			node:     &models.Node{Hostname: "some-hostname"},
			wantJSON: `{"end-system": {"end_system_hostname": "some-hostname", "hostname": "some-hostname"}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Replicator{}
			b, err := json.Marshal(r.nodeToVNCEndSystem(tt.node))
			assert.NoError(t, err)
			assert.JSONEq(t, tt.wantJSON, string(b))
		})
	}
}
