package models
// MacAddressesType



import "encoding/json"

// MacAddressesType 
//proteus:generate
type MacAddressesType struct {

    MacAddress []string `json:"mac_address,omitempty"`


}



// String returns json representation of the object
func (model *MacAddressesType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeMacAddressesType makes MacAddressesType
func MakeMacAddressesType() *MacAddressesType{
    return &MacAddressesType{
    //TODO(nati): Apply default
    MacAddress: []string{},
        
    }
}



// MakeMacAddressesTypeSlice() makes a slice of MacAddressesType
func MakeMacAddressesTypeSlice() []*MacAddressesType {
    return []*MacAddressesType{}
}
