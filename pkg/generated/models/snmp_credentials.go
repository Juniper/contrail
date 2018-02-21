package models


// MakeSNMPCredentials makes SNMPCredentials
func MakeSNMPCredentials() *SNMPCredentials{
    return &SNMPCredentials{
    //TODO(nati): Apply default
    V3PrivacyProtocol: "",
        Retries: 0,
        V3AuthenticationPassword: "",
        V3EngineTime: 0,
        V3EngineID: "",
        LocalPort: 0,
        V3SecurityLevel: "",
        V3Context: "",
        V3SecurityName: "",
        V3AuthenticationProtocol: "",
        V2Community: "",
        V3SecurityEngineID: "",
        V3ContextEngineID: "",
        Version: 0,
        Timeout: 0,
        V3PrivacyPassword: "",
        V3EngineBoots: 0,
        
    }
}

// MakeSNMPCredentialsSlice() makes a slice of SNMPCredentials
func MakeSNMPCredentialsSlice() []*SNMPCredentials {
    return []*SNMPCredentials{}
}


