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

const insertQosQueueQuery = "insert into `qos_queue` (`max_bandwidth`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`qos_queue_identifier`,`min_bandwidth`,`fq_name`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateQosQueueQuery = "update `qos_queue` set `max_bandwidth` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`qos_queue_identifier` = ?,`min_bandwidth` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteQosQueueQuery = "delete from `qos_queue` where uuid = ?"
const listQosQueueQuery = "select `qos_queue`.`max_bandwidth`,`qos_queue`.`owner`,`qos_queue`.`owner_access`,`qos_queue`.`global_access`,`qos_queue`.`share`,`qos_queue`.`uuid`,`qos_queue`.`qos_queue_identifier`,`qos_queue`.`min_bandwidth`,`qos_queue`.`fq_name`,`qos_queue`.`user_visible`,`qos_queue`.`last_modified`,`qos_queue`.`permissions_owner_access`,`qos_queue`.`other_access`,`qos_queue`.`group`,`qos_queue`.`group_access`,`qos_queue`.`permissions_owner`,`qos_queue`.`enable`,`qos_queue`.`description`,`qos_queue`.`created`,`qos_queue`.`creator`,`qos_queue`.`display_name`,`qos_queue`.`key_value_pair` from `qos_queue`"
const showQosQueueQuery = "select `qos_queue`.`max_bandwidth`,`qos_queue`.`owner`,`qos_queue`.`owner_access`,`qos_queue`.`global_access`,`qos_queue`.`share`,`qos_queue`.`uuid`,`qos_queue`.`qos_queue_identifier`,`qos_queue`.`min_bandwidth`,`qos_queue`.`fq_name`,`qos_queue`.`user_visible`,`qos_queue`.`last_modified`,`qos_queue`.`permissions_owner_access`,`qos_queue`.`other_access`,`qos_queue`.`group`,`qos_queue`.`group_access`,`qos_queue`.`permissions_owner`,`qos_queue`.`enable`,`qos_queue`.`description`,`qos_queue`.`created`,`qos_queue`.`creator`,`qos_queue`.`display_name`,`qos_queue`.`key_value_pair` from `qos_queue` where uuid = ?"

func CreateQosQueue(tx *sql.Tx, model *models.QosQueue) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertQosQueueQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.MaxBandwidth),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID),
		int(model.QosQueueIdentifier),
		int(model.MinBandwidth),
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
		utils.MustJSON(model.Annotations.KeyValuePair))

	return err
}

func scanQosQueue(rows *sql.Rows) (*models.QosQueue, error) {
	m := models.MakeQosQueue()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.MaxBandwidth,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&m.QosQueueIdentifier,
		&m.MinBandwidth,
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
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildQosQueueWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListQosQueue(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.QosQueue, error) {
	result := models.MakeQosQueueSlice()
	whereQuery, values := buildQosQueueWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listQosQueueQuery)
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
		m, _ := scanQosQueue(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowQosQueue(tx *sql.Tx, uuid string) (*models.QosQueue, error) {
	rows, err := tx.Query(showQosQueueQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanQosQueue(rows)
	}
	return nil, nil
}

func UpdateQosQueue(tx *sql.Tx, uuid string, model *models.QosQueue) error {
	return nil
}

func DeleteQosQueue(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteQosQueueQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
