package models

// UveKeysType

// UveKeysType
//proteus:generate
type UveKeysType struct {
	UveKey []string `json:"uve_key,omitempty"`
}

// MakeUveKeysType makes UveKeysType
func MakeUveKeysType() *UveKeysType {
	return &UveKeysType{
		//TODO(nati): Apply default
		UveKey: []string{},
	}
}

// MakeUveKeysTypeSlice() makes a slice of UveKeysType
func MakeUveKeysTypeSlice() []*UveKeysType {
	return []*UveKeysType{}
}
