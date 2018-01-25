package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"
)

func TestNetworkIpam(t *testing.T) {
	t.Parallel()
	db := testDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	common.UseTable(db, "metadata")
	common.UseTable(db, "network_ipam")
	defer func() {
		common.ClearTable(db, "network_ipam")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeNetworkIpam()
	model.UUID = "network_ipam_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "network_ipam_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var VirtualDNScreateref []*models.NetworkIpamVirtualDNSRef
	var VirtualDNSrefModel *models.VirtualDNS
	VirtualDNSrefModel = models.MakeVirtualDNS()
	VirtualDNSrefModel.UUID = "network_ipam_virtual_DNS_ref_uuid"
	VirtualDNSrefModel.FQName = []string{"test", "network_ipam_virtual_DNS_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualDNS(ctx, tx, &models.CreateVirtualDNSRequest{
			VirtualDNS: VirtualDNSrefModel,
		})
	})
	VirtualDNSrefModel.UUID = "network_ipam_virtual_DNS_ref_uuid1"
	VirtualDNSrefModel.FQName = []string{"test", "network_ipam_virtual_DNS_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualDNS(ctx, tx, &models.CreateVirtualDNSRequest{
			VirtualDNS: VirtualDNSrefModel,
		})
	})
	VirtualDNSrefModel.UUID = "network_ipam_virtual_DNS_ref_uuid2"
	VirtualDNSrefModel.FQName = []string{"test", "network_ipam_virtual_DNS_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualDNS(ctx, tx, &models.CreateVirtualDNSRequest{
			VirtualDNS: VirtualDNSrefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualDNScreateref = append(VirtualDNScreateref, &models.NetworkIpamVirtualDNSRef{UUID: "network_ipam_virtual_DNS_ref_uuid", To: []string{"test", "network_ipam_virtual_DNS_ref_uuid"}})
	VirtualDNScreateref = append(VirtualDNScreateref, &models.NetworkIpamVirtualDNSRef{UUID: "network_ipam_virtual_DNS_ref_uuid2", To: []string{"test", "network_ipam_virtual_DNS_ref_uuid2"}})
	model.VirtualDNSRefs = VirtualDNScreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "network_ipam_admin_project_uuid"
	projectModel.FQName = []string{"default-domain-test", "admin-test"}
	projectModel.Perms2.Owner = "admin"
	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:admin-test", TenantAccess: 7})
	model.Perms2.Share = createShare
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateProject(ctx, tx, &models.CreateProjectRequest{
			Project: projectModel,
		})
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	//    //populate update map
	//    updateMap := map[string]interface{}{}
	//
	//
	//    common.SetValueByPath(updateMap, ".UUID", ".", "test")
	//
	//
	//
	//    if ".Perms2.Share" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Perms2.Share", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ParentType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".NetworkIpamMGMT.IpamMethod", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".NetworkIpamMGMT.IpamDNSServer.VirtualDNSServerName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".NetworkIpamMGMT.IpamDNSServer.TenantDNSServerAddress.IPAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".NetworkIpamMGMT.IpamDNSMethod", ".", "test")
	//
	//
	//
	//    if ".NetworkIpamMGMT.HostRoutes.Route" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".NetworkIpamMGMT.HostRoutes.Route", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".NetworkIpamMGMT.HostRoutes.Route", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    if ".NetworkIpamMGMT.DHCPOptionList.DHCPOption" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".NetworkIpamMGMT.DHCPOptionList.DHCPOption", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".NetworkIpamMGMT.DHCPOptionList.DHCPOption", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".NetworkIpamMGMT.CidrBlock.IPPrefixLen", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".NetworkIpamMGMT.CidrBlock.IPPrefix", ".", "test")
	//
	//
	//
	//    if ".IpamSubnets.Subnets" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".IpamSubnets.Subnets", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".IpamSubnets.Subnets", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IpamSubnetMethod", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.UserVisible", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.OwnerAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.Owner", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.OtherAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.GroupAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.Group", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.LastModified", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Enable", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Description", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Creator", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Created", ".", "test")
	//
	//
	//
	//    if ".FQName" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".FQName", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".FQName", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DisplayName", ".", "test")
	//
	//
	//
	//    if ".Annotations.KeyValuePair" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test": "test"}`)
	//    }
	//
	//
	//    common.SetValueByPath(updateMap, "uuid", ".", "network_ipam_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var VirtualDNSref []interface{}
	//    VirtualDNSref = append(VirtualDNSref, map[string]interface{}{"operation":"delete", "uuid":"network_ipam_virtual_DNS_ref_uuid", "to": []string{"test", "network_ipam_virtual_DNS_ref_uuid"}})
	//    VirtualDNSref = append(VirtualDNSref, map[string]interface{}{"operation":"add", "uuid":"network_ipam_virtual_DNS_ref_uuid1", "to": []string{"test", "network_ipam_virtual_DNS_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualDNSRefs", ".", VirtualDNSref)
	//
	//
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkIpam(ctx, tx,
			&models.CreateNetworkIpamRequest{
				NetworkIpam: model,
			})
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateNetworkIpam(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_network_ipam_virtual_DNS` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualDNSRefs delete statement failed")
		}
		_, err = stmt.Exec("network_ipam_dummy_uuid", "network_ipam_virtual_DNS_ref_uuid")
		_, err = stmt.Exec("network_ipam_dummy_uuid", "network_ipam_virtual_DNS_ref_uuid1")
		_, err = stmt.Exec("network_ipam_dummy_uuid", "network_ipam_virtual_DNS_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualDNSRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualDNS(ctx, tx,
			&models.DeleteVirtualDNSRequest{
				ID: "network_ipam_virtual_DNS_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref network_ipam_virtual_DNS_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualDNS(ctx, tx,
			&models.DeleteVirtualDNSRequest{
				ID: "network_ipam_virtual_DNS_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref network_ipam_virtual_DNS_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualDNS(
			ctx,
			tx,
			&models.DeleteVirtualDNSRequest{
				ID: "network_ipam_virtual_DNS_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref network_ipam_virtual_DNS_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(ctx, tx, &models.DeleteProjectRequest{
			ID: projectModel.UUID})
	})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListNetworkIpam(ctx, tx, &models.ListNetworkIpamRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.NetworkIpams) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkIpam(ctxDemo, tx,
			&models.DeleteNetworkIpamRequest{
				ID: model.UUID},
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteNetworkIpam(ctx, tx,
			&models.DeleteNetworkIpamRequest{
				ID: model.UUID})
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateNetworkIpam(ctx, tx,
			&models.CreateNetworkIpamRequest{
				NetworkIpam: model})
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListNetworkIpam(ctx, tx, &models.ListNetworkIpamRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.NetworkIpams) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
