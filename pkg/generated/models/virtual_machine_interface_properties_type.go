
package models
// VirtualMachineInterfacePropertiesType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVirtualMachineInterfacePropertiesType_sub_interface_vlan_tag int = iota
    propVirtualMachineInterfacePropertiesType_local_preference int = iota
    propVirtualMachineInterfacePropertiesType_interface_mirror int = iota
    propVirtualMachineInterfacePropertiesType_service_interface_type int = iota
)

// VirtualMachineInterfacePropertiesType 
type VirtualMachineInterfacePropertiesType struct {

    LocalPreference int `json:"local_preference,omitempty"`
    InterfaceMirror *InterfaceMirrorType `json:"interface_mirror,omitempty"`
    ServiceInterfaceType ServiceInterfaceType `json:"service_interface_type,omitempty"`
    SubInterfaceVlanTag int `json:"sub_interface_vlan_tag,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *VirtualMachineInterfacePropertiesType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVirtualMachineInterfacePropertiesType makes VirtualMachineInterfacePropertiesType
func MakeVirtualMachineInterfacePropertiesType() *VirtualMachineInterfacePropertiesType{
    return &VirtualMachineInterfacePropertiesType{
    //TODO(nati): Apply default
    SubInterfaceVlanTag: 0,
        LocalPreference: 0,
        InterfaceMirror: MakeInterfaceMirrorType(),
        ServiceInterfaceType: MakeServiceInterfaceType(),
        
        modified: big.NewInt(0),
    }
}



// MakeVirtualMachineInterfacePropertiesTypeSlice makes a slice of VirtualMachineInterfacePropertiesType
func MakeVirtualMachineInterfacePropertiesTypeSlice() []*VirtualMachineInterfacePropertiesType {
    return []*VirtualMachineInterfacePropertiesType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualMachineInterfacePropertiesType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VirtualMachineInterfacePropertiesType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualMachineInterfacePropertiesType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VirtualMachineInterfacePropertiesType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualMachineInterfacePropertiesType) GetFQName() []string {
    return model.FQName
}

func (model *VirtualMachineInterfacePropertiesType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualMachineInterfacePropertiesType) GetParentType() string {
    return model.ParentType
}

func (model *VirtualMachineInterfacePropertiesType) GetUuid() string {
    return model.UUID
}

func (model *VirtualMachineInterfacePropertiesType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualMachineInterfacePropertiesType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualMachineInterfacePropertiesType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualMachineInterfacePropertiesType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualMachineInterfacePropertiesType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVirtualMachineInterfacePropertiesType_sub_interface_vlan_tag) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SubInterfaceVlanTag); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SubInterfaceVlanTag as sub_interface_vlan_tag")
        }
        msg["sub_interface_vlan_tag"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterfacePropertiesType_local_preference) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LocalPreference); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LocalPreference as local_preference")
        }
        msg["local_preference"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterfacePropertiesType_interface_mirror) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.InterfaceMirror); err != nil {
            return nil, errors.Wrap(err, "Marshal of: InterfaceMirror as interface_mirror")
        }
        msg["interface_mirror"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterfacePropertiesType_service_interface_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceInterfaceType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceInterfaceType as service_interface_type")
        }
        msg["service_interface_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VirtualMachineInterfacePropertiesType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualMachineInterfacePropertiesType) UpdateReferences() error {
    return nil
}


