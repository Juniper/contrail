package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertVirtualMachineQuery = "insert into `virtual_machine` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteVirtualMachineQuery = "delete from `virtual_machine` where uuid = ?"

// VirtualMachineFields is db columns for VirtualMachine
var VirtualMachineFields = []string{
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
}

// VirtualMachineRefFields is db reference fields for VirtualMachine
var VirtualMachineRefFields = map[string][]string{

	"service_instance": {
	// <common.Schema Value>

	},
}

// VirtualMachineBackRefFields is db back reference fields for VirtualMachine
var VirtualMachineBackRefFields = map[string][]string{

	"virtual_machine_interface": {
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
		"annotations_key_value_pair",
	},
}

// VirtualMachineParentTypes is possible parents for VirtualMachine
var VirtualMachineParents = []string{}

const insertVirtualMachineServiceInstanceQuery = "insert into `ref_virtual_machine_service_instance` (`from`, `to` ) values (?, ?);"

// CreateVirtualMachine inserts VirtualMachine to DB
func CreateVirtualMachine(tx *sql.Tx, model *models.VirtualMachine) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualMachineQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertVirtualMachineQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.UUID),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
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
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtServiceInstanceRef, err := tx.Prepare(insertVirtualMachineServiceInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceInstanceRefs create statement failed")
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {

		_, err = stmtServiceInstanceRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "virtual_machine",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "virtual_machine", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanVirtualMachine(values map[string]interface{}) (*models.VirtualMachine, error) {
	m := models.MakeVirtualMachine()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

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

	if value, ok := values["parent_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentUUID = castedValue

	}

	if value, ok := values["parent_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentType = castedValue

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

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_service_instance"]; ok {
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
			referenceModel := &models.VirtualMachineServiceInstanceRef{}
			referenceModel.UUID = uuid
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

		}
	}

	if value, ok := values["backref_virtual_machine_interface"]; ok {
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
			childModel := models.MakeVirtualMachineInterface()
			m.VirtualMachineInterfaces = append(m.VirtualMachineInterfaces, childModel)

			if propertyValue, ok := childResourceMap["vrf_assign_rule"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.VRFAssignTable.VRFAssignRule)

			}

			if propertyValue, ok := childResourceMap["vlan_tag_based_bridge_domain"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.VlanTagBasedBridgeDomain = castedValue

			}

			if propertyValue, ok := childResourceMap["sub_interface_vlan_tag"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.VirtualMachineInterfaceProperties.SubInterfaceVlanTag = castedValue

			}

			if propertyValue, ok := childResourceMap["service_interface_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualMachineInterfaceProperties.ServiceInterfaceType = models.ServiceInterfaceType(castedValue)

			}

			if propertyValue, ok := childResourceMap["local_preference"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.VirtualMachineInterfaceProperties.LocalPreference = castedValue

			}

			if propertyValue, ok := childResourceMap["traffic_direction"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection = models.TrafficDirectionType(castedValue)

			}

			if propertyValue, ok := childResourceMap["udp_port"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort = castedValue

			}

			if propertyValue, ok := childResourceMap["vtep_dst_mac_address"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress = castedValue

			}

			if propertyValue, ok := childResourceMap["vtep_dst_ip_address"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress = castedValue

			}

			if propertyValue, ok := childResourceMap["vni"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni = models.VxlanNetworkIdentifierType(castedValue)

			}

			if propertyValue, ok := childResourceMap["routing_instance"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance = castedValue

			}

			if propertyValue, ok := childResourceMap["nic_assisted_mirroring_vlan"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan = models.VlanIdType(castedValue)

			}

			if propertyValue, ok := childResourceMap["nic_assisted_mirroring"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring = castedValue

			}

			if propertyValue, ok := childResourceMap["nh_mode"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode = models.NHModeType(castedValue)

			}

			if propertyValue, ok := childResourceMap["juniper_header"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader = castedValue

			}

			if propertyValue, ok := childResourceMap["encapsulation"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation = castedValue

			}

			if propertyValue, ok := childResourceMap["analyzer_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName = castedValue

			}

			if propertyValue, ok := childResourceMap["analyzer_mac_address"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress = castedValue

			}

			if propertyValue, ok := childResourceMap["analyzer_ip_address"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress = castedValue

			}

			if propertyValue, ok := childResourceMap["mac_address"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceMacAddresses.MacAddress)

			}

			if propertyValue, ok := childResourceMap["route"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceHostRoutes.Route)

			}

			if propertyValue, ok := childResourceMap["fat_flow_protocol"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceFatFlowProtocols.FatFlowProtocol)

			}

			if propertyValue, ok := childResourceMap["virtual_machine_interface_disable_policy"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.VirtualMachineInterfaceDisablePolicy = castedValue

			}

			if propertyValue, ok := childResourceMap["dhcp_option"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceDHCPOptionList.DHCPOption)

			}

			if propertyValue, ok := childResourceMap["virtual_machine_interface_device_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualMachineInterfaceDeviceOwner = castedValue

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceBindings.KeyValuePair)

			}

			if propertyValue, ok := childResourceMap["allowed_address_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair)

			}

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["port_security_enabled"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.PortSecurityEnabled = castedValue

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

			if propertyValue, ok := childResourceMap["source_port"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.EcmpHashingIncludeFields.SourcePort = castedValue

			}

			if propertyValue, ok := childResourceMap["source_ip"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.EcmpHashingIncludeFields.SourceIP = castedValue

			}

			if propertyValue, ok := childResourceMap["ip_protocol"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.EcmpHashingIncludeFields.IPProtocol = castedValue

			}

			if propertyValue, ok := childResourceMap["hashing_configured"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.EcmpHashingIncludeFields.HashingConfigured = castedValue

			}

			if propertyValue, ok := childResourceMap["destination_port"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.EcmpHashingIncludeFields.DestinationPort = castedValue

			}

			if propertyValue, ok := childResourceMap["destination_ip"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.EcmpHashingIncludeFields.DestinationIP = castedValue

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

			}

			if propertyValue, ok := childResourceMap["annotations_key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListVirtualMachine lists VirtualMachine with list spec.
func ListVirtualMachine(tx *sql.Tx, spec *common.ListSpec) ([]*models.VirtualMachine, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "virtual_machine"
	spec.Fields = VirtualMachineFields
	spec.RefFields = VirtualMachineRefFields
	spec.BackRefFields = VirtualMachineBackRefFields
	result := models.MakeVirtualMachineSlice()

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
		m, err := scanVirtualMachine(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateVirtualMachine updates a resource
func UpdateVirtualMachine(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	// Prepare statement for updating data
	var updateVirtualMachineQuery = "update `virtual_machine` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateVirtualMachineQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateVirtualMachineQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateVirtualMachineQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateVirtualMachineQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateVirtualMachineQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateVirtualMachineQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateVirtualMachineQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateVirtualMachineQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateVirtualMachineQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateVirtualMachineQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateVirtualMachineQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateVirtualMachineQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateVirtualMachineQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateVirtualMachineQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateVirtualMachineQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateVirtualMachineQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateVirtualMachineQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateVirtualMachineQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateVirtualMachineQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateVirtualMachineQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateVirtualMachineQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateVirtualMachineQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateVirtualMachineQuery += ","
	}

	updateVirtualMachineQuery =
		updateVirtualMachineQuery[:len(updateVirtualMachineQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateVirtualMachineQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateVirtualMachineQuery,
	}).Debug("update query")
	_, err = stmt.Exec(updatedValues...)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}

	if value, ok := common.GetValueByPath(model, "ServiceInstanceRefs", "."); ok {
		for _, ref := range value.([]interface{}) {
			refQuery := ""
			refValues := make([]interface{}, 0)
			refKeys := make([]string, 0)
			refUUID, ok := common.GetValueByPath(ref.(map[string]interface{}), "UUID", ".")
			if !ok {
				return errors.Wrap(err, "UUID is missing for referred resource. Failed to update Refs")
			}

			refValues = append(refValues, uuid)
			refValues = append(refValues, refUUID)
			operation, ok := common.GetValueByPath(ref.(map[string]interface{}), common.OPERATION, ".")
			switch operation {
			case common.ADD:
				refQuery = "insert into `ref_virtual_machine_service_instance` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_virtual_machine_service_instance` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref ServiceInstanceRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_virtual_machine_service_instance` where `from` = ? AND `to`= ?;"
				refValues = refValues[len(refValues)-2:]
			default:
				return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
			}
			stmt, err := tx.Prepare(refQuery)
			if err != nil {
				return errors.Wrap(err, "preparing ServiceInstanceRefs update statement failed")
			}
			_, err = stmt.Exec(refValues...)
			if err != nil {
				return errors.Wrap(err, "ServiceInstanceRefs update failed")
			}
		}
	}

	share, ok := common.GetValueByPath(model, ".Perms2.Share", ".")
	if ok {
		err = common.UpdateSharing(tx, "virtual_machine", string(uuid), share.([]interface{}))
		if err != nil {
			return err
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("updated")
	return err
}

// DeleteVirtualMachine deletes a resource
func DeleteVirtualMachine(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteVirtualMachineQuery
	selectQuery := "select count(uuid) from virtual_machine where uuid = ?"
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
