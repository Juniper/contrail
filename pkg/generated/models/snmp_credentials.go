package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeSNMPCredentials makes SNMPCredentials
func MakeSNMPCredentials() *SNMPCredentials {
	return &SNMPCredentials{
		//TODO(nati): Apply default
		V3PrivacyProtocol:        "",
		Retries:                  0,
		V3AuthenticationPassword: "",
		V3EngineTime:             0,
		V3EngineID:               "",
		LocalPort:                0,
		V3SecurityLevel:          "",
		V3Context:                "",
		V3SecurityName:           "",
		V3AuthenticationProtocol: "",
		V2Community:              "",
		V3SecurityEngineID:       "",
		V3ContextEngineID:        "",
		Version:                  0,
		Timeout:                  0,
		V3PrivacyPassword:        "",
		V3EngineBoots:            0,
	}
}

// MakeSNMPCredentials makes SNMPCredentials
func InterfaceToSNMPCredentials(i interface{}) *SNMPCredentials {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &SNMPCredentials{
		//TODO(nati): Apply default
		V3PrivacyProtocol:        schema.InterfaceToString(m["v3_privacy_protocol"]),
		Retries:                  schema.InterfaceToInt64(m["retries"]),
		V3AuthenticationPassword: schema.InterfaceToString(m["v3_authentication_password"]),
		V3EngineTime:             schema.InterfaceToInt64(m["v3_engine_time"]),
		V3EngineID:               schema.InterfaceToString(m["v3_engine_id"]),
		LocalPort:                schema.InterfaceToInt64(m["local_port"]),
		V3SecurityLevel:          schema.InterfaceToString(m["v3_security_level"]),
		V3Context:                schema.InterfaceToString(m["v3_context"]),
		V3SecurityName:           schema.InterfaceToString(m["v3_security_name"]),
		V3AuthenticationProtocol: schema.InterfaceToString(m["v3_authentication_protocol"]),
		V2Community:              schema.InterfaceToString(m["v2_community"]),
		V3SecurityEngineID:       schema.InterfaceToString(m["v3_security_engine_id"]),
		V3ContextEngineID:        schema.InterfaceToString(m["v3_context_engine_id"]),
		Version:                  schema.InterfaceToInt64(m["version"]),
		Timeout:                  schema.InterfaceToInt64(m["timeout"]),
		V3PrivacyPassword:        schema.InterfaceToString(m["v3_privacy_password"]),
		V3EngineBoots:            schema.InterfaceToInt64(m["v3_engine_boots"]),
	}
}

// MakeSNMPCredentialsSlice() makes a slice of SNMPCredentials
func MakeSNMPCredentialsSlice() []*SNMPCredentials {
	return []*SNMPCredentials{}
}

// InterfaceToSNMPCredentialsSlice() makes a slice of SNMPCredentials
func InterfaceToSNMPCredentialsSlice(i interface{}) []*SNMPCredentials {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*SNMPCredentials{}
	for _, item := range list {
		result = append(result, InterfaceToSNMPCredentials(item))
	}
	return result
}
