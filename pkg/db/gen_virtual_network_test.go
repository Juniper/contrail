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

func TestVirtualNetwork(t *testing.T) {
	t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model := models.MakeVirtualNetwork()
	model.UUID = uuid.NewV4().String()
	model.FQName = []string{"default", "default-domain", model.UUID}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var SecurityLoggingObjectCreateRef []*models.VirtualNetworkSecurityLoggingObjectRef
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
		&models.VirtualNetworkSecurityLoggingObjectRef{UUID: SecurityLoggingObjectRefUUID, To: []string{"test", SecurityLoggingObjectRefUUID}})
	SecurityLoggingObjectCreateRef = append(SecurityLoggingObjectCreateRef,
		&models.VirtualNetworkSecurityLoggingObjectRef{UUID: SecurityLoggingObjectRefUUID2, To: []string{"test", SecurityLoggingObjectRefUUID2}})
	model.SecurityLoggingObjectRefs = SecurityLoggingObjectCreateRef

	var NetworkPolicyCreateRef []*models.VirtualNetworkNetworkPolicyRef
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
		&models.VirtualNetworkNetworkPolicyRef{UUID: NetworkPolicyRefUUID, To: []string{"test", NetworkPolicyRefUUID}})
	NetworkPolicyCreateRef = append(NetworkPolicyCreateRef,
		&models.VirtualNetworkNetworkPolicyRef{UUID: NetworkPolicyRefUUID2, To: []string{"test", NetworkPolicyRefUUID2}})
	model.NetworkPolicyRefs = NetworkPolicyCreateRef

	var QosConfigCreateRef []*models.VirtualNetworkQosConfigRef
	var QosConfigRefModel *models.QosConfig

	QosConfigRefUUID := uuid.NewV4().String()
	QosConfigRefUUID1 := uuid.NewV4().String()
	QosConfigRefUUID2 := uuid.NewV4().String()

	QosConfigRefModel = models.MakeQosConfig()
	QosConfigRefModel.UUID = QosConfigRefUUID
	QosConfigRefModel.FQName = []string{"test", QosConfigRefUUID}
	_, err = db.CreateQosConfig(ctx, &models.CreateQosConfigRequest{
		QosConfig: QosConfigRefModel,
	})
	QosConfigRefModel.UUID = QosConfigRefUUID1
	QosConfigRefModel.FQName = []string{"test", QosConfigRefUUID1}
	_, err = db.CreateQosConfig(ctx, &models.CreateQosConfigRequest{
		QosConfig: QosConfigRefModel,
	})
	QosConfigRefModel.UUID = QosConfigRefUUID2
	QosConfigRefModel.FQName = []string{"test", QosConfigRefUUID2}
	_, err = db.CreateQosConfig(ctx, &models.CreateQosConfigRequest{
		QosConfig: QosConfigRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	QosConfigCreateRef = append(QosConfigCreateRef,
		&models.VirtualNetworkQosConfigRef{UUID: QosConfigRefUUID, To: []string{"test", QosConfigRefUUID}})
	QosConfigCreateRef = append(QosConfigCreateRef,
		&models.VirtualNetworkQosConfigRef{UUID: QosConfigRefUUID2, To: []string{"test", QosConfigRefUUID2}})
	model.QosConfigRefs = QosConfigCreateRef

	var RouteTableCreateRef []*models.VirtualNetworkRouteTableRef
	var RouteTableRefModel *models.RouteTable

	RouteTableRefUUID := uuid.NewV4().String()
	RouteTableRefUUID1 := uuid.NewV4().String()
	RouteTableRefUUID2 := uuid.NewV4().String()

	RouteTableRefModel = models.MakeRouteTable()
	RouteTableRefModel.UUID = RouteTableRefUUID
	RouteTableRefModel.FQName = []string{"test", RouteTableRefUUID}
	_, err = db.CreateRouteTable(ctx, &models.CreateRouteTableRequest{
		RouteTable: RouteTableRefModel,
	})
	RouteTableRefModel.UUID = RouteTableRefUUID1
	RouteTableRefModel.FQName = []string{"test", RouteTableRefUUID1}
	_, err = db.CreateRouteTable(ctx, &models.CreateRouteTableRequest{
		RouteTable: RouteTableRefModel,
	})
	RouteTableRefModel.UUID = RouteTableRefUUID2
	RouteTableRefModel.FQName = []string{"test", RouteTableRefUUID2}
	_, err = db.CreateRouteTable(ctx, &models.CreateRouteTableRequest{
		RouteTable: RouteTableRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	RouteTableCreateRef = append(RouteTableCreateRef,
		&models.VirtualNetworkRouteTableRef{UUID: RouteTableRefUUID, To: []string{"test", RouteTableRefUUID}})
	RouteTableCreateRef = append(RouteTableCreateRef,
		&models.VirtualNetworkRouteTableRef{UUID: RouteTableRefUUID2, To: []string{"test", RouteTableRefUUID2}})
	model.RouteTableRefs = RouteTableCreateRef

	var VirtualNetworkCreateRef []*models.VirtualNetworkVirtualNetworkRef
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
		&models.VirtualNetworkVirtualNetworkRef{UUID: VirtualNetworkRefUUID, To: []string{"test", VirtualNetworkRefUUID}})
	VirtualNetworkCreateRef = append(VirtualNetworkCreateRef,
		&models.VirtualNetworkVirtualNetworkRef{UUID: VirtualNetworkRefUUID2, To: []string{"test", VirtualNetworkRefUUID2}})
	model.VirtualNetworkRefs = VirtualNetworkCreateRef

	var BGPVPNCreateRef []*models.VirtualNetworkBGPVPNRef
	var BGPVPNRefModel *models.BGPVPN

	BGPVPNRefUUID := uuid.NewV4().String()
	BGPVPNRefUUID1 := uuid.NewV4().String()
	BGPVPNRefUUID2 := uuid.NewV4().String()

	BGPVPNRefModel = models.MakeBGPVPN()
	BGPVPNRefModel.UUID = BGPVPNRefUUID
	BGPVPNRefModel.FQName = []string{"test", BGPVPNRefUUID}
	_, err = db.CreateBGPVPN(ctx, &models.CreateBGPVPNRequest{
		BGPVPN: BGPVPNRefModel,
	})
	BGPVPNRefModel.UUID = BGPVPNRefUUID1
	BGPVPNRefModel.FQName = []string{"test", BGPVPNRefUUID1}
	_, err = db.CreateBGPVPN(ctx, &models.CreateBGPVPNRequest{
		BGPVPN: BGPVPNRefModel,
	})
	BGPVPNRefModel.UUID = BGPVPNRefUUID2
	BGPVPNRefModel.FQName = []string{"test", BGPVPNRefUUID2}
	_, err = db.CreateBGPVPN(ctx, &models.CreateBGPVPNRequest{
		BGPVPN: BGPVPNRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	BGPVPNCreateRef = append(BGPVPNCreateRef,
		&models.VirtualNetworkBGPVPNRef{UUID: BGPVPNRefUUID, To: []string{"test", BGPVPNRefUUID}})
	BGPVPNCreateRef = append(BGPVPNCreateRef,
		&models.VirtualNetworkBGPVPNRef{UUID: BGPVPNRefUUID2, To: []string{"test", BGPVPNRefUUID2}})
	model.BGPVPNRefs = BGPVPNCreateRef

	var NetworkIpamCreateRef []*models.VirtualNetworkNetworkIpamRef
	var NetworkIpamRefModel *models.NetworkIpam

	NetworkIpamRefUUID := uuid.NewV4().String()
	NetworkIpamRefUUID1 := uuid.NewV4().String()
	NetworkIpamRefUUID2 := uuid.NewV4().String()

	NetworkIpamRefModel = models.MakeNetworkIpam()
	NetworkIpamRefModel.UUID = NetworkIpamRefUUID
	NetworkIpamRefModel.FQName = []string{"test", NetworkIpamRefUUID}
	_, err = db.CreateNetworkIpam(ctx, &models.CreateNetworkIpamRequest{
		NetworkIpam: NetworkIpamRefModel,
	})
	NetworkIpamRefModel.UUID = NetworkIpamRefUUID1
	NetworkIpamRefModel.FQName = []string{"test", NetworkIpamRefUUID1}
	_, err = db.CreateNetworkIpam(ctx, &models.CreateNetworkIpamRequest{
		NetworkIpam: NetworkIpamRefModel,
	})
	NetworkIpamRefModel.UUID = NetworkIpamRefUUID2
	NetworkIpamRefModel.FQName = []string{"test", NetworkIpamRefUUID2}
	_, err = db.CreateNetworkIpam(ctx, &models.CreateNetworkIpamRequest{
		NetworkIpam: NetworkIpamRefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	NetworkIpamCreateRef = append(NetworkIpamCreateRef,
		&models.VirtualNetworkNetworkIpamRef{UUID: NetworkIpamRefUUID, To: []string{"test", NetworkIpamRefUUID}})
	NetworkIpamCreateRef = append(NetworkIpamCreateRef,
		&models.VirtualNetworkNetworkIpamRef{UUID: NetworkIpamRefUUID2, To: []string{"test", NetworkIpamRefUUID2}})
	model.NetworkIpamRefs = NetworkIpamCreateRef

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

	_, err = db.CreateVirtualNetwork(ctx,
		&models.CreateVirtualNetworkRequest{
			VirtualNetwork: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	response, err := db.ListVirtualNetwork(ctx, &models.ListVirtualNetworkRequest{
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
	if len(response.VirtualNetworks) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteVirtualNetwork(ctxDemo,
		&models.DeleteVirtualNetworkRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateVirtualNetwork(ctx,
		&models.CreateVirtualNetworkRequest{
			VirtualNetwork: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteVirtualNetwork(ctx,
		&models.DeleteVirtualNetworkRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	_, err = db.GetVirtualNetwork(ctx, &models.GetVirtualNetworkRequest{
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
