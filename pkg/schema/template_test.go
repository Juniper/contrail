package schema

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/testutil"
)

const (
	schemaPath         = "test_data/schema"
	templateConfigPath = "test_data/templates/template_config.yaml"
	templatesPath      = "test_data/templates"

	allOutputPath           = "test_output/all.yml"
	ipamSubnetTypeTypePath  = "test_output/ipam_subnet_type_type.yml"
	ipamSubnetsTypePath     = "test_output/ipam_subnets_type.yml"
	networkIPAMResourcePath = "test_output/network_ipam_resource.yml"
	networkIPAMTypePath     = "test_output/network_ipam_type.yml"
	projectResourcePath     = "test_output/project_resource.yml"
	projectTypePath         = "test_output/project_type.yml"
	vnIDTypeTypePath        = "test_output/virtual_network_id_type_type.yml"
	vnResourcePath          = "test_output/virtual_network_resource.yml"
	vnTypePath              = "test_output/virtual_network_type.yml"
	vnSubnetsTypeTypePath   = "test_output/vn_subnets_type_type.yml"
)

func TestApplyTemplatesGeneratesFilesFilledWithData(t *testing.T) {
	err := ApplyTemplates(makeAPI(t), filepath.Dir(templatesPath), loadTemplates(t), &TemplateOption{})

	assert.Nil(t, err)
	assert.Equal(t, []string{"base", "network_ipam", "project", "virtual_network"}, loadFile(t, allOutputPath))
	assert.Equal(t, []string{"subnet_name"}, loadFile(t, ipamSubnetTypeTypePath))
	assert.Equal(t, []string{"subnets"}, loadFile(t, ipamSubnetsTypePath))
	testutil.AssertContainsStrings(t, []string{"ipam_subnets", "uuid", "display_name"}, loadFile(t, networkIPAMResourcePath))
	testutil.AssertContainsStrings(t, []string{"display_name", "ipam_subnets", "uuid"}, loadFile(t, networkIPAMTypePath))
	testutil.AssertContainsStrings(t, []string{"uuid", "display_name"}, loadFile(t, projectResourcePath))
	testutil.AssertContainsStrings(t, []string{"display_name", "uuid"}, loadFile(t, projectTypePath))
	assert.Equal(t, 0, len(loadFile(t, vnIDTypeTypePath)))
	testutil.AssertContainsStrings(t, []string{"uuid", "display_name", "virtual_network_network_id"}, loadFile(t, vnResourcePath))
	testutil.AssertContainsStrings(t, []string{"virtual_network_network_id", "uuid", "display_name"}, loadFile(t, vnTypePath))
	assert.Equal(t, []string{"ipam_subnets"}, loadFile(t, vnSubnetsTypeTypePath))
}

func TestApplyTemplatesPrependsSuffix(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "given YAML file",
		},
		{
			name: "given Go file",
		},
		{
			name: "given Proto file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ApplyTemplates(makeAPI(t), filepath.Dir(templatesPath), loadTemplates(t), &TemplateOption{})

			assert.Nil(t, err)

		})
	}
}

func makeAPI(t *testing.T) *API {
	api, err := MakeAPI([]string{schemaPath})
	assert.Nil(t, err)

	return api
}

func loadTemplates(t *testing.T) []*TemplateConfig {
	c, err := LoadTemplates(templateConfigPath)
	assert.Nil(t, err)

	return c
}

func loadFile(t *testing.T, path string) []string {
	var data []string
	err := common.LoadFile(path, &data)
	assert.Nil(t, err)

	return data
}
