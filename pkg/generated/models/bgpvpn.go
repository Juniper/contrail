
package models
// BGPVPN



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propBGPVPN_export_route_target_list int = iota
    propBGPVPN_perms2 int = iota
    propBGPVPN_parent_uuid int = iota
    propBGPVPN_parent_type int = iota
    propBGPVPN_display_name int = iota
    propBGPVPN_route_target_list int = iota
    propBGPVPN_import_route_target_list int = iota
    propBGPVPN_bgpvpn_type int = iota
    propBGPVPN_annotations int = iota
    propBGPVPN_uuid int = iota
    propBGPVPN_fq_name int = iota
    propBGPVPN_id_perms int = iota
)

// BGPVPN 
type BGPVPN struct {

    ParentType string `json:"parent_type,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    ExportRouteTargetList *RouteTargetList `json:"export_route_target_list,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    RouteTargetList *RouteTargetList `json:"route_target_list,omitempty"`
    ImportRouteTargetList *RouteTargetList `json:"import_route_target_list,omitempty"`
    BGPVPNType VpnType `json:"bgpvpn_type,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *BGPVPN) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeBGPVPN makes BGPVPN
func MakeBGPVPN() *BGPVPN{
    return &BGPVPN{
    //TODO(nati): Apply default
    Perms2: MakePermType2(),
        ParentUUID: "",
        ParentType: "",
        DisplayName: "",
        ExportRouteTargetList: MakeRouteTargetList(),
        ImportRouteTargetList: MakeRouteTargetList(),
        BGPVPNType: MakeVpnType(),
        Annotations: MakeKeyValuePairs(),
        UUID: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        RouteTargetList: MakeRouteTargetList(),
        
        modified: big.NewInt(0),
    }
}



// MakeBGPVPNSlice makes a slice of BGPVPN
func MakeBGPVPNSlice() []*BGPVPN {
    return []*BGPVPN{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *BGPVPN) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc4201834a0])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *BGPVPN) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *BGPVPN) GetDefaultName() string {
    return strings.Replace("default-bgpvpn", "_", "-", -1)
}

func (model *BGPVPN) GetType() string {
    return strings.Replace("bgpvpn", "_", "-", -1)
}

func (model *BGPVPN) GetFQName() []string {
    return model.FQName
}

func (model *BGPVPN) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *BGPVPN) GetParentType() string {
    return model.ParentType
}

func (model *BGPVPN) GetUuid() string {
    return model.UUID
}

func (model *BGPVPN) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *BGPVPN) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *BGPVPN) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *BGPVPN) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *BGPVPN) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propBGPVPN_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propBGPVPN_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propBGPVPN_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propBGPVPN_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propBGPVPN_export_route_target_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ExportRouteTargetList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ExportRouteTargetList as export_route_target_list")
        }
        msg["export_route_target_list"] = &val
    }
    
    if model.modified.Bit(propBGPVPN_import_route_target_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ImportRouteTargetList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ImportRouteTargetList as import_route_target_list")
        }
        msg["import_route_target_list"] = &val
    }
    
    if model.modified.Bit(propBGPVPN_bgpvpn_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.BGPVPNType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: BGPVPNType as bgpvpn_type")
        }
        msg["bgpvpn_type"] = &val
    }
    
    if model.modified.Bit(propBGPVPN_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propBGPVPN_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propBGPVPN_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propBGPVPN_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propBGPVPN_route_target_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RouteTargetList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RouteTargetList as route_target_list")
        }
        msg["route_target_list"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *BGPVPN) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *BGPVPN) UpdateReferences() error {
    return nil
}


