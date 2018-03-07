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

func TestProviderAttachment(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeProviderAttachment()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var VirtualRouterCreateRef []*models.ProviderAttachmentVirtualRouterRef
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
		&models.ProviderAttachmentVirtualRouterRef{UUID: VirtualRouterRefUUID, To: []string{"test", VirtualRouterRefUUID}})
	VirtualRouterCreateRef = append(VirtualRouterCreateRef,
		&models.ProviderAttachmentVirtualRouterRef{UUID: VirtualRouterRefUUID2, To: []string{"test", VirtualRouterRefUUID2}})
	model.VirtualRouterRefs = VirtualRouterCreateRef

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

	_, err = db.CreateProviderAttachment(ctx,
		&models.CreateProviderAttachmentRequest{
			ProviderAttachment: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListProviderAttachment(ctx, &models.ListProviderAttachmentRequest{
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
	if len(response.ProviderAttachments) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteProviderAttachment(ctxDemo,
		&models.DeleteProviderAttachmentRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateProviderAttachment(ctx,
		&models.CreateProviderAttachmentRequest{
			ProviderAttachment: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteProviderAttachment(ctx,
		&models.DeleteProviderAttachmentRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetProviderAttachment(ctx, &models.GetProviderAttachmentRequest{
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
