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

const insertTagTypeQuery = "insert into `tag_type` (`uuid`,`fq_name`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`tag_type_id`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateTagTypeQuery = "update `tag_type` set `uuid` = ?,`fq_name` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`tag_type_id` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteTagTypeQuery = "delete from `tag_type` where uuid = ?"
const listTagTypeQuery = "select `tag_type`.`uuid`,`tag_type`.`fq_name`,`tag_type`.`last_modified`,`tag_type`.`group`,`tag_type`.`group_access`,`tag_type`.`owner`,`tag_type`.`owner_access`,`tag_type`.`other_access`,`tag_type`.`enable`,`tag_type`.`description`,`tag_type`.`created`,`tag_type`.`creator`,`tag_type`.`user_visible`,`tag_type`.`display_name`,`tag_type`.`tag_type_id`,`tag_type`.`key_value_pair`,`tag_type`.`perms2_owner`,`tag_type`.`perms2_owner_access`,`tag_type`.`global_access`,`tag_type`.`share` from `tag_type`"
const showTagTypeQuery = "select `tag_type`.`uuid`,`tag_type`.`fq_name`,`tag_type`.`last_modified`,`tag_type`.`group`,`tag_type`.`group_access`,`tag_type`.`owner`,`tag_type`.`owner_access`,`tag_type`.`other_access`,`tag_type`.`enable`,`tag_type`.`description`,`tag_type`.`created`,`tag_type`.`creator`,`tag_type`.`user_visible`,`tag_type`.`display_name`,`tag_type`.`tag_type_id`,`tag_type`.`key_value_pair`,`tag_type`.`perms2_owner`,`tag_type`.`perms2_owner_access`,`tag_type`.`global_access`,`tag_type`.`share` from `tag_type` where uuid = ?"

func CreateTagType(tx *sql.Tx, model *models.TagType) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertTagTypeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.DisplayName),
		string(model.TagTypeID),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share))

	return err
}

func scanTagType(rows *sql.Rows) (*models.TagType, error) {
	m := models.MakeTagType()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
		&jsonFQName,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.DisplayName,
		&m.TagTypeID,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildTagTypeWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["tag_type_id"]; ok {
		results = append(results, "tag_type_id = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListTagType(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.TagType, error) {
	result := models.MakeTagTypeSlice()
	whereQuery, values := buildTagTypeWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listTagTypeQuery)
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
		m, _ := scanTagType(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowTagType(tx *sql.Tx, uuid string) (*models.TagType, error) {
	rows, err := tx.Query(showTagTypeQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanTagType(rows)
	}
	return nil, nil
}

func UpdateTagType(tx *sql.Tx, uuid string, model *models.TagType) error {
	return nil
}

func DeleteTagType(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteTagTypeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
