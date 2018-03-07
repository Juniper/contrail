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

func TestLoadbalancerListener(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeLoadbalancerListener()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var LoadbalancerCreateRef []*models.LoadbalancerListenerLoadbalancerRef
	var LoadbalancerRefModel *models.Loadbalancer

	LoadbalancerRefUUID := uuid.NewV4().String()
	LoadbalancerRefUUID1 := uuid.NewV4().String()
	LoadbalancerRefUUID2 := uuid.NewV4().String()

	LoadbalancerRefModel = models.MakeLoadbalancer()
	LoadbalancerRefModel.UUID = LoadbalancerRefUUID
	LoadbalancerRefModel.FQName = []string{"test", LoadbalancerRefUUID}
	_, err = db.CreateLoadbalancer(ctx, &models.CreateLoadbalancerRequest{
		Loadbalancer: LoadbalancerRefModel,
	})
	LoadbalancerRefModel.UUID = LoadbalancerRefUUID1
	LoadbalancerRefModel.FQName = []string{"test", LoadbalancerRefUUID1}
	_, err = db.CreateLoadbalancer(ctx, &models.CreateLoadbalancerRequest{
		Loadbalancer: LoadbalancerRefModel,
	})
	LoadbalancerRefModel.UUID = LoadbalancerRefUUID2
	LoadbalancerRefModel.FQName = []string{"test", LoadbalancerRefUUID2}
	_, err = db.CreateLoadbalancer(ctx, &models.CreateLoadbalancerRequest{
		Loadbalancer: LoadbalancerRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	LoadbalancerCreateRef = append(LoadbalancerCreateRef,
		&models.LoadbalancerListenerLoadbalancerRef{UUID: LoadbalancerRefUUID, To: []string{"test", LoadbalancerRefUUID}})
	LoadbalancerCreateRef = append(LoadbalancerCreateRef,
		&models.LoadbalancerListenerLoadbalancerRef{UUID: LoadbalancerRefUUID2, To: []string{"test", LoadbalancerRefUUID2}})
	model.LoadbalancerRefs = LoadbalancerCreateRef

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

	_, err = db.CreateLoadbalancerListener(ctx,
		&models.CreateLoadbalancerListenerRequest{
			LoadbalancerListener: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListLoadbalancerListener(ctx, &models.ListLoadbalancerListenerRequest{
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
	if len(response.LoadbalancerListeners) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteLoadbalancerListener(ctxDemo,
		&models.DeleteLoadbalancerListenerRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateLoadbalancerListener(ctx,
		&models.CreateLoadbalancerListenerRequest{
			LoadbalancerListener: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteLoadbalancerListener(ctx,
		&models.DeleteLoadbalancerListenerRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetLoadbalancerListener(ctx, &models.GetLoadbalancerListenerRequest{
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
