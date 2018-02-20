package models

// PluginProperties

// PluginProperties
//proteus:generate
type PluginProperties struct {
	PluginProperty []*PluginProperty `json:"plugin_property,omitempty"`
}

// MakePluginProperties makes PluginProperties
func MakePluginProperties() *PluginProperties {
	return &PluginProperties{
		//TODO(nati): Apply default

		PluginProperty: MakePluginPropertySlice(),
	}
}

// MakePluginPropertiesSlice() makes a slice of PluginProperties
func MakePluginPropertiesSlice() []*PluginProperties {
	return []*PluginProperties{}
}
