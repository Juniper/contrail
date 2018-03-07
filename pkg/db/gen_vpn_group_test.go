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

func TestVPNGroup(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeVPNGroup()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var LocationCreateRef []*models.VPNGroupLocationRef
	var LocationRefModel *models.Location

	LocationRefUUID := uuid.NewV4().String()
	LocationRefUUID1 := uuid.NewV4().String()
	LocationRefUUID2 := uuid.NewV4().String()

	LocationRefModel = models.MakeLocation()
	LocationRefModel.UUID = LocationRefUUID
	LocationRefModel.FQName = []string{"test", LocationRefUUID}
	_, err = db.CreateLocation(ctx, &models.CreateLocationRequest{
		Location: LocationRefModel,
	})
	LocationRefModel.UUID = LocationRefUUID1
	LocationRefModel.FQName = []string{"test", LocationRefUUID1}
	_, err = db.CreateLocation(ctx, &models.CreateLocationRequest{
		Location: LocationRefModel,
	})
	LocationRefModel.UUID = LocationRefUUID2
	LocationRefModel.FQName = []string{"test", LocationRefUUID2}
	_, err = db.CreateLocation(ctx, &models.CreateLocationRequest{
		Location: LocationRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	LocationCreateRef = append(LocationCreateRef,
		&models.VPNGroupLocationRef{UUID: LocationRefUUID, To: []string{"test", LocationRefUUID}})
	LocationCreateRef = append(LocationCreateRef,
		&models.VPNGroupLocationRef{UUID: LocationRefUUID2, To: []string{"test", LocationRefUUID2}})
	model.LocationRefs = LocationCreateRef

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

	_, err = db.CreateVPNGroup(ctx,
		&models.CreateVPNGroupRequest{
			VPNGroup: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListVPNGroup(ctx, &models.ListVPNGroupRequest{
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
	if len(response.VPNGroups) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteVPNGroup(ctxDemo,
		&models.DeleteVPNGroupRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateVPNGroup(ctx,
		&models.CreateVPNGroupRequest{
			VPNGroup: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteVPNGroup(ctx,
		&models.DeleteVPNGroupRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetVPNGroup(ctx, &models.GetVPNGroupRequest{
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
