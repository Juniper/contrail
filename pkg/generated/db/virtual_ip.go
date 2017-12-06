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

const insertVirtualIPQuery = "insert into `virtual_ip` (`fq_name`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`protocol_port`,`status`,`protocol`,`address`,`subnet_id`,`persistence_cookie_name`,`persistence_type`,`connection_limit`,`admin_state`,`status_description`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualIPQuery = "update `virtual_ip` set `fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`protocol_port` = ?,`status` = ?,`protocol` = ?,`address` = ?,`subnet_id` = ?,`persistence_cookie_name` = ?,`persistence_type` = ?,`connection_limit` = ?,`admin_state` = ?,`status_description` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?;"
const deleteVirtualIPQuery = "delete from `virtual_ip` where uuid = ?"

// VirtualIPFields is db columns for VirtualIP
var VirtualIPFields = []string{
	"fq_name",
	"user_visible",
	"last_modified",
	"group_access",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"enable",
	"description",
	"created",
	"creator",
	"protocol_port",
	"status",
	"protocol",
	"address",
	"subnet_id",
	"persistence_cookie_name",
	"persistence_type",
	"connection_limit",
	"admin_state",
	"status_description",
	"display_name",
	"key_value_pair",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"share",
	"uuid",
}

// VirtualIPRefFields is db reference fields for VirtualIP
var VirtualIPRefFields = map[string][]string{

	"loadbalancer_pool": {
	// <utils.Schema Value>

	},

	"virtual_machine_interface": {
	// <utils.Schema Value>

	},
}

const insertVirtualIPLoadbalancerPoolQuery = "insert into `ref_virtual_ip_loadbalancer_pool` (`from`, `to` ) values (?, ?);"

const insertVirtualIPVirtualMachineInterfaceQuery = "insert into `ref_virtual_ip_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

// CreateVirtualIP inserts VirtualIP to DB
func CreateVirtualIP(tx *sql.Tx, model *models.VirtualIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualIPQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertVirtualIPQuery,
	}).Debug("create query")
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		int(model.VirtualIPProperties.ProtocolPort),
		string(model.VirtualIPProperties.Status),
		string(model.VirtualIPProperties.Protocol),
		string(model.VirtualIPProperties.Address),
		string(model.VirtualIPProperties.SubnetID),
		string(model.VirtualIPProperties.PersistenceCookieName),
		string(model.VirtualIPProperties.PersistenceType),
		int(model.VirtualIPProperties.ConnectionLimit),
		bool(model.VirtualIPProperties.AdminState),
		string(model.VirtualIPProperties.StatusDescription),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtLoadbalancerPoolRef, err := tx.Prepare(insertVirtualIPLoadbalancerPoolQuery)
	if err != nil {
		return errors.Wrap(err, "preparing LoadbalancerPoolRefs create statement failed")
	}
	defer stmtLoadbalancerPoolRef.Close()
	for _, ref := range model.LoadbalancerPoolRefs {
		_, err = stmtLoadbalancerPoolRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "LoadbalancerPoolRefs create failed")
		}
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertVirtualIPVirtualMachineInterfaceQuery)
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

func scanVirtualIP(values map[string]interface{}) (*models.VirtualIP, error) {
	m := models.MakeVirtualIP()

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

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

	if value, ok := values["protocol_port"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.VirtualIPProperties.ProtocolPort = castedValue

	}

	if value, ok := values["status"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualIPProperties.Status = castedValue

	}

	if value, ok := values["protocol"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualIPProperties.Protocol = models.LoadbalancerProtocolType(castedValue)

	}

	if value, ok := values["address"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualIPProperties.Address = models.IpAddressType(castedValue)

	}

	if value, ok := values["subnet_id"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualIPProperties.SubnetID = models.UuidStringType(castedValue)

	}

	if value, ok := values["persistence_cookie_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualIPProperties.PersistenceCookieName = castedValue

	}

	if value, ok := values["persistence_type"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualIPProperties.PersistenceType = models.SessionPersistenceType(castedValue)

	}

	if value, ok := values["connection_limit"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.VirtualIPProperties.ConnectionLimit = castedValue

	}

	if value, ok := values["admin_state"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.VirtualIPProperties.AdminState = castedValue

	}

	if value, ok := values["status_description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualIPProperties.StatusDescription = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

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

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["ref_loadbalancer_pool"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualIPLoadbalancerPoolRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.LoadbalancerPoolRefs = append(m.LoadbalancerPoolRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.VirtualIPVirtualMachineInterfaceRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	return m, nil
}

// ListVirtualIP lists VirtualIP with list spec.
func ListVirtualIP(tx *sql.Tx, spec *db.ListSpec) ([]*models.VirtualIP, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "virtual_ip"
	spec.Fields = VirtualIPFields
	spec.RefFields = VirtualIPRefFields
	result := models.MakeVirtualIPSlice()
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
		m, err := scanVirtualIP(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowVirtualIP shows VirtualIP resource
func ShowVirtualIP(tx *sql.Tx, uuid string) (*models.VirtualIP, error) {
	list, err := ListVirtualIP(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateVirtualIP updates a resource
func UpdateVirtualIP(tx *sql.Tx, uuid string, model *models.VirtualIP) error {
	//TODO(nati) support update
	return nil
}

// DeleteVirtualIP deletes a resource
func DeleteVirtualIP(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualIPQuery)
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
