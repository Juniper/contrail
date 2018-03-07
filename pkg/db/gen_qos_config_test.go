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

func TestQosConfig(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeQosConfig()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var GlobalSystemConfigCreateRef []*models.QosConfigGlobalSystemConfigRef
	var GlobalSystemConfigRefModel *models.GlobalSystemConfig

	GlobalSystemConfigRefUUID := uuid.NewV4().String()
	GlobalSystemConfigRefUUID1 := uuid.NewV4().String()
	GlobalSystemConfigRefUUID2 := uuid.NewV4().String()

	GlobalSystemConfigRefModel = models.MakeGlobalSystemConfig()
	GlobalSystemConfigRefModel.UUID = GlobalSystemConfigRefUUID
	GlobalSystemConfigRefModel.FQName = []string{"test", GlobalSystemConfigRefUUID}
	_, err = db.CreateGlobalSystemConfig(ctx, &models.CreateGlobalSystemConfigRequest{
		GlobalSystemConfig: GlobalSystemConfigRefModel,
	})
	GlobalSystemConfigRefModel.UUID = GlobalSystemConfigRefUUID1
	GlobalSystemConfigRefModel.FQName = []string{"test", GlobalSystemConfigRefUUID1}
	_, err = db.CreateGlobalSystemConfig(ctx, &models.CreateGlobalSystemConfigRequest{
		GlobalSystemConfig: GlobalSystemConfigRefModel,
	})
	GlobalSystemConfigRefModel.UUID = GlobalSystemConfigRefUUID2
	GlobalSystemConfigRefModel.FQName = []string{"test", GlobalSystemConfigRefUUID2}
	_, err = db.CreateGlobalSystemConfig(ctx, &models.CreateGlobalSystemConfigRequest{
		GlobalSystemConfig: GlobalSystemConfigRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	GlobalSystemConfigCreateRef = append(GlobalSystemConfigCreateRef,
		&models.QosConfigGlobalSystemConfigRef{UUID: GlobalSystemConfigRefUUID, To: []string{"test", GlobalSystemConfigRefUUID}})
	GlobalSystemConfigCreateRef = append(GlobalSystemConfigCreateRef,
		&models.QosConfigGlobalSystemConfigRef{UUID: GlobalSystemConfigRefUUID2, To: []string{"test", GlobalSystemConfigRefUUID2}})
	model.GlobalSystemConfigRefs = GlobalSystemConfigCreateRef

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

	_, err = db.CreateQosConfig(ctx,
		&models.CreateQosConfigRequest{
			QosConfig: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListQosConfig(ctx, &models.ListQosConfigRequest{
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
	if len(response.QosConfigs) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteQosConfig(ctxDemo,
		&models.DeleteQosConfigRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateQosConfig(ctx,
		&models.CreateQosConfigRequest{
			QosConfig: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteQosConfig(ctx,
		&models.DeleteQosConfigRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetQosConfig(ctx, &models.GetQosConfigRequest{
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
