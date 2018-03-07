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

// VirtualNetworkFields is db columns for VirtualNetwork
var VirtualNetworkFields = []string{
	"vxlan_network_identifier",
	"rpf",
	"network_id",
	"mirror_destination",
	"forwarding_mode",
	"allow_transit",
	"virtual_network_network_id",
	"uuid",
	"router_external",
	"route_target",
	"segmentation_id",
	"physical_network",
	"port_security_enabled",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"pbb_evpn_enable",
	"pbb_etree_enable",
	"parent_uuid",
	"parent_type",
	"multi_policy_service_chains_enabled",
	"mac_move_time_window",
	"mac_move_limit_action",
	"mac_move_limit",
	"mac_limit_action",
	"mac_limit",
	"mac_learning_enabled",
	"mac_aging_time",
	"layer2_control_word",
	"is_shared",
	"import_route_target_list_route_target",
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
	"flood_unknown_unicast",
	"external_ipam",
	"export_route_target_list_route_target",
	"source_port",
	"source_ip",
	"ip_protocol",
	"hashing_configured",
	"destination_port",
	"destination_ip",
	"display_name",
	"configuration_version",
	"key_value_pair",
	"address_allocation_mode",
}

// VirtualNetworkRefFields is db reference fields for VirtualNetwork
var VirtualNetworkRefFields = map[string][]string{

	"virtual_network": []string{
	// <schema.Schema Value>

	},

	"bgpvpn": []string{
	// <schema.Schema Value>

	},

	"network_ipam": []string{
		// <schema.Schema Value>
		"ipam_subnets",
		"route",
	},

	"security_logging_object": []string{
	// <schema.Schema Value>

	},

	"network_policy": []string{
		// <schema.Schema Value>
		"end_time",
		"start_time",
		"off_interval",
		"on_interval",
		"major",
		"minor",
	},

	"qos_config": []string{
	// <schema.Schema Value>

	},

	"route_table": []string{
	// <schema.Schema Value>

	},
}

// VirtualNetworkBackRefFields is db back reference fields for VirtualNetwork
var VirtualNetworkBackRefFields = map[string][]string{

	"access_control_list": []string{
		"uuid",
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
		"display_name",
		"configuration_version",
		"key_value_pair",
		"access_control_list_hash",
		"dynamic",
		"acl_rule",
	},

	"alias_ip_pool": []string{
		"uuid",
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
		"display_name",
		"configuration_version",
		"key_value_pair",
	},

	"bridge_domain": []string{
		"uuid",
		"share",
		"owner_access",
		"owner",
		"global_access",
		"parent_uuid",
		"parent_type",
		"mac_move_time_window",
		"mac_move_limit_action",
		"mac_move_limit",
		"mac_limit_action",
		"mac_limit",
		"mac_learning_enabled",
		"mac_aging_time",
		"isid",
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
		"display_name",
		"configuration_version",
		"key_value_pair",
	},

	"floating_ip_pool": []string{
		"uuid",
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
		"subnet_uuid",
		"display_name",
		"configuration_version",
		"key_value_pair",
	},

	"routing_instance": []string{
		"uuid",
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
		"display_name",
		"configuration_version",
		"key_value_pair",
	},
}

// VirtualNetworkParentTypes is possible parents for VirtualNetwork
var VirtualNetworkParents = []string{

	"project",
}

