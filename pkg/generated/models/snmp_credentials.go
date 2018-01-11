package models

// SNMPCredentials

import "encoding/json"

// SNMPCredentials
type SNMPCredentials struct {
	V3SecurityEngineID       string `json:"v3_security_engine_id"`
	V3EngineBoots            int    `json:"v3_engine_boots"`
	V3SecurityLevel          string `json:"v3_security_level"`
	V3SecurityName           string `json:"v3_security_name"`
	V2Community              string `json:"v2_community"`
	LocalPort                int    `json:"local_port"`
	V3AuthenticationProtocol string `json:"v3_authentication_protocol"`
	Version                  int    `json:"version"`
	Retries                  int    `json:"retries"`
	V3AuthenticationPassword string `json:"v3_authentication_password"`
	V3EngineTime             int    `json:"v3_engine_time"`
	V3PrivacyProtocol        string `json:"v3_privacy_protocol"`
	V3Context                string `json:"v3_context"`
	Timeout                  int    `json:"timeout"`
	V3EngineID               string `json:"v3_engine_id"`
	V3ContextEngineID        string `json:"v3_context_engine_id"`
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
		V3PrivacyProtocol:        "",
		V3Context:                "",
		Timeout:                  0,
		V3EngineID:               "",
		V3ContextEngineID:        "",
		V3PrivacyPassword:        "",
		V3SecurityLevel:          "",
		V3SecurityName:           "",
		V2Community:              "",
		V3SecurityEngineID:       "",
		V3EngineBoots:            0,
		Version:                  0,
		Retries:                  0,
		V3AuthenticationPassword: "",
		V3EngineTime:             0,
		LocalPort:                0,
		V3AuthenticationProtocol: "",
	}
}

// MakeSNMPCredentialsSlice() makes a slice of SNMPCredentials
func MakeSNMPCredentialsSlice() []*SNMPCredentials {
	return []*SNMPCredentials{}
}
