package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertFloatingIPQuery = "insert into `floating_ip` (`floating_ip_address_family`,`floating_ip_port_mappings`,`floating_ip_port_mappings_enable`,`floating_ip_traffic_direction`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`floating_ip_is_virtual_ip`,`floating_ip_address`,`floating_ip_fixed_ip_address`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFloatingIPQuery = "update `floating_ip` set `floating_ip_address_family` = ?,`floating_ip_port_mappings` = ?,`floating_ip_port_mappings_enable` = ?,`floating_ip_traffic_direction` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`fq_name` = ?,`floating_ip_is_virtual_ip` = ?,`floating_ip_address` = ?,`floating_ip_fixed_ip_address` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteFloatingIPQuery = "delete from `floating_ip` where uuid = ?"

// FloatingIPFields is db columns for FloatingIP
var FloatingIPFields = []string{
	"floating_ip_address_family",
	"floating_ip_port_mappings",
	"floating_ip_port_mappings_enable",
	"floating_ip_traffic_direction",
	"global_access",
	"share",
	"owner",
	"owner_access",
	"uuid",
	"fq_name",
	"floating_ip_is_virtual_ip",
	"floating_ip_address",
	"floating_ip_fixed_ip_address",
	"creator",
	"user_visible",
	"last_modified",
	"group",
	"group_access",
	"permissions_owner",
	"permissions_owner_access",
	"other_access",
	"enable",
	"description",
	"created",
	"display_name",
	"key_value_pair",
}

// FloatingIPRefFields is db reference fields for FloatingIP
var FloatingIPRefFields = map[string][]string{

	"project": {
	// <common.Schema Value>

	},

	"virtual_machine_interface": {
	// <common.Schema Value>

	},
}

const insertFloatingIPProjectQuery = "insert into `ref_floating_ip_project` (`from`, `to` ) values (?, ?);"

const insertFloatingIPVirtualMachineInterfaceQuery = "insert into `ref_floating_ip_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

// CreateFloatingIP inserts FloatingIP to DB
func CreateFloatingIP(tx *sql.Tx, model *models.FloatingIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFloatingIPQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertFloatingIPQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.FloatingIPAddressFamily),
		common.MustJSON(model.FloatingIPPortMappings),
		bool(model.FloatingIPPortMappingsEnable),
		string(model.FloatingIPTrafficDirection),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID),
		common.MustJSON(model.FQName),
		bool(model.FloatingIPIsVirtualIP),
		string(model.FloatingIPAddress),
		string(model.FloatingIPFixedIPAddress),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtProjectRef, err := tx.Prepare(insertFloatingIPProjectQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ProjectRefs create statement failed")
	}
	defer stmtProjectRef.Close()
	for _, ref := range model.ProjectRefs {

		_, err = stmtProjectRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ProjectRefs create failed")
		}
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertFloatingIPVirtualMachineInterfaceQuery)
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

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanFloatingIP(values map[string]interface{}) (*models.FloatingIP, error) {
	m := models.MakeFloatingIP()

	if value, ok := values["floating_ip_address_family"]; ok {

		castedValue := common.InterfaceToString(value)

		m.FloatingIPAddressFamily = models.IpAddressFamilyType(castedValue)

	}

	if value, ok := values["floating_ip_port_mappings"]; ok {

		json.Unmarshal(value.([]byte), &m.FloatingIPPortMappings)

	}

	if value, ok := values["floating_ip_port_mappings_enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.FloatingIPPortMappingsEnable = castedValue

	}

	if value, ok := values["floating_ip_traffic_direction"]; ok {

		castedValue := common.InterfaceToString(value)

		m.FloatingIPTrafficDirection = models.TrafficDirectionType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["floating_ip_is_virtual_ip"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.FloatingIPIsVirtualIP = castedValue

	}

	if value, ok := values["floating_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.FloatingIPAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["floating_ip_fixed_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.FloatingIPFixedIPAddress = models.IpAddressType(castedValue)

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

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

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

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_project"]; ok {
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
			referenceModel := &models.FloatingIPProjectRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ProjectRefs = append(m.ProjectRefs, referenceModel)

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
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.FloatingIPVirtualMachineInterfaceRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	return m, nil
}

// ListFloatingIP lists FloatingIP with list spec.
func ListFloatingIP(tx *sql.Tx, spec *common.ListSpec) ([]*models.FloatingIP, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "floating_ip"
	spec.Fields = FloatingIPFields
	spec.RefFields = FloatingIPRefFields
	result := models.MakeFloatingIPSlice()
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
		m, err := scanFloatingIP(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowFloatingIP shows FloatingIP resource
func ShowFloatingIP(tx *sql.Tx, uuid string) (*models.FloatingIP, error) {
	list, err := ListFloatingIP(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateFloatingIP updates a resource
func UpdateFloatingIP(tx *sql.Tx, uuid string, model *models.FloatingIP) error {
	//TODO(nati) support update
	return nil
}

// DeleteFloatingIP deletes a resource
func DeleteFloatingIP(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteFloatingIPQuery)
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
