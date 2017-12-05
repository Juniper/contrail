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

const insertVirtualDNSRecordQuery = "insert into `virtual_DNS_record` (`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`display_name`,`record_name`,`record_class`,`record_data`,`record_type`,`record_ttl_seconds`,`record_mx_preference`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualDNSRecordQuery = "update `virtual_DNS_record` set `uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`record_name` = ?,`record_class` = ?,`record_data` = ?,`record_type` = ?,`record_ttl_seconds` = ?,`record_mx_preference` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?;"
const deleteVirtualDNSRecordQuery = "delete from `virtual_DNS_record` where uuid = ?"
const listVirtualDNSRecordQuery = "select `virtual_DNS_record`.`uuid`,`virtual_DNS_record`.`fq_name`,`virtual_DNS_record`.`created`,`virtual_DNS_record`.`creator`,`virtual_DNS_record`.`user_visible`,`virtual_DNS_record`.`last_modified`,`virtual_DNS_record`.`owner`,`virtual_DNS_record`.`owner_access`,`virtual_DNS_record`.`other_access`,`virtual_DNS_record`.`group`,`virtual_DNS_record`.`group_access`,`virtual_DNS_record`.`enable`,`virtual_DNS_record`.`description`,`virtual_DNS_record`.`display_name`,`virtual_DNS_record`.`record_name`,`virtual_DNS_record`.`record_class`,`virtual_DNS_record`.`record_data`,`virtual_DNS_record`.`record_type`,`virtual_DNS_record`.`record_ttl_seconds`,`virtual_DNS_record`.`record_mx_preference`,`virtual_DNS_record`.`key_value_pair`,`virtual_DNS_record`.`global_access`,`virtual_DNS_record`.`share`,`virtual_DNS_record`.`perms2_owner`,`virtual_DNS_record`.`perms2_owner_access` from `virtual_DNS_record`"
const showVirtualDNSRecordQuery = "select `virtual_DNS_record`.`uuid`,`virtual_DNS_record`.`fq_name`,`virtual_DNS_record`.`created`,`virtual_DNS_record`.`creator`,`virtual_DNS_record`.`user_visible`,`virtual_DNS_record`.`last_modified`,`virtual_DNS_record`.`owner`,`virtual_DNS_record`.`owner_access`,`virtual_DNS_record`.`other_access`,`virtual_DNS_record`.`group`,`virtual_DNS_record`.`group_access`,`virtual_DNS_record`.`enable`,`virtual_DNS_record`.`description`,`virtual_DNS_record`.`display_name`,`virtual_DNS_record`.`record_name`,`virtual_DNS_record`.`record_class`,`virtual_DNS_record`.`record_data`,`virtual_DNS_record`.`record_type`,`virtual_DNS_record`.`record_ttl_seconds`,`virtual_DNS_record`.`record_mx_preference`,`virtual_DNS_record`.`key_value_pair`,`virtual_DNS_record`.`global_access`,`virtual_DNS_record`.`share`,`virtual_DNS_record`.`perms2_owner`,`virtual_DNS_record`.`perms2_owner_access` from `virtual_DNS_record` where uuid = ?"

func CreateVirtualDNSRecord(tx *sql.Tx, model *models.VirtualDNSRecord) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualDNSRecordQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		utils.MustJSON(model.FQName),
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
		string(model.DisplayName),
		string(model.VirtualDNSRecordData.RecordName),
		string(model.VirtualDNSRecordData.RecordClass),
		string(model.VirtualDNSRecordData.RecordData),
		string(model.VirtualDNSRecordData.RecordType),
		int(model.VirtualDNSRecordData.RecordTTLSeconds),
		int(model.VirtualDNSRecordData.RecordMXPreference),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess))

	return err
}

func scanVirtualDNSRecord(rows *sql.Rows) (*models.VirtualDNSRecord, error) {
	m := models.MakeVirtualDNSRecord()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
		&jsonFQName,
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
		&m.DisplayName,
		&m.VirtualDNSRecordData.RecordName,
		&m.VirtualDNSRecordData.RecordClass,
		&m.VirtualDNSRecordData.RecordData,
		&m.VirtualDNSRecordData.RecordType,
		&m.VirtualDNSRecordData.RecordTTLSeconds,
		&m.VirtualDNSRecordData.RecordMXPreference,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildVirtualDNSRecordWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["record_name"]; ok {
		results = append(results, "record_name = ?")
		values = append(values, value)
	}

	if value, ok := where["record_class"]; ok {
		results = append(results, "record_class = ?")
		values = append(values, value)
	}

	if value, ok := where["record_data"]; ok {
		results = append(results, "record_data = ?")
		values = append(values, value)
	}

	if value, ok := where["record_type"]; ok {
		results = append(results, "record_type = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListVirtualDNSRecord(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.VirtualDNSRecord, error) {
	result := models.MakeVirtualDNSRecordSlice()
	whereQuery, values := buildVirtualDNSRecordWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listVirtualDNSRecordQuery)
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
		m, _ := scanVirtualDNSRecord(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowVirtualDNSRecord(tx *sql.Tx, uuid string) (*models.VirtualDNSRecord, error) {
	rows, err := tx.Query(showVirtualDNSRecordQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanVirtualDNSRecord(rows)
	}
	return nil, nil
}

func UpdateVirtualDNSRecord(tx *sql.Tx, uuid string, model *models.VirtualDNSRecord) error {
	return nil
}

func DeleteVirtualDNSRecord(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualDNSRecordQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
