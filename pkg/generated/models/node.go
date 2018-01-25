package models
// Node



import "encoding/json"

// Node 
//proteus:generate
type Node struct {

    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    Hostname string `json:"hostname,omitempty"`
    IPAddress string `json:"ip_address,omitempty"`
    MacAddress string `json:"mac_address,omitempty"`
    Type string `json:"type,omitempty"`
    Password string `json:"password,omitempty"`
    SSHKey string `json:"ssh_key,omitempty"`
    Username string `json:"username,omitempty"`
    AwsAmi string `json:"aws_ami,omitempty"`
    AwsInstanceType string `json:"aws_instance_type,omitempty"`
    GCPImage string `json:"gcp_image,omitempty"`
    GCPMachineType string `json:"gcp_machine_type,omitempty"`
    PrivateMachineProperties string `json:"private_machine_properties,omitempty"`
    PrivateMachineState string `json:"private_machine_state,omitempty"`
    PrivatePowerManagementIP string `json:"private_power_management_ip,omitempty"`
    PrivatePowerManagementPassword string `json:"private_power_management_password,omitempty"`
    PrivatePowerManagementUsername string `json:"private_power_management_username,omitempty"`


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
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Hostname: "",
        IPAddress: "",
        MacAddress: "",
        Type: "",
        Password: "",
        SSHKey: "",
        Username: "",
        AwsAmi: "",
        AwsInstanceType: "",
        GCPImage: "",
        GCPMachineType: "",
        PrivateMachineProperties: "",
        PrivateMachineState: "",
        PrivatePowerManagementIP: "",
        PrivatePowerManagementPassword: "",
        PrivatePowerManagementUsername: "",
        
    }
}



// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
    return []*Node{}
}
