package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertVirtualNetworkQuery = "insert into `virtual_network` (`virtual_network_network_id`,`flood_unknown_unicast`,`external_ipam`,`fq_name`,`route_target`,`route_target_list_route_target`,`export_route_target_list_route_target`,`pbb_evpn_enable`,`router_external`,`mac_aging_time`,`uuid`,`display_name`,`network_id`,`mirror_destination`,`vxlan_network_identifier`,`rpf`,`forwarding_mode`,`allow_transit`,`segmentation_id`,`physical_network`,`mac_learning_enabled`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`mac_move_limit`,`mac_move_limit_action`,`mac_move_time_window`,`mac_limit`,`mac_limit_action`,`is_shared`,`destination_ip`,`ip_protocol`,`source_ip`,`hashing_configured`,`source_port`,`destination_port`,`address_allocation_mode`,`layer2_control_word`,`multi_policy_service_chains_enabled`,`pbb_etree_enable`,`port_security_enabled`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualNetworkQuery = "update `virtual_network` set `virtual_network_network_id` = ?,`flood_unknown_unicast` = ?,`external_ipam` = ?,`fq_name` = ?,`route_target` = ?,`route_target_list_route_target` = ?,`export_route_target_list_route_target` = ?,`pbb_evpn_enable` = ?,`router_external` = ?,`mac_aging_time` = ?,`uuid` = ?,`display_name` = ?,`network_id` = ?,`mirror_destination` = ?,`vxlan_network_identifier` = ?,`rpf` = ?,`forwarding_mode` = ?,`allow_transit` = ?,`segmentation_id` = ?,`physical_network` = ?,`mac_learning_enabled` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`mac_move_limit` = ?,`mac_move_limit_action` = ?,`mac_move_time_window` = ?,`mac_limit` = ?,`mac_limit_action` = ?,`is_shared` = ?,`destination_ip` = ?,`ip_protocol` = ?,`source_ip` = ?,`hashing_configured` = ?,`source_port` = ?,`destination_port` = ?,`address_allocation_mode` = ?,`layer2_control_word` = ?,`multi_policy_service_chains_enabled` = ?,`pbb_etree_enable` = ?,`port_security_enabled` = ?;"
const deleteVirtualNetworkQuery = "delete from `virtual_network` where uuid = ?"

// VirtualNetworkFields is db columns for VirtualNetwork
var VirtualNetworkFields = []string{
	"virtual_network_network_id",
	"flood_unknown_unicast",
	"external_ipam",
	"fq_name",
	"route_target",
	"route_target_list_route_target",
	"export_route_target_list_route_target",
	"pbb_evpn_enable",
	"router_external",
	"mac_aging_time",
	"uuid",
	"display_name",
	"network_id",
	"mirror_destination",
	"vxlan_network_identifier",
	"rpf",
	"forwarding_mode",
	"allow_transit",
	"segmentation_id",
	"physical_network",
	"mac_learning_enabled",
	"creator",
	"user_visible",
	"last_modified",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"enable",
	"description",
	"created",
	"key_value_pair",
	"perms2_owner_access",
	"global_access",
	"share",
	"perms2_owner",
	"mac_move_limit",
	"mac_move_limit_action",
	"mac_move_time_window",
	"mac_limit",
	"mac_limit_action",
	"is_shared",
	"destination_ip",
	"ip_protocol",
	"source_ip",
	"hashing_configured",
	"source_port",
	"destination_port",
	"address_allocation_mode",
	"layer2_control_word",
	"multi_policy_service_chains_enabled",
	"pbb_etree_enable",
	"port_security_enabled",
}

// VirtualNetworkRefFields is db reference fields for VirtualNetwork
var VirtualNetworkRefFields = map[string][]string{

	"network_policy": {
		// <common.Schema Value>
		"on_interval",
		"end_time",
		"start_time",
		"off_interval",
		"major",
		"minor",
	},

	"qos_config": {
	// <common.Schema Value>

	},

	"route_table": {
	// <common.Schema Value>

	},

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
}

const insertVirtualNetworkRouteTableQuery = "insert into `ref_virtual_network_route_table` (`from`, `to` ) values (?, ?);"

const insertVirtualNetworkVirtualNetworkQuery = "insert into `ref_virtual_network_virtual_network` (`from`, `to` ) values (?, ?);"

const insertVirtualNetworkBGPVPNQuery = "insert into `ref_virtual_network_bgpvpn` (`from`, `to` ) values (?, ?);"

