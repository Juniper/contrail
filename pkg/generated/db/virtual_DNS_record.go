package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertVirtualDNSRecordQuery = "insert into `virtual_DNS_record` (`record_type`,`record_ttl_seconds`,`record_name`,`record_mx_preference`,`record_data`,`record_class`,`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualDNSRecordQuery = "update `virtual_DNS_record` set `record_type` = ?,`record_ttl_seconds` = ?,`record_name` = ?,`record_mx_preference` = ?,`record_data` = ?,`record_class` = ?,`uuid` = ?,`share` = ?,`owner_access` = ?,`owner` = ?,`global_access` = ?,`parent_uuid` = ?,`parent_type` = ?,`user_visible` = ?,`permissions_owner_access` = ?,`permissions_owner` = ?,`other_access` = ?,`group_access` = ?,`group` = ?,`last_modified` = ?,`enable` = ?,`description` = ?,`creator` = ?,`created` = ?,`fq_name` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteVirtualDNSRecordQuery = "delete from `virtual_DNS_record` where uuid = ?"

// VirtualDNSRecordFields is db columns for VirtualDNSRecord
var VirtualDNSRecordFields = []string{
	"record_type",
	"record_ttl_seconds",
	"record_name",
	"record_mx_preference",
	"record_data",
	"record_class",
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"user_visible",
	"permissions_owner_access",
	"permissions_owner",
	"other_access",
	"group_access",
	"group",
	"last_modified",
	"enable",
	"description",
	"creator",
	"created",
	"fq_name",
	"display_name",
	"key_value_pair",
}

// VirtualDNSRecordRefFields is db reference fields for VirtualDNSRecord
var VirtualDNSRecordRefFields = map[string][]string{}

// VirtualDNSRecordBackRefFields is db back reference fields for VirtualDNSRecord
var VirtualDNSRecordBackRefFields = map[string][]string{}

// CreateVirtualDNSRecord inserts VirtualDNSRecord to DB
func CreateVirtualDNSRecord(tx *sql.Tx, model *models.VirtualDNSRecord) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualDNSRecordQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertVirtualDNSRecordQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.VirtualDNSRecordData.RecordType),
		int(model.VirtualDNSRecordData.RecordTTLSeconds),
		string(model.VirtualDNSRecordData.RecordName),
		int(model.VirtualDNSRecordData.RecordMXPreference),
		string(model.VirtualDNSRecordData.RecordData),
		string(model.VirtualDNSRecordData.RecordClass),
		string(model.UUID),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		bool(model.IDPerms.UserVisible),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OtherAccess),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Group),
		string(model.IDPerms.LastModified),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Creator),
		string(model.IDPerms.Created),
		common.MustJSON(model.FQName),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanVirtualDNSRecord(values map[string]interface{}) (*models.VirtualDNSRecord, error) {
	m := models.MakeVirtualDNSRecord()

	if value, ok := values["record_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualDNSRecordData.RecordType = models.DnsRecordTypeType(castedValue)

	}

	if value, ok := values["record_ttl_seconds"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualDNSRecordData.RecordTTLSeconds = castedValue

	}

	if value, ok := values["record_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualDNSRecordData.RecordName = castedValue

	}

	if value, ok := values["record_mx_preference"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualDNSRecordData.RecordMXPreference = castedValue

	}

	if value, ok := values["record_data"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualDNSRecordData.RecordData = castedValue

	}

	if value, ok := values["record_class"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualDNSRecordData.RecordClass = models.DnsRecordClassType(castedValue)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["parent_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentUUID = castedValue

	}

	if value, ok := values["parent_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentType = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListVirtualDNSRecord lists VirtualDNSRecord with list spec.
func ListVirtualDNSRecord(tx *sql.Tx, spec *common.ListSpec) ([]*models.VirtualDNSRecord, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "virtual_DNS_record"
	spec.Fields = VirtualDNSRecordFields
	spec.RefFields = VirtualDNSRecordRefFields
	spec.BackRefFields = VirtualDNSRecordBackRefFields
	result := models.MakeVirtualDNSRecordSlice()
	query, columns, values := common.BuildListQuery(spec)
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.Query(query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "select query failed")
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "row error")
	}
	for rows.Next() {
		valuesMap := map[string]interface{}{}
		values := make([]interface{}, len(columns))
		valuesPointers := make([]interface{}, len(columns))
		for _, index := range columns {
			valuesPointers[index] = &values[index]
		}
		if err := rows.Scan(valuesPointers...); err != nil {
			return nil, errors.Wrap(err, "scan failed")
		}
		for column, index := range columns {
			val := valuesPointers[index].(*interface{})
			valuesMap[column] = *val
		}
		log.WithFields(log.Fields{
			"valuesMap": valuesMap,
		}).Debug("valueMap")
		m, err := scanVirtualDNSRecord(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowVirtualDNSRecord shows VirtualDNSRecord resource
func ShowVirtualDNSRecord(tx *sql.Tx, uuid string) (*models.VirtualDNSRecord, error) {
	list, err := ListVirtualDNSRecord(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateVirtualDNSRecord updates a resource
func UpdateVirtualDNSRecord(tx *sql.Tx, uuid string, model *models.VirtualDNSRecord) error {
	//TODO(nati) support update
	return nil
}

// DeleteVirtualDNSRecord deletes a resource
func DeleteVirtualDNSRecord(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualDNSRecordQuery)
	if err != nil {
		return errors.Wrap(err, "preparing delete query failed")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	if err != nil {
		return errors.Wrap(err, "delete failed")
	}
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}
