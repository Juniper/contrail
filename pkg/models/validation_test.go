package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceInterfaceTypeValidation(t *testing.T) {
	tests := []struct {
		name                string
		serviceInterfaceTag *ServiceInterfaceTag
		fails               bool
	}{
		{
			name: "Bad standard value - validation fails",
			serviceInterfaceTag: &ServiceInterfaceTag{
				InterfaceType: "hogehoge",
			},
			fails: true,
		},
		{
			name: "Bad regex value - validation fails",
			serviceInterfaceTag: &ServiceInterfaceTag{
				InterfaceType: "othe",
			},
			fails: true,
		},
		{
			name: "Standard value",
			serviceInterfaceTag: &ServiceInterfaceTag{
				InterfaceType: "management",
			},
			fails: false,
		},
		{
			name: "Regex value",
			serviceInterfaceTag: &ServiceInterfaceTag{
				InterfaceType: "other0",
			},
			fails: false,
		},
	}

	tv, err := NewTypeValidatorWithFormat()

	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tv.ValidateServiceInterfaceTag(tt.serviceInterfaceTag)
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSchemaValidationJobTemplate(t *testing.T) {
	tests := []struct {
		name        string
		jobTemplate *JobTemplate
		fails       bool
	}{
		{
			name:        "Empty struct - validation fails",
			jobTemplate: &JobTemplate{},
			fails:       true,
		},
		{
			name: "Validation pass",
			jobTemplate: &JobTemplate{
				UUID:                      "beef",
				ParentUUID:                "beef-beef",
				ParentType:                "global-system-config",
				JobTemplatePlaybooks:      &PlaybookInfoListType{},
				JobTemplateMultiDeviceJob: true,
				FQName: []string{"a", "b"},
			},
			fails: false,
		},
		{
			name: "Missing required string property",
			jobTemplate: &JobTemplate{
				UUID:                      "",
				ParentUUID:                "beef-beef",
				ParentType:                "global-system-config",
				JobTemplatePlaybooks:      &PlaybookInfoListType{},
				JobTemplateMultiDeviceJob: true,
				FQName: []string{"a", "b"},
			},
			fails: true,
		},
		// {
		// 	name: "Missing required integer property",
		// 	jobTemplate: &JobTemplate{
		// 		UUID:                      "beef",
		// 		ParentUUID:                "beef-beef",
		// 		ParentType:                "global-system-config",
		// 		JobTemplatePlaybooks:      &PlaybookInfoListType{},
		// 		JobTemplateMultiDeviceJob: true,
		// 		FQName: []string{"a", "b"},
		// 	},
		// 	fails: true,
		// },
		// {
		// 	name: "Missing required boolean property",
		// 	jobTemplate: &JobTemplate{
		// 		UUID:                      "beef",
		// 		ParentUUID:                "beef-beef",
		// 		ParentType:                "global-system-config",
		// 		JobTemplatePlaybooks:      &PlaybookInfoListType{},
		// 		JobTemplateMultiDeviceJob: false,
		// 		FQName: []string{"a", "b"},
		// 	},
		// 	fails: true,
		// },
		{
			name: "Missing required object property",
			jobTemplate: &JobTemplate{
				UUID:                      "beef",
				ParentUUID:                "beef-beef",
				ParentType:                "global-system-config",
				JobTemplatePlaybooks:      nil,
				JobTemplateMultiDeviceJob: true,
				FQName: []string{"a", "b"},
			},
			fails: true,
		},
		{
			name: "Bad parent type",
			jobTemplate: &JobTemplate{
				UUID:                      "beef",
				ParentUUID:                "beef-beef",
				ParentType:                "hogehoge",
				JobTemplatePlaybooks:      &PlaybookInfoListType{},
				JobTemplateMultiDeviceJob: true,
				FQName: []string{"a", "b"},
			},
			fails: true,
		},
	}

	tv, err := NewTypeValidatorWithFormat()

	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tv.ValidateJobTemplate(tt.jobTemplate)
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSchemaValidationBgpFamilyAttributes(t *testing.T) {
	tests := []struct {
		name                string
		bgpFamilyAttributes *BgpFamilyAttributes
		fails               bool
	}{
		{
			name:                "Empty struct - validation fails",
			bgpFamilyAttributes: &BgpFamilyAttributes{},
			fails:               true,
		},
		{
			name: "Validation pass",
			bgpFamilyAttributes: &BgpFamilyAttributes{
				AddressFamily: "inet",
				LoopCount:     0,
			},
			fails: false,
		},
		{
			name: "Bad string value",
			bgpFamilyAttributes: &BgpFamilyAttributes{
				AddressFamily: "hogehoge",
			},
			fails: true,
		},
		{
			name: "Number value too small",
			bgpFamilyAttributes: &BgpFamilyAttributes{
				AddressFamily: "inet",
				LoopCount:     -1,
			},
			fails: true,
		},
		{
			name: "Number value too big",
			bgpFamilyAttributes: &BgpFamilyAttributes{
				AddressFamily: "inet",
				LoopCount:     17,
			},
			fails: true,
		},
	}

	tv, err := NewTypeValidatorWithFormat()

	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tv.ValidateBgpFamilyAttributes(tt.bgpFamilyAttributes)
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
