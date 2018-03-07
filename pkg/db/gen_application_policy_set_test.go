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

func TestApplicationPolicySet(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeApplicationPolicySet()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var FirewallPolicyCreateRef []*models.ApplicationPolicySetFirewallPolicyRef
	var FirewallPolicyRefModel *models.FirewallPolicy

	FirewallPolicyRefUUID := uuid.NewV4().String()
	FirewallPolicyRefUUID1 := uuid.NewV4().String()
	FirewallPolicyRefUUID2 := uuid.NewV4().String()

	FirewallPolicyRefModel = models.MakeFirewallPolicy()
	FirewallPolicyRefModel.UUID = FirewallPolicyRefUUID
	FirewallPolicyRefModel.FQName = []string{"test", FirewallPolicyRefUUID}
	_, err = db.CreateFirewallPolicy(ctx, &models.CreateFirewallPolicyRequest{
		FirewallPolicy: FirewallPolicyRefModel,
	})
	FirewallPolicyRefModel.UUID = FirewallPolicyRefUUID1
	FirewallPolicyRefModel.FQName = []string{"test", FirewallPolicyRefUUID1}
	_, err = db.CreateFirewallPolicy(ctx, &models.CreateFirewallPolicyRequest{
		FirewallPolicy: FirewallPolicyRefModel,
	})
	FirewallPolicyRefModel.UUID = FirewallPolicyRefUUID2
	FirewallPolicyRefModel.FQName = []string{"test", FirewallPolicyRefUUID2}
	_, err = db.CreateFirewallPolicy(ctx, &models.CreateFirewallPolicyRequest{
		FirewallPolicy: FirewallPolicyRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	FirewallPolicyCreateRef = append(FirewallPolicyCreateRef,
		&models.ApplicationPolicySetFirewallPolicyRef{UUID: FirewallPolicyRefUUID, To: []string{"test", FirewallPolicyRefUUID}})
	FirewallPolicyCreateRef = append(FirewallPolicyCreateRef,
		&models.ApplicationPolicySetFirewallPolicyRef{UUID: FirewallPolicyRefUUID2, To: []string{"test", FirewallPolicyRefUUID2}})
	model.FirewallPolicyRefs = FirewallPolicyCreateRef

	var GlobalVrouterConfigCreateRef []*models.ApplicationPolicySetGlobalVrouterConfigRef
	var GlobalVrouterConfigRefModel *models.GlobalVrouterConfig

	GlobalVrouterConfigRefUUID := uuid.NewV4().String()
	GlobalVrouterConfigRefUUID1 := uuid.NewV4().String()
	GlobalVrouterConfigRefUUID2 := uuid.NewV4().String()

	GlobalVrouterConfigRefModel = models.MakeGlobalVrouterConfig()
	GlobalVrouterConfigRefModel.UUID = GlobalVrouterConfigRefUUID
	GlobalVrouterConfigRefModel.FQName = []string{"test", GlobalVrouterConfigRefUUID}
	_, err = db.CreateGlobalVrouterConfig(ctx, &models.CreateGlobalVrouterConfigRequest{
		GlobalVrouterConfig: GlobalVrouterConfigRefModel,
	})
	GlobalVrouterConfigRefModel.UUID = GlobalVrouterConfigRefUUID1
	GlobalVrouterConfigRefModel.FQName = []string{"test", GlobalVrouterConfigRefUUID1}
	_, err = db.CreateGlobalVrouterConfig(ctx, &models.CreateGlobalVrouterConfigRequest{
		GlobalVrouterConfig: GlobalVrouterConfigRefModel,
	})
	GlobalVrouterConfigRefModel.UUID = GlobalVrouterConfigRefUUID2
	GlobalVrouterConfigRefModel.FQName = []string{"test", GlobalVrouterConfigRefUUID2}
	_, err = db.CreateGlobalVrouterConfig(ctx, &models.CreateGlobalVrouterConfigRequest{
		GlobalVrouterConfig: GlobalVrouterConfigRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	GlobalVrouterConfigCreateRef = append(GlobalVrouterConfigCreateRef,
		&models.ApplicationPolicySetGlobalVrouterConfigRef{UUID: GlobalVrouterConfigRefUUID, To: []string{"test", GlobalVrouterConfigRefUUID}})
	GlobalVrouterConfigCreateRef = append(GlobalVrouterConfigCreateRef,
		&models.ApplicationPolicySetGlobalVrouterConfigRef{UUID: GlobalVrouterConfigRefUUID2, To: []string{"test", GlobalVrouterConfigRefUUID2}})
	model.GlobalVrouterConfigRefs = GlobalVrouterConfigCreateRef

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

	_, err = db.CreateApplicationPolicySet(ctx,
		&models.CreateApplicationPolicySetRequest{
			ApplicationPolicySet: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListApplicationPolicySet(ctx, &models.ListApplicationPolicySetRequest{
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
	if len(response.ApplicationPolicySets) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteApplicationPolicySet(ctxDemo,
		&models.DeleteApplicationPolicySetRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateApplicationPolicySet(ctx,
		&models.CreateApplicationPolicySetRequest{
			ApplicationPolicySet: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteApplicationPolicySet(ctx,
		&models.DeleteApplicationPolicySetRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetApplicationPolicySet(ctx, &models.GetApplicationPolicySetRequest{
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
