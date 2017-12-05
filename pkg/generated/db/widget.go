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

const insertWidgetQuery = "insert into `widget` (`fq_name`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`content_config`,`layout_config`,`uuid`,`container_config`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateWidgetQuery = "update `widget` set `fq_name` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`content_config` = ?,`layout_config` = ?,`uuid` = ?,`container_config` = ?,`key_value_pair` = ?;"
const deleteWidgetQuery = "delete from `widget` where uuid = ?"
const listWidgetQuery = "select `widget`.`fq_name`,`widget`.`last_modified`,`widget`.`owner`,`widget`.`owner_access`,`widget`.`other_access`,`widget`.`group`,`widget`.`group_access`,`widget`.`enable`,`widget`.`description`,`widget`.`created`,`widget`.`creator`,`widget`.`user_visible`,`widget`.`display_name`,`widget`.`share`,`widget`.`perms2_owner`,`widget`.`perms2_owner_access`,`widget`.`global_access`,`widget`.`content_config`,`widget`.`layout_config`,`widget`.`uuid`,`widget`.`container_config`,`widget`.`key_value_pair` from `widget`"
const showWidgetQuery = "select `widget`.`fq_name`,`widget`.`last_modified`,`widget`.`owner`,`widget`.`owner_access`,`widget`.`other_access`,`widget`.`group`,`widget`.`group_access`,`widget`.`enable`,`widget`.`description`,`widget`.`created`,`widget`.`creator`,`widget`.`user_visible`,`widget`.`display_name`,`widget`.`share`,`widget`.`perms2_owner`,`widget`.`perms2_owner_access`,`widget`.`global_access`,`widget`.`content_config`,`widget`.`layout_config`,`widget`.`uuid`,`widget`.`container_config`,`widget`.`key_value_pair` from `widget` where uuid = ?"

func CreateWidget(tx *sql.Tx, model *models.Widget) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertWidgetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
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
		string(model.DisplayName),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.ContentConfig),
		string(model.LayoutConfig),
		string(model.UUID),
		string(model.ContainerConfig),
		utils.MustJSON(model.Annotations.KeyValuePair))

	return err
}

func scanWidget(rows *sql.Rows) (*models.Widget, error) {
	m := models.MakeWidget()

	var jsonFQName string

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&jsonFQName,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.DisplayName,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.ContentConfig,
		&m.LayoutConfig,
		&m.UUID,
		&m.ContainerConfig,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildWidgetWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
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

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["content_config"]; ok {
		results = append(results, "content_config = ?")
		values = append(values, value)
	}

	if value, ok := where["layout_config"]; ok {
		results = append(results, "layout_config = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["container_config"]; ok {
		results = append(results, "container_config = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListWidget(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Widget, error) {
	result := models.MakeWidgetSlice()
	whereQuery, values := buildWidgetWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listWidgetQuery)
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
		m, _ := scanWidget(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowWidget(tx *sql.Tx, uuid string) (*models.Widget, error) {
	rows, err := tx.Query(showWidgetQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanWidget(rows)
	}
	return nil, nil
}

func UpdateWidget(tx *sql.Tx, uuid string, model *models.Widget) error {
	return nil
}

func DeleteWidget(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteWidgetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
