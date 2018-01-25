package models
// MACLimitControlType



import "encoding/json"

// MACLimitControlType 
//proteus:generate
type MACLimitControlType struct {

    MacLimit int `json:"mac_limit,omitempty"`
    MacLimitAction MACLimitExceedActionType `json:"mac_limit_action,omitempty"`


}



// String returns json representation of the object
func (model *MACLimitControlType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeMACLimitControlType makes MACLimitControlType
func MakeMACLimitControlType() *MACLimitControlType{
    return &MACLimitControlType{
    //TODO(nati): Apply default
    MacLimit: 0,
        MacLimitAction: MakeMACLimitExceedActionType(),
        
    }
}



// MakeMACLimitControlTypeSlice() makes a slice of MACLimitControlType
func MakeMACLimitControlTypeSlice() []*MACLimitControlType {
    return []*MACLimitControlType{}
}
