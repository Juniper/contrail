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

const insertTagQuery = "insert into `tag` (`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`tag_type_name`,`tag_value`,`uuid`,`fq_name`,`tag_id`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateTagQuery = "update `tag` set `enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`tag_type_name` = ?,`tag_value` = ?,`uuid` = ?,`fq_name` = ?,`tag_id` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteTagQuery = "delete from `tag` where uuid = ?"

// TagFields is db columns for Tag
var TagFields = []string{
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"group",
	"group_access",
	"owner",
	"owner_access",
	"other_access",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"share",
	"tag_type_name",
	"tag_value",
	"uuid",
	"fq_name",
	"tag_id",
	"display_name",
	"key_value_pair",
}

// TagRefFields is db reference fields for Tag
var TagRefFields = map[string][]string{

	"tag_type": {
	// <utils.Schema Value>

	},
}

const insertTagTagTypeQuery = "insert into `ref_tag_tag_type` (`from`, `to` ) values (?, ?);"

// CreateTag inserts Tag to DB
func CreateTag(tx *sql.Tx, model *models.Tag) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertTagQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertTagQuery,
	}).Debug("create query")
	_, err = stmt.Exec(bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.TagTypeName),
		string(model.TagValue),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.TagID),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtTagTypeRef, err := tx.Prepare(insertTagTagTypeQuery)
	if err != nil {
		return errors.Wrap(err, "preparing TagTypeRefs create statement failed")
	}
	defer stmtTagTypeRef.Close()
	for _, ref := range model.TagTypeRefs {
		_, err = stmtTagTypeRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "TagTypeRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanTag(values map[string]interface{}) (*models.Tag, error) {
	m := models.MakeTag()

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

	if value, ok := values["group"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

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

	if value, ok := values["tag_type_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.TagTypeName = castedValue

	}

	if value, ok := values["tag_value"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.TagValue = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["tag_id"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.TagID = models.U32BitHexInt(castedValue)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_tag_type"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.TagTagTypeRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.TagTypeRefs = append(m.TagTypeRefs, referenceModel)

		}
	}

	return m, nil
}

// ListTag lists Tag with list spec.
func ListTag(tx *sql.Tx, spec *db.ListSpec) ([]*models.Tag, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "tag"
	spec.Fields = TagFields
	spec.RefFields = TagRefFields
	result := models.MakeTagSlice()
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
		m, err := scanTag(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowTag shows Tag resource
func ShowTag(tx *sql.Tx, uuid string) (*models.Tag, error) {
	list, err := ListTag(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateTag updates a resource
func UpdateTag(tx *sql.Tx, uuid string, model *models.Tag) error {
	//TODO(nati) support update
	return nil
}

// DeleteTag deletes a resource
func DeleteTag(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteTagQuery)
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
