
package models
// SNMPCredentials



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propSNMPCredentials_v3_engine_id int = iota
    propSNMPCredentials_v3_engine_boots int = iota
    propSNMPCredentials_v3_privacy_protocol int = iota
    propSNMPCredentials_v3_engine_time int = iota
    propSNMPCredentials_v3_authentication_protocol int = iota
    propSNMPCredentials_v2_community int = iota
    propSNMPCredentials_local_port int = iota
    propSNMPCredentials_v3_security_level int = iota
    propSNMPCredentials_v3_security_name int = iota
    propSNMPCredentials_v3_context_engine_id int = iota
    propSNMPCredentials_timeout int = iota
    propSNMPCredentials_retries int = iota
    propSNMPCredentials_v3_authentication_password int = iota
    propSNMPCredentials_version int = iota
    propSNMPCredentials_v3_privacy_password int = iota
    propSNMPCredentials_v3_context int = iota
    propSNMPCredentials_v3_security_engine_id int = iota
)

// SNMPCredentials 
type SNMPCredentials struct {

    V3PrivacyProtocol string `json:"v3_privacy_protocol,omitempty"`
    V3EngineTime int `json:"v3_engine_time,omitempty"`
    V3EngineID string `json:"v3_engine_id,omitempty"`
    V3EngineBoots int `json:"v3_engine_boots,omitempty"`
    V3AuthenticationProtocol string `json:"v3_authentication_protocol,omitempty"`
    V2Community string `json:"v2_community,omitempty"`
    V3SecurityName string `json:"v3_security_name,omitempty"`
    V3ContextEngineID string `json:"v3_context_engine_id,omitempty"`
    Timeout int `json:"timeout,omitempty"`
    Retries int `json:"retries,omitempty"`
    V3AuthenticationPassword string `json:"v3_authentication_password,omitempty"`
    LocalPort int `json:"local_port,omitempty"`
    V3SecurityLevel string `json:"v3_security_level,omitempty"`
    V3Context string `json:"v3_context,omitempty"`
    V3SecurityEngineID string `json:"v3_security_engine_id,omitempty"`
    Version int `json:"version,omitempty"`
    V3PrivacyPassword string `json:"v3_privacy_password,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *SNMPCredentials) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeSNMPCredentials makes SNMPCredentials
func MakeSNMPCredentials() *SNMPCredentials{
    return &SNMPCredentials{
    //TODO(nati): Apply default
    Retries: 0,
        V3AuthenticationPassword: "",
        LocalPort: 0,
        V3SecurityLevel: "",
        V3SecurityName: "",
        V3ContextEngineID: "",
        Timeout: 0,
        V3Context: "",
        V3SecurityEngineID: "",
        Version: 0,
        V3PrivacyPassword: "",
        V3PrivacyProtocol: "",
        V3EngineTime: 0,
        V3EngineID: "",
        V3EngineBoots: 0,
        V3AuthenticationProtocol: "",
        V2Community: "",
        
        modified: big.NewInt(0),
    }
}



// MakeSNMPCredentialsSlice makes a slice of SNMPCredentials
func MakeSNMPCredentialsSlice() []*SNMPCredentials {
    return []*SNMPCredentials{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *SNMPCredentials) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *SNMPCredentials) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *SNMPCredentials) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *SNMPCredentials) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *SNMPCredentials) GetFQName() []string {
    return model.FQName
}

func (model *SNMPCredentials) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *SNMPCredentials) GetParentType() string {
    return model.ParentType
}

func (model *SNMPCredentials) GetUuid() string {
    return model.UUID
}

func (model *SNMPCredentials) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *SNMPCredentials) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *SNMPCredentials) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *SNMPCredentials) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *SNMPCredentials) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propSNMPCredentials_v3_privacy_protocol) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.V3PrivacyProtocol); err != nil {
            return nil, errors.Wrap(err, "Marshal of: V3PrivacyProtocol as v3_privacy_protocol")
        }
        msg["v3_privacy_protocol"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_v3_engine_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.V3EngineTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: V3EngineTime as v3_engine_time")
        }
        msg["v3_engine_time"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_v3_engine_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.V3EngineID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: V3EngineID as v3_engine_id")
        }
        msg["v3_engine_id"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_v3_engine_boots) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.V3EngineBoots); err != nil {
            return nil, errors.Wrap(err, "Marshal of: V3EngineBoots as v3_engine_boots")
        }
        msg["v3_engine_boots"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_v3_authentication_protocol) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.V3AuthenticationProtocol); err != nil {
            return nil, errors.Wrap(err, "Marshal of: V3AuthenticationProtocol as v3_authentication_protocol")
        }
        msg["v3_authentication_protocol"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_v2_community) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.V2Community); err != nil {
            return nil, errors.Wrap(err, "Marshal of: V2Community as v2_community")
        }
        msg["v2_community"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_retries) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Retries); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Retries as retries")
        }
        msg["retries"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_v3_authentication_password) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.V3AuthenticationPassword); err != nil {
            return nil, errors.Wrap(err, "Marshal of: V3AuthenticationPassword as v3_authentication_password")
        }
        msg["v3_authentication_password"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_local_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LocalPort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LocalPort as local_port")
        }
        msg["local_port"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_v3_security_level) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.V3SecurityLevel); err != nil {
            return nil, errors.Wrap(err, "Marshal of: V3SecurityLevel as v3_security_level")
        }
        msg["v3_security_level"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_v3_security_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.V3SecurityName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: V3SecurityName as v3_security_name")
        }
        msg["v3_security_name"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_v3_context_engine_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.V3ContextEngineID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: V3ContextEngineID as v3_context_engine_id")
        }
        msg["v3_context_engine_id"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_timeout) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Timeout); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Timeout as timeout")
        }
        msg["timeout"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_v3_context) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.V3Context); err != nil {
            return nil, errors.Wrap(err, "Marshal of: V3Context as v3_context")
        }
        msg["v3_context"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_v3_security_engine_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.V3SecurityEngineID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: V3SecurityEngineID as v3_security_engine_id")
        }
        msg["v3_security_engine_id"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_version) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Version); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Version as version")
        }
        msg["version"] = &val
    }
    
    if model.modified.Bit(propSNMPCredentials_v3_privacy_password) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.V3PrivacyPassword); err != nil {
            return nil, errors.Wrap(err, "Marshal of: V3PrivacyPassword as v3_privacy_password")
        }
        msg["v3_privacy_password"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *SNMPCredentials) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *SNMPCredentials) UpdateReferences() error {
    return nil
}


