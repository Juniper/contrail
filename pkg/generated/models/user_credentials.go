package models

// UserCredentials

import "encoding/json"

// UserCredentials
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//  parents relation object

// String returns json representation of the object
func (model *UserCredentials) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeUserCredentials makes UserCredentials
func MakeUserCredentials() *UserCredentials {
	return &UserCredentials{
		//TODO(nati): Apply default
		Username: "",
		Password: "",
	}
}

// InterfaceToUserCredentials makes UserCredentials from interface
func InterfaceToUserCredentials(iData interface{}) *UserCredentials {
	data := iData.(map[string]interface{})
	return &UserCredentials{
		Username: data["username"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Username","GoType":"string","GoPremitive":true}
		Password: data["password"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Password","GoType":"string","GoPremitive":true}

	}
}

// InterfaceToUserCredentialsSlice makes a slice of UserCredentials from interface
func InterfaceToUserCredentialsSlice(data interface{}) []*UserCredentials {
	list := data.([]interface{})
	result := MakeUserCredentialsSlice()
	for _, item := range list {
		result = append(result, InterfaceToUserCredentials(item))
	}
	return result
}

// MakeUserCredentialsSlice() makes a slice of UserCredentials
func MakeUserCredentialsSlice() []*UserCredentials {
	return []*UserCredentials{}
}
