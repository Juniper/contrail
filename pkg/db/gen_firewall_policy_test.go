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

func TestFirewallPolicy(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeFirewallPolicy()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var FirewallRuleCreateRef []*models.FirewallPolicyFirewallRuleRef
	var FirewallRuleRefModel *models.FirewallRule

	FirewallRuleRefUUID := uuid.NewV4().String()
	FirewallRuleRefUUID1 := uuid.NewV4().String()
	FirewallRuleRefUUID2 := uuid.NewV4().String()

	FirewallRuleRefModel = models.MakeFirewallRule()
	FirewallRuleRefModel.UUID = FirewallRuleRefUUID
	FirewallRuleRefModel.FQName = []string{"test", FirewallRuleRefUUID}
	_, err = db.CreateFirewallRule(ctx, &models.CreateFirewallRuleRequest{
		FirewallRule: FirewallRuleRefModel,
	})
	FirewallRuleRefModel.UUID = FirewallRuleRefUUID1
	FirewallRuleRefModel.FQName = []string{"test", FirewallRuleRefUUID1}
	_, err = db.CreateFirewallRule(ctx, &models.CreateFirewallRuleRequest{
		FirewallRule: FirewallRuleRefModel,
	})
	FirewallRuleRefModel.UUID = FirewallRuleRefUUID2
	FirewallRuleRefModel.FQName = []string{"test", FirewallRuleRefUUID2}
	_, err = db.CreateFirewallRule(ctx, &models.CreateFirewallRuleRequest{
		FirewallRule: FirewallRuleRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	FirewallRuleCreateRef = append(FirewallRuleCreateRef,
		&models.FirewallPolicyFirewallRuleRef{UUID: FirewallRuleRefUUID, To: []string{"test", FirewallRuleRefUUID}})
	FirewallRuleCreateRef = append(FirewallRuleCreateRef,
		&models.FirewallPolicyFirewallRuleRef{UUID: FirewallRuleRefUUID2, To: []string{"test", FirewallRuleRefUUID2}})
	model.FirewallRuleRefs = FirewallRuleCreateRef

	var SecurityLoggingObjectCreateRef []*models.FirewallPolicySecurityLoggingObjectRef
	var SecurityLoggingObjectRefModel *models.SecurityLoggingObject

	SecurityLoggingObjectRefUUID := uuid.NewV4().String()
	SecurityLoggingObjectRefUUID1 := uuid.NewV4().String()
	SecurityLoggingObjectRefUUID2 := uuid.NewV4().String()

	SecurityLoggingObjectRefModel = models.MakeSecurityLoggingObject()
	SecurityLoggingObjectRefModel.UUID = SecurityLoggingObjectRefUUID
	SecurityLoggingObjectRefModel.FQName = []string{"test", SecurityLoggingObjectRefUUID}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectRefModel,
	})
	SecurityLoggingObjectRefModel.UUID = SecurityLoggingObjectRefUUID1
	SecurityLoggingObjectRefModel.FQName = []string{"test", SecurityLoggingObjectRefUUID1}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectRefModel,
	})
	SecurityLoggingObjectRefModel.UUID = SecurityLoggingObjectRefUUID2
	SecurityLoggingObjectRefModel.FQName = []string{"test", SecurityLoggingObjectRefUUID2}
	_, err = db.CreateSecurityLoggingObject(ctx, &models.CreateSecurityLoggingObjectRequest{
		SecurityLoggingObject: SecurityLoggingObjectRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	SecurityLoggingObjectCreateRef = append(SecurityLoggingObjectCreateRef,
		&models.FirewallPolicySecurityLoggingObjectRef{UUID: SecurityLoggingObjectRefUUID, To: []string{"test", SecurityLoggingObjectRefUUID}})
	SecurityLoggingObjectCreateRef = append(SecurityLoggingObjectCreateRef,
		&models.FirewallPolicySecurityLoggingObjectRef{UUID: SecurityLoggingObjectRefUUID2, To: []string{"test", SecurityLoggingObjectRefUUID2}})
	model.SecurityLoggingObjectRefs = SecurityLoggingObjectCreateRef

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

	_, err = db.CreateFirewallPolicy(ctx,
		&models.CreateFirewallPolicyRequest{
			FirewallPolicy: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListFirewallPolicy(ctx, &models.ListFirewallPolicyRequest{
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
	if len(response.FirewallPolicys) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteFirewallPolicy(ctxDemo,
		&models.DeleteFirewallPolicyRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateFirewallPolicy(ctx,
		&models.CreateFirewallPolicyRequest{
			FirewallPolicy: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteFirewallPolicy(ctx,
		&models.DeleteFirewallPolicyRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetFirewallPolicy(ctx, &models.GetFirewallPolicyRequest{
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
