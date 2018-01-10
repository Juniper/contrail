package models

// SNMPCredentials

import "encoding/json"

// SNMPCredentials
type SNMPCredentials struct {
	V3SecurityName           string `json:"v3_security_name"`
	V3AuthenticationProtocol string `json:"v3_authentication_protocol"`
	V2Community              string `json:"v2_community"`
	V3PrivacyProtocol        string `json:"v3_privacy_protocol"`
	Retries                  int    `json:"retries"`
	V3AuthenticationPassword string `json:"v3_authentication_password"`
	V3EngineID               string `json:"v3_engine_id"`
	LocalPort                int    `json:"local_port"`
	V3Context                string `json:"v3_context"`
	V3EngineBoots            int    `json:"v3_engine_boots"`
	V3EngineTime             int    `json:"v3_engine_time"`
	V3SecurityLevel          string `json:"v3_security_level"`
	V3SecurityEngineID       string `json:"v3_security_engine_id"`
	V3ContextEngineID        string `json:"v3_context_engine_id"`
	Version                  int    `json:"version"`
	Timeout                  int    `json:"timeout"`
	V3PrivacyPassword        string `json:"v3_privacy_password"`
}

// String returns json representation of the object
func (model *SNMPCredentials) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeSNMPCredentials makes SNMPCredentials
func MakeSNMPCredentials() *SNMPCredentials {
	return &SNMPCredentials{
		//TODO(nati): Apply default
		V3EngineBoots:            0,
		V3PrivacyProtocol:        "",
		Retries:                  0,
		V3AuthenticationPassword: "",
		V3EngineID:               "",
		LocalPort:                0,
		V3Context:                "",
		V3EngineTime:             0,
		V3SecurityLevel:          "",
		V3SecurityEngineID:       "",
		V3ContextEngineID:        "",
		Version:                  0,
		Timeout:                  0,
		V3PrivacyPassword:        "",
		V3SecurityName:           "",
		V3AuthenticationProtocol: "",
		V2Community:              "",
	}
}

// InterfaceToSNMPCredentials makes SNMPCredentials from interface
func InterfaceToSNMPCredentials(iData interface{}) *SNMPCredentials {
	data := iData.(map[string]interface{})
	return &SNMPCredentials{
		V3EngineTime: data["v3_engine_time"].(int),

		//{"type":"integer"}
		V3SecurityLevel: data["v3_security_level"].(string),

		//{"type":"string"}
		V3SecurityEngineID: data["v3_security_engine_id"].(string),

		//{"type":"string"}
		V3ContextEngineID: data["v3_context_engine_id"].(string),

		//{"type":"string"}
		Version: data["version"].(int),

		//{"type":"integer"}
		Timeout: data["timeout"].(int),

		//{"type":"integer"}
		V3PrivacyPassword: data["v3_privacy_password"].(string),

		//{"type":"string"}
		V3SecurityName: data["v3_security_name"].(string),

		//{"type":"string"}
		V3AuthenticationProtocol: data["v3_authentication_protocol"].(string),

		//{"type":"string"}
		V2Community: data["v2_community"].(string),

		//{"type":"string"}
		V3EngineBoots: data["v3_engine_boots"].(int),

		//{"type":"integer"}
		V3PrivacyProtocol: data["v3_privacy_protocol"].(string),

		//{"type":"string"}
		Retries: data["retries"].(int),

		//{"type":"integer"}
		V3AuthenticationPassword: data["v3_authentication_password"].(string),

		//{"type":"string"}
		V3EngineID: data["v3_engine_id"].(string),

		//{"type":"string"}
		LocalPort: data["local_port"].(int),

		//{"type":"integer"}
		V3Context: data["v3_context"].(string),

		//{"type":"string"}

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
