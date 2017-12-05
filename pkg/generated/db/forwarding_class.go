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

const insertForwardingClassQuery = "insert into `forwarding_class` (`forwarding_class_dscp`,`forwarding_class_id`,`fq_name`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`forwarding_class_vlan_priority`,`forwarding_class_mpls_exp`,`uuid`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateForwardingClassQuery = "update `forwarding_class` set `forwarding_class_dscp` = ?,`forwarding_class_id` = ?,`fq_name` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`forwarding_class_vlan_priority` = ?,`forwarding_class_mpls_exp` = ?,`uuid` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?;"
const deleteForwardingClassQuery = "delete from `forwarding_class` where uuid = ?"
const listForwardingClassQuery = "select `forwarding_class`.`forwarding_class_dscp`,`forwarding_class`.`forwarding_class_id`,`forwarding_class`.`fq_name`,`forwarding_class`.`last_modified`,`forwarding_class`.`owner`,`forwarding_class`.`owner_access`,`forwarding_class`.`other_access`,`forwarding_class`.`group`,`forwarding_class`.`group_access`,`forwarding_class`.`enable`,`forwarding_class`.`description`,`forwarding_class`.`created`,`forwarding_class`.`creator`,`forwarding_class`.`user_visible`,`forwarding_class`.`display_name`,`forwarding_class`.`key_value_pair`,`forwarding_class`.`forwarding_class_vlan_priority`,`forwarding_class`.`forwarding_class_mpls_exp`,`forwarding_class`.`uuid`,`forwarding_class`.`perms2_owner_access`,`forwarding_class`.`global_access`,`forwarding_class`.`share`,`forwarding_class`.`perms2_owner` from `forwarding_class`"
const showForwardingClassQuery = "select `forwarding_class`.`forwarding_class_dscp`,`forwarding_class`.`forwarding_class_id`,`forwarding_class`.`fq_name`,`forwarding_class`.`last_modified`,`forwarding_class`.`owner`,`forwarding_class`.`owner_access`,`forwarding_class`.`other_access`,`forwarding_class`.`group`,`forwarding_class`.`group_access`,`forwarding_class`.`enable`,`forwarding_class`.`description`,`forwarding_class`.`created`,`forwarding_class`.`creator`,`forwarding_class`.`user_visible`,`forwarding_class`.`display_name`,`forwarding_class`.`key_value_pair`,`forwarding_class`.`forwarding_class_vlan_priority`,`forwarding_class`.`forwarding_class_mpls_exp`,`forwarding_class`.`uuid`,`forwarding_class`.`perms2_owner_access`,`forwarding_class`.`global_access`,`forwarding_class`.`share`,`forwarding_class`.`perms2_owner` from `forwarding_class` where uuid = ?"

const insertForwardingClassQosQueueQuery = "insert into `ref_forwarding_class_qos_queue` (`from`, `to` ) values (?, ?);"

func CreateForwardingClass(tx *sql.Tx, model *models.ForwardingClass) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertForwardingClassQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.ForwardingClassDSCP),
		int(model.ForwardingClassID),
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
		bool(model.IDPerms.UserVisible),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.ForwardingClassVlanPriority),
		int(model.ForwardingClassMPLSExp),
		string(model.UUID),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner))

	stmtQosQueueRef, err := tx.Prepare(insertForwardingClassQosQueueQuery)
	if err != nil {
		return err
	}
	defer stmtQosQueueRef.Close()
	for _, ref := range model.QosQueueRefs {
		_, err = stmtQosQueueRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanForwardingClass(rows *sql.Rows) (*models.ForwardingClass, error) {
	m := models.MakeForwardingClass()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.ForwardingClassDSCP,
		&m.ForwardingClassID,
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
		&m.IDPerms.UserVisible,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.ForwardingClassVlanPriority,
		&m.ForwardingClassMPLSExp,
		&m.UUID,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildForwardingClassWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListForwardingClass(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ForwardingClass, error) {
	result := models.MakeForwardingClassSlice()
	whereQuery, values := buildForwardingClassWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listForwardingClassQuery)
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
		m, _ := scanForwardingClass(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowForwardingClass(tx *sql.Tx, uuid string) (*models.ForwardingClass, error) {
	rows, err := tx.Query(showForwardingClassQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanForwardingClass(rows)
	}
	return nil, nil
}

func UpdateForwardingClass(tx *sql.Tx, uuid string, model *models.ForwardingClass) error {
	return nil
}

func DeleteForwardingClass(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteForwardingClassQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
