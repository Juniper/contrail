
package models
// Node



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propNode_aws_instance_type int = iota
    propNode_private_power_management_ip int = iota
    propNode_uuid int = iota
    propNode_ip_address int = iota
    propNode_username int = iota
    propNode_gcp_image int = iota
    propNode_private_machine_properties int = iota
    propNode_private_machine_state int = iota
    propNode_fq_name int = iota
    propNode_display_name int = iota
    propNode_mac_address int = iota
    propNode_type int = iota
    propNode_private_power_management_username int = iota
    propNode_perms2 int = iota
    propNode_password int = iota
    propNode_ssh_key int = iota
    propNode_gcp_machine_type int = iota
    propNode_private_power_management_password int = iota
    propNode_parent_uuid int = iota
    propNode_parent_type int = iota
    propNode_id_perms int = iota
    propNode_annotations int = iota
    propNode_hostname int = iota
    propNode_aws_ami int = iota
)

// Node 
type Node struct {

    IPAddress string `json:"ip_address,omitempty"`
    Username string `json:"username,omitempty"`
    AwsInstanceType string `json:"aws_instance_type,omitempty"`
    PrivatePowerManagementIP string `json:"private_power_management_ip,omitempty"`
    UUID string `json:"uuid,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    MacAddress string `json:"mac_address,omitempty"`
    Type string `json:"type,omitempty"`
    GCPImage string `json:"gcp_image,omitempty"`
    PrivateMachineProperties string `json:"private_machine_properties,omitempty"`
    PrivateMachineState string `json:"private_machine_state,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    Password string `json:"password,omitempty"`
    SSHKey string `json:"ssh_key,omitempty"`
    PrivatePowerManagementUsername string `json:"private_power_management_username,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Hostname string `json:"hostname,omitempty"`
    AwsAmi string `json:"aws_ami,omitempty"`
    GCPMachineType string `json:"gcp_machine_type,omitempty"`
    PrivatePowerManagementPassword string `json:"private_power_management_password,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *Node) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeNode makes Node
func MakeNode() *Node{
    return &Node{
    //TODO(nati): Apply default
    PrivatePowerManagementPassword: "",
        ParentUUID: "",
        ParentType: "",
        IDPerms: MakeIdPermsType(),
        Annotations: MakeKeyValuePairs(),
        Hostname: "",
        AwsAmi: "",
        GCPMachineType: "",
        PrivatePowerManagementIP: "",
        UUID: "",
        IPAddress: "",
        Username: "",
        AwsInstanceType: "",
        PrivateMachineProperties: "",
        PrivateMachineState: "",
        FQName: []string{},
        DisplayName: "",
        MacAddress: "",
        Type: "",
        GCPImage: "",
        Perms2: MakePermType2(),
        Password: "",
        SSHKey: "",
        PrivatePowerManagementUsername: "",
        
        modified: big.NewInt(0),
    }
}



// MakeNodeSlice makes a slice of Node
func MakeNodeSlice() []*Node {
    return []*Node{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *Node) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *Node) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *Node) GetDefaultName() string {
    return strings.Replace("default-node", "_", "-", -1)
}

func (model *Node) GetType() string {
    return strings.Replace("node", "_", "-", -1)
}

func (model *Node) GetFQName() []string {
    return model.FQName
}

func (model *Node) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *Node) GetParentType() string {
    return model.ParentType
}

func (model *Node) GetUuid() string {
    return model.UUID
}

func (model *Node) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *Node) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *Node) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *Node) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *Node) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propNode_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propNode_password) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Password); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Password as password")
        }
        msg["password"] = &val
    }
    
    if model.modified.Bit(propNode_ssh_key) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SSHKey); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SSHKey as ssh_key")
        }
        msg["ssh_key"] = &val
    }
    
    if model.modified.Bit(propNode_private_power_management_username) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivatePowerManagementUsername); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivatePowerManagementUsername as private_power_management_username")
        }
        msg["private_power_management_username"] = &val
    }
    
    if model.modified.Bit(propNode_private_power_management_password) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivatePowerManagementPassword); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivatePowerManagementPassword as private_power_management_password")
        }
        msg["private_power_management_password"] = &val
    }
    
    if model.modified.Bit(propNode_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propNode_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propNode_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propNode_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propNode_hostname) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Hostname); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Hostname as hostname")
        }
        msg["hostname"] = &val
    }
    
    if model.modified.Bit(propNode_aws_ami) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AwsAmi); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AwsAmi as aws_ami")
        }
        msg["aws_ami"] = &val
    }
    
    if model.modified.Bit(propNode_gcp_machine_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.GCPMachineType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: GCPMachineType as gcp_machine_type")
        }
        msg["gcp_machine_type"] = &val
    }
    
    if model.modified.Bit(propNode_private_power_management_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivatePowerManagementIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivatePowerManagementIP as private_power_management_ip")
        }
        msg["private_power_management_ip"] = &val
    }
    
    if model.modified.Bit(propNode_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propNode_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IPAddress as ip_address")
        }
        msg["ip_address"] = &val
    }
    
    if model.modified.Bit(propNode_username) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Username); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Username as username")
        }
        msg["username"] = &val
    }
    
    if model.modified.Bit(propNode_aws_instance_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AwsInstanceType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AwsInstanceType as aws_instance_type")
        }
        msg["aws_instance_type"] = &val
    }
    
    if model.modified.Bit(propNode_private_machine_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateMachineProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateMachineProperties as private_machine_properties")
        }
        msg["private_machine_properties"] = &val
    }
    
    if model.modified.Bit(propNode_private_machine_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateMachineState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateMachineState as private_machine_state")
        }
        msg["private_machine_state"] = &val
    }
    
    if model.modified.Bit(propNode_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propNode_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propNode_mac_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacAddress as mac_address")
        }
        msg["mac_address"] = &val
    }
    
    if model.modified.Bit(propNode_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Type); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Type as type")
        }
        msg["type"] = &val
    }
    
    if model.modified.Bit(propNode_gcp_image) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.GCPImage); err != nil {
            return nil, errors.Wrap(err, "Marshal of: GCPImage as gcp_image")
        }
        msg["gcp_image"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *Node) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *Node) UpdateReferences() error {
    return nil
}


