package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestGlobalVrouterConfig(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "metadata")
	common.UseTable(db, "global_vrouter_config")
	defer func() {
		common.ClearTable(db, "global_vrouter_config")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeGlobalVrouterConfig()
	model.UUID = "global_vrouter_config_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "global_vrouter_config_dummy"}
	model.Perms2.Owner = "admin"
	updateMap := map[string]interface{}{}

	common.SetValueByPath(updateMap, ".VxlanNetworkIdentifierMode", ".", "test")

	common.SetValueByPath(updateMap, ".UUID", ".", "test")

	common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")

	common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")

	common.SetValueByPath(updateMap, ".ParentType", ".", "test")

	common.SetValueByPath(updateMap, ".LinklocalServices.LinklocalServiceEntry", ".", `{"test":"test"}`)

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

	common.SetValueByPath(updateMap, ".ForwardingMode", ".", "test")

	common.SetValueByPath(updateMap, ".FlowExportRate", ".", 1.0)

	common.SetValueByPath(updateMap, ".FlowAgingTimeoutList.FlowAgingTimeout", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".EncapsulationPriorities.Encapsulation", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".EnableSecurityLogging", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.SourcePort", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.SourceIP", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.IPProtocol", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.HashingConfigured", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.DestinationPort", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.DestinationIP", ".", true)

	common.SetValueByPath(updateMap, ".DisplayName", ".", "test")

	common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, "uuid", ".", "global_vrouter_config_dummy_uuid")

	common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})

	common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateGlobalVrouterConfig(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return UpdateGlobalVrouterConfig(tx, model.UUID, updateMap)
	})
	if err != nil {
		t.Fatal("update failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListGlobalVrouterConfig(tx, &common.ListSpec{Limit: 1})
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
		return DeleteGlobalVrouterConfig(tx, model.UUID,
			common.NewAuthContext("default", "demo", "demo", []string{}),
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteGlobalVrouterConfig(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListGlobalVrouterConfig(tx, &common.ListSpec{Limit: 1})
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