// CreateVirtualNetwork inserts VirtualNetwork to DB
// nolint
func (db *DB) createVirtualNetwork(
	ctx context.Context,
	request *models.CreateVirtualNetworkRequest) error {
	qb := db.queryBuilders["virtual_network"]
	tx := GetTransaction(ctx)
	model := request.VirtualNetwork
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), int(model.GetVirtualNetworkProperties().GetVxlanNetworkIdentifier()),
		string(model.GetVirtualNetworkProperties().GetRPF()),
		int(model.GetVirtualNetworkProperties().GetNetworkID()),
		bool(model.GetVirtualNetworkProperties().GetMirrorDestination()),
		string(model.GetVirtualNetworkProperties().GetForwardingMode()),
		bool(model.GetVirtualNetworkProperties().GetAllowTransit()),
		int(model.GetVirtualNetworkNetworkID()),
		string(model.GetUUID()),
		bool(model.GetRouterExternal()),
		common.MustJSON(model.GetRouteTargetList().GetRouteTarget()),
		int(model.GetProviderProperties().GetSegmentationID()),
		string(model.GetProviderProperties().GetPhysicalNetwork()),
		bool(model.GetPortSecurityEnabled()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		bool(model.GetPBBEvpnEnable()),
		bool(model.GetPBBEtreeEnable()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		bool(model.GetMultiPolicyServiceChainsEnabled()),
		int(model.GetMacMoveControl().GetMacMoveTimeWindow()),
		string(model.GetMacMoveControl().GetMacMoveLimitAction()),
		int(model.GetMacMoveControl().GetMacMoveLimit()),
		string(model.GetMacLimitControl().GetMacLimitAction()),
		int(model.GetMacLimitControl().GetMacLimit()),
		bool(model.GetMacLearningEnabled()),
		int(model.GetMacAgingTime()),
		bool(model.GetLayer2ControlWord()),
		bool(model.GetIsShared()),
		common.MustJSON(model.GetImportRouteTargetList().GetRouteTarget()),
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
		bool(model.GetFloodUnknownUnicast()),
		bool(model.GetExternalIpam()),
		common.MustJSON(model.GetExportRouteTargetList().GetRouteTarget()),
		bool(model.GetEcmpHashingIncludeFields().GetSourcePort()),
		bool(model.GetEcmpHashingIncludeFields().GetSourceIP()),
		bool(model.GetEcmpHashingIncludeFields().GetIPProtocol()),
		bool(model.GetEcmpHashingIncludeFields().GetHashingConfigured()),
		bool(model.GetEcmpHashingIncludeFields().GetDestinationPort()),
		bool(model.GetEcmpHashingIncludeFields().GetDestinationIP()),
		string(model.GetDisplayName()),
		int(model.GetConfigurationVersion()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()),
		string(model.GetAddressAllocationMode()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	for _, ref := range model.BGPVPNRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("bgpvpn"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "BGPVPNRefs create failed")
		}
	}

	for _, ref := range model.NetworkIpamRefs {

		if ref.Attr == nil {
			ref.Attr = &models.VnSubnetsType{}
		}

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("network_ipam"), model.UUID, ref.UUID, common.MustJSON(ref.Attr.GetIpamSubnets()),
			common.MustJSON(ref.Attr.GetHostRoutes().GetRoute()))
		if err != nil {
			return errors.Wrap(err, "NetworkIpamRefs create failed")
		}
	}

	for _, ref := range model.SecurityLoggingObjectRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("security_logging_object"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "SecurityLoggingObjectRefs create failed")
		}
	}

	for _, ref := range model.NetworkPolicyRefs {

		if ref.Attr == nil {
			ref.Attr = &models.VirtualNetworkPolicyType{}
		}

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("network_policy"), model.UUID, ref.UUID, string(ref.Attr.GetTimer().GetEndTime()),
			string(ref.Attr.GetTimer().GetStartTime()),
			string(ref.Attr.GetTimer().GetOffInterval()),
			string(ref.Attr.GetTimer().GetOnInterval()),
			int(ref.Attr.GetSequence().GetMajor()),
			int(ref.Attr.GetSequence().GetMinor()))
		if err != nil {
			return errors.Wrap(err, "NetworkPolicyRefs create failed")
		}
	}

	for _, ref := range model.QosConfigRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("qos_config"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "QosConfigRefs create failed")
		}
	}

	for _, ref := range model.RouteTableRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("route_table"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "RouteTableRefs create failed")
		}
	}

	for _, ref := range model.VirtualNetworkRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("virtual_network"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "virtual_network",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "virtual_network", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanVirtualNetwork(values map[string]interface{}) (*models.VirtualNetwork, error) {
	m := models.MakeVirtualNetwork()

	if value, ok := values["vxlan_network_identifier"]; ok {

		m.VirtualNetworkProperties.VxlanNetworkIdentifier = common.InterfaceToInt64(value)

	}

	if value, ok := values["rpf"]; ok {

		m.VirtualNetworkProperties.RPF = common.InterfaceToString(value)

	}

	if value, ok := values["network_id"]; ok {

		m.VirtualNetworkProperties.NetworkID = common.InterfaceToInt64(value)

	}

	if value, ok := values["mirror_destination"]; ok {

		m.VirtualNetworkProperties.MirrorDestination = common.InterfaceToBool(value)

	}

	if value, ok := values["forwarding_mode"]; ok {

		m.VirtualNetworkProperties.ForwardingMode = common.InterfaceToString(value)

	}

	if value, ok := values["allow_transit"]; ok {

		m.VirtualNetworkProperties.AllowTransit = common.InterfaceToBool(value)

	}

	if value, ok := values["virtual_network_network_id"]; ok {

		m.VirtualNetworkNetworkID = common.InterfaceToInt64(value)

	}

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["router_external"]; ok {

		m.RouterExternal = common.InterfaceToBool(value)

	}

	if value, ok := values["route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.RouteTargetList.RouteTarget)

	}

	if value, ok := values["segmentation_id"]; ok {

		m.ProviderProperties.SegmentationID = common.InterfaceToInt64(value)

	}

	if value, ok := values["physical_network"]; ok {

		m.ProviderProperties.PhysicalNetwork = common.InterfaceToString(value)

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

	if value, ok := values["pbb_evpn_enable"]; ok {

		m.PBBEvpnEnable = common.InterfaceToBool(value)

	}

	if value, ok := values["pbb_etree_enable"]; ok {

		m.PBBEtreeEnable = common.InterfaceToBool(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = common.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = common.InterfaceToString(value)

	}

	if value, ok := values["multi_policy_service_chains_enabled"]; ok {

		m.MultiPolicyServiceChainsEnabled = common.InterfaceToBool(value)

	}

	if value, ok := values["mac_move_time_window"]; ok {

		m.MacMoveControl.MacMoveTimeWindow = common.InterfaceToInt64(value)

	}

	if value, ok := values["mac_move_limit_action"]; ok {

		m.MacMoveControl.MacMoveLimitAction = common.InterfaceToString(value)

	}

	if value, ok := values["mac_move_limit"]; ok {

		m.MacMoveControl.MacMoveLimit = common.InterfaceToInt64(value)

	}

	if value, ok := values["mac_limit_action"]; ok {

		m.MacLimitControl.MacLimitAction = common.InterfaceToString(value)

	}

	if value, ok := values["mac_limit"]; ok {

		m.MacLimitControl.MacLimit = common.InterfaceToInt64(value)

	}

	if value, ok := values["mac_learning_enabled"]; ok {

		m.MacLearningEnabled = common.InterfaceToBool(value)

	}

	if value, ok := values["mac_aging_time"]; ok {

		m.MacAgingTime = common.InterfaceToInt64(value)

	}

	if value, ok := values["layer2_control_word"]; ok {

		m.Layer2ControlWord = common.InterfaceToBool(value)

	}

	if value, ok := values["is_shared"]; ok {

		m.IsShared = common.InterfaceToBool(value)

	}

	if value, ok := values["import_route_target_list_route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.ImportRouteTargetList.RouteTarget)

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

	if value, ok := values["flood_unknown_unicast"]; ok {

		m.FloodUnknownUnicast = common.InterfaceToBool(value)

	}

	if value, ok := values["external_ipam"]; ok {

		m.ExternalIpam = common.InterfaceToBool(value)

	}

	if value, ok := values["export_route_target_list_route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.ExportRouteTargetList.RouteTarget)

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

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["address_allocation_mode"]; ok {

		m.AddressAllocationMode = common.InterfaceToString(value)

	}

	if value, ok := values["ref_network_ipam"]; ok {
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
			referenceModel := &models.VirtualNetworkNetworkIpamRef{}
			referenceModel.UUID = uuid
			m.NetworkIpamRefs = append(m.NetworkIpamRefs, referenceModel)

			attr := models.MakeVnSubnetsType()
			referenceModel.Attr = attr

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
			referenceModel := &models.VirtualNetworkSecurityLoggingObjectRef{}
			referenceModel.UUID = uuid
			m.SecurityLoggingObjectRefs = append(m.SecurityLoggingObjectRefs, referenceModel)

		}
	}

	if value, ok := values["ref_network_policy"]; ok {
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
			referenceModel := &models.VirtualNetworkNetworkPolicyRef{}
			referenceModel.UUID = uuid
			m.NetworkPolicyRefs = append(m.NetworkPolicyRefs, referenceModel)

			attr := models.MakeVirtualNetworkPolicyType()
			referenceModel.Attr = attr

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
			referenceModel := &models.VirtualNetworkQosConfigRef{}
			referenceModel.UUID = uuid
			m.QosConfigRefs = append(m.QosConfigRefs, referenceModel)

		}
	}

	if value, ok := values["ref_route_table"]; ok {
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
			referenceModel := &models.VirtualNetworkRouteTableRef{}
			referenceModel.UUID = uuid
			m.RouteTableRefs = append(m.RouteTableRefs, referenceModel)

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
			referenceModel := &models.VirtualNetworkVirtualNetworkRef{}
			referenceModel.UUID = uuid
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
	}

	if value, ok := values["ref_bgpvpn"]; ok {
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
			referenceModel := &models.VirtualNetworkBGPVPNRef{}
			referenceModel.UUID = uuid
			m.BGPVPNRefs = append(m.BGPVPNRefs, referenceModel)

		}
	}

	if value, ok := values["backref_access_control_list"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeAccessControlList()
			m.AccessControlLists = append(m.AccessControlLists, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["configuration_version"]; ok && propertyValue != nil {

				childModel.ConfigurationVersion = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

			if propertyValue, ok := childResourceMap["access_control_list_hash"]; ok && propertyValue != nil {

				childModel.AccessControlListHash = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["dynamic"]; ok && propertyValue != nil {

				childModel.AccessControlListEntries.Dynamic = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["acl_rule"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.AccessControlListEntries.ACLRule)

			}

		}
	}

	if value, ok := values["backref_alias_ip_pool"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeAliasIPPool()
			m.AliasIPPools = append(m.AliasIPPools, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["configuration_version"]; ok && propertyValue != nil {

				childModel.ConfigurationVersion = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	if value, ok := values["backref_bridge_domain"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeBridgeDomain()
			m.BridgeDomains = append(m.BridgeDomains, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["mac_move_time_window"]; ok && propertyValue != nil {

				childModel.MacMoveControl.MacMoveTimeWindow = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["mac_move_limit_action"]; ok && propertyValue != nil {

				childModel.MacMoveControl.MacMoveLimitAction = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["mac_move_limit"]; ok && propertyValue != nil {

				childModel.MacMoveControl.MacMoveLimit = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["mac_limit_action"]; ok && propertyValue != nil {

				childModel.MacLimitControl.MacLimitAction = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["mac_limit"]; ok && propertyValue != nil {

				childModel.MacLimitControl.MacLimit = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["mac_learning_enabled"]; ok && propertyValue != nil {

				childModel.MacLearningEnabled = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["mac_aging_time"]; ok && propertyValue != nil {

				childModel.MacAgingTime = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["isid"]; ok && propertyValue != nil {

				childModel.Isid = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["configuration_version"]; ok && propertyValue != nil {

				childModel.ConfigurationVersion = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	if value, ok := values["backref_floating_ip_pool"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeFloatingIPPool()
			m.FloatingIPPools = append(m.FloatingIPPools, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["subnet_uuid"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FloatingIPPoolSubnets.SubnetUUID)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["configuration_version"]; ok && propertyValue != nil {

				childModel.ConfigurationVersion = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	if value, ok := values["backref_routing_instance"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeRoutingInstance()
			m.RoutingInstances = append(m.RoutingInstances, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["configuration_version"]; ok && propertyValue != nil {

				childModel.ConfigurationVersion = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListVirtualNetwork lists VirtualNetwork with list spec.
func (db *DB) listVirtualNetwork(ctx context.Context, request *models.ListVirtualNetworkRequest) (response *models.ListVirtualNetworkResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["virtual_network"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.VirtualNetwork{}

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
		m, err := scanVirtualNetwork(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListVirtualNetworkResponse{
		VirtualNetworks: result,
	}
	return response, nil
}

// UpdateVirtualNetwork updates a resource
func (db *DB) updateVirtualNetwork(
	ctx context.Context,
	request *models.UpdateVirtualNetworkRequest,
) error {
	//TODO
	return nil
}

// DeleteVirtualNetwork deletes a resource
func (db *DB) deleteVirtualNetwork(
	ctx context.Context,
	request *models.DeleteVirtualNetworkRequest) error {
	qb := db.queryBuilders["virtual_network"]

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

//CreateVirtualNetwork handle a Create API
// nolint
func (db *DB) CreateVirtualNetwork(
	ctx context.Context,
	request *models.CreateVirtualNetworkRequest) (*models.CreateVirtualNetworkResponse, error) {
	model := request.VirtualNetwork
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createVirtualNetwork(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_network",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateVirtualNetworkResponse{
		VirtualNetwork: request.VirtualNetwork,
	}, nil
}

//UpdateVirtualNetwork handles a Update request.
func (db *DB) UpdateVirtualNetwork(
	ctx context.Context,
	request *models.UpdateVirtualNetworkRequest) (*models.UpdateVirtualNetworkResponse, error) {
	model := request.VirtualNetwork
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateVirtualNetwork(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_network",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateVirtualNetworkResponse{
		VirtualNetwork: model,
	}, nil
}

//DeleteVirtualNetwork delete a resource.
func (db *DB) DeleteVirtualNetwork(ctx context.Context, request *models.DeleteVirtualNetworkRequest) (*models.DeleteVirtualNetworkResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteVirtualNetwork(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteVirtualNetworkResponse{
		ID: request.ID,
	}, nil
}

//GetVirtualNetwork a Get request.
func (db *DB) GetVirtualNetwork(ctx context.Context, request *models.GetVirtualNetworkRequest) (response *models.GetVirtualNetworkResponse, err error) {
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
	listRequest := &models.ListVirtualNetworkRequest{
		Spec: spec,
	}
	var result *models.ListVirtualNetworkResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listVirtualNetwork(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.VirtualNetworks) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetVirtualNetworkResponse{
		VirtualNetwork: result.VirtualNetworks[0],
	}
	return response, nil
}

//ListVirtualNetwork handles a List service Request.
// nolint
func (db *DB) ListVirtualNetwork(
	ctx context.Context,
	request *models.ListVirtualNetworkRequest) (response *models.ListVirtualNetworkResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listVirtualNetwork(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
