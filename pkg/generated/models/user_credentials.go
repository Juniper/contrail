package models

// UserCredentials

import "encoding/json"

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (model *UserCredentials) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeUserCredentials() *UserCredentials {
	return &UserCredentials{
		//TODO(nati): Apply default
		Password: "",
		Username: "",
	}
}

func InterfaceToUserCredentials(iData interface{}) *UserCredentials {
	data := iData.(map[string]interface{})
	return &UserCredentials{
		Username: data["username"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Username","GoType":"string"}
		Password: data["password"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Password","GoType":"string"}

	}
}

func InterfaceToUserCredentialsSlice(data interface{}) []*UserCredentials {
	list := data.([]interface{})
	result := MakeUserCredentialsSlice()
	for _, item := range list {
		result = append(result, InterfaceToUserCredentials(item))
	}
	return result
}

func MakeUserCredentialsSlice() []*UserCredentials {
	return []*UserCredentials{}
}
