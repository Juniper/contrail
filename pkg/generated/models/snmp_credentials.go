package models

// SNMPCredentials

import "encoding/json"

// SNMPCredentials
type SNMPCredentials struct {
	LocalPort                int    `json:"local_port,omitempty"`
	V3Context                string `json:"v3_context,omitempty"`
	V3ContextEngineID        string `json:"v3_context_engine_id,omitempty"`
	Version                  int    `json:"version,omitempty"`
	Timeout                  int    `json:"timeout,omitempty"`
	V3AuthenticationPassword string `json:"v3_authentication_password,omitempty"`
	V2Community              string `json:"v2_community,omitempty"`
	V3PrivacyPassword        string `json:"v3_privacy_password,omitempty"`
	V3SecurityLevel          string `json:"v3_security_level,omitempty"`
	V3AuthenticationProtocol string `json:"v3_authentication_protocol,omitempty"`
	V3PrivacyProtocol        string `json:"v3_privacy_protocol,omitempty"`
	Retries                  int    `json:"retries,omitempty"`
	V3EngineTime             int    `json:"v3_engine_time,omitempty"`
	V3EngineID               string `json:"v3_engine_id,omitempty"`
	V3SecurityName           string `json:"v3_security_name,omitempty"`
	V3SecurityEngineID       string `json:"v3_security_engine_id,omitempty"`
	V3EngineBoots            int    `json:"v3_engine_boots,omitempty"`
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
		V3AuthenticationPassword: "",
		V2Community:              "",
		V3PrivacyPassword:        "",
		V3SecurityLevel:          "",
		V3AuthenticationProtocol: "",
		V3PrivacyProtocol:        "",
		Retries:                  0,
		V3EngineTime:             0,
		V3EngineID:               "",
		V3SecurityName:           "",
		V3SecurityEngineID:       "",
		V3EngineBoots:            0,
		LocalPort:                0,
		V3Context:                "",
		V3ContextEngineID:        "",
		Version:                  0,
		Timeout:                  0,
	}
}

// MakeSNMPCredentialsSlice() makes a slice of SNMPCredentials
func MakeSNMPCredentialsSlice() []*SNMPCredentials {
	return []*SNMPCredentials{}
}
