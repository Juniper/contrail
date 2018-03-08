// nolint
package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

// VirtualMachineInterfaceFields is db columns for VirtualMachineInterface
var VirtualMachineInterfaceFields = []string{
	"vrf_assign_rule",
	"vlan_tag_based_bridge_domain",
	"sub_interface_vlan_tag",
	"service_interface_type",
	"local_preference",
	"traffic_direction",
	"udp_port",
	"vtep_dst_mac_address",
	"vtep_dst_ip_address",
	"vni",
	"routing_instance",
	"nic_assisted_mirroring_vlan",
	"nic_assisted_mirroring",
	"nh_mode",
	"juniper_header",
	"encapsulation",
	"analyzer_name",
	"analyzer_mac_address",
	"analyzer_ip_address",
	"mac_address",
	"route",
	"fat_flow_protocol",
	"virtual_machine_interface_disable_policy",
	"dhcp_option",
	"virtual_machine_interface_device_owner",
	"key_value_pair",
	"allowed_address_pair",
	"uuid",
	"port_security_enabled",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"user_visible",
	"permissions_owner_access",
	"permissions_owner",
	"other_access",
	"group_access",
	"group",
	"last_modified",
	"enable",
	"description",
	"creator",
	"created",
	"fq_name",
	"source_port",
	"source_ip",
	"ip_protocol",
	"hashing_configured",
	"destination_port",
	"destination_ip",
	"display_name",
	"configuration_version",
	"annotations_key_value_pair",
}

// VirtualMachineInterfaceRefFields is db reference fields for VirtualMachineInterface
var VirtualMachineInterfaceRefFields = map[string][]string{

	"routing_instance": []string{
		// <schema.Schema Value>
		"ipv6_service_chain_address",
		"direction",
		"mpls_label",
		"vlan_tag",
		"src_mac",
		"service_chain_address",
		"dst_mac",
		"protocol",
	},

	"bridge_domain": []string{
		// <schema.Schema Value>
		"vlan_tag",
	},

	"bgp_router": []string{
	// <schema.Schema Value>

	},

	"security_logging_object": []string{
	// <schema.Schema Value>

	},

	"qos_config": []string{
	// <schema.Schema Value>

	},

	"port_tuple": []string{
	// <schema.Schema Value>

	},

	"physical_interface": []string{
	// <schema.Schema Value>

	},

	"virtual_machine": []string{
	// <schema.Schema Value>

	},

	"security_group": []string{
	// <schema.Schema Value>

	},

	"service_endpoint": []string{
	// <schema.Schema Value>

	},

	"virtual_machine_interface": []string{
	// <schema.Schema Value>

	},

	"interface_route_table": []string{
	// <schema.Schema Value>

	},

	"service_health_check": []string{
	// <schema.Schema Value>

	},

	"virtual_network": []string{
	// <schema.Schema Value>

	},
}

// VirtualMachineInterfaceBackRefFields is db back reference fields for VirtualMachineInterface
var VirtualMachineInterfaceBackRefFields = map[string][]string{}

// VirtualMachineInterfaceParentTypes is possible parents for VirtualMachineInterface
var VirtualMachineInterfaceParents = []string{

	"project",

	"virtual_machine",

	"virtual_router",
}

