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

const insertAlarmQuery = "insert into `alarm` (`uve_key`,`alarm_severity`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`key_value_pair`,`alarm_rules`,`fq_name`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAlarmQuery = "update `alarm` set `uve_key` = ?,`alarm_severity` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`key_value_pair` = ?,`alarm_rules` = ?,`fq_name` = ?,`display_name` = ?;"
const deleteAlarmQuery = "delete from `alarm` where uuid = ?"
const listAlarmQuery = "select `alarm`.`uve_key`,`alarm`.`alarm_severity`,`alarm`.`global_access`,`alarm`.`share`,`alarm`.`owner`,`alarm`.`owner_access`,`alarm`.`uuid`,`alarm`.`created`,`alarm`.`creator`,`alarm`.`user_visible`,`alarm`.`last_modified`,`alarm`.`permissions_owner`,`alarm`.`permissions_owner_access`,`alarm`.`other_access`,`alarm`.`group`,`alarm`.`group_access`,`alarm`.`enable`,`alarm`.`description`,`alarm`.`key_value_pair`,`alarm`.`alarm_rules`,`alarm`.`fq_name`,`alarm`.`display_name` from `alarm`"
const showAlarmQuery = "select `alarm`.`uve_key`,`alarm`.`alarm_severity`,`alarm`.`global_access`,`alarm`.`share`,`alarm`.`owner`,`alarm`.`owner_access`,`alarm`.`uuid`,`alarm`.`created`,`alarm`.`creator`,`alarm`.`user_visible`,`alarm`.`last_modified`,`alarm`.`permissions_owner`,`alarm`.`permissions_owner_access`,`alarm`.`other_access`,`alarm`.`group`,`alarm`.`group_access`,`alarm`.`enable`,`alarm`.`description`,`alarm`.`key_value_pair`,`alarm`.`alarm_rules`,`alarm`.`fq_name`,`alarm`.`display_name` from `alarm` where uuid = ?"

func CreateAlarm(tx *sql.Tx, model *models.Alarm) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAlarmQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.UveKeys.UveKey),
		int(model.AlarmSeverity),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.AlarmRules),
		utils.MustJSON(model.FQName),
		string(model.DisplayName))

	return err
}

func scanAlarm(rows *sql.Rows) (*models.Alarm, error) {
	m := models.MakeAlarm()

	var jsonUveKeysUveKey string

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	var jsonAlarmRules string

	var jsonFQName string

	if err := rows.Scan(&jsonUveKeysUveKey,
		&m.AlarmSeverity,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.UUID,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&jsonAnnotationsKeyValuePair,
		&jsonAlarmRules,
		&jsonFQName,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonUveKeysUveKey), &m.UveKeys.UveKey)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonAlarmRules), &m.AlarmRules)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildAlarmWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListAlarm(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Alarm, error) {
	result := models.MakeAlarmSlice()
	whereQuery, values := buildAlarmWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listAlarmQuery)
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
		m, _ := scanAlarm(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAlarm(tx *sql.Tx, uuid string) (*models.Alarm, error) {
	rows, err := tx.Query(showAlarmQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAlarm(rows)
	}
	return nil, nil
}

func UpdateAlarm(tx *sql.Tx, uuid string, model *models.Alarm) error {
	return nil
}

func DeleteAlarm(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAlarmQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
