package models

// PluginProperty

// PluginProperty
//proteus:generate
type PluginProperty struct {
	Property string `json:"property,omitempty"`
	Value    string `json:"value,omitempty"`
}

// MakePluginProperty makes PluginProperty
func MakePluginProperty() *PluginProperty {
	return &PluginProperty{
		//TODO(nati): Apply default
		Property: "",
		Value:    "",
	}
}

// MakePluginPropertySlice() makes a slice of PluginProperty
func MakePluginPropertySlice() []*PluginProperty {
	return []*PluginProperty{}
}
