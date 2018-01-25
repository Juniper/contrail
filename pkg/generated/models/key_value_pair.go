package models
// KeyValuePair



import "encoding/json"

// KeyValuePair 
//proteus:generate
type KeyValuePair struct {

    Value string `json:"value,omitempty"`
    Key string `json:"key,omitempty"`


}



// String returns json representation of the object
func (model *KeyValuePair) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeKeyValuePair makes KeyValuePair
func MakeKeyValuePair() *KeyValuePair{
    return &KeyValuePair{
    //TODO(nati): Apply default
    Value: "",
        Key: "",
        
    }
}



// MakeKeyValuePairSlice() makes a slice of KeyValuePair
func MakeKeyValuePairSlice() []*KeyValuePair {
    return []*KeyValuePair{}
}