// CreateVirtualMachineInterface inserts VirtualMachineInterface to DB
// nolint
func (db *DB) createVirtualMachineInterface(
	ctx context.Context,
	request *models.CreateVirtualMachineInterfaceRequest) error {
	qb := db.queryBuilders["virtual_machine_interface"]
	tx := GetTransaction(ctx)
	model := request.VirtualMachineInterface
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), common.MustJSON(model.GetVRFAssignTable().GetVRFAssignRule()),
		bool(model.GetVlanTagBasedBridgeDomain()),
		int(model.GetVirtualMachineInterfaceProperties().GetSubInterfaceVlanTag()),
		string(model.GetVirtualMachineInterfaceProperties().GetServiceInterfaceType()),
		int(model.GetVirtualMachineInterfaceProperties().GetLocalPreference()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetTrafficDirection()),
		int(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetUDPPort()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetStaticNHHeader().GetVtepDSTMacAddress()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetStaticNHHeader().GetVtepDSTIPAddress()),
		int(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetStaticNHHeader().GetVni()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetRoutingInstance()),
		int(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetNicAssistedMirroringVlan()),
		bool(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetNicAssistedMirroring()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetNHMode()),
		bool(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetJuniperHeader()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetEncapsulation()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetAnalyzerName()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetAnalyzerMacAddress()),
		string(model.GetVirtualMachineInterfaceProperties().GetInterfaceMirror().GetMirrorTo().GetAnalyzerIPAddress()),
		common.MustJSON(model.GetVirtualMachineInterfaceMacAddresses().GetMacAddress()),
		common.MustJSON(model.GetVirtualMachineInterfaceHostRoutes().GetRoute()),
		common.MustJSON(model.GetVirtualMachineInterfaceFatFlowProtocols().GetFatFlowProtocol()),
		bool(model.GetVirtualMachineInterfaceDisablePolicy()),
		common.MustJSON(model.GetVirtualMachineInterfaceDHCPOptionList().GetDHCPOption()),
		string(model.GetVirtualMachineInterfaceDeviceOwner()),
		common.MustJSON(model.GetVirtualMachineInterfaceBindings().GetKeyValuePair()),
		common.MustJSON(model.GetVirtualMachineInterfaceAllowedAddressPairs().GetAllowedAddressPair()),
		string(model.GetUUID()),
		bool(model.GetPortSecurityEnabled()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		bool(model.GetIDPerms().GetUserVisible()),
		int(model.GetIDPerms().GetPermissions().GetOwnerAccess()),
		string(model.GetIDPerms().GetPermissions().GetOwner()),
		int(model.GetIDPerms().GetPermissions().GetOtherAccess()),
		int(model.GetIDPerms().GetPermissions().GetGroupAccess()),
		string(model.GetIDPerms().GetPermissions().GetGroup()),
		string(model.GetIDPerms().GetLastModified()),
		bool(model.GetIDPerms().GetEnable()),
		string(model.GetIDPerms().GetDescription()),
		string(model.GetIDPerms().GetCreator()),
		string(model.GetIDPerms().GetCreated()),
		common.MustJSON(model.GetFQName()),
		bool(model.GetEcmpHashingIncludeFields().GetSourcePort()),
		bool(model.GetEcmpHashingIncludeFields().GetSourceIP()),
		bool(model.GetEcmpHashingIncludeFields().GetIPProtocol()),
		bool(model.GetEcmpHashingIncludeFields().GetHashingConfigured()),
		bool(model.GetEcmpHashingIncludeFields().GetDestinationPort()),
		bool(model.GetEcmpHashingIncludeFields().GetDestinationIP()),
		string(model.GetDisplayName()),
		int(model.GetConfigurationVersion()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	for _, ref := range model.VirtualMachineRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("virtual_machine"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineRefs create failed")
		}
	}

	for _, ref := range model.SecurityGroupRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("security_group"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "SecurityGroupRefs create failed")
		}
	}

	for _, ref := range model.ServiceEndpointRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("service_endpoint"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceEndpointRefs create failed")
		}
	}

	for _, ref := range model.VirtualMachineInterfaceRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("virtual_machine_interface"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	for _, ref := range model.InterfaceRouteTableRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("interface_route_table"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "InterfaceRouteTableRefs create failed")
		}
	}

	for _, ref := range model.ServiceHealthCheckRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("service_health_check"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceHealthCheckRefs create failed")
		}
	}

	for _, ref := range model.VirtualNetworkRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("virtual_network"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs create failed")
		}
	}

	for _, ref := range model.RoutingInstanceRefs {

		if ref.Attr == nil {
			ref.Attr = &models.PolicyBasedForwardingRuleType{}
		}

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("routing_instance"), model.UUID, ref.UUID, string(ref.Attr.GetIpv6ServiceChainAddress()),
			string(ref.Attr.GetDirection()),
			int(ref.Attr.GetMPLSLabel()),
			int(ref.Attr.GetVlanTag()),
			string(ref.Attr.GetSRCMac()),
			string(ref.Attr.GetServiceChainAddress()),
			string(ref.Attr.GetDSTMac()),
			string(ref.Attr.GetProtocol()))
		if err != nil {
			return errors.Wrap(err, "RoutingInstanceRefs create failed")
		}
	}

	for _, ref := range model.BridgeDomainRefs {

		if ref.Attr == nil {
			ref.Attr = &models.BridgeDomainMembershipType{}
		}

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("bridge_domain"), model.UUID, ref.UUID, int(ref.Attr.GetVlanTag()))
		if err != nil {
			return errors.Wrap(err, "BridgeDomainRefs create failed")
		}
	}

	for _, ref := range model.BGPRouterRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("bgp_router"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "BGPRouterRefs create failed")
		}
	}

	for _, ref := range model.SecurityLoggingObjectRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("security_logging_object"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "SecurityLoggingObjectRefs create failed")
		}
	}

	for _, ref := range model.QosConfigRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("qos_config"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "QosConfigRefs create failed")
		}
	}

	for _, ref := range model.PortTupleRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("port_tuple"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "PortTupleRefs create failed")
		}
	}

	for _, ref := range model.PhysicalInterfaceRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("physical_interface"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "PhysicalInterfaceRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "virtual_machine_interface",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "virtual_machine_interface", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanVirtualMachineInterface(values map[string]interface{}) (*models.VirtualMachineInterface, error) {
	m := models.MakeVirtualMachineInterface()

	if value, ok := values["vrf_assign_rule"]; ok {

		json.Unmarshal(value.([]byte), &m.VRFAssignTable.VRFAssignRule)

	}

	if value, ok := values["vlan_tag_based_bridge_domain"]; ok {

		m.VlanTagBasedBridgeDomain = common.InterfaceToBool(value)

	}

	if value, ok := values["sub_interface_vlan_tag"]; ok {

		m.VirtualMachineInterfaceProperties.SubInterfaceVlanTag = common.InterfaceToInt64(value)

	}

	if value, ok := values["service_interface_type"]; ok {

		m.VirtualMachineInterfaceProperties.ServiceInterfaceType = common.InterfaceToString(value)

	}

	if value, ok := values["local_preference"]; ok {

		m.VirtualMachineInterfaceProperties.LocalPreference = common.InterfaceToInt64(value)

	}

	if value, ok := values["traffic_direction"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection = common.InterfaceToString(value)

	}

	if value, ok := values["udp_port"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort = common.InterfaceToInt64(value)

	}

	if value, ok := values["vtep_dst_mac_address"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress = common.InterfaceToString(value)

	}

	if value, ok := values["vtep_dst_ip_address"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress = common.InterfaceToString(value)

	}

	if value, ok := values["vni"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni = common.InterfaceToInt64(value)

	}

	if value, ok := values["routing_instance"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance = common.InterfaceToString(value)

	}

	if value, ok := values["nic_assisted_mirroring_vlan"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan = common.InterfaceToInt64(value)

	}

	if value, ok := values["nic_assisted_mirroring"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring = common.InterfaceToBool(value)

	}

	if value, ok := values["nh_mode"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode = common.InterfaceToString(value)

	}

	if value, ok := values["juniper_header"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader = common.InterfaceToBool(value)

	}

	if value, ok := values["encapsulation"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation = common.InterfaceToString(value)

	}

	if value, ok := values["analyzer_name"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName = common.InterfaceToString(value)

	}

	if value, ok := values["analyzer_mac_address"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress = common.InterfaceToString(value)

	}

	if value, ok := values["analyzer_ip_address"]; ok {

		m.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress = common.InterfaceToString(value)

	}

	if value, ok := values["mac_address"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceMacAddresses.MacAddress)

	}

	if value, ok := values["route"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceHostRoutes.Route)

	}

	if value, ok := values["fat_flow_protocol"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceFatFlowProtocols.FatFlowProtocol)

	}

	if value, ok := values["virtual_machine_interface_disable_policy"]; ok {

		m.VirtualMachineInterfaceDisablePolicy = common.InterfaceToBool(value)

	}

	if value, ok := values["dhcp_option"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceDHCPOptionList.DHCPOption)

	}

	if value, ok := values["virtual_machine_interface_device_owner"]; ok {

		m.VirtualMachineInterfaceDeviceOwner = common.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceBindings.KeyValuePair)

	}

	if value, ok := values["allowed_address_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair)

	}

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["port_security_enabled"]; ok {

		m.PortSecurityEnabled = common.InterfaceToBool(value)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		m.Perms2.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["owner"]; ok {

		m.Perms2.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["global_access"]; ok {

		m.Perms2.GlobalAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = common.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = common.InterfaceToString(value)

	}

	if value, ok := values["user_visible"]; ok {

		m.IDPerms.UserVisible = common.InterfaceToBool(value)

	}

	if value, ok := values["permissions_owner_access"]; ok {

		m.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["permissions_owner"]; ok {

		m.IDPerms.Permissions.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["other_access"]; ok {

		m.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group_access"]; ok {

		m.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group"]; ok {

		m.IDPerms.Permissions.Group = common.InterfaceToString(value)

	}

	if value, ok := values["last_modified"]; ok {

		m.IDPerms.LastModified = common.InterfaceToString(value)

	}

	if value, ok := values["enable"]; ok {

		m.IDPerms.Enable = common.InterfaceToBool(value)

	}

	if value, ok := values["description"]; ok {

		m.IDPerms.Description = common.InterfaceToString(value)

	}

	if value, ok := values["creator"]; ok {

		m.IDPerms.Creator = common.InterfaceToString(value)

	}

	if value, ok := values["created"]; ok {

		m.IDPerms.Created = common.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["source_port"]; ok {

		m.EcmpHashingIncludeFields.SourcePort = common.InterfaceToBool(value)

	}

	if value, ok := values["source_ip"]; ok {

		m.EcmpHashingIncludeFields.SourceIP = common.InterfaceToBool(value)

	}

	if value, ok := values["ip_protocol"]; ok {

		m.EcmpHashingIncludeFields.IPProtocol = common.InterfaceToBool(value)

	}

	if value, ok := values["hashing_configured"]; ok {

		m.EcmpHashingIncludeFields.HashingConfigured = common.InterfaceToBool(value)

	}

	if value, ok := values["destination_port"]; ok {

		m.EcmpHashingIncludeFields.DestinationPort = common.InterfaceToBool(value)

	}

	if value, ok := values["destination_ip"]; ok {

		m.EcmpHashingIncludeFields.DestinationIP = common.InterfaceToBool(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["configuration_version"]; ok {

		m.ConfigurationVersion = common.InterfaceToInt64(value)

	}

	if value, ok := values["annotations_key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_virtual_machine"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceVirtualMachineRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineRefs = append(m.VirtualMachineRefs, referenceModel)

		}
	}

	if value, ok := values["ref_security_group"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceSecurityGroupRef{}
			referenceModel.UUID = uuid
			m.SecurityGroupRefs = append(m.SecurityGroupRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_endpoint"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceServiceEndpointRef{}
			referenceModel.UUID = uuid
			m.ServiceEndpointRefs = append(m.ServiceEndpointRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_interface_route_table"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceInterfaceRouteTableRef{}
			referenceModel.UUID = uuid
			m.InterfaceRouteTableRefs = append(m.InterfaceRouteTableRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_health_check"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceServiceHealthCheckRef{}
			referenceModel.UUID = uuid
			m.ServiceHealthCheckRefs = append(m.ServiceHealthCheckRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_network"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceVirtualNetworkRef{}
			referenceModel.UUID = uuid
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
	}

	if value, ok := values["ref_routing_instance"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceRoutingInstanceRef{}
			referenceModel.UUID = uuid
			m.RoutingInstanceRefs = append(m.RoutingInstanceRefs, referenceModel)

			attr := models.MakePolicyBasedForwardingRuleType()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_bridge_domain"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceBridgeDomainRef{}
			referenceModel.UUID = uuid
			m.BridgeDomainRefs = append(m.BridgeDomainRefs, referenceModel)

			attr := models.MakeBridgeDomainMembershipType()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_bgp_router"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceBGPRouterRef{}
			referenceModel.UUID = uuid
			m.BGPRouterRefs = append(m.BGPRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_security_logging_object"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceSecurityLoggingObjectRef{}
			referenceModel.UUID = uuid
			m.SecurityLoggingObjectRefs = append(m.SecurityLoggingObjectRefs, referenceModel)

		}
	}

	if value, ok := values["ref_qos_config"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfaceQosConfigRef{}
			referenceModel.UUID = uuid
			m.QosConfigRefs = append(m.QosConfigRefs, referenceModel)

		}
	}

	if value, ok := values["ref_port_tuple"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfacePortTupleRef{}
			referenceModel.UUID = uuid
			m.PortTupleRefs = append(m.PortTupleRefs, referenceModel)

		}
	}

	if value, ok := values["ref_physical_interface"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualMachineInterfacePhysicalInterfaceRef{}
			referenceModel.UUID = uuid
			m.PhysicalInterfaceRefs = append(m.PhysicalInterfaceRefs, referenceModel)

		}
	}

	return m, nil
}

// ListVirtualMachineInterface lists VirtualMachineInterface with list spec.
func (db *DB) listVirtualMachineInterface(ctx context.Context, request *models.ListVirtualMachineInterfaceRequest) (response *models.ListVirtualMachineInterfaceResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["virtual_machine_interface"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.VirtualMachineInterface{}

	if spec.ParentFQName != nil {
		parentMetaData, err := db.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = models.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}
	query, columns, values := qb.ListQuery(auth, spec)
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "select query failed")
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "row error")
	}

	for rows.Next() {
		valuesMap := map[string]interface{}{}
		values := make([]interface{}, len(columns))
		valuesPointers := make([]interface{}, len(columns))
		for _, index := range columns {
			valuesPointers[index] = &values[index]
		}
		if err := rows.Scan(valuesPointers...); err != nil {
			return nil, errors.Wrap(err, "scan failed")
		}
		for column, index := range columns {
			val := valuesPointers[index].(*interface{})
			valuesMap[column] = *val
		}
		m, err := scanVirtualMachineInterface(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListVirtualMachineInterfaceResponse{
		VirtualMachineInterfaces: result,
	}
	return response, nil
}

// UpdateVirtualMachineInterface updates a resource
func (db *DB) updateVirtualMachineInterface(
	ctx context.Context,
	request *models.UpdateVirtualMachineInterfaceRequest,
) error {
	//TODO
	return nil
}

// DeleteVirtualMachineInterface deletes a resource
func (db *DB) deleteVirtualMachineInterface(
	ctx context.Context,
	request *models.DeleteVirtualMachineInterfaceRequest) error {
	qb := db.queryBuilders["virtual_machine_interface"]

	selectQuery := qb.SelectForDeleteQuery()
	deleteQuery := qb.DeleteQuery()

	var err error
	var count int
	uuid := request.ID
	tx := GetTransaction(ctx)
	auth := common.GetAuthCTX(ctx)
	if auth.IsAdmin() {
		row := tx.QueryRowContext(ctx, selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	err = db.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}

//CreateVirtualMachineInterface handle a Create API
// nolint
func (db *DB) CreateVirtualMachineInterface(
	ctx context.Context,
	request *models.CreateVirtualMachineInterfaceRequest) (*models.CreateVirtualMachineInterfaceResponse, error) {
	model := request.VirtualMachineInterface
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createVirtualMachineInterface(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine_interface",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateVirtualMachineInterfaceResponse{
		VirtualMachineInterface: request.VirtualMachineInterface,
	}, nil
}

//UpdateVirtualMachineInterface handles a Update request.
func (db *DB) UpdateVirtualMachineInterface(
	ctx context.Context,
	request *models.UpdateVirtualMachineInterfaceRequest) (*models.UpdateVirtualMachineInterfaceResponse, error) {
	model := request.VirtualMachineInterface
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateVirtualMachineInterface(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine_interface",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateVirtualMachineInterfaceResponse{
		VirtualMachineInterface: model,
	}, nil
}

//DeleteVirtualMachineInterface delete a resource.
func (db *DB) DeleteVirtualMachineInterface(ctx context.Context, request *models.DeleteVirtualMachineInterfaceRequest) (*models.DeleteVirtualMachineInterfaceResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteVirtualMachineInterface(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteVirtualMachineInterfaceResponse{
		ID: request.ID,
	}, nil
}

//GetVirtualMachineInterface a Get request.
func (db *DB) GetVirtualMachineInterface(ctx context.Context, request *models.GetVirtualMachineInterfaceRequest) (response *models.GetVirtualMachineInterfaceResponse, err error) {
	spec := &models.ListSpec{
		Limit:  1,
		Detail: true,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListVirtualMachineInterfaceRequest{
		Spec: spec,
	}
	var result *models.ListVirtualMachineInterfaceResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listVirtualMachineInterface(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.VirtualMachineInterfaces) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetVirtualMachineInterfaceResponse{
		VirtualMachineInterface: result.VirtualMachineInterfaces[0],
	}
	return response, nil
}

//ListVirtualMachineInterface handles a List service Request.
// nolint
func (db *DB) ListVirtualMachineInterface(
	ctx context.Context,
	request *models.ListVirtualMachineInterfaceRequest) (response *models.ListVirtualMachineInterfaceResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listVirtualMachineInterface(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
