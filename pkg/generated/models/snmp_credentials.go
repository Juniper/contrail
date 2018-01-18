package models

// SNMPCredentials

import "encoding/json"

// SNMPCredentials
type SNMPCredentials struct {
	V3ContextEngineID        string `json:"v3_context_engine_id,omitempty"`
	V3AuthenticationPassword string `json:"v3_authentication_password,omitempty"`
	V3EngineTime             int    `json:"v3_engine_time,omitempty"`
	V3SecurityName           string `json:"v3_security_name,omitempty"`
	V3SecurityEngineID       string `json:"v3_security_engine_id,omitempty"`
	Version                  int    `json:"version,omitempty"`
	Timeout                  int    `json:"timeout,omitempty"`
	V3PrivacyPassword        string `json:"v3_privacy_password,omitempty"`
	Retries                  int    `json:"retries,omitempty"`
	V3Context                string `json:"v3_context,omitempty"`
	V3AuthenticationProtocol string `json:"v3_authentication_protocol,omitempty"`
	V3EngineBoots            int    `json:"v3_engine_boots,omitempty"`
	V3PrivacyProtocol        string `json:"v3_privacy_protocol,omitempty"`
	V3EngineID               string `json:"v3_engine_id,omitempty"`
	V3SecurityLevel          string `json:"v3_security_level,omitempty"`
	LocalPort                int    `json:"local_port,omitempty"`
	V2Community              string `json:"v2_community,omitempty"`
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
		V3EngineTime:             0,
		V3SecurityName:           "",
		V3ContextEngineID:        "",
		Timeout:                  0,
		V3PrivacyPassword:        "",
		Retries:                  0,
		V3Context:                "",
		V3AuthenticationProtocol: "",
		V3SecurityEngineID:       "",
		Version:                  0,
		V3PrivacyProtocol:        "",
		V3EngineID:               "",
		V3SecurityLevel:          "",
		V3EngineBoots:            0,
		LocalPort:                0,
		V2Community:              "",
	}
}

// MakeSNMPCredentialsSlice() makes a slice of SNMPCredentials
func MakeSNMPCredentialsSlice() []*SNMPCredentials {
	return []*SNMPCredentials{}
}
