package db

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"strings"
)

const insertTagQuery = "insert into `tag` (`tag_type_name`,`tag_id`,`tag_value`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`fq_name`,`display_name`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateTagQuery = "update `tag` set `tag_type_name` = ?,`tag_id` = ?,`tag_value` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`fq_name` = ?,`display_name` = ?,`uuid` = ?;"
const deleteTagQuery = "delete from `tag` where uuid = ?"
const listTagQuery = "select `tag`.`tag_type_name`,`tag`.`tag_id`,`tag`.`tag_value`,`tag`.`last_modified`,`tag`.`owner_access`,`tag`.`other_access`,`tag`.`group`,`tag`.`group_access`,`tag`.`owner`,`tag`.`enable`,`tag`.`description`,`tag`.`created`,`tag`.`creator`,`tag`.`user_visible`,`tag`.`key_value_pair`,`tag`.`perms2_owner`,`tag`.`perms2_owner_access`,`tag`.`global_access`,`tag`.`share`,`tag`.`fq_name`,`tag`.`display_name`,`tag`.`uuid` from `tag`"
const showTagQuery = "select `tag`.`tag_type_name`,`tag`.`tag_id`,`tag`.`tag_value`,`tag`.`last_modified`,`tag`.`owner_access`,`tag`.`other_access`,`tag`.`group`,`tag`.`group_access`,`tag`.`owner`,`tag`.`enable`,`tag`.`description`,`tag`.`created`,`tag`.`creator`,`tag`.`user_visible`,`tag`.`key_value_pair`,`tag`.`perms2_owner`,`tag`.`perms2_owner_access`,`tag`.`global_access`,`tag`.`share`,`tag`.`fq_name`,`tag`.`display_name`,`tag`.`uuid` from `tag` where uuid = ?"

const insertTagTagTypeQuery = "insert into `ref_tag_tag_type` (`from`, `to` ) values (?, ?);"

func CreateTag(tx *sql.Tx, model *models.Tag) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertTagQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.TagTypeName),
		string(model.TagID),
		string(model.TagValue),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		utils.MustJSON(model.FQName),
		string(model.DisplayName),
		string(model.UUID))

	stmtTagTypeRef, err := tx.Prepare(insertTagTagTypeQuery)
	if err != nil {
		return err
	}
	defer stmtTagTypeRef.Close()
	for _, ref := range model.TagTypeRefs {
		_, err = stmtTagTypeRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanTag(rows *sql.Rows) (*models.Tag, error) {
	m := models.MakeTag()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.TagTypeName,
		&m.TagID,
		&m.TagValue,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&jsonFQName,
		&m.DisplayName,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildTagWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["tag_type_name"]; ok {
		results = append(results, "tag_type_name = ?")
		values = append(values, value)
	}

	if value, ok := where["tag_id"]; ok {
		results = append(results, "tag_id = ?")
		values = append(values, value)
	}

	if value, ok := where["tag_value"]; ok {
		results = append(results, "tag_value = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListTag(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Tag, error) {
	result := models.MakeTagSlice()
	whereQuery, values := buildTagWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listTagQuery)
	query.WriteRune(' ')
	query.WriteString(whereQuery)
	query.WriteRune(' ')
	query.WriteString(pagenationQuery)
	rows, err = tx.Query(query.String(), values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanTag(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowTag(tx *sql.Tx, uuid string) (*models.Tag, error) {
	rows, err := tx.Query(showTagQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanTag(rows)
	}
	return nil, nil
}

func UpdateTag(tx *sql.Tx, uuid string, model *models.Tag) error {
	return nil
}

func DeleteTag(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteTagQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
