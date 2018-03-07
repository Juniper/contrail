// nolint
package db

import (
	"context"
	"github.com/satori/go.uuid"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"
)

//For skip import error.
var _ = errors.New("")

func TestNetworkIpam(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeNetworkIpam()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var VirtualDNSCreateRef []*models.NetworkIpamVirtualDNSRef
	var VirtualDNSRefModel *models.VirtualDNS

	VirtualDNSRefUUID := uuid.NewV4().String()
	VirtualDNSRefUUID1 := uuid.NewV4().String()
	VirtualDNSRefUUID2 := uuid.NewV4().String()

	VirtualDNSRefModel = models.MakeVirtualDNS()
	VirtualDNSRefModel.UUID = VirtualDNSRefUUID
	VirtualDNSRefModel.FQName = []string{"test", VirtualDNSRefUUID}
	_, err = db.CreateVirtualDNS(ctx, &models.CreateVirtualDNSRequest{
		VirtualDNS: VirtualDNSRefModel,
	})
	VirtualDNSRefModel.UUID = VirtualDNSRefUUID1
	VirtualDNSRefModel.FQName = []string{"test", VirtualDNSRefUUID1}
	_, err = db.CreateVirtualDNS(ctx, &models.CreateVirtualDNSRequest{
		VirtualDNS: VirtualDNSRefModel,
	})
	VirtualDNSRefModel.UUID = VirtualDNSRefUUID2
	VirtualDNSRefModel.FQName = []string{"test", VirtualDNSRefUUID2}
	_, err = db.CreateVirtualDNS(ctx, &models.CreateVirtualDNSRequest{
		VirtualDNS: VirtualDNSRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualDNSCreateRef = append(VirtualDNSCreateRef,
		&models.NetworkIpamVirtualDNSRef{UUID: VirtualDNSRefUUID, To: []string{"test", VirtualDNSRefUUID}})
	VirtualDNSCreateRef = append(VirtualDNSCreateRef,
		&models.NetworkIpamVirtualDNSRef{UUID: VirtualDNSRefUUID2, To: []string{"test", VirtualDNSRefUUID2}})
	model.VirtualDNSRefs = VirtualDNSCreateRef

	//create project to which resource is shared
	projectModel := models.MakeProject()

	projectModel.UUID = uuid.NewV4().String()
	projectModel.FQName = []string{"default-domain-test", projectModel.UUID}
	projectModel.Perms2.Owner = "admin"

	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:" + projectModel.UUID, TenantAccess: 7})
	model.Perms2.Share = createShare

	_, err = db.CreateProject(ctx, &models.CreateProjectRequest{
		Project: projectModel,
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	_, err = db.CreateNetworkIpam(ctx,
		&models.CreateNetworkIpamRequest{
			NetworkIpam: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListNetworkIpam(ctx, &models.ListNetworkIpamRequest{
		Spec: &models.ListSpec{Limit: 1,
			Filters: []*models.Filter{
				&models.Filter{
					Key:    "uuid",
					Values: []string{model.UUID},
				},
			},
		}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.NetworkIpams) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteNetworkIpam(ctxDemo,
		&models.DeleteNetworkIpamRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateNetworkIpam(ctx,
		&models.CreateNetworkIpamRequest{
			NetworkIpam: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteNetworkIpam(ctx,
		&models.DeleteNetworkIpamRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetNetworkIpam(ctx, &models.GetNetworkIpamRequest{
		ID: model.UUID})
	if err == nil {
		t.Fatal("expected not found error")
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}
	return
}
