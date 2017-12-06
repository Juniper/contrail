package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertVirtualDNSRecordQuery = "insert into `virtual_DNS_record` (`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`record_name`,`record_class`,`record_data`,`record_type`,`record_ttl_seconds`,`record_mx_preference`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualDNSRecordQuery = "update `virtual_DNS_record` set `owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`record_name` = ?,`record_class` = ?,`record_data` = ?,`record_type` = ?,`record_ttl_seconds` = ?,`record_mx_preference` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteVirtualDNSRecordQuery = "delete from `virtual_DNS_record` where uuid = ?"

// VirtualDNSRecordFields is db columns for VirtualDNSRecord
var VirtualDNSRecordFields = []string{
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"owner",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"display_name",
	"key_value_pair",
	"share",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"record_name",
	"record_class",
	"record_data",
	"record_type",
	"record_ttl_seconds",
	"record_mx_preference",
	"uuid",
	"fq_name",
}

// VirtualDNSRecordRefFields is db reference fields for VirtualDNSRecord
var VirtualDNSRecordRefFields = map[string][]string{}

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
	_, err = stmt.Exec(int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.VirtualDNSRecordData.RecordName),
		string(model.VirtualDNSRecordData.RecordClass),
		string(model.VirtualDNSRecordData.RecordData),
		string(model.VirtualDNSRecordData.RecordType),
		int(model.VirtualDNSRecordData.RecordTTLSeconds),
		int(model.VirtualDNSRecordData.RecordMXPreference),
		string(model.UUID),
		utils.MustJSON(model.FQName))
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

	if value, ok := values["owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["record_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualDNSRecordData.RecordName = castedValue

	}

	if value, ok := values["record_class"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualDNSRecordData.RecordClass = models.DnsRecordClassType(castedValue)

	}

	if value, ok := values["record_data"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualDNSRecordData.RecordData = castedValue

	}

	if value, ok := values["record_type"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualDNSRecordData.RecordType = models.DnsRecordTypeType(castedValue)

	}

	if value, ok := values["record_ttl_seconds"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.VirtualDNSRecordData.RecordTTLSeconds = castedValue

	}

	if value, ok := values["record_mx_preference"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.VirtualDNSRecordData.RecordMXPreference = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	return m, nil
}

// ListVirtualDNSRecord lists VirtualDNSRecord with list spec.
func ListVirtualDNSRecord(tx *sql.Tx, spec *db.ListSpec) ([]*models.VirtualDNSRecord, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "virtual_DNS_record"
	spec.Fields = VirtualDNSRecordFields
	spec.RefFields = VirtualDNSRecordRefFields
	result := models.MakeVirtualDNSRecordSlice()
	query, columns, values := db.BuildListQuery(spec)
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
	list, err := ListVirtualDNSRecord(tx, &db.ListSpec{
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
