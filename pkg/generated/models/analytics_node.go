
package models
// AnalyticsNode



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAnalyticsNode_display_name int = iota
    propAnalyticsNode_annotations int = iota
    propAnalyticsNode_parent_uuid int = iota
    propAnalyticsNode_parent_type int = iota
    propAnalyticsNode_fq_name int = iota
    propAnalyticsNode_id_perms int = iota
    propAnalyticsNode_analytics_node_ip_address int = iota
    propAnalyticsNode_perms2 int = iota
    propAnalyticsNode_uuid int = iota
)

// AnalyticsNode 
type AnalyticsNode struct {

    AnalyticsNodeIPAddress IpAddressType `json:"analytics_node_ip_address,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AnalyticsNode) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAnalyticsNode makes AnalyticsNode
func MakeAnalyticsNode() *AnalyticsNode{
    return &AnalyticsNode{
    //TODO(nati): Apply default
    AnalyticsNodeIPAddress: MakeIpAddressType(),
        Perms2: MakePermType2(),
        UUID: "",
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeAnalyticsNodeSlice makes a slice of AnalyticsNode
func MakeAnalyticsNodeSlice() []*AnalyticsNode {
    return []*AnalyticsNode{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AnalyticsNode) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[global_system_config:0xc420182dc0])
    fqn := []string{}
    
    fqn = GlobalSystemConfig{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *AnalyticsNode) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-global_system_config", "_", "-", -1)
}

func (model *AnalyticsNode) GetDefaultName() string {
    return strings.Replace("default-analytics_node", "_", "-", -1)
}

func (model *AnalyticsNode) GetType() string {
    return strings.Replace("analytics_node", "_", "-", -1)
}

func (model *AnalyticsNode) GetFQName() []string {
    return model.FQName
}

func (model *AnalyticsNode) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AnalyticsNode) GetParentType() string {
    return model.ParentType
}

func (model *AnalyticsNode) GetUuid() string {
    return model.UUID
}

func (model *AnalyticsNode) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AnalyticsNode) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AnalyticsNode) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AnalyticsNode) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AnalyticsNode) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAnalyticsNode_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propAnalyticsNode_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propAnalyticsNode_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propAnalyticsNode_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propAnalyticsNode_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propAnalyticsNode_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propAnalyticsNode_analytics_node_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AnalyticsNodeIPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AnalyticsNodeIPAddress as analytics_node_ip_address")
        }
        msg["analytics_node_ip_address"] = &val
    }
    
    if model.modified.Bit(propAnalyticsNode_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propAnalyticsNode_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AnalyticsNode) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AnalyticsNode) UpdateReferences() error {
    return nil
}


