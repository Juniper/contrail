package models


// MakePluginProperty makes PluginProperty
func MakePluginProperty() *PluginProperty{
    return &PluginProperty{
    //TODO(nati): Apply default
    Property: "",
        Value: "",
        
    }
}

// MakePluginPropertySlice() makes a slice of PluginProperty
func MakePluginPropertySlice() []*PluginProperty {
    return []*PluginProperty{}
}


