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

const insertQosConfigQuery = "insert into `qos_config` (`default_forwarding_class_id`,`dscp_entries`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`vlan_priority_entries`,`mpls_exp_entries`,`uuid`,`fq_name`,`display_name`,`qos_config_type`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateQosConfigQuery = "update `qos_config` set `default_forwarding_class_id` = ?,`dscp_entries` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`vlan_priority_entries` = ?,`mpls_exp_entries` = ?,`uuid` = ?,`fq_name` = ?,`display_name` = ?,`qos_config_type` = ?;"
const deleteQosConfigQuery = "delete from `qos_config` where uuid = ?"
const listQosConfigQuery = "select `qos_config`.`default_forwarding_class_id`,`qos_config`.`dscp_entries`,`qos_config`.`creator`,`qos_config`.`user_visible`,`qos_config`.`last_modified`,`qos_config`.`owner_access`,`qos_config`.`other_access`,`qos_config`.`group`,`qos_config`.`group_access`,`qos_config`.`owner`,`qos_config`.`enable`,`qos_config`.`description`,`qos_config`.`created`,`qos_config`.`key_value_pair`,`qos_config`.`global_access`,`qos_config`.`share`,`qos_config`.`perms2_owner`,`qos_config`.`perms2_owner_access`,`qos_config`.`vlan_priority_entries`,`qos_config`.`mpls_exp_entries`,`qos_config`.`uuid`,`qos_config`.`fq_name`,`qos_config`.`display_name`,`qos_config`.`qos_config_type` from `qos_config`"
const showQosConfigQuery = "select `qos_config`.`default_forwarding_class_id`,`qos_config`.`dscp_entries`,`qos_config`.`creator`,`qos_config`.`user_visible`,`qos_config`.`last_modified`,`qos_config`.`owner_access`,`qos_config`.`other_access`,`qos_config`.`group`,`qos_config`.`group_access`,`qos_config`.`owner`,`qos_config`.`enable`,`qos_config`.`description`,`qos_config`.`created`,`qos_config`.`key_value_pair`,`qos_config`.`global_access`,`qos_config`.`share`,`qos_config`.`perms2_owner`,`qos_config`.`perms2_owner_access`,`qos_config`.`vlan_priority_entries`,`qos_config`.`mpls_exp_entries`,`qos_config`.`uuid`,`qos_config`.`fq_name`,`qos_config`.`display_name`,`qos_config`.`qos_config_type` from `qos_config` where uuid = ?"

const insertQosConfigGlobalSystemConfigQuery = "insert into `ref_qos_config_global_system_config` (`from`, `to` ) values (?, ?);"

func CreateQosConfig(tx *sql.Tx, model *models.QosConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertQosConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.DefaultForwardingClassID),
		utils.MustJSON(model.DSCPEntries),
		string(model.IDPerms.Creator),
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
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		utils.MustJSON(model.VlanPriorityEntries),
		utils.MustJSON(model.MPLSExpEntries),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.DisplayName),
		string(model.QosConfigType))

	stmtGlobalSystemConfigRef, err := tx.Prepare(insertQosConfigGlobalSystemConfigQuery)
	if err != nil {
		return err
	}
	defer stmtGlobalSystemConfigRef.Close()
	for _, ref := range model.GlobalSystemConfigRefs {
		_, err = stmtGlobalSystemConfigRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanQosConfig(rows *sql.Rows) (*models.QosConfig, error) {
	m := models.MakeQosConfig()

	var jsonDSCPEntries string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonVlanPriorityEntries string

	var jsonMPLSExpEntries string

	var jsonFQName string

	if err := rows.Scan(&m.DefaultForwardingClassID,
		&jsonDSCPEntries,
		&m.IDPerms.Creator,
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
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&jsonVlanPriorityEntries,
		&jsonMPLSExpEntries,
		&m.UUID,
		&jsonFQName,
		&m.DisplayName,
		&m.QosConfigType); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonDSCPEntries), &m.DSCPEntries)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonVlanPriorityEntries), &m.VlanPriorityEntries)

	json.Unmarshal([]byte(jsonMPLSExpEntries), &m.MPLSExpEntries)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildQosConfigWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
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

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["qos_config_type"]; ok {
		results = append(results, "qos_config_type = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListQosConfig(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.QosConfig, error) {
	result := models.MakeQosConfigSlice()
	whereQuery, values := buildQosConfigWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listQosConfigQuery)
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
		m, _ := scanQosConfig(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowQosConfig(tx *sql.Tx, uuid string) (*models.QosConfig, error) {
	rows, err := tx.Query(showQosConfigQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanQosConfig(rows)
	}
	return nil, nil
}

func UpdateQosConfig(tx *sql.Tx, uuid string, model *models.QosConfig) error {
	return nil
}

func DeleteQosConfig(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteQosConfigQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
