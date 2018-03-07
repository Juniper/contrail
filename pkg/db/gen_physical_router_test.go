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

func TestPhysicalRouter(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakePhysicalRouter()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var BGPRouterCreateRef []*models.PhysicalRouterBGPRouterRef
	var BGPRouterRefModel *models.BGPRouter

	BGPRouterRefUUID := uuid.NewV4().String()
	BGPRouterRefUUID1 := uuid.NewV4().String()
	BGPRouterRefUUID2 := uuid.NewV4().String()

	BGPRouterRefModel = models.MakeBGPRouter()
	BGPRouterRefModel.UUID = BGPRouterRefUUID
	BGPRouterRefModel.FQName = []string{"test", BGPRouterRefUUID}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterRefModel,
	})
	BGPRouterRefModel.UUID = BGPRouterRefUUID1
	BGPRouterRefModel.FQName = []string{"test", BGPRouterRefUUID1}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterRefModel,
	})
	BGPRouterRefModel.UUID = BGPRouterRefUUID2
	BGPRouterRefModel.FQName = []string{"test", BGPRouterRefUUID2}
	_, err = db.CreateBGPRouter(ctx, &models.CreateBGPRouterRequest{
		BGPRouter: BGPRouterRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	BGPRouterCreateRef = append(BGPRouterCreateRef,
		&models.PhysicalRouterBGPRouterRef{UUID: BGPRouterRefUUID, To: []string{"test", BGPRouterRefUUID}})
	BGPRouterCreateRef = append(BGPRouterCreateRef,
		&models.PhysicalRouterBGPRouterRef{UUID: BGPRouterRefUUID2, To: []string{"test", BGPRouterRefUUID2}})
	model.BGPRouterRefs = BGPRouterCreateRef

	var VirtualRouterCreateRef []*models.PhysicalRouterVirtualRouterRef
	var VirtualRouterRefModel *models.VirtualRouter

	VirtualRouterRefUUID := uuid.NewV4().String()
	VirtualRouterRefUUID1 := uuid.NewV4().String()
	VirtualRouterRefUUID2 := uuid.NewV4().String()

	VirtualRouterRefModel = models.MakeVirtualRouter()
	VirtualRouterRefModel.UUID = VirtualRouterRefUUID
	VirtualRouterRefModel.FQName = []string{"test", VirtualRouterRefUUID}
	_, err = db.CreateVirtualRouter(ctx, &models.CreateVirtualRouterRequest{
		VirtualRouter: VirtualRouterRefModel,
	})
	VirtualRouterRefModel.UUID = VirtualRouterRefUUID1
	VirtualRouterRefModel.FQName = []string{"test", VirtualRouterRefUUID1}
	_, err = db.CreateVirtualRouter(ctx, &models.CreateVirtualRouterRequest{
		VirtualRouter: VirtualRouterRefModel,
	})
	VirtualRouterRefModel.UUID = VirtualRouterRefUUID2
	VirtualRouterRefModel.FQName = []string{"test", VirtualRouterRefUUID2}
	_, err = db.CreateVirtualRouter(ctx, &models.CreateVirtualRouterRequest{
		VirtualRouter: VirtualRouterRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualRouterCreateRef = append(VirtualRouterCreateRef,
		&models.PhysicalRouterVirtualRouterRef{UUID: VirtualRouterRefUUID, To: []string{"test", VirtualRouterRefUUID}})
	VirtualRouterCreateRef = append(VirtualRouterCreateRef,
		&models.PhysicalRouterVirtualRouterRef{UUID: VirtualRouterRefUUID2, To: []string{"test", VirtualRouterRefUUID2}})
	model.VirtualRouterRefs = VirtualRouterCreateRef

	var VirtualNetworkCreateRef []*models.PhysicalRouterVirtualNetworkRef
	var VirtualNetworkRefModel *models.VirtualNetwork

	VirtualNetworkRefUUID := uuid.NewV4().String()
	VirtualNetworkRefUUID1 := uuid.NewV4().String()
	VirtualNetworkRefUUID2 := uuid.NewV4().String()

	VirtualNetworkRefModel = models.MakeVirtualNetwork()
	VirtualNetworkRefModel.UUID = VirtualNetworkRefUUID
	VirtualNetworkRefModel.FQName = []string{"test", VirtualNetworkRefUUID}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkRefModel,
	})
	VirtualNetworkRefModel.UUID = VirtualNetworkRefUUID1
	VirtualNetworkRefModel.FQName = []string{"test", VirtualNetworkRefUUID1}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkRefModel,
	})
	VirtualNetworkRefModel.UUID = VirtualNetworkRefUUID2
	VirtualNetworkRefModel.FQName = []string{"test", VirtualNetworkRefUUID2}
	_, err = db.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{
		VirtualNetwork: VirtualNetworkRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualNetworkCreateRef = append(VirtualNetworkCreateRef,
		&models.PhysicalRouterVirtualNetworkRef{UUID: VirtualNetworkRefUUID, To: []string{"test", VirtualNetworkRefUUID}})
	VirtualNetworkCreateRef = append(VirtualNetworkCreateRef,
		&models.PhysicalRouterVirtualNetworkRef{UUID: VirtualNetworkRefUUID2, To: []string{"test", VirtualNetworkRefUUID2}})
	model.VirtualNetworkRefs = VirtualNetworkCreateRef

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

	_, err = db.CreatePhysicalRouter(ctx,
		&models.CreatePhysicalRouterRequest{
			PhysicalRouter: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListPhysicalRouter(ctx, &models.ListPhysicalRouterRequest{
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
	if len(response.PhysicalRouters) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeletePhysicalRouter(ctxDemo,
		&models.DeletePhysicalRouterRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreatePhysicalRouter(ctx,
		&models.CreatePhysicalRouterRequest{
			PhysicalRouter: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeletePhysicalRouter(ctx,
		&models.DeletePhysicalRouterRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetPhysicalRouter(ctx, &models.GetPhysicalRouterRequest{
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
