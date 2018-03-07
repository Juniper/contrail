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

func TestConfigRoot(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeConfigRoot()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var TagCreateRef []*models.ConfigRootTagRef
	var TagRefModel *models.Tag

	TagRefUUID := uuid.NewV4().String()
	TagRefUUID1 := uuid.NewV4().String()
	TagRefUUID2 := uuid.NewV4().String()

	TagRefModel = models.MakeTag()
	TagRefModel.UUID = TagRefUUID
	TagRefModel.FQName = []string{"test", TagRefUUID}
	_, err = db.CreateTag(ctx, &models.CreateTagRequest{
		Tag: TagRefModel,
	})
	TagRefModel.UUID = TagRefUUID1
	TagRefModel.FQName = []string{"test", TagRefUUID1}
	_, err = db.CreateTag(ctx, &models.CreateTagRequest{
		Tag: TagRefModel,
	})
	TagRefModel.UUID = TagRefUUID2
	TagRefModel.FQName = []string{"test", TagRefUUID2}
	_, err = db.CreateTag(ctx, &models.CreateTagRequest{
		Tag: TagRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	TagCreateRef = append(TagCreateRef,
		&models.ConfigRootTagRef{UUID: TagRefUUID, To: []string{"test", TagRefUUID}})
	TagCreateRef = append(TagCreateRef,
		&models.ConfigRootTagRef{UUID: TagRefUUID2, To: []string{"test", TagRefUUID2}})
	model.TagRefs = TagCreateRef

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

	_, err = db.CreateConfigRoot(ctx,
		&models.CreateConfigRootRequest{
			ConfigRoot: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListConfigRoot(ctx, &models.ListConfigRootRequest{
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
	if len(response.ConfigRoots) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteConfigRoot(ctxDemo,
		&models.DeleteConfigRootRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateConfigRoot(ctx,
		&models.CreateConfigRootRequest{
			ConfigRoot: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteConfigRoot(ctx,
		&models.DeleteConfigRootRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetConfigRoot(ctx, &models.GetConfigRootRequest{
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
