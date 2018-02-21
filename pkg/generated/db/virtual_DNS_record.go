package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertVirtualDNSRecordQuery = "insert into `virtual_DNS_record` (`record_type`,`record_ttl_seconds`,`record_name`,`record_mx_preference`,`record_data`,`record_class`,`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
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

// VirtualDNSRecordParentTypes is possible parents for VirtualDNSRecord
var VirtualDNSRecordParents = []string{

	"virtual_DNS",
}

// CreateVirtualDNSRecord inserts VirtualDNSRecord to DB
func CreateVirtualDNSRecord(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateVirtualDNSRecordRequest) error {
	model := request.VirtualDNSRecord
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
	_, err = stmt.ExecContext(ctx, string(model.GetVirtualDNSRecordData().GetRecordType()),
		int(model.GetVirtualDNSRecordData().GetRecordTTLSeconds()),
		string(model.GetVirtualDNSRecordData().GetRecordName()),
		int(model.GetVirtualDNSRecordData().GetRecordMXPreference()),
		string(model.GetVirtualDNSRecordData().GetRecordData()),
		string(model.GetVirtualDNSRecordData().GetRecordClass()),
		string(model.GetUUID()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		bool(model.GetIDPerms().GetUserVisible()),
		int(model.GetIDPerms().GetPermissions().GetOwnerAccess()),
		string(model.GetIDPerms().GetPermissions().GetOwner()),
		int(model.GetIDPerms().GetPermissions().GetOtherAccess()),
		int(model.GetIDPerms().GetPermissions().GetGroupAccess()),
		string(model.GetIDPerms().GetPermissions().GetGroup()),
		string(model.GetIDPerms().GetLastModified()),
		bool(model.GetIDPerms().GetEnable()),
		string(model.GetIDPerms().GetDescription()),
		string(model.GetIDPerms().GetCreator()),
		string(model.GetIDPerms().GetCreated()),
		common.MustJSON(model.GetFQName()),
		string(model.GetDisplayName()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "virtual_DNS_record",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "virtual_DNS_record", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanVirtualDNSRecord(values map[string]interface{}) (*models.VirtualDNSRecord, error) {
	m := models.MakeVirtualDNSRecord()

	if value, ok := values["record_type"]; ok {

		m.VirtualDNSRecordData.RecordType = common.InterfaceToString(value)

	}

	if value, ok := values["record_ttl_seconds"]; ok {

		m.VirtualDNSRecordData.RecordTTLSeconds = common.InterfaceToInt64(value)

	}

	if value, ok := values["record_name"]; ok {

		m.VirtualDNSRecordData.RecordName = common.InterfaceToString(value)

	}

	if value, ok := values["record_mx_preference"]; ok {

		m.VirtualDNSRecordData.RecordMXPreference = common.InterfaceToInt64(value)

	}

	if value, ok := values["record_data"]; ok {

		m.VirtualDNSRecordData.RecordData = common.InterfaceToString(value)

	}

	if value, ok := values["record_class"]; ok {

		m.VirtualDNSRecordData.RecordClass = common.InterfaceToString(value)

	}

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		m.Perms2.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["owner"]; ok {

		m.Perms2.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["global_access"]; ok {

		m.Perms2.GlobalAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = common.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = common.InterfaceToString(value)

	}

	if value, ok := values["user_visible"]; ok {

		m.IDPerms.UserVisible = common.InterfaceToBool(value)

	}

	if value, ok := values["permissions_owner_access"]; ok {

		m.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["permissions_owner"]; ok {

		m.IDPerms.Permissions.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["other_access"]; ok {

		m.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group_access"]; ok {

		m.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group"]; ok {

		m.IDPerms.Permissions.Group = common.InterfaceToString(value)

	}

	if value, ok := values["last_modified"]; ok {

		m.IDPerms.LastModified = common.InterfaceToString(value)

	}

	if value, ok := values["enable"]; ok {

		m.IDPerms.Enable = common.InterfaceToBool(value)

	}

	if value, ok := values["description"]; ok {

		m.IDPerms.Description = common.InterfaceToString(value)

	}

	if value, ok := values["creator"]; ok {

		m.IDPerms.Creator = common.InterfaceToString(value)

	}

	if value, ok := values["created"]; ok {

		m.IDPerms.Created = common.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListVirtualDNSRecord lists VirtualDNSRecord with list spec.
func ListVirtualDNSRecord(ctx context.Context, tx *sql.Tx, request *models.ListVirtualDNSRecordRequest) (response *models.ListVirtualDNSRecordResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "virtual_DNS_record"
	qb.Fields = VirtualDNSRecordFields
	qb.RefFields = VirtualDNSRecordRefFields
	qb.BackRefFields = VirtualDNSRecordBackRefFields
	result := []*models.VirtualDNSRecord{}

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = common.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}

	query := qb.BuildQuery()
	columns := qb.Columns
	values := qb.Values
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.QueryContext(ctx, query, values...)
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
		m, err := scanVirtualDNSRecord(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListVirtualDNSRecordResponse{
		VirtualDNSRecords: result,
	}
	return response, nil
}

// UpdateVirtualDNSRecord updates a resource
func UpdateVirtualDNSRecord(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateVirtualDNSRecordRequest,
) error {
	//TODO
	return nil
}

// DeleteVirtualDNSRecord deletes a resource
func DeleteVirtualDNSRecord(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteVirtualDNSRecordRequest) error {
	deleteQuery := deleteVirtualDNSRecordQuery
	selectQuery := "select count(uuid) from virtual_DNS_record where uuid = ?"
	var err error
	var count int
	uuid := request.ID
	auth := common.GetAuthCTX(ctx)
	if auth.IsAdmin() {
		row := tx.QueryRowContext(ctx, selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	err = common.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}
