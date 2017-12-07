package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertInstanceIPQuery = "insert into `instance_ip` (`instance_ip_mode`,`instance_ip_local_ip`,`instance_ip_address`,`subnet_uuid`,`service_instance_ip`,`owner`,`owner_access`,`global_access`,`share`,`fq_name`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`key_value_pair`,`display_name`,`service_health_check_ip`,`ip_prefix`,`ip_prefix_len`,`instance_ip_family`,`instance_ip_secondary`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateInstanceIPQuery = "update `instance_ip` set `instance_ip_mode` = ?,`instance_ip_local_ip` = ?,`instance_ip_address` = ?,`subnet_uuid` = ?,`service_instance_ip` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`fq_name` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`key_value_pair` = ?,`display_name` = ?,`service_health_check_ip` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`instance_ip_family` = ?,`instance_ip_secondary` = ?,`uuid` = ?;"
const deleteInstanceIPQuery = "delete from `instance_ip` where uuid = ?"

// InstanceIPFields is db columns for InstanceIP
var InstanceIPFields = []string{
	"instance_ip_mode",
	"instance_ip_local_ip",
	"instance_ip_address",
	"subnet_uuid",
	"service_instance_ip",
	"owner",
	"owner_access",
	"global_access",
	"share",
	"fq_name",
	"permissions_owner",
	"permissions_owner_access",
	"other_access",
	"group",
	"group_access",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"key_value_pair",
	"display_name",
	"service_health_check_ip",
	"ip_prefix",
	"ip_prefix_len",
	"instance_ip_family",
	"instance_ip_secondary",
	"uuid",
}

// InstanceIPRefFields is db reference fields for InstanceIP
var InstanceIPRefFields = map[string][]string{

	"network_ipam": {
	// <utils.Schema Value>

	},

	"virtual_network": {
	// <utils.Schema Value>

	},

	"virtual_machine_interface": {
	// <utils.Schema Value>

	},

	"physical_router": {
	// <utils.Schema Value>

	},

	"virtual_router": {
	// <utils.Schema Value>

	},
}

const insertInstanceIPNetworkIpamQuery = "insert into `ref_instance_ip_network_ipam` (`from`, `to` ) values (?, ?);"

const insertInstanceIPVirtualNetworkQuery = "insert into `ref_instance_ip_virtual_network` (`from`, `to` ) values (?, ?);"

const insertInstanceIPVirtualMachineInterfaceQuery = "insert into `ref_instance_ip_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertInstanceIPPhysicalRouterQuery = "insert into `ref_instance_ip_physical_router` (`from`, `to` ) values (?, ?);"

const insertInstanceIPVirtualRouterQuery = "insert into `ref_instance_ip_virtual_router` (`from`, `to` ) values (?, ?);"

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
	_, err = stmt.Exec(string(model.InstanceIPMode),
		bool(model.InstanceIPLocalIP),
		string(model.InstanceIPAddress),
		string(model.SubnetUUID),
		bool(model.ServiceInstanceIP),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.DisplayName),
		bool(model.ServiceHealthCheckIP),
		string(model.SecondaryIPTrackingIP.IPPrefix),
		int(model.SecondaryIPTrackingIP.IPPrefixLen),
		string(model.InstanceIPFamily),
		bool(model.InstanceIPSecondary),
		string(model.UUID))
	if err != nil {
		return errors.Wrap(err, "create failed")
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

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanInstanceIP(values map[string]interface{}) (*models.InstanceIP, error) {
	m := models.MakeInstanceIP()

	if value, ok := values["instance_ip_mode"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.InstanceIPMode = models.AddressMode(castedValue)

	}

	if value, ok := values["instance_ip_local_ip"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.InstanceIPLocalIP = castedValue

	}

	if value, ok := values["instance_ip_address"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.InstanceIPAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["subnet_uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.SubnetUUID = castedValue

	}

	if value, ok := values["service_instance_ip"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.ServiceInstanceIP = castedValue

	}

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["enable"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["service_health_check_ip"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.ServiceHealthCheckIP = castedValue

	}

	if value, ok := values["ip_prefix"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.SecondaryIPTrackingIP.IPPrefix = castedValue

	}

	if value, ok := values["ip_prefix_len"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.SecondaryIPTrackingIP.IPPrefixLen = castedValue

	}

	if value, ok := values["instance_ip_family"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.InstanceIPFamily = models.IpAddressFamilyType(castedValue)

	}

	if value, ok := values["instance_ip_secondary"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.InstanceIPSecondary = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["ref_virtual_router"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.InstanceIPVirtualRouterRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.VirtualRouterRefs = append(m.VirtualRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_network_ipam"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.InstanceIPNetworkIpamRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.NetworkIpamRefs = append(m.NetworkIpamRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_network"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.InstanceIPVirtualNetworkRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.InstanceIPVirtualMachineInterfaceRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_physical_router"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.InstanceIPPhysicalRouterRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.PhysicalRouterRefs = append(m.PhysicalRouterRefs, referenceModel)

		}
	}

	return m, nil
}

// ListInstanceIP lists InstanceIP with list spec.
func ListInstanceIP(tx *sql.Tx, spec *db.ListSpec) ([]*models.InstanceIP, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "instance_ip"
	spec.Fields = InstanceIPFields
	spec.RefFields = InstanceIPRefFields
	result := models.MakeInstanceIPSlice()
	query, columns, values := db.BuildListQuery(spec)
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
		m, err := scanInstanceIP(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowInstanceIP shows InstanceIP resource
func ShowInstanceIP(tx *sql.Tx, uuid string) (*models.InstanceIP, error) {
	list, err := ListInstanceIP(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateInstanceIP updates a resource
func UpdateInstanceIP(tx *sql.Tx, uuid string, model *models.InstanceIP) error {
	//TODO(nati) support update
	return nil
}

// DeleteInstanceIP deletes a resource
func DeleteInstanceIP(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteInstanceIPQuery)
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
