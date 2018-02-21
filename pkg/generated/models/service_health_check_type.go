
package models
// ServiceHealthCheckType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceHealthCheckType_delay int = iota
    propServiceHealthCheckType_delayUsecs int = iota
    propServiceHealthCheckType_timeoutUsecs int = iota
    propServiceHealthCheckType_max_retries int = iota
    propServiceHealthCheckType_health_check_type int = iota
    propServiceHealthCheckType_http_method int = iota
    propServiceHealthCheckType_timeout int = iota
    propServiceHealthCheckType_url_path int = iota
    propServiceHealthCheckType_monitor_type int = iota
    propServiceHealthCheckType_enabled int = iota
    propServiceHealthCheckType_expected_codes int = iota
)

// ServiceHealthCheckType 
type ServiceHealthCheckType struct {

    DelayUsecs int `json:"delayUsecs,omitempty"`
    TimeoutUsecs int `json:"timeoutUsecs,omitempty"`
    Delay int `json:"delay,omitempty"`
    Enabled bool `json:"enabled"`
    ExpectedCodes string `json:"expected_codes,omitempty"`
    MaxRetries int `json:"max_retries,omitempty"`
    HealthCheckType HealthCheckType `json:"health_check_type,omitempty"`
    HTTPMethod string `json:"http_method,omitempty"`
    Timeout int `json:"timeout,omitempty"`
    URLPath string `json:"url_path,omitempty"`
    MonitorType HealthCheckProtocolType `json:"monitor_type,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ServiceHealthCheckType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceHealthCheckType makes ServiceHealthCheckType
func MakeServiceHealthCheckType() *ServiceHealthCheckType{
    return &ServiceHealthCheckType{
    //TODO(nati): Apply default
    URLPath: "",
        MonitorType: MakeHealthCheckProtocolType(),
        Enabled: false,
        ExpectedCodes: "",
        MaxRetries: 0,
        HealthCheckType: MakeHealthCheckType(),
        HTTPMethod: "",
        Timeout: 0,
        DelayUsecs: 0,
        TimeoutUsecs: 0,
        Delay: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeServiceHealthCheckTypeSlice makes a slice of ServiceHealthCheckType
func MakeServiceHealthCheckTypeSlice() []*ServiceHealthCheckType {
    return []*ServiceHealthCheckType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceHealthCheckType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ServiceHealthCheckType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceHealthCheckType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ServiceHealthCheckType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceHealthCheckType) GetFQName() []string {
    return model.FQName
}

func (model *ServiceHealthCheckType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceHealthCheckType) GetParentType() string {
    return model.ParentType
}

func (model *ServiceHealthCheckType) GetUuid() string {
    return model.UUID
}

func (model *ServiceHealthCheckType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceHealthCheckType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceHealthCheckType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceHealthCheckType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceHealthCheckType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceHealthCheckType_timeout) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Timeout); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Timeout as timeout")
        }
        msg["timeout"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheckType_url_path) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.URLPath); err != nil {
            return nil, errors.Wrap(err, "Marshal of: URLPath as url_path")
        }
        msg["url_path"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheckType_monitor_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MonitorType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MonitorType as monitor_type")
        }
        msg["monitor_type"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheckType_enabled) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Enabled); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Enabled as enabled")
        }
        msg["enabled"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheckType_expected_codes) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ExpectedCodes); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ExpectedCodes as expected_codes")
        }
        msg["expected_codes"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheckType_max_retries) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MaxRetries); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MaxRetries as max_retries")
        }
        msg["max_retries"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheckType_health_check_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.HealthCheckType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: HealthCheckType as health_check_type")
        }
        msg["health_check_type"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheckType_http_method) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.HTTPMethod); err != nil {
            return nil, errors.Wrap(err, "Marshal of: HTTPMethod as http_method")
        }
        msg["http_method"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheckType_delayUsecs) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DelayUsecs); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DelayUsecs as delayUsecs")
        }
        msg["delayUsecs"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheckType_timeoutUsecs) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.TimeoutUsecs); err != nil {
            return nil, errors.Wrap(err, "Marshal of: TimeoutUsecs as timeoutUsecs")
        }
        msg["timeoutUsecs"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheckType_delay) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Delay); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Delay as delay")
        }
        msg["delay"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceHealthCheckType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceHealthCheckType) UpdateReferences() error {
    return nil
}


