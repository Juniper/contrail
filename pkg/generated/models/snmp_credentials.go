package models

// SNMPCredentials

import "encoding/json"

// SNMPCredentials
type SNMPCredentials struct {
	V3Context                string `json:"v3_context"`
	Version                  int    `json:"version"`
	V3SecurityEngineID       string `json:"v3_security_engine_id"`
	V3PrivacyPassword        string `json:"v3_privacy_password"`
	V3PrivacyProtocol        string `json:"v3_privacy_protocol"`
	V3AuthenticationProtocol string `json:"v3_authentication_protocol"`
	V3SecurityLevel          string `json:"v3_security_level"`
	V3ContextEngineID        string `json:"v3_context_engine_id"`
	V3EngineBoots            int    `json:"v3_engine_boots"`
	V3AuthenticationPassword string `json:"v3_authentication_password"`
	LocalPort                int    `json:"local_port"`
	V3EngineID               string `json:"v3_engine_id"`
	V3SecurityName           string `json:"v3_security_name"`
	V2Community              string `json:"v2_community"`
	Timeout                  int    `json:"timeout"`
	Retries                  int    `json:"retries"`
	V3EngineTime             int    `json:"v3_engine_time"`
}

//  parents relation object

// String returns json representation of the object
func (model *SNMPCredentials) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeSNMPCredentials makes SNMPCredentials
func MakeSNMPCredentials() *SNMPCredentials {
	return &SNMPCredentials{
		//TODO(nati): Apply default
		V3SecurityName:           "",
		V2Community:              "",
		Timeout:                  0,
		Retries:                  0,
		V3EngineTime:             0,
		V3EngineID:               "",
		V3Context:                "",
		Version:                  0,
		V3PrivacyPassword:        "",
		V3PrivacyProtocol:        "",
		V3AuthenticationProtocol: "",
		V3SecurityEngineID:       "",
		V3ContextEngineID:        "",
		V3EngineBoots:            0,
		V3AuthenticationPassword: "",
		LocalPort:                0,
		V3SecurityLevel:          "",
	}
}

// InterfaceToSNMPCredentials makes SNMPCredentials from interface
func InterfaceToSNMPCredentials(iData interface{}) *SNMPCredentials {
	data := iData.(map[string]interface{})
	return &SNMPCredentials{
		V3EngineBoots: data["v3_engine_boots"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3EngineBoots","GoType":"int","GoPremitive":true}
		V3AuthenticationPassword: data["v3_authentication_password"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3AuthenticationPassword","GoType":"string","GoPremitive":true}
		LocalPort: data["local_port"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LocalPort","GoType":"int","GoPremitive":true}
		V3SecurityLevel: data["v3_security_level"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3SecurityLevel","GoType":"string","GoPremitive":true}
		V3ContextEngineID: data["v3_context_engine_id"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3ContextEngineID","GoType":"string","GoPremitive":true}
		V2Community: data["v2_community"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V2Community","GoType":"string","GoPremitive":true}
		Timeout: data["timeout"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Timeout","GoType":"int","GoPremitive":true}
		Retries: data["retries"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Retries","GoType":"int","GoPremitive":true}
		V3EngineTime: data["v3_engine_time"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3EngineTime","GoType":"int","GoPremitive":true}
		V3EngineID: data["v3_engine_id"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3EngineID","GoType":"string","GoPremitive":true}
		V3SecurityName: data["v3_security_name"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3SecurityName","GoType":"string","GoPremitive":true}
		V3Context: data["v3_context"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3Context","GoType":"string","GoPremitive":true}
		Version: data["version"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Version","GoType":"int","GoPremitive":true}
		V3PrivacyProtocol: data["v3_privacy_protocol"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3PrivacyProtocol","GoType":"string","GoPremitive":true}
		V3AuthenticationProtocol: data["v3_authentication_protocol"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3AuthenticationProtocol","GoType":"string","GoPremitive":true}
		V3SecurityEngineID: data["v3_security_engine_id"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3SecurityEngineID","GoType":"string","GoPremitive":true}
		V3PrivacyPassword: data["v3_privacy_password"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"V3PrivacyPassword","GoType":"string","GoPremitive":true}

	}
}

// InterfaceToSNMPCredentialsSlice makes a slice of SNMPCredentials from interface
func InterfaceToSNMPCredentialsSlice(data interface{}) []*SNMPCredentials {
	list := data.([]interface{})
	result := MakeSNMPCredentialsSlice()
	for _, item := range list {
		result = append(result, InterfaceToSNMPCredentials(item))
	}
	return result
}

// MakeSNMPCredentialsSlice() makes a slice of SNMPCredentials
func MakeSNMPCredentialsSlice() []*SNMPCredentials {
	return []*SNMPCredentials{}
}
