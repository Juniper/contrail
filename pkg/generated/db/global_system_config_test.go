package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestGlobalSystemConfig(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "metadata")
	common.UseTable(db, "global_system_config")
	defer func() {
		common.ClearTable(db, "global_system_config")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeGlobalSystemConfig()
	model.UUID = "global_system_config_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "global_system_config_dummy"}
	model.Perms2.Owner = "admin"
	updateMap := map[string]interface{}{}

	common.SetValueByPath(updateMap, ".UUID", ".", "test")

	common.SetValueByPath(updateMap, ".UserDefinedLogStatistics.Statlist", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".PluginTuning.PluginProperty", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")

	common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")

	common.SetValueByPath(updateMap, ".ParentType", ".", "test")

	common.SetValueByPath(updateMap, ".MacMoveControl.MacMoveTimeWindow", ".", 1.0)

	common.SetValueByPath(updateMap, ".MacMoveControl.MacMoveLimitAction", ".", "test")

	common.SetValueByPath(updateMap, ".MacMoveControl.MacMoveLimit", ".", 1.0)

	common.SetValueByPath(updateMap, ".MacLimitControl.MacLimitAction", ".", "test")

	common.SetValueByPath(updateMap, ".MacLimitControl.MacLimit", ".", 1.0)

	common.SetValueByPath(updateMap, ".MacAgingTime", ".", 1.0)

	common.SetValueByPath(updateMap, ".IPFabricSubnets.Subnet", ".", `{"test":"test"}`)

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

	common.SetValueByPath(updateMap, ".IbgpAutoMesh", ".", true)

	common.SetValueByPath(updateMap, ".GracefulRestartParameters.XMPPHelperEnable", ".", true)

	common.SetValueByPath(updateMap, ".GracefulRestartParameters.RestartTime", ".", 1.0)

	common.SetValueByPath(updateMap, ".GracefulRestartParameters.LongLivedRestartTime", ".", 1.0)

	common.SetValueByPath(updateMap, ".GracefulRestartParameters.EndOfRibTimeout", ".", 1.0)

	common.SetValueByPath(updateMap, ".GracefulRestartParameters.Enable", ".", true)

	common.SetValueByPath(updateMap, ".GracefulRestartParameters.BGPHelperEnable", ".", true)

	common.SetValueByPath(updateMap, ".FQName", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".DisplayName", ".", "test")

	common.SetValueByPath(updateMap, ".ConfigVersion", ".", "test")

	common.SetValueByPath(updateMap, ".BgpaasParameters.PortStart", ".", 1.0)

	common.SetValueByPath(updateMap, ".BgpaasParameters.PortEnd", ".", 1.0)

	common.SetValueByPath(updateMap, ".BGPAlwaysCompareMed", ".", true)

	common.SetValueByPath(updateMap, ".AutonomousSystem", ".", 1.0)

	common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".AlarmEnable", ".", true)

	common.SetValueByPath(updateMap, "uuid", ".", "global_system_config_dummy_uuid")

	common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})

	common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateGlobalSystemConfig(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return UpdateGlobalSystemConfig(tx, model.UUID, updateMap)
	})
	if err != nil {
		t.Fatal("update failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListGlobalSystemConfig(tx, &common.ListSpec{Limit: 1})
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
		return DeleteGlobalSystemConfig(tx, model.UUID,
			common.NewAuthContext("default", "demo", "demo", []string{}),
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteGlobalSystemConfig(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListGlobalSystemConfig(tx, &common.ListSpec{Limit: 1})
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
