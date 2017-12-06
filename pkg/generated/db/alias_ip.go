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

const insertAliasIPQuery = "insert into `alias_ip` (`uuid`,`fq_name`,`alias_ip_address`,`alias_ip_address_family`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAliasIPQuery = "update `alias_ip` set `uuid` = ?,`fq_name` = ?,`alias_ip_address` = ?,`alias_ip_address_family` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?;"
const deleteAliasIPQuery = "delete from `alias_ip` where uuid = ?"

// AliasIPFields is db columns for AliasIP
var AliasIPFields = []string{
	"uuid",
	"fq_name",
	"alias_ip_address",
	"alias_ip_address_family",
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
	"display_name",
	"key_value_pair",
	"perms2_owner_access",
	"global_access",
	"share",
	"perms2_owner",
}

// AliasIPRefFields is db reference fields for AliasIP
var AliasIPRefFields = map[string][]string{

	"project": {
	// <utils.Schema Value>

	},

	"virtual_machine_interface": {
	// <utils.Schema Value>

	},
}

const insertAliasIPProjectQuery = "insert into `ref_alias_ip_project` (`from`, `to` ) values (?, ?);"

const insertAliasIPVirtualMachineInterfaceQuery = "insert into `ref_alias_ip_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

// CreateAliasIP inserts AliasIP to DB
func CreateAliasIP(tx *sql.Tx, model *models.AliasIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAliasIPQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertAliasIPQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.AliasIPAddress),
		string(model.AliasIPAddressFamily),
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
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertAliasIPVirtualMachineInterfaceQuery)
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

	stmtProjectRef, err := tx.Prepare(insertAliasIPProjectQuery)
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

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanAliasIP(values map[string]interface{}) (*models.AliasIP, error) {
	m := models.MakeAliasIP()

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["alias_ip_address"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.AliasIPAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["alias_ip_address_family"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.AliasIPAddressFamily = models.IpAddressFamilyType(castedValue)

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

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

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

	if value, ok := values["perms2_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["ref_project"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.AliasIPProjectRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.ProjectRefs = append(m.ProjectRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.AliasIPVirtualMachineInterfaceRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	return m, nil
}

// ListAliasIP lists AliasIP with list spec.
func ListAliasIP(tx *sql.Tx, spec *db.ListSpec) ([]*models.AliasIP, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "alias_ip"
	spec.Fields = AliasIPFields
	spec.RefFields = AliasIPRefFields
	result := models.MakeAliasIPSlice()
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
		m, err := scanAliasIP(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowAliasIP shows AliasIP resource
func ShowAliasIP(tx *sql.Tx, uuid string) (*models.AliasIP, error) {
	list, err := ListAliasIP(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateAliasIP updates a resource
func UpdateAliasIP(tx *sql.Tx, uuid string, model *models.AliasIP) error {
	//TODO(nati) support update
	return nil
}

// DeleteAliasIP deletes a resource
func DeleteAliasIP(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAliasIPQuery)
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
