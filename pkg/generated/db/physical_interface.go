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

const insertPhysicalInterfaceQuery = "insert into `physical_interface` (`ethernet_segment_identifier`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePhysicalInterfaceQuery = "update `physical_interface` set `ethernet_segment_identifier` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?;"
const deletePhysicalInterfaceQuery = "delete from `physical_interface` where uuid = ?"

// PhysicalInterfaceFields is db columns for PhysicalInterface
var PhysicalInterfaceFields = []string{
	"ethernet_segment_identifier",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"display_name",
	"key_value_pair",
	"share",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"uuid",
	"fq_name",
}

// PhysicalInterfaceRefFields is db reference fields for PhysicalInterface
var PhysicalInterfaceRefFields = map[string][]string{

	"physical_interface": {
	// <utils.Schema Value>

	},
}

const insertPhysicalInterfacePhysicalInterfaceQuery = "insert into `ref_physical_interface_physical_interface` (`from`, `to` ) values (?, ?);"

// CreatePhysicalInterface inserts PhysicalInterface to DB
func CreatePhysicalInterface(tx *sql.Tx, model *models.PhysicalInterface) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPhysicalInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertPhysicalInterfaceQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.EthernetSegmentIdentifier),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		utils.MustJSON(model.FQName))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtPhysicalInterfaceRef, err := tx.Prepare(insertPhysicalInterfacePhysicalInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing PhysicalInterfaceRefs create statement failed")
	}
	defer stmtPhysicalInterfaceRef.Close()
	for _, ref := range model.PhysicalInterfaceRefs {
		_, err = stmtPhysicalInterfaceRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "PhysicalInterfaceRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanPhysicalInterface(values map[string]interface{}) (*models.PhysicalInterface, error) {
	m := models.MakePhysicalInterface()

	if value, ok := values["ethernet_segment_identifier"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.EthernetSegmentIdentifier = castedValue

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

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

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

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["ref_physical_interface"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.PhysicalInterfacePhysicalInterfaceRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.PhysicalInterfaceRefs = append(m.PhysicalInterfaceRefs, referenceModel)

		}
	}

	return m, nil
}

// ListPhysicalInterface lists PhysicalInterface with list spec.
func ListPhysicalInterface(tx *sql.Tx, spec *db.ListSpec) ([]*models.PhysicalInterface, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "physical_interface"
	spec.Fields = PhysicalInterfaceFields
	spec.RefFields = PhysicalInterfaceRefFields
	result := models.MakePhysicalInterfaceSlice()
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
		m, err := scanPhysicalInterface(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowPhysicalInterface shows PhysicalInterface resource
func ShowPhysicalInterface(tx *sql.Tx, uuid string) (*models.PhysicalInterface, error) {
	list, err := ListPhysicalInterface(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdatePhysicalInterface updates a resource
func UpdatePhysicalInterface(tx *sql.Tx, uuid string, model *models.PhysicalInterface) error {
	//TODO(nati) support update
	return nil
}

// DeletePhysicalInterface deletes a resource
func DeletePhysicalInterface(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deletePhysicalInterfaceQuery)
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
