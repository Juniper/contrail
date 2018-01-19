package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestPhysicalRouter(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "metadata")
	common.UseTable(db, "physical_router")
	defer func() {
		common.ClearTable(db, "physical_router")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakePhysicalRouter()
	model.UUID = "physical_router_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "physical_router_dummy"}
	model.Perms2.Owner = "admin"
	updateMap := map[string]interface{}{}

	common.SetValueByPath(updateMap, ".UUID", ".", "test")

	common.SetValueByPath(updateMap, ".TelemetryInfo.ServerPort", ".", 1.0)

	common.SetValueByPath(updateMap, ".TelemetryInfo.ServerIP", ".", "test")

	common.SetValueByPath(updateMap, ".TelemetryInfo.Resource", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".PhysicalRouterVNCManaged", ".", true)

	common.SetValueByPath(updateMap, ".PhysicalRouterVendorName", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterUserCredentials.Username", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterUserCredentials.Password", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.Version", ".", 1.0)

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3SecurityName", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3SecurityLevel", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3SecurityEngineID", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3PrivacyProtocol", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3PrivacyPassword", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3EngineTime", ".", 1.0)

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3EngineID", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3EngineBoots", ".", 1.0)

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3ContextEngineID", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3Context", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3AuthenticationProtocol", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V3AuthenticationPassword", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.V2Community", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.Timeout", ".", 1.0)

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.Retries", ".", 1.0)

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMPCredentials.LocalPort", ".", 1.0)

	common.SetValueByPath(updateMap, ".PhysicalRouterSNMP", ".", true)

	common.SetValueByPath(updateMap, ".PhysicalRouterRole", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterProductName", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterManagementIP", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterLoopbackIP", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterLLDP", ".", true)

	common.SetValueByPath(updateMap, ".PhysicalRouterJunosServicePorts.ServicePort", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".PhysicalRouterImageURI", ".", "test")

	common.SetValueByPath(updateMap, ".PhysicalRouterDataplaneIP", ".", "test")

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

	common.SetValueByPath(updateMap, ".FQName", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".DisplayName", ".", "test")

	common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, "uuid", ".", "physical_router_dummy_uuid")

	common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})

	common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreatePhysicalRouter(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return UpdatePhysicalRouter(tx, model.UUID, updateMap)
	})
	if err != nil {
		t.Fatal("update failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListPhysicalRouter(tx, &common.ListSpec{Limit: 1})
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
		return DeletePhysicalRouter(tx, model.UUID,
			common.NewAuthContext("default", "demo", "demo", []string{}),
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeletePhysicalRouter(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListPhysicalRouter(tx, &common.ListSpec{Limit: 1})
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
