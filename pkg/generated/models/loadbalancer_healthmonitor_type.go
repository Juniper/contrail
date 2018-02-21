
package models
// LoadbalancerHealthmonitorType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLoadbalancerHealthmonitorType_max_retries int = iota
    propLoadbalancerHealthmonitorType_http_method int = iota
    propLoadbalancerHealthmonitorType_admin_state int = iota
    propLoadbalancerHealthmonitorType_timeout int = iota
    propLoadbalancerHealthmonitorType_url_path int = iota
    propLoadbalancerHealthmonitorType_monitor_type int = iota
    propLoadbalancerHealthmonitorType_delay int = iota
    propLoadbalancerHealthmonitorType_expected_codes int = iota
)

// LoadbalancerHealthmonitorType 
type LoadbalancerHealthmonitorType struct {

    URLPath string `json:"url_path,omitempty"`
    MonitorType HealthmonitorType `json:"monitor_type,omitempty"`
    Delay int `json:"delay,omitempty"`
    ExpectedCodes string `json:"expected_codes,omitempty"`
    MaxRetries int `json:"max_retries,omitempty"`
    HTTPMethod string `json:"http_method,omitempty"`
    AdminState bool `json:"admin_state"`
    Timeout int `json:"timeout,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *LoadbalancerHealthmonitorType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLoadbalancerHealthmonitorType makes LoadbalancerHealthmonitorType
func MakeLoadbalancerHealthmonitorType() *LoadbalancerHealthmonitorType{
    return &LoadbalancerHealthmonitorType{
    //TODO(nati): Apply default
    ExpectedCodes: "",
        MaxRetries: 0,
        HTTPMethod: "",
        AdminState: false,
        Timeout: 0,
        URLPath: "",
        MonitorType: MakeHealthmonitorType(),
        Delay: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeLoadbalancerHealthmonitorTypeSlice makes a slice of LoadbalancerHealthmonitorType
func MakeLoadbalancerHealthmonitorTypeSlice() []*LoadbalancerHealthmonitorType {
    return []*LoadbalancerHealthmonitorType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LoadbalancerHealthmonitorType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *LoadbalancerHealthmonitorType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *LoadbalancerHealthmonitorType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *LoadbalancerHealthmonitorType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *LoadbalancerHealthmonitorType) GetFQName() []string {
    return model.FQName
}

func (model *LoadbalancerHealthmonitorType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LoadbalancerHealthmonitorType) GetParentType() string {
    return model.ParentType
}

func (model *LoadbalancerHealthmonitorType) GetUuid() string {
    return model.UUID
}

func (model *LoadbalancerHealthmonitorType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LoadbalancerHealthmonitorType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LoadbalancerHealthmonitorType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LoadbalancerHealthmonitorType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LoadbalancerHealthmonitorType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLoadbalancerHealthmonitorType_http_method) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.HTTPMethod); err != nil {
            return nil, errors.Wrap(err, "Marshal of: HTTPMethod as http_method")
        }
        msg["http_method"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitorType_admin_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AdminState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AdminState as admin_state")
        }
        msg["admin_state"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitorType_timeout) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Timeout); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Timeout as timeout")
        }
        msg["timeout"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitorType_url_path) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.URLPath); err != nil {
            return nil, errors.Wrap(err, "Marshal of: URLPath as url_path")
        }
        msg["url_path"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitorType_monitor_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MonitorType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MonitorType as monitor_type")
        }
        msg["monitor_type"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitorType_delay) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Delay); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Delay as delay")
        }
        msg["delay"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitorType_expected_codes) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ExpectedCodes); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ExpectedCodes as expected_codes")
        }
        msg["expected_codes"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitorType_max_retries) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MaxRetries); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MaxRetries as max_retries")
        }
        msg["max_retries"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *LoadbalancerHealthmonitorType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LoadbalancerHealthmonitorType) UpdateReferences() error {
    return nil
}


