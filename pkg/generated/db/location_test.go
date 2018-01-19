package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestLocation(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "metadata")
	common.UseTable(db, "location")
	defer func() {
		common.ClearTable(db, "location")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeLocation()
	model.UUID = "location_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "location_dummy"}
	model.Perms2.Owner = "admin"
	updateMap := map[string]interface{}{}

	common.SetValueByPath(updateMap, ".UUID", ".", "test")

	common.SetValueByPath(updateMap, ".Type", ".", "test")

	common.SetValueByPath(updateMap, ".ProvisioningState", ".", "test")

	common.SetValueByPath(updateMap, ".ProvisioningStartTime", ".", "test")

	common.SetValueByPath(updateMap, ".ProvisioningProgressStage", ".", "test")

	common.SetValueByPath(updateMap, ".ProvisioningProgress", ".", 1.0)

	common.SetValueByPath(updateMap, ".ProvisioningLog", ".", "test")

	common.SetValueByPath(updateMap, ".PrivateRedhatSubscriptionUser", ".", "test")

	common.SetValueByPath(updateMap, ".PrivateRedhatSubscriptionPasword", ".", "test")

	common.SetValueByPath(updateMap, ".PrivateRedhatSubscriptionKey", ".", "test")

	common.SetValueByPath(updateMap, ".PrivateRedhatPoolID", ".", "test")

	common.SetValueByPath(updateMap, ".PrivateOspdVMVcpus", ".", "test")

	common.SetValueByPath(updateMap, ".PrivateOspdVMRAMMB", ".", "test")

	common.SetValueByPath(updateMap, ".PrivateOspdVMName", ".", "test")

	common.SetValueByPath(updateMap, ".PrivateOspdVMDiskGB", ".", "test")

	common.SetValueByPath(updateMap, ".PrivateOspdUserPassword", ".", "test")

	common.SetValueByPath(updateMap, ".PrivateOspdUserName", ".", "test")

	common.SetValueByPath(updateMap, ".PrivateOspdPackageURL", ".", "test")

	common.SetValueByPath(updateMap, ".PrivateNTPHosts", ".", "test")

	common.SetValueByPath(updateMap, ".PrivateDNSServers", ".", "test")

	common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test":"test"}`)

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

	common.SetValueByPath(updateMap, ".GCPSubnet", ".", "test")

	common.SetValueByPath(updateMap, ".GCPRegion", ".", "test")

	common.SetValueByPath(updateMap, ".GCPAsn", ".", 1.0)

	common.SetValueByPath(updateMap, ".GCPAccountInfo", ".", "test")

	common.SetValueByPath(updateMap, ".FQName", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".DisplayName", ".", "test")

	common.SetValueByPath(updateMap, ".AwsSubnet", ".", "test")

	common.SetValueByPath(updateMap, ".AwsSecretKey", ".", "test")

	common.SetValueByPath(updateMap, ".AwsRegion", ".", "test")

	common.SetValueByPath(updateMap, ".AwsAccessKey", ".", "test")

	common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, "uuid", ".", "location_dummy_uuid")

	common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})

	common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateLocation(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return UpdateLocation(tx, model.UUID, updateMap)
	})
	if err != nil {
		t.Fatal("update failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListLocation(tx, &common.ListSpec{Limit: 1})
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
		return DeleteLocation(tx, model.UUID,
			common.NewAuthContext("default", "demo", "demo", []string{}),
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteLocation(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListLocation(tx, &common.ListSpec{Limit: 1})
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
