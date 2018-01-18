package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestVirtualNetwork(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "metadata")
	common.UseTable(db, "virtual_network")
	defer func() {
		common.ClearTable(db, "virtual_network")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeVirtualNetwork()
	model.UUID = "virtual_network_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "virtual_network_dummy"}
	model.Perms2.Owner = "admin"
	updateMap := map[string]interface{}{}

	common.SetValueByPath(updateMap, ".VirtualNetworkProperties.VxlanNetworkIdentifier", ".", 1.0)

	common.SetValueByPath(updateMap, ".VirtualNetworkProperties.RPF", ".", "test")

	common.SetValueByPath(updateMap, ".VirtualNetworkProperties.NetworkID", ".", 1.0)

	common.SetValueByPath(updateMap, ".VirtualNetworkProperties.MirrorDestination", ".", true)

	common.SetValueByPath(updateMap, ".VirtualNetworkProperties.ForwardingMode", ".", "test")

	common.SetValueByPath(updateMap, ".VirtualNetworkProperties.AllowTransit", ".", true)

	common.SetValueByPath(updateMap, ".VirtualNetworkNetworkID", ".", 1.0)

	common.SetValueByPath(updateMap, ".UUID", ".", "test")

	common.SetValueByPath(updateMap, ".RouterExternal", ".", true)

	common.SetValueByPath(updateMap, ".RouteTargetList.RouteTarget", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".ProviderProperties.SegmentationID", ".", 1.0)

	common.SetValueByPath(updateMap, ".ProviderProperties.PhysicalNetwork", ".", "test")

	common.SetValueByPath(updateMap, ".PortSecurityEnabled", ".", true)

	common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")

	common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".PBBEvpnEnable", ".", true)

	common.SetValueByPath(updateMap, ".PBBEtreeEnable", ".", true)

	common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")

	common.SetValueByPath(updateMap, ".ParentType", ".", "test")

	common.SetValueByPath(updateMap, ".MultiPolicyServiceChainsEnabled", ".", true)

	common.SetValueByPath(updateMap, ".MacMoveControl.MacMoveTimeWindow", ".", 1.0)

	common.SetValueByPath(updateMap, ".MacMoveControl.MacMoveLimitAction", ".", "test")

	common.SetValueByPath(updateMap, ".MacMoveControl.MacMoveLimit", ".", 1.0)

	common.SetValueByPath(updateMap, ".MacLimitControl.MacLimitAction", ".", "test")

	common.SetValueByPath(updateMap, ".MacLimitControl.MacLimit", ".", 1.0)

	common.SetValueByPath(updateMap, ".MacLearningEnabled", ".", true)

	common.SetValueByPath(updateMap, ".MacAgingTime", ".", 1.0)

	common.SetValueByPath(updateMap, ".Layer2ControlWord", ".", true)

	common.SetValueByPath(updateMap, ".IsShared", ".", true)

	common.SetValueByPath(updateMap, ".ImportRouteTargetList.RouteTarget", ".", `{"test":"test"}`)

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

	common.SetValueByPath(updateMap, ".FloodUnknownUnicast", ".", true)

	common.SetValueByPath(updateMap, ".ExternalIpam", ".", true)

	common.SetValueByPath(updateMap, ".ExportRouteTargetList.RouteTarget", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.SourcePort", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.SourceIP", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.IPProtocol", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.HashingConfigured", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.DestinationPort", ".", true)

	common.SetValueByPath(updateMap, ".EcmpHashingIncludeFields.DestinationIP", ".", true)

	common.SetValueByPath(updateMap, ".DisplayName", ".", "test")

	common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".AddressAllocationMode", ".", "test")

	common.SetValueByPath(updateMap, "uuid", ".", "virtual_network_dummy_uuid")

	common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})

	common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualNetwork(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return UpdateVirtualNetwork(tx, model.UUID, updateMap)
	})
	if err != nil {
		t.Fatal("update failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListVirtualNetwork(tx, &common.ListSpec{Limit: 1})
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
		return DeleteVirtualNetwork(tx, model.UUID,
			common.NewAuthContext("default", "demo", "demo", []string{}),
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteVirtualNetwork(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListVirtualNetwork(tx, &common.ListSpec{Limit: 1})
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
