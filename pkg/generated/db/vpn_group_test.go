package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"
)

func TestVPNGroup(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "metadata")
	common.UseTable(db, "vpn_group")
	defer func() {
		common.ClearTable(db, "vpn_group")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeVPNGroup()
	model.UUID = "vpn_group_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "vpn_group_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var Locationcreateref []*models.VPNGroupLocationRef
	var LocationrefModel *models.Location
	LocationrefModel = models.MakeLocation()
	LocationrefModel.UUID = "vpn_group_location_ref_uuid"
	LocationrefModel.FQName = []string{"test", "vpn_group_location_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLocation(tx, LocationrefModel)
	})
	LocationrefModel.UUID = "vpn_group_location_ref_uuid1"
	LocationrefModel.FQName = []string{"test", "vpn_group_location_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLocation(tx, LocationrefModel)
	})
	LocationrefModel.UUID = "vpn_group_location_ref_uuid2"
	LocationrefModel.FQName = []string{"test", "vpn_group_location_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLocation(tx, LocationrefModel)
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	Locationcreateref = append(Locationcreateref, &models.VPNGroupLocationRef{UUID: "vpn_group_location_ref_uuid", To: []string{"test", "vpn_group_location_ref_uuid"}})
	Locationcreateref = append(Locationcreateref, &models.VPNGroupLocationRef{UUID: "vpn_group_location_ref_uuid2", To: []string{"test", "vpn_group_location_ref_uuid2"}})
	model.LocationRefs = Locationcreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "vpn_group_admin_project_uuid"
	projectModel.FQName = []string{"default-domain-test", "admin-test"}
	projectModel.Perms2.Owner = "admin"
	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:admin-test", TenantAccess: 7})
	model.Perms2.Share = createShare
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateProject(tx, projectModel)
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	//populate update map
	updateMap := map[string]interface{}{}

	common.SetValueByPath(updateMap, ".UUID", ".", "test")

	common.SetValueByPath(updateMap, ".Type", ".", "test")

	common.SetValueByPath(updateMap, ".ProvisioningState", ".", "test")

	common.SetValueByPath(updateMap, ".ProvisioningStartTime", ".", "test")

	common.SetValueByPath(updateMap, ".ProvisioningProgressStage", ".", "test")

	common.SetValueByPath(updateMap, ".ProvisioningProgress", ".", 1.0)

	common.SetValueByPath(updateMap, ".ProvisioningLog", ".", "test")

	if ".Perms2.Share" == ".Perms2.Share" {
		var share []interface{}
		share = append(share, map[string]interface{}{"tenant": "default-domain-test:admin-test", "tenant_access": 7})
		common.SetValueByPath(updateMap, ".Perms2.Share", ".", share)
	} else {
		common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test": "test"}`)
	}

	common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")

	common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")

	common.SetValueByPath(updateMap, ".ParentType", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.UserVisible", ".", true)

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.OwnerAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.Owner", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.OtherAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.GroupAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".IDPerms.Permissions.Group", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.LastModified", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.Enable", ".", true)

	common.SetValueByPath(updateMap, ".IDPerms.Description", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.Creator", ".", "test")

	common.SetValueByPath(updateMap, ".IDPerms.Created", ".", "test")

	if ".FQName" == ".Perms2.Share" {
		var share []interface{}
		share = append(share, map[string]interface{}{"tenant": "default-domain-test:admin-test", "tenant_access": 7})
		common.SetValueByPath(updateMap, ".FQName", ".", share)
	} else {
		common.SetValueByPath(updateMap, ".FQName", ".", `{"test": "test"}`)
	}

	common.SetValueByPath(updateMap, ".DisplayName", ".", "test")

	if ".Annotations.KeyValuePair" == ".Perms2.Share" {
		var share []interface{}
		share = append(share, map[string]interface{}{"tenant": "default-domain-test:admin-test", "tenant_access": 7})
		common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", share)
	} else {
		common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test": "test"}`)
	}

	common.SetValueByPath(updateMap, "uuid", ".", "vpn_group_dummy_uuid")
	common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

	// Create Attr values for testing ref update(ADD,UPDATE,DELETE)

	var Locationref []interface{}
	Locationref = append(Locationref, map[string]interface{}{"operation": "delete", "uuid": "vpn_group_location_ref_uuid", "to": []string{"test", "vpn_group_location_ref_uuid"}})
	Locationref = append(Locationref, map[string]interface{}{"operation": "add", "uuid": "vpn_group_location_ref_uuid1", "to": []string{"test", "vpn_group_location_ref_uuid1"}})

	common.SetValueByPath(updateMap, "LocationRefs", ".", Locationref)

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVPNGroup(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return UpdateVPNGroup(tx, model.UUID, updateMap)
	})
	if err != nil {
		t.Fatal("update failed", err)
	}

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_vpn_group_location` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing LocationRefs delete statement failed")
		}
		_, err = stmt.Exec("vpn_group_dummy_uuid", "vpn_group_location_ref_uuid")
		_, err = stmt.Exec("vpn_group_dummy_uuid", "vpn_group_location_ref_uuid1")
		_, err = stmt.Exec("vpn_group_dummy_uuid", "vpn_group_location_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "LocationRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLocation(tx, "vpn_group_location_ref_uuid", nil)
	})
	if err != nil {
		t.Fatal("delete ref vpn_group_location_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLocation(tx, "vpn_group_location_ref_uuid1", nil)
	})
	if err != nil {
		t.Fatal("delete ref vpn_group_location_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLocation(tx, "vpn_group_location_ref_uuid2", nil)
	})
	if err != nil {
		t.Fatal("delete ref vpn_group_location_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(tx, projectModel.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListVPNGroup(tx, &common.ListSpec{Limit: 1})
		if err != nil {
			return err
		}
		if len(models) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVPNGroup(tx, model.UUID,
			common.NewAuthContext("default", "demo", "demo", []string{}),
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVPNGroup(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVPNGroup(tx, model)
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListVPNGroup(tx, &common.ListSpec{Limit: 1})
		if err != nil {
			return err
		}
		if len(models) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