const insertVirtualNetworkNetworkIpamQuery = "insert into `ref_virtual_network_network_ipam` (`from`, `to` ,`ipam_subnets`,`route`) values (?, ?,?,?);"

const insertVirtualNetworkSecurityLoggingObjectQuery = "insert into `ref_virtual_network_security_logging_object` (`from`, `to` ) values (?, ?);"

const insertVirtualNetworkNetworkPolicyQuery = "insert into `ref_virtual_network_network_policy` (`from`, `to` ,`on_interval`,`end_time`,`start_time`,`off_interval`,`major`,`minor`) values (?, ?,?,?,?,?,?,?);"

const insertVirtualNetworkQosConfigQuery = "insert into `ref_virtual_network_qos_config` (`from`, `to` ) values (?, ?);"

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
	_, err = stmt.Exec(int(model.VirtualNetworkNetworkID),
		bool(model.FloodUnknownUnicast),
		bool(model.ExternalIpam),
		common.MustJSON(model.FQName),
		common.MustJSON(model.ImportRouteTargetList.RouteTarget),
		common.MustJSON(model.RouteTargetList.RouteTarget),
		common.MustJSON(model.ExportRouteTargetList.RouteTarget),
		bool(model.PBBEvpnEnable),
		bool(model.RouterExternal),
		int(model.MacAgingTime),
		string(model.UUID),
		string(model.DisplayName),
		int(model.VirtualNetworkProperties.NetworkID),
		bool(model.VirtualNetworkProperties.MirrorDestination),
		int(model.VirtualNetworkProperties.VxlanNetworkIdentifier),
		string(model.VirtualNetworkProperties.RPF),
		string(model.VirtualNetworkProperties.ForwardingMode),
		bool(model.VirtualNetworkProperties.AllowTransit),
		int(model.ProviderProperties.SegmentationID),
		string(model.ProviderProperties.PhysicalNetwork),
		bool(model.MacLearningEnabled),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		common.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.MacMoveControl.MacMoveLimit),
		string(model.MacMoveControl.MacMoveLimitAction),
		int(model.MacMoveControl.MacMoveTimeWindow),
		int(model.MacLimitControl.MacLimit),
		string(model.MacLimitControl.MacLimitAction),
		bool(model.IsShared),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		string(model.AddressAllocationMode),
		bool(model.Layer2ControlWord),
		bool(model.MultiPolicyServiceChainsEnabled),
		bool(model.PBBEtreeEnable),
		bool(model.PortSecurityEnabled))
	if err != nil {
		return errors.Wrap(err, "create failed")
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

		_, err = stmtNetworkPolicyRef.Exec(model.UUID, ref.UUID, string(ref.Attr.Timer.OnInterval),
			string(ref.Attr.Timer.EndTime),
			string(ref.Attr.Timer.StartTime),
			string(ref.Attr.Timer.OffInterval),
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

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanVirtualNetwork(values map[string]interface{}) (*models.VirtualNetwork, error) {
	m := models.MakeVirtualNetwork()

	if value, ok := values["virtual_network_network_id"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualNetworkNetworkID = models.VirtualNetworkIdType(castedValue)

	}

	if value, ok := values["flood_unknown_unicast"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.FloodUnknownUnicast = castedValue

	}

	if value, ok := values["external_ipam"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ExternalIpam = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.ImportRouteTargetList.RouteTarget)

	}

	if value, ok := values["route_target_list_route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.RouteTargetList.RouteTarget)

	}

	if value, ok := values["export_route_target_list_route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.ExportRouteTargetList.RouteTarget)

	}

	if value, ok := values["pbb_evpn_enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.PBBEvpnEnable = castedValue

	}

	if value, ok := values["router_external"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.RouterExternal = castedValue

	}

	if value, ok := values["mac_aging_time"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacAgingTime = models.MACAgingTime(castedValue)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["network_id"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualNetworkProperties.NetworkID = castedValue

	}

	if value, ok := values["mirror_destination"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.VirtualNetworkProperties.MirrorDestination = castedValue

	}

	if value, ok := values["vxlan_network_identifier"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualNetworkProperties.VxlanNetworkIdentifier = models.VxlanNetworkIdentifierType(castedValue)

	}

	if value, ok := values["rpf"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualNetworkProperties.RPF = models.RpfModeType(castedValue)

	}

	if value, ok := values["forwarding_mode"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualNetworkProperties.ForwardingMode = models.ForwardingModeType(castedValue)

	}

	if value, ok := values["allow_transit"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.VirtualNetworkProperties.AllowTransit = castedValue

	}

	if value, ok := values["segmentation_id"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ProviderProperties.SegmentationID = models.VlanIdType(castedValue)

	}

	if value, ok := values["physical_network"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProviderProperties.PhysicalNetwork = castedValue

	}

	if value, ok := values["mac_learning_enabled"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.MacLearningEnabled = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["perms2_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["mac_move_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacMoveControl.MacMoveLimit = castedValue

	}

	if value, ok := values["mac_move_limit_action"]; ok {

		castedValue := common.InterfaceToString(value)

		m.MacMoveControl.MacMoveLimitAction = models.MACLimitExceedActionType(castedValue)

	}

	if value, ok := values["mac_move_time_window"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacMoveControl.MacMoveTimeWindow = models.MACMoveTimeWindow(castedValue)

	}

	if value, ok := values["mac_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacLimitControl.MacLimit = castedValue

	}

	if value, ok := values["mac_limit_action"]; ok {

		castedValue := common.InterfaceToString(value)

		m.MacLimitControl.MacLimitAction = models.MACLimitExceedActionType(castedValue)

	}

	if value, ok := values["is_shared"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IsShared = castedValue

	}

	if value, ok := values["destination_ip"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.DestinationIP = castedValue

	}

	if value, ok := values["ip_protocol"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.IPProtocol = castedValue

	}

	if value, ok := values["source_ip"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.SourceIP = castedValue

	}

	if value, ok := values["hashing_configured"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.HashingConfigured = castedValue

	}

	if value, ok := values["source_port"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.SourcePort = castedValue

	}

	if value, ok := values["destination_port"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.DestinationPort = castedValue

	}

	if value, ok := values["address_allocation_mode"]; ok {

		castedValue := common.InterfaceToString(value)

		m.AddressAllocationMode = models.AddressAllocationModeType(castedValue)

	}

	if value, ok := values["layer2_control_word"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.Layer2ControlWord = castedValue

	}

	if value, ok := values["multi_policy_service_chains_enabled"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.MultiPolicyServiceChainsEnabled = castedValue

	}

	if value, ok := values["pbb_etree_enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.PBBEtreeEnable = castedValue

	}

	if value, ok := values["port_security_enabled"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.PortSecurityEnabled = castedValue

	}

	if value, ok := values["ref_bgpvpn"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		log.Debug(stringValue)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualNetworkBGPVPNRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			log.Debug(referenceMap)
			log.Debug(reference)
			log.Debug(referenceMap["to"])
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
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualNetworkNetworkIpamRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
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
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualNetworkSecurityLoggingObjectRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.SecurityLoggingObjectRefs = append(m.SecurityLoggingObjectRefs, referenceModel)

		}
	}

	if value, ok := values["ref_network_policy"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		log.Debug(stringValue)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualNetworkNetworkPolicyRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.NetworkPolicyRefs = append(m.NetworkPolicyRefs, referenceModel)
			log.Debug(referenceMap)
			log.Debug(reference)
			log.Debug(referenceMap["to"])
			log.Debug(referenceModel.UUID)
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
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualNetworkQosConfigRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
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
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualNetworkRouteTableRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
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
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualNetworkVirtualNetworkRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

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
	result := models.MakeVirtualNetworkSlice()
	query, columns, values := common.BuildListQuery(spec)
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
		log.WithFields(log.Fields{
			"valuesMap": valuesMap,
		}).Debug("valueMap")
		m, err := scanVirtualNetwork(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowVirtualNetwork shows VirtualNetwork resource
func ShowVirtualNetwork(tx *sql.Tx, uuid string) (*models.VirtualNetwork, error) {
	list, err := ListVirtualNetwork(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateVirtualNetwork updates a resource
func UpdateVirtualNetwork(tx *sql.Tx, uuid string, model *models.VirtualNetwork) error {
	//TODO(nati) support update
	return nil
}

// DeleteVirtualNetwork deletes a resource
func DeleteVirtualNetwork(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualNetworkQuery)
	if err != nil {
		return errors.Wrap(err, "preparing delete query failed")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	if err != nil {
		return errors.Wrap(err, "delete failed")
	}
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}
