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

const insertVirtualDNSQuery = "insert into `virtual_DNS` (`external_visible`,`next_virtual_DNS`,`dynamic_records_from_client`,`reverse_resolution`,`default_ttl_seconds`,`record_order`,`floating_ip_record`,`domain_name`,`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualDNSQuery = "update `virtual_DNS` set `external_visible` = ?,`next_virtual_DNS` = ?,`dynamic_records_from_client` = ?,`reverse_resolution` = ?,`default_ttl_seconds` = ?,`record_order` = ?,`floating_ip_record` = ?,`domain_name` = ?,`uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?;"
const deleteVirtualDNSQuery = "delete from `virtual_DNS` where uuid = ?"

// VirtualDNSFields is db columns for VirtualDNS
var VirtualDNSFields = []string{
	"external_visible",
	"next_virtual_DNS",
	"dynamic_records_from_client",
	"reverse_resolution",
	"default_ttl_seconds",
	"record_order",
	"floating_ip_record",
	"domain_name",
	"uuid",
	"fq_name",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"other_access",
	"group",
	"group_access",
	"owner",
	"owner_access",
	"enable",
	"display_name",
	"key_value_pair",
	"global_access",
	"share",
	"perms2_owner",
	"perms2_owner_access",
}

// VirtualDNSRefFields is db reference fields for VirtualDNS
var VirtualDNSRefFields = map[string][]string{}

// CreateVirtualDNS inserts VirtualDNS to DB
func CreateVirtualDNS(tx *sql.Tx, model *models.VirtualDNS) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualDNSQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertVirtualDNSQuery,
	}).Debug("create query")
	_, err = stmt.Exec(bool(model.VirtualDNSData.ExternalVisible),
		string(model.VirtualDNSData.NextVirtualDNS),
		bool(model.VirtualDNSData.DynamicRecordsFromClient),
		bool(model.VirtualDNSData.ReverseResolution),
		int(model.VirtualDNSData.DefaultTTLSeconds),
		string(model.VirtualDNSData.RecordOrder),
		string(model.VirtualDNSData.FloatingIPRecord),
		string(model.VirtualDNSData.DomainName),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.Description),
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
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanVirtualDNS(values map[string]interface{}) (*models.VirtualDNS, error) {
	m := models.MakeVirtualDNS()

	if value, ok := values["external_visible"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.VirtualDNSData.ExternalVisible = castedValue

	}

	if value, ok := values["next_virtual_DNS"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualDNSData.NextVirtualDNS = castedValue

	}

	if value, ok := values["dynamic_records_from_client"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.VirtualDNSData.DynamicRecordsFromClient = castedValue

	}

	if value, ok := values["reverse_resolution"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.VirtualDNSData.ReverseResolution = castedValue

	}

	if value, ok := values["default_ttl_seconds"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.VirtualDNSData.DefaultTTLSeconds = castedValue

	}

	if value, ok := values["record_order"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualDNSData.RecordOrder = models.DnsRecordOrderType(castedValue)

	}

	if value, ok := values["floating_ip_record"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualDNSData.FloatingIPRecord = models.FloatingIpDnsNotation(castedValue)

	}

	if value, ok := values["domain_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VirtualDNSData.DomainName = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

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

	if value, ok := values["owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["enable"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

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

	return m, nil
}

// ListVirtualDNS lists VirtualDNS with list spec.
func ListVirtualDNS(tx *sql.Tx, spec *db.ListSpec) ([]*models.VirtualDNS, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "virtual_DNS"
	spec.Fields = VirtualDNSFields
	spec.RefFields = VirtualDNSRefFields
	result := models.MakeVirtualDNSSlice()
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
		m, err := scanVirtualDNS(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowVirtualDNS shows VirtualDNS resource
func ShowVirtualDNS(tx *sql.Tx, uuid string) (*models.VirtualDNS, error) {
	list, err := ListVirtualDNS(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateVirtualDNS updates a resource
func UpdateVirtualDNS(tx *sql.Tx, uuid string, model *models.VirtualDNS) error {
	//TODO(nati) support update
	return nil
}

// DeleteVirtualDNS deletes a resource
func DeleteVirtualDNS(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualDNSQuery)
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
