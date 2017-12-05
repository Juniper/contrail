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

const insertInterfaceRouteTableQuery = "insert into `interface_route_table` (`uuid`,`fq_name`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`route`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateInterfaceRouteTableQuery = "update `interface_route_table` set `uuid` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`route` = ?;"
const deleteInterfaceRouteTableQuery = "delete from `interface_route_table` where uuid = ?"
const listInterfaceRouteTableQuery = "select `interface_route_table`.`uuid`,`interface_route_table`.`fq_name`,`interface_route_table`.`user_visible`,`interface_route_table`.`last_modified`,`interface_route_table`.`owner_access`,`interface_route_table`.`other_access`,`interface_route_table`.`group`,`interface_route_table`.`group_access`,`interface_route_table`.`owner`,`interface_route_table`.`enable`,`interface_route_table`.`description`,`interface_route_table`.`created`,`interface_route_table`.`creator`,`interface_route_table`.`display_name`,`interface_route_table`.`key_value_pair`,`interface_route_table`.`perms2_owner`,`interface_route_table`.`perms2_owner_access`,`interface_route_table`.`global_access`,`interface_route_table`.`share`,`interface_route_table`.`route` from `interface_route_table`"
const showInterfaceRouteTableQuery = "select `interface_route_table`.`uuid`,`interface_route_table`.`fq_name`,`interface_route_table`.`user_visible`,`interface_route_table`.`last_modified`,`interface_route_table`.`owner_access`,`interface_route_table`.`other_access`,`interface_route_table`.`group`,`interface_route_table`.`group_access`,`interface_route_table`.`owner`,`interface_route_table`.`enable`,`interface_route_table`.`description`,`interface_route_table`.`created`,`interface_route_table`.`creator`,`interface_route_table`.`display_name`,`interface_route_table`.`key_value_pair`,`interface_route_table`.`perms2_owner`,`interface_route_table`.`perms2_owner_access`,`interface_route_table`.`global_access`,`interface_route_table`.`share`,`interface_route_table`.`route` from `interface_route_table` where uuid = ?"

const insertInterfaceRouteTableServiceInstanceQuery = "insert into `ref_interface_route_table_service_instance` (`from`, `to` ,`interface_type`) values (?, ?,?);"

func CreateInterfaceRouteTable(tx *sql.Tx, model *models.InterfaceRouteTable) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertInterfaceRouteTableQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		utils.MustJSON(model.FQName),
		bool(model.IDPerms.UserVisible),
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
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		utils.MustJSON(model.InterfaceRouteTableRoutes.Route))

	stmtServiceInstanceRef, err := tx.Prepare(insertInterfaceRouteTableServiceInstanceQuery)
	if err != nil {
		return err
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {
		_, err = stmtServiceInstanceRef.Exec(model.UUID, ref.UUID, string(ref.Attr.InterfaceType))
	}

	return err
}

func scanInterfaceRouteTable(rows *sql.Rows) (*models.InterfaceRouteTable, error) {
	m := models.MakeInterfaceRouteTable()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonInterfaceRouteTableRoutesRoute string

	if err := rows.Scan(&m.UUID,
		&jsonFQName,
		&m.IDPerms.UserVisible,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&jsonInterfaceRouteTableRoutesRoute); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonInterfaceRouteTableRoutesRoute), &m.InterfaceRouteTableRoutes.Route)

	return m, nil
}

func buildInterfaceRouteTableWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListInterfaceRouteTable(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.InterfaceRouteTable, error) {
	result := models.MakeInterfaceRouteTableSlice()
	whereQuery, values := buildInterfaceRouteTableWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listInterfaceRouteTableQuery)
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
		m, _ := scanInterfaceRouteTable(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowInterfaceRouteTable(tx *sql.Tx, uuid string) (*models.InterfaceRouteTable, error) {
	rows, err := tx.Query(showInterfaceRouteTableQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanInterfaceRouteTable(rows)
	}
	return nil, nil
}

func UpdateInterfaceRouteTable(tx *sql.Tx, uuid string, model *models.InterfaceRouteTable) error {
	return nil
}

func DeleteInterfaceRouteTable(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteInterfaceRouteTableQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
