package models

// SubnetType

// SubnetType
//proteus:generate
type SubnetType struct {
	IPPrefix    string `json:"ip_prefix,omitempty"`
	IPPrefixLen int    `json:"ip_prefix_len,omitempty"`
}

// MakeSubnetType makes SubnetType
func MakeSubnetType() *SubnetType {
	return &SubnetType{
		//TODO(nati): Apply default
		IPPrefix:    "",
		IPPrefixLen: 0,
	}
}

// MakeSubnetTypeSlice() makes a slice of SubnetType
func MakeSubnetTypeSlice() []*SubnetType {
	return []*SubnetType{}
}
