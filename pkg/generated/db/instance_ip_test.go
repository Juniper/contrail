package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestInstanceIP(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "metadata")
	common.UseTable(db, "instance_ip")
	defer func() {
		common.ClearTable(db, "instance_ip")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeInstanceIP()
	model.UUID = "instance_ip_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "instance_ip_dummy"}
	model.Perms2.Owner = "admin"
	updateMap := map[string]interface{}{}

	common.SetValueByPath(updateMap, ".UUID", ".", "test")

	common.SetValueByPath(updateMap, ".SubnetUUID", ".", "test")

	common.SetValueByPath(updateMap, ".ServiceInstanceIP", ".", true)

	common.SetValueByPath(updateMap, ".ServiceHealthCheckIP", ".", true)

	common.SetValueByPath(updateMap, ".SecondaryIPTrackingIP.IPPrefixLen", ".", 1.0)

	common.SetValueByPath(updateMap, ".SecondaryIPTrackingIP.IPPrefix", ".", "test")

	common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")

	common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")

	common.SetValueByPath(updateMap, ".ParentType", ".", "test")

	common.SetValueByPath(updateMap, ".InstanceIPSecondary", ".", true)

	common.SetValueByPath(updateMap, ".InstanceIPMode", ".", "test")

	common.SetValueByPath(updateMap, ".InstanceIPLocalIP", ".", true)

	common.SetValueByPath(updateMap, ".InstanceIPFamily", ".", "test")

	common.SetValueByPath(updateMap, ".InstanceIPAddress", ".", "test")

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

	common.SetValueByPath(updateMap, ".FQName", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".DisplayName", ".", "test")

	common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, "uuid", ".", "instance_ip_dummy_uuid")

	common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})

	common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateInstanceIP(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return UpdateInstanceIP(tx, model.UUID, updateMap)
	})
	if err != nil {
		t.Fatal("update failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListInstanceIP(tx, &common.ListSpec{Limit: 1})
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
		return DeleteInstanceIP(tx, model.UUID,
			common.NewAuthContext("default", "demo", "demo", []string{}),
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteInstanceIP(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListInstanceIP(tx, &common.ListSpec{Limit: 1})
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
