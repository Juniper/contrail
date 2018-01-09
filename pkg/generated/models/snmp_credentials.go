package models

// SNMPCredentials

import "encoding/json"

// SNMPCredentials
type SNMPCredentials struct {
	V3SecurityLevel          string `json:"v3_security_level"`
	V3PrivacyProtocol        string `json:"v3_privacy_protocol"`
	Retries                  int    `json:"retries"`
	LocalPort                int    `json:"local_port"`
	V3AuthenticationProtocol string `json:"v3_authentication_protocol"`
	V3ContextEngineID        string `json:"v3_context_engine_id"`
	V3EngineBoots            int    `json:"v3_engine_boots"`
	V3AuthenticationPassword string `json:"v3_authentication_password"`
	V3SecurityName           string `json:"v3_security_name"`
	Version                  int    `json:"version"`
	V3PrivacyPassword        string `json:"v3_privacy_password"`
	V3EngineTime             int    `json:"v3_engine_time"`
	V3EngineID               string `json:"v3_engine_id"`
	V3Context                string `json:"v3_context"`
	V2Community              string `json:"v2_community"`
	V3SecurityEngineID       string `json:"v3_security_engine_id"`
	Timeout                  int    `json:"timeout"`
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
		LocalPort:                0,
		V3AuthenticationProtocol: "",
		V3ContextEngineID:        "",
		V3AuthenticationPassword: "",
		V3SecurityName:           "",
		Version:                  0,
		Timeout:                  0,
		V3PrivacyPassword:        "",
		V3EngineTime:             0,
		V3EngineID:               "",
		V3Context:                "",
		V2Community:              "",
		V3SecurityEngineID:       "",
		V3SecurityLevel:          "",
	}
}

// InterfaceToSNMPCredentials makes SNMPCredentials from interface
func InterfaceToSNMPCredentials(iData interface{}) *SNMPCredentials {
	data := iData.(map[string]interface{})
	return &SNMPCredentials{
		V3EngineTime: data["v3_engine_time"].(int),

		//{"type":"integer"}
		V3EngineID: data["v3_engine_id"].(string),

		//{"type":"string"}
		V3Context: data["v3_context"].(string),

		//{"type":"string"}
		V2Community: data["v2_community"].(string),

		//{"type":"string"}
		V3SecurityEngineID: data["v3_security_engine_id"].(string),

		//{"type":"string"}
		Timeout: data["timeout"].(int),

		//{"type":"integer"}
		V3PrivacyPassword: data["v3_privacy_password"].(string),

		//{"type":"string"}
		V3SecurityLevel: data["v3_security_level"].(string),

		//{"type":"string"}
		V3PrivacyProtocol: data["v3_privacy_protocol"].(string),

		//{"type":"string"}
		Retries: data["retries"].(int),

		//{"type":"integer"}
		LocalPort: data["local_port"].(int),

		//{"type":"integer"}
		V3AuthenticationProtocol: data["v3_authentication_protocol"].(string),

		//{"type":"string"}
		V3ContextEngineID: data["v3_context_engine_id"].(string),

		//{"type":"string"}
		V3EngineBoots: data["v3_engine_boots"].(int),

		//{"type":"integer"}
		V3AuthenticationPassword: data["v3_authentication_password"].(string),

		//{"type":"string"}
		V3SecurityName: data["v3_security_name"].(string),

		//{"type":"string"}
		Version: data["version"].(int),

		//{"type":"integer"}

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
