package models


// MakePluginProperties makes PluginProperties
func MakePluginProperties() *PluginProperties{
    return &PluginProperties{
    //TODO(nati): Apply default
    
            
                PluginProperty:  MakePluginPropertySlice(),
            
        
    }
}

// MakePluginPropertiesSlice() makes a slice of PluginProperties
func MakePluginPropertiesSlice() []*PluginProperties {
    return []*PluginProperties{}
}


