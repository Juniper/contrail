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

const insertFloatingIPPoolQuery = "insert into `floating_ip_pool` (`uuid`,`fq_name`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`subnet_uuid`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFloatingIPPoolQuery = "update `floating_ip_pool` set `uuid` = ?,`fq_name` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`subnet_uuid` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?;"
const deleteFloatingIPPoolQuery = "delete from `floating_ip_pool` where uuid = ?"
const listFloatingIPPoolQuery = "select `floating_ip_pool`.`uuid`,`floating_ip_pool`.`fq_name`,`floating_ip_pool`.`group`,`floating_ip_pool`.`group_access`,`floating_ip_pool`.`owner`,`floating_ip_pool`.`owner_access`,`floating_ip_pool`.`other_access`,`floating_ip_pool`.`enable`,`floating_ip_pool`.`description`,`floating_ip_pool`.`created`,`floating_ip_pool`.`creator`,`floating_ip_pool`.`user_visible`,`floating_ip_pool`.`last_modified`,`floating_ip_pool`.`display_name`,`floating_ip_pool`.`subnet_uuid`,`floating_ip_pool`.`key_value_pair`,`floating_ip_pool`.`share`,`floating_ip_pool`.`perms2_owner`,`floating_ip_pool`.`perms2_owner_access`,`floating_ip_pool`.`global_access` from `floating_ip_pool`"
const showFloatingIPPoolQuery = "select `floating_ip_pool`.`uuid`,`floating_ip_pool`.`fq_name`,`floating_ip_pool`.`group`,`floating_ip_pool`.`group_access`,`floating_ip_pool`.`owner`,`floating_ip_pool`.`owner_access`,`floating_ip_pool`.`other_access`,`floating_ip_pool`.`enable`,`floating_ip_pool`.`description`,`floating_ip_pool`.`created`,`floating_ip_pool`.`creator`,`floating_ip_pool`.`user_visible`,`floating_ip_pool`.`last_modified`,`floating_ip_pool`.`display_name`,`floating_ip_pool`.`subnet_uuid`,`floating_ip_pool`.`key_value_pair`,`floating_ip_pool`.`share`,`floating_ip_pool`.`perms2_owner`,`floating_ip_pool`.`perms2_owner_access`,`floating_ip_pool`.`global_access` from `floating_ip_pool` where uuid = ?"

func CreateFloatingIPPool(tx *sql.Tx, model *models.FloatingIPPool) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFloatingIPPoolQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		utils.MustJSON(model.FQName),
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
		string(model.IDPerms.LastModified),
		string(model.DisplayName),
		utils.MustJSON(model.FloatingIPPoolSubnets.SubnetUUID),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess))

	return err
}

func scanFloatingIPPool(rows *sql.Rows) (*models.FloatingIPPool, error) {
	m := models.MakeFloatingIPPool()

	var jsonFQName string

	var jsonFloatingIPPoolSubnetsSubnetUUID string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.LastModified,
		&m.DisplayName,
		&jsonFloatingIPPoolSubnetsSubnetUUID,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonFloatingIPPoolSubnetsSubnetUUID), &m.FloatingIPPoolSubnets.SubnetUUID)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildFloatingIPPoolWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
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

func ListFloatingIPPool(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.FloatingIPPool, error) {
	result := models.MakeFloatingIPPoolSlice()
	whereQuery, values := buildFloatingIPPoolWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listFloatingIPPoolQuery)
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
		m, _ := scanFloatingIPPool(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowFloatingIPPool(tx *sql.Tx, uuid string) (*models.FloatingIPPool, error) {
	rows, err := tx.Query(showFloatingIPPoolQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanFloatingIPPool(rows)
	}
	return nil, nil
}

func UpdateFloatingIPPool(tx *sql.Tx, uuid string, model *models.FloatingIPPool) error {
	return nil
}

func DeleteFloatingIPPool(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteFloatingIPPoolQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
