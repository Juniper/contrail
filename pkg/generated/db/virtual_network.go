package db

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertVirtualNetworkQuery = "insert into `virtual_network` (`vxlan_network_identifier`,`rpf`,`network_id`,`mirror_destination`,`forwarding_mode`,`allow_transit`,`virtual_network_network_id`,`uuid`,`router_external`,`route_target`,`segmentation_id`,`physical_network`,`port_security_enabled`,`share`,`owner_access`,`owner`,`global_access`,`pbb_evpn_enable`,`pbb_etree_enable`,`parent_uuid`,`parent_type`,`multi_policy_service_chains_enabled`,`mac_move_time_window`,`mac_move_limit_action`,`mac_move_limit`,`mac_limit_action`,`mac_limit`,`mac_learning_enabled`,`mac_aging_time`,`layer2_control_word`,`is_shared`,`import_route_target_list_route_target`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`flood_unknown_unicast`,`external_ipam`,`export_route_target_list_route_target`,`source_port`,`source_ip`,`ip_protocol`,`hashing_configured`,`destination_port`,`destination_ip`,`display_name`,`key_value_pair`,`address_allocation_mode`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteVirtualNetworkQuery = "delete from `virtual_network` where uuid = ?"

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
	"key_value_pair",
	"address_allocation_mode",
}

// VirtualNetworkRefFields is db reference fields for VirtualNetwork
var VirtualNetworkRefFields = map[string][]string{

	"virtual_network": {
	// <common.Schema Value>

	},

	"bgpvpn": {
	// <common.Schema Value>

	},

	"network_ipam": {
		// <common.Schema Value>
		"ipam_subnets",
		"route",
	},

	"security_logging_object": {
	// <common.Schema Value>

	},

	"network_policy": {
		// <common.Schema Value>
		"end_time",
		"start_time",
		"off_interval",
		"on_interval",
		"major",
		"minor",
	},

	"qos_config": {
	// <common.Schema Value>

	},

	"route_table": {
	// <common.Schema Value>

	},
}

// VirtualNetworkBackRefFields is db back reference fields for VirtualNetwork
var VirtualNetworkBackRefFields = map[string][]string{

	"access_control_list": {
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
		"key_value_pair",
		"access_control_list_hash",
		"dynamic",
		"acl_rule",
	},

	"alias_ip_pool": {
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
		"key_value_pair",
	},

	"bridge_domain": {
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
		"key_value_pair",
	},

	"floating_ip_pool": {
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
		"key_value_pair",
	},

	"routing_instance": {
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
		"key_value_pair",
	},
}

// VirtualNetworkParentTypes is possible parents for VirtualNetwork
var VirtualNetworkParents = []string{

	"project",
}

const insertVirtualNetworkNetworkPolicyQuery = "insert into `ref_virtual_network_network_policy` (`from`, `to` ,`end_time`,`start_time`,`off_interval`,`on_interval`,`major`,`minor`) values (?, ?,?,?,?,?,?,?);"

const insertVirtualNetworkQosConfigQuery = "insert into `ref_virtual_network_qos_config` (`from`, `to` ) values (?, ?);"

const insertVirtualNetworkRouteTableQuery = "insert into `ref_virtual_network_route_table` (`from`, `to` ) values (?, ?);"

const insertVirtualNetworkVirtualNetworkQuery = "insert into `ref_virtual_network_virtual_network` (`from`, `to` ) values (?, ?);"

const insertVirtualNetworkBGPVPNQuery = "insert into `ref_virtual_network_bgpvpn` (`from`, `to` ) values (?, ?);"

const insertVirtualNetworkNetworkIpamQuery = "insert into `ref_virtual_network_network_ipam` (`from`, `to` ,`ipam_subnets`,`route`) values (?, ?,?,?);"

const insertVirtualNetworkSecurityLoggingObjectQuery = "insert into `ref_virtual_network_security_logging_object` (`from`, `to` ) values (?, ?);"

