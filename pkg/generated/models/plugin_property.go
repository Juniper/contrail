package models
// PluginProperty



import "encoding/json"

// PluginProperty 
//proteus:generate
type PluginProperty struct {

    Property string `json:"property,omitempty"`
    Value string `json:"value,omitempty"`


}



// String returns json representation of the object
func (model *PluginProperty) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

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
