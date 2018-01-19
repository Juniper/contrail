package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertInstanceIPQuery = "insert into `instance_ip` (`uuid`,`subnet_uuid`,`service_instance_ip`,`service_health_check_ip`,`ip_prefix_len`,`ip_prefix`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`instance_ip_secondary`,`instance_ip_mode`,`instance_ip_local_ip`,`instance_ip_family`,`instance_ip_address`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteInstanceIPQuery = "delete from `instance_ip` where uuid = ?"

// InstanceIPFields is db columns for InstanceIP
var InstanceIPFields = []string{
	"uuid",
	"subnet_uuid",
	"service_instance_ip",
	"service_health_check_ip",
	"ip_prefix_len",
	"ip_prefix",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"instance_ip_secondary",
	"instance_ip_mode",
	"instance_ip_local_ip",
	"instance_ip_family",
	"instance_ip_address",
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

// InstanceIPRefFields is db reference fields for InstanceIP
var InstanceIPRefFields = map[string][]string{

	"network_ipam": {
	// <common.Schema Value>

	},

	"virtual_network": {
	// <common.Schema Value>

	},

	"virtual_machine_interface": {
	// <common.Schema Value>

	},

	"physical_router": {
	// <common.Schema Value>

	},

	"virtual_router": {
	// <common.Schema Value>

	},
}

// InstanceIPBackRefFields is db back reference fields for InstanceIP
var InstanceIPBackRefFields = map[string][]string{

	"floating_ip": {
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
		"floating_ip_traffic_direction",
		"port_mappings",
		"floating_ip_port_mappings_enable",
		"floating_ip_is_virtual_ip",
		"floating_ip_fixed_ip_address",
		"floating_ip_address_family",
		"floating_ip_address",
		"display_name",
		"key_value_pair",
	},
}

// InstanceIPParentTypes is possible parents for InstanceIP
var InstanceIPParents = []string{}

const insertInstanceIPVirtualNetworkQuery = "insert into `ref_instance_ip_virtual_network` (`from`, `to` ) values (?, ?);"

const insertInstanceIPVirtualMachineInterfaceQuery = "insert into `ref_instance_ip_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertInstanceIPPhysicalRouterQuery = "insert into `ref_instance_ip_physical_router` (`from`, `to` ) values (?, ?);"

const insertInstanceIPVirtualRouterQuery = "insert into `ref_instance_ip_virtual_router` (`from`, `to` ) values (?, ?);"

const insertInstanceIPNetworkIpamQuery = "insert into `ref_instance_ip_network_ipam` (`from`, `to` ) values (?, ?);"

// CreateInstanceIP inserts InstanceIP to DB
func CreateInstanceIP(tx *sql.Tx, model *models.InstanceIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertInstanceIPQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertInstanceIPQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.UUID),
		string(model.SubnetUUID),
		bool(model.ServiceInstanceIP),
		bool(model.ServiceHealthCheckIP),
		int(model.SecondaryIPTrackingIP.IPPrefixLen),
		string(model.SecondaryIPTrackingIP.IPPrefix),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		bool(model.InstanceIPSecondary),
		string(model.InstanceIPMode),
		bool(model.InstanceIPLocalIP),
		string(model.InstanceIPFamily),
		string(model.InstanceIPAddress),
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

	stmtNetworkIpamRef, err := tx.Prepare(insertInstanceIPNetworkIpamQuery)
	if err != nil {
		return errors.Wrap(err, "preparing NetworkIpamRefs create statement failed")
	}
	defer stmtNetworkIpamRef.Close()
	for _, ref := range model.NetworkIpamRefs {

		_, err = stmtNetworkIpamRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "NetworkIpamRefs create failed")
		}
	}

	stmtVirtualNetworkRef, err := tx.Prepare(insertInstanceIPVirtualNetworkQuery)
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

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertInstanceIPVirtualMachineInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs create statement failed")
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {

		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	stmtPhysicalRouterRef, err := tx.Prepare(insertInstanceIPPhysicalRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing PhysicalRouterRefs create statement failed")
	}
	defer stmtPhysicalRouterRef.Close()
	for _, ref := range model.PhysicalRouterRefs {

		_, err = stmtPhysicalRouterRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "PhysicalRouterRefs create failed")
		}
	}

	stmtVirtualRouterRef, err := tx.Prepare(insertInstanceIPVirtualRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualRouterRefs create statement failed")
	}
	defer stmtVirtualRouterRef.Close()
	for _, ref := range model.VirtualRouterRefs {

		_, err = stmtVirtualRouterRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualRouterRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "instance_ip",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "instance_ip", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanInstanceIP(values map[string]interface{}) (*models.InstanceIP, error) {
	m := models.MakeInstanceIP()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["subnet_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.SubnetUUID = castedValue

	}

	if value, ok := values["service_instance_ip"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ServiceInstanceIP = castedValue

	}

	if value, ok := values["service_health_check_ip"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ServiceHealthCheckIP = castedValue

	}

	if value, ok := values["ip_prefix_len"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.SecondaryIPTrackingIP.IPPrefixLen = castedValue

	}

	if value, ok := values["ip_prefix"]; ok {

		castedValue := common.InterfaceToString(value)

		m.SecondaryIPTrackingIP.IPPrefix = castedValue

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

	if value, ok := values["instance_ip_secondary"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.InstanceIPSecondary = castedValue

	}

	if value, ok := values["instance_ip_mode"]; ok {

		castedValue := common.InterfaceToString(value)

		m.InstanceIPMode = models.AddressMode(castedValue)

	}

	if value, ok := values["instance_ip_local_ip"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.InstanceIPLocalIP = castedValue

	}

	if value, ok := values["instance_ip_family"]; ok {

		castedValue := common.InterfaceToString(value)

		m.InstanceIPFamily = models.IpAddressFamilyType(castedValue)

	}

	if value, ok := values["instance_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.InstanceIPAddress = models.IpAddressType(castedValue)

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
			referenceModel := &models.InstanceIPNetworkIpamRef{}
			referenceModel.UUID = uuid
			m.NetworkIpamRefs = append(m.NetworkIpamRefs, referenceModel)

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
			referenceModel := &models.InstanceIPVirtualNetworkRef{}
			referenceModel.UUID = uuid
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

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
			referenceModel := &models.InstanceIPVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_physical_router"]; ok {
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
			referenceModel := &models.InstanceIPPhysicalRouterRef{}
			referenceModel.UUID = uuid
			m.PhysicalRouterRefs = append(m.PhysicalRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_router"]; ok {
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
			referenceModel := &models.InstanceIPVirtualRouterRef{}
			referenceModel.UUID = uuid
			m.VirtualRouterRefs = append(m.VirtualRouterRefs, referenceModel)

		}
	}

	if value, ok := values["backref_floating_ip"]; ok {
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
			childModel := models.MakeFloatingIP()
			m.FloatingIPs = append(m.FloatingIPs, childModel)

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

			if propertyValue, ok := childResourceMap["floating_ip_traffic_direction"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.FloatingIPTrafficDirection = models.TrafficDirectionType(castedValue)

			}

			if propertyValue, ok := childResourceMap["port_mappings"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FloatingIPPortMappings.PortMappings)

			}

			if propertyValue, ok := childResourceMap["floating_ip_port_mappings_enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.FloatingIPPortMappingsEnable = castedValue

			}

			if propertyValue, ok := childResourceMap["floating_ip_is_virtual_ip"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.FloatingIPIsVirtualIP = castedValue

			}

			if propertyValue, ok := childResourceMap["floating_ip_fixed_ip_address"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.FloatingIPFixedIPAddress = models.IpAddressType(castedValue)

			}

			if propertyValue, ok := childResourceMap["floating_ip_address_family"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.FloatingIPAddressFamily = models.IpAddressFamilyType(castedValue)

			}

			if propertyValue, ok := childResourceMap["floating_ip_address"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.FloatingIPAddress = models.IpAddressType(castedValue)

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

// ListInstanceIP lists InstanceIP with list spec.
func ListInstanceIP(tx *sql.Tx, spec *common.ListSpec) ([]*models.InstanceIP, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "instance_ip"
	spec.Fields = InstanceIPFields
	spec.RefFields = InstanceIPRefFields
	spec.BackRefFields = InstanceIPBackRefFields
	result := models.MakeInstanceIPSlice()

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
		m, err := scanInstanceIP(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateInstanceIP updates a resource
func UpdateInstanceIP(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	//TODO (handle references)
	// Prepare statement for updating data
	var updateInstanceIPQuery = "update `instance_ip` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateInstanceIPQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".SubnetUUID", "."); ok {
		updateInstanceIPQuery += "`subnet_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ServiceInstanceIP", "."); ok {
		updateInstanceIPQuery += "`service_instance_ip` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ServiceHealthCheckIP", "."); ok {
		updateInstanceIPQuery += "`service_health_check_ip` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".SecondaryIPTrackingIP.IPPrefixLen", "."); ok {
		updateInstanceIPQuery += "`ip_prefix_len` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".SecondaryIPTrackingIP.IPPrefix", "."); ok {
		updateInstanceIPQuery += "`ip_prefix` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateInstanceIPQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateInstanceIPQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateInstanceIPQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateInstanceIPQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateInstanceIPQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateInstanceIPQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".InstanceIPSecondary", "."); ok {
		updateInstanceIPQuery += "`instance_ip_secondary` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".InstanceIPMode", "."); ok {
		updateInstanceIPQuery += "`instance_ip_mode` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".InstanceIPLocalIP", "."); ok {
		updateInstanceIPQuery += "`instance_ip_local_ip` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".InstanceIPFamily", "."); ok {
		updateInstanceIPQuery += "`instance_ip_family` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".InstanceIPAddress", "."); ok {
		updateInstanceIPQuery += "`instance_ip_address` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateInstanceIPQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateInstanceIPQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateInstanceIPQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateInstanceIPQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateInstanceIPQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateInstanceIPQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateInstanceIPQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateInstanceIPQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateInstanceIPQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateInstanceIPQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateInstanceIPQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateInstanceIPQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateInstanceIPQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateInstanceIPQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateInstanceIPQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateInstanceIPQuery += ","
	}

	updateInstanceIPQuery =
		updateInstanceIPQuery[:len(updateInstanceIPQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateInstanceIPQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateInstanceIPQuery,
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

// DeleteInstanceIP deletes a resource
func DeleteInstanceIP(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteInstanceIPQuery
	selectQuery := "select count(uuid) from instance_ip where uuid = ?"
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
