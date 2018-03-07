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

func TestSecurityLoggingObject(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeSecurityLoggingObject()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var SecurityGroupCreateRef []*models.SecurityLoggingObjectSecurityGroupRef
	var SecurityGroupRefModel *models.SecurityGroup

	SecurityGroupRefUUID := uuid.NewV4().String()
	SecurityGroupRefUUID1 := uuid.NewV4().String()
	SecurityGroupRefUUID2 := uuid.NewV4().String()

	SecurityGroupRefModel = models.MakeSecurityGroup()
	SecurityGroupRefModel.UUID = SecurityGroupRefUUID
	SecurityGroupRefModel.FQName = []string{"test", SecurityGroupRefUUID}
	_, err = db.CreateSecurityGroup(ctx, &models.CreateSecurityGroupRequest{
		SecurityGroup: SecurityGroupRefModel,
	})
	SecurityGroupRefModel.UUID = SecurityGroupRefUUID1
	SecurityGroupRefModel.FQName = []string{"test", SecurityGroupRefUUID1}
	_, err = db.CreateSecurityGroup(ctx, &models.CreateSecurityGroupRequest{
		SecurityGroup: SecurityGroupRefModel,
	})
	SecurityGroupRefModel.UUID = SecurityGroupRefUUID2
	SecurityGroupRefModel.FQName = []string{"test", SecurityGroupRefUUID2}
	_, err = db.CreateSecurityGroup(ctx, &models.CreateSecurityGroupRequest{
		SecurityGroup: SecurityGroupRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	SecurityGroupCreateRef = append(SecurityGroupCreateRef,
		&models.SecurityLoggingObjectSecurityGroupRef{UUID: SecurityGroupRefUUID, To: []string{"test", SecurityGroupRefUUID}})
	SecurityGroupCreateRef = append(SecurityGroupCreateRef,
		&models.SecurityLoggingObjectSecurityGroupRef{UUID: SecurityGroupRefUUID2, To: []string{"test", SecurityGroupRefUUID2}})
	model.SecurityGroupRefs = SecurityGroupCreateRef

	var NetworkPolicyCreateRef []*models.SecurityLoggingObjectNetworkPolicyRef
	var NetworkPolicyRefModel *models.NetworkPolicy

	NetworkPolicyRefUUID := uuid.NewV4().String()
	NetworkPolicyRefUUID1 := uuid.NewV4().String()
	NetworkPolicyRefUUID2 := uuid.NewV4().String()

	NetworkPolicyRefModel = models.MakeNetworkPolicy()
	NetworkPolicyRefModel.UUID = NetworkPolicyRefUUID
	NetworkPolicyRefModel.FQName = []string{"test", NetworkPolicyRefUUID}
	_, err = db.CreateNetworkPolicy(ctx, &models.CreateNetworkPolicyRequest{
		NetworkPolicy: NetworkPolicyRefModel,
	})
	NetworkPolicyRefModel.UUID = NetworkPolicyRefUUID1
	NetworkPolicyRefModel.FQName = []string{"test", NetworkPolicyRefUUID1}
	_, err = db.CreateNetworkPolicy(ctx, &models.CreateNetworkPolicyRequest{
		NetworkPolicy: NetworkPolicyRefModel,
	})
	NetworkPolicyRefModel.UUID = NetworkPolicyRefUUID2
	NetworkPolicyRefModel.FQName = []string{"test", NetworkPolicyRefUUID2}
	_, err = db.CreateNetworkPolicy(ctx, &models.CreateNetworkPolicyRequest{
		NetworkPolicy: NetworkPolicyRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	NetworkPolicyCreateRef = append(NetworkPolicyCreateRef,
		&models.SecurityLoggingObjectNetworkPolicyRef{UUID: NetworkPolicyRefUUID, To: []string{"test", NetworkPolicyRefUUID}})
	NetworkPolicyCreateRef = append(NetworkPolicyCreateRef,
		&models.SecurityLoggingObjectNetworkPolicyRef{UUID: NetworkPolicyRefUUID2, To: []string{"test", NetworkPolicyRefUUID2}})
	model.NetworkPolicyRefs = NetworkPolicyCreateRef

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

	_, err = db.CreateSecurityLoggingObject(ctx,
		&models.CreateSecurityLoggingObjectRequest{
			SecurityLoggingObject: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListSecurityLoggingObject(ctx, &models.ListSecurityLoggingObjectRequest{
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
	if len(response.SecurityLoggingObjects) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteSecurityLoggingObject(ctxDemo,
		&models.DeleteSecurityLoggingObjectRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateSecurityLoggingObject(ctx,
		&models.CreateSecurityLoggingObjectRequest{
			SecurityLoggingObject: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteSecurityLoggingObject(ctx,
		&models.DeleteSecurityLoggingObjectRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetSecurityLoggingObject(ctx, &models.GetSecurityLoggingObjectRequest{
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
