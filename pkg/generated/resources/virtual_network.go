package resources

import (
	"encoding/json"
	"fmt"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail/pkg/generated/models"
)

type VirtualNetworkResource struct {
	//contrail.ObjectBase
	obj       models.VirtualNetwork
	Href      string `json:"href"`
	clientPtr contrail.ObjectInterface
}

func (this VirtualNetworkResource) MarshalJSON() ([]byte, error) {
	fmt.Println("\n- - - - - Custom json marshaller")
	/*
		return json.Marshal(&struct {
			Obj  *models.VirtualNetwork `json:"virtual-network"`
			Blah bool                   `json:"blahfield"`
		}{
			Obj:  &this.obj,
			Blah: false,
		})
	*/
	return json.Marshal(this.obj)
}

func (this *VirtualNetworkResource) GetDefaultParent() []string {
	return []string{"default-domain", "default-project"}
}

func (this *VirtualNetworkResource) GetDefaultParentType() string {
	return "project"
}

func (this *VirtualNetworkResource) GetFQName() []string {
	return this.obj.FQName
}

func (this *VirtualNetworkResource) GetName() string {
	return this.obj.FQName[len(this.obj.FQName)-1]
}

func (this *VirtualNetworkResource) GetType() string {
	return "virtual-network"
}

func (this *VirtualNetworkResource) GetParentType() string {
	return this.obj.ParentType
}

func (this *VirtualNetworkResource) GetUuid() string {
	return this.obj.UUID
}

func (this *VirtualNetworkResource) GetHref() string {
	return this.Href
}

func (this *VirtualNetworkResource) SetName(s string) {
	if len(this.obj.FQName) == 0 {
		fqn := []string{s}
		this.SetFQName(this.GetParentType(), fqn)
		fmt.Printf("New FQ Name: %s, %+v -- %+v\n", s, this.obj.FQName, fqn)
	} else {
		this.obj.FQName[len(this.obj.FQName)-1] = s
		fmt.Println("Modifying FQ Name")
	}
}

func (this *VirtualNetworkResource) SetUuid(s string) {
	if this.clientPtr != nil {
		panic(fmt.Sprintf("Attempt to override uuid for %s", this.obj.UUID))
	}
	this.obj.UUID = s
}

func (this *VirtualNetworkResource) SetFQName(parentType string, fqn []string) {
	fmt.Printf("input FQN: %+v [%v]\n", fqn, len(fqn))
	this.obj.FQName = make([]string, len(fqn))
	n := copy(this.obj.FQName, fqn)
	fmt.Printf("Copied %v\n", n)
	fmt.Printf("Set fqn: %+v\n", this.obj.FQName)
	this.obj.ParentType = parentType
}

func (this *VirtualNetworkResource) SetClient(c contrail.ObjectInterface) {
	this.clientPtr = c
}

func (this *VirtualNetworkResource) UpdateObject() ([]byte, error) {
	return []byte{}, nil
}

func (this *VirtualNetworkResource) UpdateReferences() error {
	return nil
}

func (this *VirtualNetworkResource) UpdateDone() {
}

// SetDisplayName will be auto-geenrated
func (this *VirtualNetworkResource) SetDisplayName(x string) {
	this.obj.DisplayName = x
}

// SetIsShared will be auto-geenrated
func (this *VirtualNetworkResource) SetIsShared(x bool) {
	this.obj.IsShared = x
}