// CreateVirtualNetwork inserts VirtualNetwork to DB
func CreateVirtualNetwork(tx *sql.Tx, model *models.VirtualNetwork) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualNetworkQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertVirtualNetworkQuery,
	}).Debug("create query")
	_, err = stmt.Exec(int(model.VirtualNetworkProperties.VxlanNetworkIdentifier),
		string(model.VirtualNetworkProperties.RPF),
		int(model.VirtualNetworkProperties.NetworkID),
		bool(model.VirtualNetworkProperties.MirrorDestination),
		string(model.VirtualNetworkProperties.ForwardingMode),
		bool(model.VirtualNetworkProperties.AllowTransit),
		int(model.VirtualNetworkNetworkID),
		string(model.UUID),
		bool(model.RouterExternal),
		common.MustJSON(model.RouteTargetList.RouteTarget),
		int(model.ProviderProperties.SegmentationID),
		string(model.ProviderProperties.PhysicalNetwork),
		bool(model.PortSecurityEnabled),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		bool(model.PBBEvpnEnable),
		bool(model.PBBEtreeEnable),
		string(model.ParentUUID),
		string(model.ParentType),
		bool(model.MultiPolicyServiceChainsEnabled),
		int(model.MacMoveControl.MacMoveTimeWindow),
		string(model.MacMoveControl.MacMoveLimitAction),
		int(model.MacMoveControl.MacMoveLimit),
		string(model.MacLimitControl.MacLimitAction),
		int(model.MacLimitControl.MacLimit),
		bool(model.MacLearningEnabled),
		int(model.MacAgingTime),
		bool(model.Layer2ControlWord),
		bool(model.IsShared),
		common.MustJSON(model.ImportRouteTargetList.RouteTarget),
		bool(model.IDPerms.UserVisible),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OtherAccess),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Group),
		string(model.IDPerms.LastModified),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Creator),
		string(model.IDPerms.Created),
		common.MustJSON(model.FQName),
		bool(model.FloodUnknownUnicast),
		bool(model.ExternalIpam),
		common.MustJSON(model.ExportRouteTargetList.RouteTarget),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair),
		string(model.AddressAllocationMode))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtBGPVPNRef, err := tx.Prepare(insertVirtualNetworkBGPVPNQuery)
	if err != nil {
		return errors.Wrap(err, "preparing BGPVPNRefs create statement failed")
	}
	defer stmtBGPVPNRef.Close()
	for _, ref := range model.BGPVPNRefs {

		_, err = stmtBGPVPNRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "BGPVPNRefs create failed")
		}
	}

	stmtNetworkIpamRef, err := tx.Prepare(insertVirtualNetworkNetworkIpamQuery)
	if err != nil {
		return errors.Wrap(err, "preparing NetworkIpamRefs create statement failed")
	}
	defer stmtNetworkIpamRef.Close()
	for _, ref := range model.NetworkIpamRefs {

		if ref.Attr == nil {
			ref.Attr = models.MakeVnSubnetsType()
		}

		_, err = stmtNetworkIpamRef.Exec(model.UUID, ref.UUID, common.MustJSON(ref.Attr.IpamSubnets),
			common.MustJSON(ref.Attr.HostRoutes.Route))
		if err != nil {
			return errors.Wrap(err, "NetworkIpamRefs create failed")
		}
	}

	stmtSecurityLoggingObjectRef, err := tx.Prepare(insertVirtualNetworkSecurityLoggingObjectQuery)
	if err != nil {
		return errors.Wrap(err, "preparing SecurityLoggingObjectRefs create statement failed")
	}
	defer stmtSecurityLoggingObjectRef.Close()
	for _, ref := range model.SecurityLoggingObjectRefs {

		_, err = stmtSecurityLoggingObjectRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "SecurityLoggingObjectRefs create failed")
		}
	}

	stmtNetworkPolicyRef, err := tx.Prepare(insertVirtualNetworkNetworkPolicyQuery)
	if err != nil {
		return errors.Wrap(err, "preparing NetworkPolicyRefs create statement failed")
	}
	defer stmtNetworkPolicyRef.Close()
	for _, ref := range model.NetworkPolicyRefs {

		if ref.Attr == nil {
			ref.Attr = models.MakeVirtualNetworkPolicyType()
		}

		_, err = stmtNetworkPolicyRef.Exec(model.UUID, ref.UUID, string(ref.Attr.Timer.EndTime),
			string(ref.Attr.Timer.StartTime),
			string(ref.Attr.Timer.OffInterval),
			string(ref.Attr.Timer.OnInterval),
			int(ref.Attr.Sequence.Major),
			int(ref.Attr.Sequence.Minor))
		if err != nil {
			return errors.Wrap(err, "NetworkPolicyRefs create failed")
		}
	}

	stmtQosConfigRef, err := tx.Prepare(insertVirtualNetworkQosConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing QosConfigRefs create statement failed")
	}
	defer stmtQosConfigRef.Close()
	for _, ref := range model.QosConfigRefs {

		_, err = stmtQosConfigRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "QosConfigRefs create failed")
		}
	}

	stmtRouteTableRef, err := tx.Prepare(insertVirtualNetworkRouteTableQuery)
	if err != nil {
		return errors.Wrap(err, "preparing RouteTableRefs create statement failed")
	}
	defer stmtRouteTableRef.Close()
	for _, ref := range model.RouteTableRefs {

		_, err = stmtRouteTableRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "RouteTableRefs create failed")
		}
	}

	stmtVirtualNetworkRef, err := tx.Prepare(insertVirtualNetworkVirtualNetworkQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualNetworkRefs create statement failed")
	}
	defer stmtVirtualNetworkRef.Close()
	for _, ref := range model.VirtualNetworkRefs {

		_, err = stmtVirtualNetworkRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "virtual_network",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "virtual_network", model.UUID, model.Perms2.Share)
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

		castedValue := common.InterfaceToInt(value)

		m.VirtualNetworkProperties.VxlanNetworkIdentifier = models.VxlanNetworkIdentifierType(castedValue)

	}

	if value, ok := values["rpf"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualNetworkProperties.RPF = models.RpfModeType(castedValue)

	}

	if value, ok := values["network_id"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualNetworkProperties.NetworkID = castedValue

	}

	if value, ok := values["mirror_destination"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.VirtualNetworkProperties.MirrorDestination = castedValue

	}

	if value, ok := values["forwarding_mode"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualNetworkProperties.ForwardingMode = models.ForwardingModeType(castedValue)

	}

	if value, ok := values["allow_transit"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.VirtualNetworkProperties.AllowTransit = castedValue

	}

	if value, ok := values["virtual_network_network_id"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualNetworkNetworkID = models.VirtualNetworkIdType(castedValue)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["router_external"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.RouterExternal = castedValue

	}

	if value, ok := values["route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.RouteTargetList.RouteTarget)

	}

	if value, ok := values["segmentation_id"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ProviderProperties.SegmentationID = models.VlanIdType(castedValue)

	}

	if value, ok := values["physical_network"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProviderProperties.PhysicalNetwork = castedValue

	}

	if value, ok := values["port_security_enabled"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.PortSecurityEnabled = castedValue

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["pbb_evpn_enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.PBBEvpnEnable = castedValue

	}

	if value, ok := values["pbb_etree_enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.PBBEtreeEnable = castedValue

	}

	if value, ok := values["parent_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentUUID = castedValue

	}

	if value, ok := values["parent_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentType = castedValue

	}

	if value, ok := values["multi_policy_service_chains_enabled"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.MultiPolicyServiceChainsEnabled = castedValue

	}

	if value, ok := values["mac_move_time_window"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacMoveControl.MacMoveTimeWindow = models.MACMoveTimeWindow(castedValue)

	}

	if value, ok := values["mac_move_limit_action"]; ok {

		castedValue := common.InterfaceToString(value)

		m.MacMoveControl.MacMoveLimitAction = models.MACLimitExceedActionType(castedValue)

	}

	if value, ok := values["mac_move_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacMoveControl.MacMoveLimit = castedValue

	}

	if value, ok := values["mac_limit_action"]; ok {

		castedValue := common.InterfaceToString(value)

		m.MacLimitControl.MacLimitAction = models.MACLimitExceedActionType(castedValue)

	}

	if value, ok := values["mac_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacLimitControl.MacLimit = castedValue

	}

	if value, ok := values["mac_learning_enabled"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.MacLearningEnabled = castedValue

	}

	if value, ok := values["mac_aging_time"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacAgingTime = models.MACAgingTime(castedValue)

	}

	if value, ok := values["layer2_control_word"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.Layer2ControlWord = castedValue

	}

	if value, ok := values["is_shared"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IsShared = castedValue

	}

	if value, ok := values["import_route_target_list_route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.ImportRouteTargetList.RouteTarget)

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["flood_unknown_unicast"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.FloodUnknownUnicast = castedValue

	}

	if value, ok := values["external_ipam"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ExternalIpam = castedValue

	}

	if value, ok := values["export_route_target_list_route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.ExportRouteTargetList.RouteTarget)

	}

	if value, ok := values["source_port"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.SourcePort = castedValue

	}

	if value, ok := values["source_ip"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.SourceIP = castedValue

	}

	if value, ok := values["ip_protocol"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.IPProtocol = castedValue

	}

	if value, ok := values["hashing_configured"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.HashingConfigured = castedValue

	}

	if value, ok := values["destination_port"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.DestinationPort = castedValue

	}

	if value, ok := values["destination_ip"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.DestinationIP = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["address_allocation_mode"]; ok {

		castedValue := common.InterfaceToString(value)

		m.AddressAllocationMode = models.AddressAllocationModeType(castedValue)

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

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.Perms2.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.GlobalAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentUUID = castedValue

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentType = castedValue

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.UserVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Group = castedValue

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.LastModified = castedValue

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.Enable = castedValue

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Description = castedValue

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Creator = castedValue

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Created = castedValue

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

			if propertyValue, ok := childResourceMap["access_control_list_hash"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.AccessControlListHash)

			}

			if propertyValue, ok := childResourceMap["dynamic"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.AccessControlListEntries.Dynamic = castedValue

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

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.Perms2.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.GlobalAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentUUID = castedValue

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentType = castedValue

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.UserVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Group = castedValue

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.LastModified = castedValue

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.Enable = castedValue

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Description = castedValue

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Creator = castedValue

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Created = castedValue

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

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

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.Perms2.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.GlobalAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentUUID = castedValue

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentType = castedValue

			}

			if propertyValue, ok := childResourceMap["mac_move_time_window"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.MacMoveControl.MacMoveTimeWindow = models.MACMoveTimeWindow(castedValue)

			}

			if propertyValue, ok := childResourceMap["mac_move_limit_action"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.MacMoveControl.MacMoveLimitAction = models.MACLimitExceedActionType(castedValue)

			}

			if propertyValue, ok := childResourceMap["mac_move_limit"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.MacMoveControl.MacMoveLimit = castedValue

			}

			if propertyValue, ok := childResourceMap["mac_limit_action"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.MacLimitControl.MacLimitAction = models.MACLimitExceedActionType(castedValue)

			}

			if propertyValue, ok := childResourceMap["mac_limit"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.MacLimitControl.MacLimit = castedValue

			}

			if propertyValue, ok := childResourceMap["mac_learning_enabled"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.MacLearningEnabled = castedValue

			}

			if propertyValue, ok := childResourceMap["mac_aging_time"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.MacAgingTime = models.MACAgingTime(castedValue)

			}

			if propertyValue, ok := childResourceMap["isid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Isid = models.IsidType(castedValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.UserVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Group = castedValue

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.LastModified = castedValue

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.Enable = castedValue

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Description = castedValue

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Creator = castedValue

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Created = castedValue

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

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

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.Perms2.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.GlobalAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentUUID = castedValue

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentType = castedValue

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.UserVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Group = castedValue

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.LastModified = castedValue

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.Enable = castedValue

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Description = castedValue

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Creator = castedValue

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Created = castedValue

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["subnet_uuid"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FloatingIPPoolSubnets.SubnetUUID)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

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

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.Perms2.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.GlobalAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentUUID = castedValue

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentType = castedValue

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.UserVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Group = castedValue

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.LastModified = castedValue

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.Enable = castedValue

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Description = castedValue

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Creator = castedValue

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Created = castedValue

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListVirtualNetwork lists VirtualNetwork with list spec.
func ListVirtualNetwork(tx *sql.Tx, spec *common.ListSpec) ([]*models.VirtualNetwork, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "virtual_network"
	spec.Fields = VirtualNetworkFields
	spec.RefFields = VirtualNetworkRefFields
	spec.BackRefFields = VirtualNetworkBackRefFields
	result := models.MakeVirtualNetworkSlice()

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filter.AppendValues("parent_uuid", []string{parentMetaData.UUID})
	}

	query := spec.BuildQuery()
	columns := spec.Columns
	values := spec.Values
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.Query(query, values...)
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
	return result, nil
}

// UpdateVirtualNetwork updates a resource
func UpdateVirtualNetwork(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	//TODO (handle references)
	// Prepare statement for updating data
	var updateVirtualNetworkQuery = "update `virtual_network` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".VirtualNetworkProperties.VxlanNetworkIdentifier", "."); ok {
		updateVirtualNetworkQuery += "`vxlan_network_identifier` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualNetworkProperties.RPF", "."); ok {
		updateVirtualNetworkQuery += "`rpf` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualNetworkProperties.NetworkID", "."); ok {
		updateVirtualNetworkQuery += "`network_id` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualNetworkProperties.MirrorDestination", "."); ok {
		updateVirtualNetworkQuery += "`mirror_destination` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualNetworkProperties.ForwardingMode", "."); ok {
		updateVirtualNetworkQuery += "`forwarding_mode` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualNetworkProperties.AllowTransit", "."); ok {
		updateVirtualNetworkQuery += "`allow_transit` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".VirtualNetworkNetworkID", "."); ok {
		updateVirtualNetworkQuery += "`virtual_network_network_id` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateVirtualNetworkQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".RouterExternal", "."); ok {
		updateVirtualNetworkQuery += "`router_external` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".RouteTargetList.RouteTarget", "."); ok {
		updateVirtualNetworkQuery += "`route_target` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ProviderProperties.SegmentationID", "."); ok {
		updateVirtualNetworkQuery += "`segmentation_id` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ProviderProperties.PhysicalNetwork", "."); ok {
		updateVirtualNetworkQuery += "`physical_network` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PortSecurityEnabled", "."); ok {
		updateVirtualNetworkQuery += "`port_security_enabled` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateVirtualNetworkQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateVirtualNetworkQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateVirtualNetworkQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateVirtualNetworkQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PBBEvpnEnable", "."); ok {
		updateVirtualNetworkQuery += "`pbb_evpn_enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PBBEtreeEnable", "."); ok {
		updateVirtualNetworkQuery += "`pbb_etree_enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateVirtualNetworkQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateVirtualNetworkQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MultiPolicyServiceChainsEnabled", "."); ok {
		updateVirtualNetworkQuery += "`multi_policy_service_chains_enabled` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacMoveControl.MacMoveTimeWindow", "."); ok {
		updateVirtualNetworkQuery += "`mac_move_time_window` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacMoveControl.MacMoveLimitAction", "."); ok {
		updateVirtualNetworkQuery += "`mac_move_limit_action` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacMoveControl.MacMoveLimit", "."); ok {
		updateVirtualNetworkQuery += "`mac_move_limit` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacLimitControl.MacLimitAction", "."); ok {
		updateVirtualNetworkQuery += "`mac_limit_action` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacLimitControl.MacLimit", "."); ok {
		updateVirtualNetworkQuery += "`mac_limit` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacLearningEnabled", "."); ok {
		updateVirtualNetworkQuery += "`mac_learning_enabled` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacAgingTime", "."); ok {
		updateVirtualNetworkQuery += "`mac_aging_time` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Layer2ControlWord", "."); ok {
		updateVirtualNetworkQuery += "`layer2_control_word` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IsShared", "."); ok {
		updateVirtualNetworkQuery += "`is_shared` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ImportRouteTargetList.RouteTarget", "."); ok {
		updateVirtualNetworkQuery += "`import_route_target_list_route_target` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateVirtualNetworkQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateVirtualNetworkQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateVirtualNetworkQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateVirtualNetworkQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateVirtualNetworkQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateVirtualNetworkQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateVirtualNetworkQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateVirtualNetworkQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateVirtualNetworkQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateVirtualNetworkQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateVirtualNetworkQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateVirtualNetworkQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FloodUnknownUnicast", "."); ok {
		updateVirtualNetworkQuery += "`flood_unknown_unicast` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ExternalIpam", "."); ok {
		updateVirtualNetworkQuery += "`external_ipam` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ExportRouteTargetList.RouteTarget", "."); ok {
		updateVirtualNetworkQuery += "`export_route_target_list_route_target` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".EcmpHashingIncludeFields.SourcePort", "."); ok {
		updateVirtualNetworkQuery += "`source_port` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".EcmpHashingIncludeFields.SourceIP", "."); ok {
		updateVirtualNetworkQuery += "`source_ip` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".EcmpHashingIncludeFields.IPProtocol", "."); ok {
		updateVirtualNetworkQuery += "`ip_protocol` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".EcmpHashingIncludeFields.HashingConfigured", "."); ok {
		updateVirtualNetworkQuery += "`hashing_configured` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".EcmpHashingIncludeFields.DestinationPort", "."); ok {
		updateVirtualNetworkQuery += "`destination_port` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".EcmpHashingIncludeFields.DestinationIP", "."); ok {
		updateVirtualNetworkQuery += "`destination_ip` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateVirtualNetworkQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateVirtualNetworkQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualNetworkQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".AddressAllocationMode", "."); ok {
		updateVirtualNetworkQuery += "`address_allocation_mode` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualNetworkQuery += ","
	}

	updateVirtualNetworkQuery =
		updateVirtualNetworkQuery[:len(updateVirtualNetworkQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateVirtualNetworkQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateVirtualNetworkQuery,
	}).Debug("update query")
	_, err = stmt.Exec(updatedValues...)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("updated")
	return err
}

// DeleteVirtualNetwork deletes a resource
func DeleteVirtualNetwork(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteVirtualNetworkQuery
	selectQuery := "select count(uuid) from virtual_network where uuid = ?"
	var err error
	var count int

	if auth.IsAdmin() {
		row := tx.QueryRow(selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.Exec(deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRow(selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.Exec(deleteQuery, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	err = common.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}
