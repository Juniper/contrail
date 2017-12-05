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

const insertPortTupleQuery = "insert into `port_tuple` (`display_name`,`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePortTupleQuery = "update `port_tuple` set `display_name` = ?,`key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?;"
const deletePortTupleQuery = "delete from `port_tuple` where uuid = ?"
const listPortTupleQuery = "select `port_tuple`.`display_name`,`port_tuple`.`key_value_pair`,`port_tuple`.`owner`,`port_tuple`.`owner_access`,`port_tuple`.`global_access`,`port_tuple`.`share`,`port_tuple`.`uuid`,`port_tuple`.`fq_name`,`port_tuple`.`enable`,`port_tuple`.`description`,`port_tuple`.`created`,`port_tuple`.`creator`,`port_tuple`.`user_visible`,`port_tuple`.`last_modified`,`port_tuple`.`permissions_owner`,`port_tuple`.`permissions_owner_access`,`port_tuple`.`other_access`,`port_tuple`.`group`,`port_tuple`.`group_access` from `port_tuple`"
const showPortTupleQuery = "select `port_tuple`.`display_name`,`port_tuple`.`key_value_pair`,`port_tuple`.`owner`,`port_tuple`.`owner_access`,`port_tuple`.`global_access`,`port_tuple`.`share`,`port_tuple`.`uuid`,`port_tuple`.`fq_name`,`port_tuple`.`enable`,`port_tuple`.`description`,`port_tuple`.`created`,`port_tuple`.`creator`,`port_tuple`.`user_visible`,`port_tuple`.`last_modified`,`port_tuple`.`permissions_owner`,`port_tuple`.`permissions_owner_access`,`port_tuple`.`other_access`,`port_tuple`.`group`,`port_tuple`.`group_access` from `port_tuple` where uuid = ?"

func CreatePortTuple(tx *sql.Tx, model *models.PortTuple) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPortTupleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID),
		utils.MustJSON(model.FQName),
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
		int(model.IDPerms.Permissions.GroupAccess))

	return err
}

func scanPortTuple(rows *sql.Rows) (*models.PortTuple, error) {
	m := models.MakePortTuple()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildPortTupleWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	return "where " + strings.Join(results, " and "), values
}

func ListPortTuple(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.PortTuple, error) {
	result := models.MakePortTupleSlice()
	whereQuery, values := buildPortTupleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listPortTupleQuery)
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
		m, _ := scanPortTuple(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowPortTuple(tx *sql.Tx, uuid string) (*models.PortTuple, error) {
	rows, err := tx.Query(showPortTupleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanPortTuple(rows)
	}
	return nil, nil
}

func UpdatePortTuple(tx *sql.Tx, uuid string, model *models.PortTuple) error {
	return nil
}

func DeletePortTuple(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deletePortTupleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
