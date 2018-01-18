package models

// SNMPCredentials

import "encoding/json"

// SNMPCredentials
type SNMPCredentials struct {
	V3ContextEngineID        string `json:"v3_context_engine_id,omitempty"`
	Timeout                  int    `json:"timeout,omitempty"`
	V3SecurityLevel          string `json:"v3_security_level,omitempty"`
	V3SecurityName           string `json:"v3_security_name,omitempty"`
	V2Community              string `json:"v2_community,omitempty"`
	V3SecurityEngineID       string `json:"v3_security_engine_id,omitempty"`
	V3EngineBoots            int    `json:"v3_engine_boots,omitempty"`
	V3PrivacyProtocol        string `json:"v3_privacy_protocol,omitempty"`
	V3EngineID               string `json:"v3_engine_id,omitempty"`
	LocalPort                int    `json:"local_port,omitempty"`
	V3AuthenticationProtocol string `json:"v3_authentication_protocol,omitempty"`
	Version                  int    `json:"version,omitempty"`
	V3PrivacyPassword        string `json:"v3_privacy_password,omitempty"`
	Retries                  int    `json:"retries,omitempty"`
	V3AuthenticationPassword string `json:"v3_authentication_password,omitempty"`
	V3EngineTime             int    `json:"v3_engine_time,omitempty"`
	V3Context                string `json:"v3_context,omitempty"`
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
		V3SecurityEngineID:       "",
		V3ContextEngineID:        "",
		Timeout:                  0,
		V3SecurityLevel:          "",
		V3SecurityName:           "",
		V2Community:              "",
		V3AuthenticationProtocol: "",
		V3EngineBoots:            0,
		V3PrivacyProtocol:        "",
		V3EngineID:               "",
		LocalPort:                0,
		Version:                  0,
		V3PrivacyPassword:        "",
		V3Context:                "",
		Retries:                  0,
		V3AuthenticationPassword: "",
		V3EngineTime:             0,
	}
}

// MakeSNMPCredentialsSlice() makes a slice of SNMPCredentials
func MakeSNMPCredentialsSlice() []*SNMPCredentials {
	return []*SNMPCredentials{}
}
