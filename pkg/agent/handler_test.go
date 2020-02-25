package agent

import (
	"testing"

	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/deploy"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestGenerateClusterConfig(t *testing.T) {
	generateClusterConfigCases := map[string]struct {
		inputEvent     *services.Event
		inputConfig    *Config
		expectedConfig *deploy.Config
	}{
		"Contrail Creation Event": {
			inputEvent: &services.Event{
				Request: &services.Event_CreateContrailClusterRequest{
					CreateContrailClusterRequest: &services.CreateContrailClusterRequest{
						ContrailCluster: &models.ContrailCluster{
							UUID:            "1",
							ProvisionerType: "openstack",
						},
					},
				},
			},
			inputConfig: &Config{
				ID:                  "2",
				Password:            "admin",
				DomainID:            "3",
				ProjectID:           "4",
				DomainName:          "a",
				ProjectName:         "b",
				AuthURL:             "A",
				Endpoint:            "B",
				InSecure:            true,
				ServiceUserID:       "5",
				ServiceUserPassword: "minda",
			},
			expectedConfig: &deploy.Config{
				ID:                  "2",
				Password:            "admin",
				DomainID:            "3",
				ProjectID:           "4",
				DomainName:          "a",
				ProjectName:         "b",
				AuthURL:             "A",
				Endpoint:            "B",
				InSecure:            true,
				ResourceType:        basemodels.KindToSchemaID("contrail-cluster"),
				ResourceID:          "1",
				Action:              services.OperationCreate,
				ProvisionerType:     "openstack",
				LogLevel:            "debug",
				LogFile:             "/var/log/contrail/deploy.log",
				TemplateRoot:        "/usr/share/contrail/templates/",
				ServiceUserID:       "5",
				ServiceUserPassword: "minda",
			},
		},
	}
	for name, generateClusterConfigCase := range generateClusterConfigCases {
		t.Run(
			name,
			func(t *testing.T) {
				clusterConfig, err := generateClusterConfig(
					generateClusterConfigCase.inputEvent,
					generateClusterConfigCase.inputConfig,
				)

				assert.NoError(t, err)
				assert.Equal(t, generateClusterConfigCase.expectedConfig, clusterConfig)
			},
		)
	}
}

func TestGenerateCloudConfig(t *testing.T) {
	generateCloudConfigCases := map[string]struct {
		inputEvent     *services.Event
		inputConfig    *Config
		expectedConfig *cloud.Config
	}{
		"Contrail Creation Event": {
			inputEvent: &services.Event{
				Request: &services.Event_CreateContrailClusterRequest{
					CreateContrailClusterRequest: &services.CreateContrailClusterRequest{
						ContrailCluster: &models.ContrailCluster{
							UUID: "1",
						},
					},
				},
			},
			inputConfig: &Config{
				ID:          "2",
				Password:    "admin",
				DomainID:    "3",
				ProjectID:   "4",
				DomainName:  "a",
				ProjectName: "b",
				AuthURL:     "A",
				Endpoint:    "B",
				InSecure:    true,
			},
			expectedConfig: &cloud.Config{
				ID:           "2",
				Password:     "admin",
				DomainID:     "3",
				ProjectID:    "4",
				DomainName:   "a",
				ProjectName:  "b",
				AuthURL:      "A",
				Endpoint:     "B",
				InSecure:     true,
				CloudID:      "1",
				Action:       services.OperationCreate,
				LogLevel:     "debug",
				LogFile:      "/var/log/contrail/cloud.log",
				TemplateRoot: "/usr/share/contrail/templates/",
			},
		},
	}
	for name, generateCloudConfigCase := range generateCloudConfigCases {
		t.Run(
			name,
			func(t *testing.T) {
				cloudConfig := generateCloudConfig(generateCloudConfigCase.inputEvent, generateCloudConfigCase.inputConfig)

				assert.Equal(t, generateCloudConfigCase.expectedConfig, cloudConfig)
			},
		)
	}
}
