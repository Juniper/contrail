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

func TestFirewallRule(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeFirewallRule()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var ServiceGroupCreateRef []*models.FirewallRuleServiceGroupRef
	var ServiceGroupRefModel *models.ServiceGroup

	ServiceGroupRefUUID := uuid.NewV4().String()
	ServiceGroupRefUUID1 := uuid.NewV4().String()
	ServiceGroupRefUUID2 := uuid.NewV4().String()

	ServiceGroupRefModel = models.MakeServiceGroup()
	ServiceGroupRefModel.UUID = ServiceGroupRefUUID
	ServiceGroupRefModel.FQName = []string{"test", ServiceGroupRefUUID}
	_, err = db.CreateServiceGroup(ctx, &models.CreateServiceGroupRequest{
		ServiceGroup: ServiceGroupRefModel,
	})
	ServiceGroupRefModel.UUID = ServiceGroupRefUUID1
	ServiceGroupRefModel.FQName = []string{"test", ServiceGroupRefUUID1}
	_, err = db.CreateServiceGroup(ctx, &models.CreateServiceGroupRequest{
		ServiceGroup: ServiceGroupRefModel,
	})
	ServiceGroupRefModel.UUID = ServiceGroupRefUUID2
	ServiceGroupRefModel.FQName = []string{"test", ServiceGroupRefUUID2}
	_, err = db.CreateServiceGroup(ctx, &models.CreateServiceGroupRequest{
		ServiceGroup: ServiceGroupRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceGroupCreateRef = append(ServiceGroupCreateRef,
		&models.FirewallRuleServiceGroupRef{UUID: ServiceGroupRefUUID, To: []string{"test", ServiceGroupRefUUID}})
	ServiceGroupCreateRef = append(ServiceGroupCreateRef,
		&models.FirewallRuleServiceGroupRef{UUID: ServiceGroupRefUUID2, To: []string{"test", ServiceGroupRefUUID2}})
	model.ServiceGroupRefs = ServiceGroupCreateRef

	var AddressGroupCreateRef []*models.FirewallRuleAddressGroupRef
	var AddressGroupRefModel *models.AddressGroup

	AddressGroupRefUUID := uuid.NewV4().String()
	AddressGroupRefUUID1 := uuid.NewV4().String()
	AddressGroupRefUUID2 := uuid.NewV4().String()

	AddressGroupRefModel = models.MakeAddressGroup()
	AddressGroupRefModel.UUID = AddressGroupRefUUID
	AddressGroupRefModel.FQName = []string{"test", AddressGroupRefUUID}
	_, err = db.CreateAddressGroup(ctx, &models.CreateAddressGroupRequest{
		AddressGroup: AddressGroupRefModel,
	})
	AddressGroupRefModel.UUID = AddressGroupRefUUID1
	AddressGroupRefModel.FQName = []string{"test", AddressGroupRefUUID1}
	_, err = db.CreateAddressGroup(ctx, &models.CreateAddressGroupRequest{
		AddressGroup: AddressGroupRefModel,
	})
	AddressGroupRefModel.UUID = AddressGroupRefUUID2
	AddressGroupRefModel.FQName = []string{"test", AddressGroupRefUUID2}
	_, err = db.CreateAddressGroup(ctx, &models.CreateAddressGroupRequest{
		AddressGroup: AddressGroupRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	AddressGroupCreateRef = append(AddressGroupCreateRef,
		&models.FirewallRuleAddressGroupRef{UUID: AddressGroupRefUUID, To: []string{"test", AddressGroupRefUUID}})
	AddressGroupCreateRef = append(AddressGroupCreateRef,
		&models.FirewallRuleAddressGroupRef{UUID: AddressGroupRefUUID2, To: []string{"test", AddressGroupRefUUID2}})
	model.AddressGroupRefs = AddressGroupCreateRef

	var SecurityLoggingObjectCreateRef []*models.FirewallRuleSecurityLoggingObjectRef
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
		&models.FirewallRuleSecurityLoggingObjectRef{UUID: SecurityLoggingObjectRefUUID, To: []string{"test", SecurityLoggingObjectRefUUID}})
	SecurityLoggingObjectCreateRef = append(SecurityLoggingObjectCreateRef,
		&models.FirewallRuleSecurityLoggingObjectRef{UUID: SecurityLoggingObjectRefUUID2, To: []string{"test", SecurityLoggingObjectRefUUID2}})
	model.SecurityLoggingObjectRefs = SecurityLoggingObjectCreateRef

	var VirtualNetworkCreateRef []*models.FirewallRuleVirtualNetworkRef
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
		&models.FirewallRuleVirtualNetworkRef{UUID: VirtualNetworkRefUUID, To: []string{"test", VirtualNetworkRefUUID}})
	VirtualNetworkCreateRef = append(VirtualNetworkCreateRef,
		&models.FirewallRuleVirtualNetworkRef{UUID: VirtualNetworkRefUUID2, To: []string{"test", VirtualNetworkRefUUID2}})
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

	_, err = db.CreateFirewallRule(ctx,
		&models.CreateFirewallRuleRequest{
			FirewallRule: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListFirewallRule(ctx, &models.ListFirewallRuleRequest{
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
	if len(response.FirewallRules) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteFirewallRule(ctxDemo,
		&models.DeleteFirewallRuleRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateFirewallRule(ctx,
		&models.CreateFirewallRuleRequest{
			FirewallRule: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteFirewallRule(ctx,
		&models.DeleteFirewallRuleRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetFirewallRule(ctx, &models.GetFirewallRuleRequest{
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
