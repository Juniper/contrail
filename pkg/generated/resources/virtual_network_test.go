package resources_test

import (
	"os"
	"testing"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail/pkg/generated/resources"
)

func TestVirtualNetworkResourceCreate(t *testing.T) {
	server := os.Getenv("CONTRAIL_API_SERVER")
	password := os.Getenv("CONTRAIL_API_PASSWORD")
	client := contrail.NewClient(server, 8082)
	keyst := contrail.NewKeystoneClient("http://"+server+":5000/v2.0/", "admin", "admin", password, "")
	if err := keyst.Authenticate(); err != nil {
		t.Fatalf("Authentication error: %v", err)
	}
	client.SetAuthenticator(keyst)

	vn := &resources.VirtualNetworkResource{}
	vn.SetDisplayName("VN SpockNet")
	vn.SetIsShared(true)
	vn.SetName("spocknet")
	vn.SetFQName("project", []string{"default-domain", "default-project", "spocknet"})

	if uuid, err := client.UuidByName("virtual-network", "default-domain:default-project:spocknet"); err == nil {
		_ = client.DeleteByUuid("virtual-network", uuid)
	}
	if err := client.Create(vn); err != nil {
		t.Errorf("Creating VN: %v", err)
	}

	if err := client.Delete(vn); err != nil {
		t.Errorf("Deleting VN: %v", err)
	}
}
