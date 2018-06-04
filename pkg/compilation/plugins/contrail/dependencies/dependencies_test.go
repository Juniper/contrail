package dependencies

import (
	"testing"

	"github.com/Juniper/contrail/pkg/models"
)

// TestDependencies - test
func TestDependencies(t *testing.T) {
	t.Logf("Start of Dependencies Test\n")

	ObjsCache := make(map[string]map[string]interface{})

	ObjsCache["virtual_network"] = make(map[string]interface{})
	Vn1 := models.VirtualNetwork{}
	Vn1.UUID = "Virtual-Network-1"
	ObjsCache["virtual_network"][Vn1.UUID] = &Vn1

	Vn2 := models.VirtualNetwork{}
	Vn2.UUID = "Virtual-Network-2"
	ObjsCache["virtual_network"][Vn2.UUID] = &Vn2

	ObjsCache["NetworkPolicy"] = make(map[string]interface{})
	Np1 := models.NetworkPolicy{}
	Np1.UUID = "Network-policy-1"
	ObjsCache["NetworkPolicy"][Np1.UUID] = &Np1

	Np2 := models.NetworkPolicy{}
	Np2.UUID = "Network-policy-2"
	ObjsCache["NetworkPolicy"][Np2.UUID] = &Np2

	Np1Ref := models.VirtualNetworkNetworkPolicyRef{}
	Np1Ref.UUID = Np1.UUID
	Np1Ref.To = []string{"default-domain", "default-project", "Network-Policy-1"}
	Vn1.NetworkPolicyRefs = append(Vn1.NetworkPolicyRefs, &Np1Ref)

	Np2Ref := models.VirtualNetworkNetworkPolicyRef{}
	Np2Ref.UUID = Np2.UUID
	Np2Ref.To = []string{"default-domain", "default-project", "Network-Policy-2"}
	Vn1.NetworkPolicyRefs = append(Vn1.NetworkPolicyRefs, &Np2Ref)

	Vn2.NetworkPolicyRefs = append(Vn2.NetworkPolicyRefs, &Np2Ref)

	Np1.VirtualNetworkBackRefs = append(Np1.VirtualNetworkBackRefs, &Vn1)
	Np2.VirtualNetworkBackRefs = append(Np1.VirtualNetworkBackRefs, &Vn1)

	Np2.VirtualNetworkBackRefs = append(Np1.VirtualNetworkBackRefs, &Vn2)

	d := NewDependencyProcessor(ObjsCache)
	d.Evaluate(Vn1, "VirtualNetwork", "Self")
	t.Logf("Dependencies: %v", d.GetResourcesPretty())

	//t.Errorf("Unexpected number of go-subroutines %d", diffGoRoutines)

	t.Logf("End of Dependencies Test\n")
}

