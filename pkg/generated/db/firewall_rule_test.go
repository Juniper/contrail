package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestFirewallRule(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "metadata")
	common.UseTable(db, "firewall_rule")
	defer func() {
		common.ClearTable(db, "firewall_rule")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeFirewallRule()
	model.UUID = "firewall_rule_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "firewall_rule_dummy"}
	model.Perms2.Owner = "admin"
	updateMap := map[string]interface{}{}

	common.SetValueByPath(updateMap, ".UUID", ".", "test")

	common.SetValueByPath(updateMap, ".Service.SRCPorts.StartPort", ".", 1.0)

	common.SetValueByPath(updateMap, ".Service.SRCPorts.EndPort", ".", 1.0)

	common.SetValueByPath(updateMap, ".Service.ProtocolID", ".", 1.0)

	common.SetValueByPath(updateMap, ".Service.Protocol", ".", "test")

	common.SetValueByPath(updateMap, ".Service.DSTPorts.StartPort", ".", 1.0)

	common.SetValueByPath(updateMap, ".Service.DSTPorts.EndPort", ".", 1.0)

	common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")

	common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)

	common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")

	common.SetValueByPath(updateMap, ".ParentType", ".", "test")

	common.SetValueByPath(updateMap, ".MatchTags.TagList", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".MatchTagTypes.TagType", ".", `{"test":"test"}`)

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

	common.SetValueByPath(updateMap, ".Endpoint2.VirtualNetwork", ".", "test")

	common.SetValueByPath(updateMap, ".Endpoint2.Tags", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".Endpoint2.TagIds", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".Endpoint2.Subnet.IPPrefixLen", ".", 1.0)

	common.SetValueByPath(updateMap, ".Endpoint2.Subnet.IPPrefix", ".", "test")

	common.SetValueByPath(updateMap, ".Endpoint2.Any", ".", true)

	common.SetValueByPath(updateMap, ".Endpoint2.AddressGroup", ".", "test")

	common.SetValueByPath(updateMap, ".Endpoint1.VirtualNetwork", ".", "test")

	common.SetValueByPath(updateMap, ".Endpoint1.Tags", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".Endpoint1.TagIds", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".Endpoint1.Subnet.IPPrefixLen", ".", 1.0)

	common.SetValueByPath(updateMap, ".Endpoint1.Subnet.IPPrefix", ".", "test")

	common.SetValueByPath(updateMap, ".Endpoint1.Any", ".", true)

	common.SetValueByPath(updateMap, ".Endpoint1.AddressGroup", ".", "test")

	common.SetValueByPath(updateMap, ".DisplayName", ".", "test")

	common.SetValueByPath(updateMap, ".Direction", ".", "test")

	common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".ActionList.SimpleAction", ".", "test")

	common.SetValueByPath(updateMap, ".ActionList.QosAction", ".", "test")

	common.SetValueByPath(updateMap, ".ActionList.MirrorTo.UDPPort", ".", 1.0)

	common.SetValueByPath(updateMap, ".ActionList.MirrorTo.StaticNHHeader.VtepDSTMacAddress", ".", "test")

	common.SetValueByPath(updateMap, ".ActionList.MirrorTo.StaticNHHeader.VtepDSTIPAddress", ".", "test")

	common.SetValueByPath(updateMap, ".ActionList.MirrorTo.StaticNHHeader.Vni", ".", 1.0)

	common.SetValueByPath(updateMap, ".ActionList.MirrorTo.RoutingInstance", ".", "test")

	common.SetValueByPath(updateMap, ".ActionList.MirrorTo.NicAssistedMirroringVlan", ".", 1.0)

	common.SetValueByPath(updateMap, ".ActionList.MirrorTo.NicAssistedMirroring", ".", true)

	common.SetValueByPath(updateMap, ".ActionList.MirrorTo.NHMode", ".", "test")

	common.SetValueByPath(updateMap, ".ActionList.MirrorTo.JuniperHeader", ".", true)

	common.SetValueByPath(updateMap, ".ActionList.MirrorTo.Encapsulation", ".", "test")

	common.SetValueByPath(updateMap, ".ActionList.MirrorTo.AnalyzerName", ".", "test")

	common.SetValueByPath(updateMap, ".ActionList.MirrorTo.AnalyzerMacAddress", ".", "test")

	common.SetValueByPath(updateMap, ".ActionList.MirrorTo.AnalyzerIPAddress", ".", "test")

	common.SetValueByPath(updateMap, ".ActionList.Log", ".", true)

	common.SetValueByPath(updateMap, ".ActionList.GatewayName", ".", "test")

	common.SetValueByPath(updateMap, ".ActionList.AssignRoutingInstance", ".", "test")

	common.SetValueByPath(updateMap, ".ActionList.ApplyService", ".", `{"test":"test"}`)

	common.SetValueByPath(updateMap, ".ActionList.Alert", ".", true)

	common.SetValueByPath(updateMap, "uuid", ".", "firewall_rule_dummy_uuid")

	common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})

	common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateFirewallRule(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return UpdateFirewallRule(tx, model.UUID, updateMap)
	})
	if err != nil {
		t.Fatal("update failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListFirewallRule(tx, &common.ListSpec{Limit: 1})
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
		return DeleteFirewallRule(tx, model.UUID,
			common.NewAuthContext("default", "demo", "demo", []string{}),
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteFirewallRule(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListFirewallRule(tx, &common.ListSpec{Limit: 1})
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
