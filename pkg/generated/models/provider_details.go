
package models
// ProviderDetails



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propProviderDetails_segmentation_id int = iota
    propProviderDetails_physical_network int = iota
)

// ProviderDetails 
type ProviderDetails struct {

    SegmentationID VlanIdType `json:"segmentation_id,omitempty"`
    PhysicalNetwork string `json:"physical_network,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ProviderDetails) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeProviderDetails makes ProviderDetails
func MakeProviderDetails() *ProviderDetails{
    return &ProviderDetails{
    //TODO(nati): Apply default
    SegmentationID: MakeVlanIdType(),
        PhysicalNetwork: "",
        
        modified: big.NewInt(0),
    }
}



// MakeProviderDetailsSlice makes a slice of ProviderDetails
func MakeProviderDetailsSlice() []*ProviderDetails {
    return []*ProviderDetails{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ProviderDetails) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ProviderDetails) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ProviderDetails) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ProviderDetails) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ProviderDetails) GetFQName() []string {
    return model.FQName
}

func (model *ProviderDetails) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ProviderDetails) GetParentType() string {
    return model.ParentType
}

func (model *ProviderDetails) GetUuid() string {
    return model.UUID
}

func (model *ProviderDetails) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ProviderDetails) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ProviderDetails) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ProviderDetails) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ProviderDetails) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propProviderDetails_segmentation_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SegmentationID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SegmentationID as segmentation_id")
        }
        msg["segmentation_id"] = &val
    }
    
    if model.modified.Bit(propProviderDetails_physical_network) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PhysicalNetwork); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PhysicalNetwork as physical_network")
        }
        msg["physical_network"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ProviderDetails) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ProviderDetails) UpdateReferences() error {
    return nil
}


