
package models
// DatabaseNode



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propDatabaseNode_uuid int = iota
    propDatabaseNode_parent_type int = iota
    propDatabaseNode_database_node_ip_address int = iota
    propDatabaseNode_display_name int = iota
    propDatabaseNode_annotations int = iota
    propDatabaseNode_perms2 int = iota
    propDatabaseNode_parent_uuid int = iota
    propDatabaseNode_fq_name int = iota
    propDatabaseNode_id_perms int = iota
)

// DatabaseNode 
type DatabaseNode struct {

    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    DatabaseNodeIPAddress IpAddressType `json:"database_node_ip_address,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *DatabaseNode) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeDatabaseNode makes DatabaseNode
func MakeDatabaseNode() *DatabaseNode{
    return &DatabaseNode{
    //TODO(nati): Apply default
    IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ParentUUID: "",
        FQName: []string{},
        DatabaseNodeIPAddress: MakeIpAddressType(),
        UUID: "",
        ParentType: "",
        
        modified: big.NewInt(0),
    }
}



// MakeDatabaseNodeSlice makes a slice of DatabaseNode
func MakeDatabaseNodeSlice() []*DatabaseNode {
    return []*DatabaseNode{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *DatabaseNode) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[global_system_config:0xc420183860])
    fqn := []string{}
    
    fqn = GlobalSystemConfig{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *DatabaseNode) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-global_system_config", "_", "-", -1)
}

func (model *DatabaseNode) GetDefaultName() string {
    return strings.Replace("default-database_node", "_", "-", -1)
}

func (model *DatabaseNode) GetType() string {
    return strings.Replace("database_node", "_", "-", -1)
}

func (model *DatabaseNode) GetFQName() []string {
    return model.FQName
}

func (model *DatabaseNode) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *DatabaseNode) GetParentType() string {
    return model.ParentType
}

func (model *DatabaseNode) GetUuid() string {
    return model.UUID
}

func (model *DatabaseNode) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *DatabaseNode) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *DatabaseNode) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *DatabaseNode) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *DatabaseNode) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propDatabaseNode_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propDatabaseNode_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propDatabaseNode_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propDatabaseNode_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propDatabaseNode_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propDatabaseNode_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propDatabaseNode_database_node_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DatabaseNodeIPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DatabaseNodeIPAddress as database_node_ip_address")
        }
        msg["database_node_ip_address"] = &val
    }
    
    if model.modified.Bit(propDatabaseNode_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propDatabaseNode_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *DatabaseNode) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *DatabaseNode) UpdateReferences() error {
    return nil
}


