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

const insertVirtualDNSQuery = "insert into `virtual_DNS` (`next_virtual_DNS`,`dynamic_records_from_client`,`reverse_resolution`,`default_ttl_seconds`,`record_order`,`floating_ip_record`,`domain_name`,`external_visible`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualDNSQuery = "update `virtual_DNS` set `next_virtual_DNS` = ?,`dynamic_records_from_client` = ?,`reverse_resolution` = ?,`default_ttl_seconds` = ?,`record_order` = ?,`floating_ip_record` = ?,`domain_name` = ?,`external_visible` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteVirtualDNSQuery = "delete from `virtual_DNS` where uuid = ?"
const listVirtualDNSQuery = "select `virtual_DNS`.`next_virtual_DNS`,`virtual_DNS`.`dynamic_records_from_client`,`virtual_DNS`.`reverse_resolution`,`virtual_DNS`.`default_ttl_seconds`,`virtual_DNS`.`record_order`,`virtual_DNS`.`floating_ip_record`,`virtual_DNS`.`domain_name`,`virtual_DNS`.`external_visible`,`virtual_DNS`.`created`,`virtual_DNS`.`creator`,`virtual_DNS`.`user_visible`,`virtual_DNS`.`last_modified`,`virtual_DNS`.`other_access`,`virtual_DNS`.`group`,`virtual_DNS`.`group_access`,`virtual_DNS`.`owner`,`virtual_DNS`.`owner_access`,`virtual_DNS`.`enable`,`virtual_DNS`.`description`,`virtual_DNS`.`display_name`,`virtual_DNS`.`key_value_pair`,`virtual_DNS`.`share`,`virtual_DNS`.`perms2_owner`,`virtual_DNS`.`perms2_owner_access`,`virtual_DNS`.`global_access`,`virtual_DNS`.`uuid`,`virtual_DNS`.`fq_name` from `virtual_DNS`"
const showVirtualDNSQuery = "select `virtual_DNS`.`next_virtual_DNS`,`virtual_DNS`.`dynamic_records_from_client`,`virtual_DNS`.`reverse_resolution`,`virtual_DNS`.`default_ttl_seconds`,`virtual_DNS`.`record_order`,`virtual_DNS`.`floating_ip_record`,`virtual_DNS`.`domain_name`,`virtual_DNS`.`external_visible`,`virtual_DNS`.`created`,`virtual_DNS`.`creator`,`virtual_DNS`.`user_visible`,`virtual_DNS`.`last_modified`,`virtual_DNS`.`other_access`,`virtual_DNS`.`group`,`virtual_DNS`.`group_access`,`virtual_DNS`.`owner`,`virtual_DNS`.`owner_access`,`virtual_DNS`.`enable`,`virtual_DNS`.`description`,`virtual_DNS`.`display_name`,`virtual_DNS`.`key_value_pair`,`virtual_DNS`.`share`,`virtual_DNS`.`perms2_owner`,`virtual_DNS`.`perms2_owner_access`,`virtual_DNS`.`global_access`,`virtual_DNS`.`uuid`,`virtual_DNS`.`fq_name` from `virtual_DNS` where uuid = ?"

func CreateVirtualDNS(tx *sql.Tx, model *models.VirtualDNS) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualDNSQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.VirtualDNSData.NextVirtualDNS),
		bool(model.VirtualDNSData.DynamicRecordsFromClient),
		bool(model.VirtualDNSData.ReverseResolution),
		int(model.VirtualDNSData.DefaultTTLSeconds),
		string(model.VirtualDNSData.RecordOrder),
		string(model.VirtualDNSData.FloatingIPRecord),
		string(model.VirtualDNSData.DomainName),
		bool(model.VirtualDNSData.ExternalVisible),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		utils.MustJSON(model.FQName))

	return err
}

func scanVirtualDNS(rows *sql.Rows) (*models.VirtualDNS, error) {
	m := models.MakeVirtualDNS()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.VirtualDNSData.NextVirtualDNS,
		&m.VirtualDNSData.DynamicRecordsFromClient,
		&m.VirtualDNSData.ReverseResolution,
		&m.VirtualDNSData.DefaultTTLSeconds,
		&m.VirtualDNSData.RecordOrder,
		&m.VirtualDNSData.FloatingIPRecord,
		&m.VirtualDNSData.DomainName,
		&m.VirtualDNSData.ExternalVisible,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildVirtualDNSWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["next_virtual_DNS"]; ok {
		results = append(results, "next_virtual_DNS = ?")
		values = append(values, value)
	}

	if value, ok := where["record_order"]; ok {
		results = append(results, "record_order = ?")
		values = append(values, value)
	}

	if value, ok := where["floating_ip_record"]; ok {
		results = append(results, "floating_ip_record = ?")
		values = append(values, value)
	}

	if value, ok := where["domain_name"]; ok {
		results = append(results, "domain_name = ?")
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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListVirtualDNS(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.VirtualDNS, error) {
	result := models.MakeVirtualDNSSlice()
	whereQuery, values := buildVirtualDNSWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listVirtualDNSQuery)
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
		m, _ := scanVirtualDNS(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowVirtualDNS(tx *sql.Tx, uuid string) (*models.VirtualDNS, error) {
	rows, err := tx.Query(showVirtualDNSQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanVirtualDNS(rows)
	}
	return nil, nil
}

func UpdateVirtualDNS(tx *sql.Tx, uuid string, model *models.VirtualDNS) error {
	return nil
}

func DeleteVirtualDNS(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualDNSQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
