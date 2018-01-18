package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestProject(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "metadata")
	common.UseTable(db, "project")
	defer func() {
		common.ClearTable(db, "project")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeProject()
	model.UUID = "project_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "project_dummy"}
	model.Perms2.Owner = "admin"
	updateMap := map[string]interface{}{}

	common.SetValueByPath(updateMap, ".VxlanRouting", ".", true)

	common.SetValueByPath(updateMap, ".UUID", ".", "test")

	common.SetValueByPath(updateMap, ".Quota.VirtualRouter", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.VirtualNetwork", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.VirtualMachineInterface", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.VirtualIP", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.VirtualDNSRecord", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.VirtualDNS", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.Subnet", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.ServiceTemplate", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.ServiceInstance", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.SecurityLoggingObject", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.SecurityGroupRule", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.SecurityGroup", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.RouteTable", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.NetworkPolicy", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.NetworkIpam", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.LogicalRouter", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.LoadbalancerPool", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.LoadbalancerMember", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.LoadbalancerHealthmonitor", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.InstanceIP", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.GlobalVrouterConfig", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.FloatingIPPool", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.FloatingIP", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.Defaults", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.BGPRouter", ".", 1.0)

	common.SetValueByPath(updateMap, ".Quota.AccessControlList", ".", 1.0)

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

	common.SetValueByPath(updateMap, ".AlarmEnable", ".", true)

	common.SetValueByPath(updateMap, "uuid", ".", "project_dummy_uuid")

	common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})

	common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateProject(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return UpdateProject(tx, model.UUID, updateMap)
	})
	if err != nil {
		t.Fatal("update failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListProject(tx, &common.ListSpec{Limit: 1})
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
		return DeleteProject(tx, model.UUID,
			common.NewAuthContext("default", "demo", "demo", []string{}),
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListProject(tx, &common.ListSpec{Limit: 1})
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
