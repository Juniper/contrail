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

const insertServiceObjectQuery = "insert into `service_object` (`display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceObjectQuery = "update `service_object` set `display_name` = ?,`key_value_pair` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?;"
const deleteServiceObjectQuery = "delete from `service_object` where uuid = ?"
const listServiceObjectQuery = "select `service_object`.`display_name`,`service_object`.`key_value_pair`,`service_object`.`share`,`service_object`.`owner`,`service_object`.`owner_access`,`service_object`.`global_access`,`service_object`.`uuid`,`service_object`.`fq_name`,`service_object`.`last_modified`,`service_object`.`permissions_owner`,`service_object`.`permissions_owner_access`,`service_object`.`other_access`,`service_object`.`group`,`service_object`.`group_access`,`service_object`.`enable`,`service_object`.`description`,`service_object`.`created`,`service_object`.`creator`,`service_object`.`user_visible` from `service_object`"
const showServiceObjectQuery = "select `service_object`.`display_name`,`service_object`.`key_value_pair`,`service_object`.`share`,`service_object`.`owner`,`service_object`.`owner_access`,`service_object`.`global_access`,`service_object`.`uuid`,`service_object`.`fq_name`,`service_object`.`last_modified`,`service_object`.`permissions_owner`,`service_object`.`permissions_owner_access`,`service_object`.`other_access`,`service_object`.`group`,`service_object`.`group_access`,`service_object`.`enable`,`service_object`.`description`,`service_object`.`created`,`service_object`.`creator`,`service_object`.`user_visible` from `service_object` where uuid = ?"

func CreateServiceObject(tx *sql.Tx, model *models.ServiceObject) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceObjectQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		utils.MustJSON(model.FQName),
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
		bool(model.IDPerms.UserVisible))

	return err
}

func scanServiceObject(rows *sql.Rows) (*models.ServiceObject, error) {
	m := models.MakeServiceObject()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.UserVisible); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildServiceObjectWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListServiceObject(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ServiceObject, error) {
	result := models.MakeServiceObjectSlice()
	whereQuery, values := buildServiceObjectWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listServiceObjectQuery)
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
		m, _ := scanServiceObject(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceObject(tx *sql.Tx, uuid string) (*models.ServiceObject, error) {
	rows, err := tx.Query(showServiceObjectQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceObject(rows)
	}
	return nil, nil
}

func UpdateServiceObject(tx *sql.Tx, uuid string, model *models.ServiceObject) error {
	return nil
}

func DeleteServiceObject(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceObjectQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
