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

const insertBridgeDomainQuery = "insert into `bridge_domain` (`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`isid`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`mac_move_limit`,`mac_move_limit_action`,`mac_move_time_window`,`mac_limit`,`mac_limit_action`,`display_name`,`key_value_pair`,`uuid`,`fq_name`,`mac_aging_time`,`mac_learning_enabled`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBridgeDomainQuery = "update `bridge_domain` set `created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`isid` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`mac_move_limit` = ?,`mac_move_limit_action` = ?,`mac_move_time_window` = ?,`mac_limit` = ?,`mac_limit_action` = ?,`display_name` = ?,`key_value_pair` = ?,`uuid` = ?,`fq_name` = ?,`mac_aging_time` = ?,`mac_learning_enabled` = ?;"
const deleteBridgeDomainQuery = "delete from `bridge_domain` where uuid = ?"
const listBridgeDomainQuery = "select `bridge_domain`.`created`,`bridge_domain`.`creator`,`bridge_domain`.`user_visible`,`bridge_domain`.`last_modified`,`bridge_domain`.`group_access`,`bridge_domain`.`owner`,`bridge_domain`.`owner_access`,`bridge_domain`.`other_access`,`bridge_domain`.`group`,`bridge_domain`.`enable`,`bridge_domain`.`description`,`bridge_domain`.`isid`,`bridge_domain`.`share`,`bridge_domain`.`perms2_owner`,`bridge_domain`.`perms2_owner_access`,`bridge_domain`.`global_access`,`bridge_domain`.`mac_move_limit`,`bridge_domain`.`mac_move_limit_action`,`bridge_domain`.`mac_move_time_window`,`bridge_domain`.`mac_limit`,`bridge_domain`.`mac_limit_action`,`bridge_domain`.`display_name`,`bridge_domain`.`key_value_pair`,`bridge_domain`.`uuid`,`bridge_domain`.`fq_name`,`bridge_domain`.`mac_aging_time`,`bridge_domain`.`mac_learning_enabled` from `bridge_domain`"
const showBridgeDomainQuery = "select `bridge_domain`.`created`,`bridge_domain`.`creator`,`bridge_domain`.`user_visible`,`bridge_domain`.`last_modified`,`bridge_domain`.`group_access`,`bridge_domain`.`owner`,`bridge_domain`.`owner_access`,`bridge_domain`.`other_access`,`bridge_domain`.`group`,`bridge_domain`.`enable`,`bridge_domain`.`description`,`bridge_domain`.`isid`,`bridge_domain`.`share`,`bridge_domain`.`perms2_owner`,`bridge_domain`.`perms2_owner_access`,`bridge_domain`.`global_access`,`bridge_domain`.`mac_move_limit`,`bridge_domain`.`mac_move_limit_action`,`bridge_domain`.`mac_move_time_window`,`bridge_domain`.`mac_limit`,`bridge_domain`.`mac_limit_action`,`bridge_domain`.`display_name`,`bridge_domain`.`key_value_pair`,`bridge_domain`.`uuid`,`bridge_domain`.`fq_name`,`bridge_domain`.`mac_aging_time`,`bridge_domain`.`mac_learning_enabled` from `bridge_domain` where uuid = ?"

func CreateBridgeDomain(tx *sql.Tx, model *models.BridgeDomain) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBridgeDomainQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		int(model.Isid),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		int(model.MacMoveControl.MacMoveLimit),
		string(model.MacMoveControl.MacMoveLimitAction),
		int(model.MacMoveControl.MacMoveTimeWindow),
		int(model.MacLimitControl.MacLimit),
		string(model.MacLimitControl.MacLimitAction),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		int(model.MacAgingTime),
		bool(model.MacLearningEnabled))

	return err
}

func scanBridgeDomain(rows *sql.Rows) (*models.BridgeDomain, error) {
	m := models.MakeBridgeDomain()

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	if err := rows.Scan(&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.Isid,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.MacMoveControl.MacMoveLimit,
		&m.MacMoveControl.MacMoveLimitAction,
		&m.MacMoveControl.MacMoveTimeWindow,
		&m.MacLimitControl.MacLimit,
		&m.MacLimitControl.MacLimitAction,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.UUID,
		&jsonFQName,
		&m.MacAgingTime,
		&m.MacLearningEnabled); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildBridgeDomainWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

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

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["mac_move_limit_action"]; ok {
		results = append(results, "mac_move_limit_action = ?")
		values = append(values, value)
	}

	if value, ok := where["mac_limit_action"]; ok {
		results = append(results, "mac_limit_action = ?")
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

func ListBridgeDomain(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.BridgeDomain, error) {
	result := models.MakeBridgeDomainSlice()
	whereQuery, values := buildBridgeDomainWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listBridgeDomainQuery)
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
		m, _ := scanBridgeDomain(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowBridgeDomain(tx *sql.Tx, uuid string) (*models.BridgeDomain, error) {
	rows, err := tx.Query(showBridgeDomainQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanBridgeDomain(rows)
	}
	return nil, nil
}

func UpdateBridgeDomain(tx *sql.Tx, uuid string, model *models.BridgeDomain) error {
	return nil
}

func DeleteBridgeDomain(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteBridgeDomainQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
