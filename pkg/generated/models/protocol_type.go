package models
// ProtocolType



import "encoding/json"

// ProtocolType 
//proteus:generate
type ProtocolType struct {

    Protocol string `json:"protocol,omitempty"`
    Port int `json:"port,omitempty"`


}



// String returns json representation of the object
func (model *ProtocolType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeProtocolType makes ProtocolType
func MakeProtocolType() *ProtocolType{
    return &ProtocolType{
    //TODO(nati): Apply default
    Protocol: "",
        Port: 0,
        
    }
}



// MakeProtocolTypeSlice() makes a slice of ProtocolType
func MakeProtocolTypeSlice() []*ProtocolType {
    return []*ProtocolType{}
}
