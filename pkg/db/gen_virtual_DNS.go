// nolint
package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

// VirtualDNSFields is db columns for VirtualDNS
var VirtualDNSFields = []string{
	"reverse_resolution",
	"record_order",
	"next_virtual_DNS",
	"floating_ip_record",
	"external_visible",
	"dynamic_records_from_client",
	"domain_name",
	"default_ttl_seconds",
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
	"configuration_version",
	"key_value_pair",
}

// VirtualDNSRefFields is db reference fields for VirtualDNS
var VirtualDNSRefFields = map[string][]string{}

// VirtualDNSBackRefFields is db back reference fields for VirtualDNS
var VirtualDNSBackRefFields = map[string][]string{

	"virtual_DNS_record": []string{
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
		"configuration_version",
		"key_value_pair",
	},
}

// VirtualDNSParentTypes is possible parents for VirtualDNS
var VirtualDNSParents = []string{

	"domain",
}

// CreateVirtualDNS inserts VirtualDNS to DB
// nolint
func (db *DB) createVirtualDNS(
	ctx context.Context,
	request *models.CreateVirtualDNSRequest) error {
	qb := db.queryBuilders["virtual_DNS"]
	tx := GetTransaction(ctx)
	model := request.VirtualDNS
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), bool(model.GetVirtualDNSData().GetReverseResolution()),
		string(model.GetVirtualDNSData().GetRecordOrder()),
		string(model.GetVirtualDNSData().GetNextVirtualDNS()),
		string(model.GetVirtualDNSData().GetFloatingIPRecord()),
		bool(model.GetVirtualDNSData().GetExternalVisible()),
		bool(model.GetVirtualDNSData().GetDynamicRecordsFromClient()),
		string(model.GetVirtualDNSData().GetDomainName()),
		int(model.GetVirtualDNSData().GetDefaultTTLSeconds()),
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
		int(model.GetConfigurationVersion()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "virtual_DNS",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "virtual_DNS", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanVirtualDNS(values map[string]interface{}) (*models.VirtualDNS, error) {
	m := models.MakeVirtualDNS()

	if value, ok := values["reverse_resolution"]; ok {

		m.VirtualDNSData.ReverseResolution = common.InterfaceToBool(value)

	}

	if value, ok := values["record_order"]; ok {

		m.VirtualDNSData.RecordOrder = common.InterfaceToString(value)

	}

	if value, ok := values["next_virtual_DNS"]; ok {

		m.VirtualDNSData.NextVirtualDNS = common.InterfaceToString(value)

	}

	if value, ok := values["floating_ip_record"]; ok {

		m.VirtualDNSData.FloatingIPRecord = common.InterfaceToString(value)

	}

	if value, ok := values["external_visible"]; ok {

		m.VirtualDNSData.ExternalVisible = common.InterfaceToBool(value)

	}

	if value, ok := values["dynamic_records_from_client"]; ok {

		m.VirtualDNSData.DynamicRecordsFromClient = common.InterfaceToBool(value)

	}

	if value, ok := values["domain_name"]; ok {

		m.VirtualDNSData.DomainName = common.InterfaceToString(value)

	}

	if value, ok := values["default_ttl_seconds"]; ok {

		m.VirtualDNSData.DefaultTTLSeconds = common.InterfaceToInt64(value)

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

	if value, ok := values["configuration_version"]; ok {

		m.ConfigurationVersion = common.InterfaceToInt64(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["backref_virtual_DNS_record"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeVirtualDNSRecord()
			m.VirtualDNSRecords = append(m.VirtualDNSRecords, childModel)

			if propertyValue, ok := childResourceMap["record_type"]; ok && propertyValue != nil {

				childModel.VirtualDNSRecordData.RecordType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["record_ttl_seconds"]; ok && propertyValue != nil {

				childModel.VirtualDNSRecordData.RecordTTLSeconds = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["record_name"]; ok && propertyValue != nil {

				childModel.VirtualDNSRecordData.RecordName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["record_mx_preference"]; ok && propertyValue != nil {

				childModel.VirtualDNSRecordData.RecordMXPreference = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["record_data"]; ok && propertyValue != nil {

				childModel.VirtualDNSRecordData.RecordData = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["record_class"]; ok && propertyValue != nil {

				childModel.VirtualDNSRecordData.RecordClass = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["configuration_version"]; ok && propertyValue != nil {

				childModel.ConfigurationVersion = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListVirtualDNS lists VirtualDNS with list spec.
func (db *DB) listVirtualDNS(ctx context.Context, request *models.ListVirtualDNSRequest) (response *models.ListVirtualDNSResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["virtual_DNS"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.VirtualDNS{}

	if spec.ParentFQName != nil {
		parentMetaData, err := db.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = models.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}
	query, columns, values := qb.ListQuery(auth, spec)
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
		m, err := scanVirtualDNS(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListVirtualDNSResponse{
		VirtualDNSs: result,
	}
	return response, nil
}

// UpdateVirtualDNS updates a resource
func (db *DB) updateVirtualDNS(
	ctx context.Context,
	request *models.UpdateVirtualDNSRequest,
) error {
	//TODO
	return nil
}

// DeleteVirtualDNS deletes a resource
func (db *DB) deleteVirtualDNS(
	ctx context.Context,
	request *models.DeleteVirtualDNSRequest) error {
	qb := db.queryBuilders["virtual_DNS"]

	selectQuery := qb.SelectForDeleteQuery()
	deleteQuery := qb.DeleteQuery()

	var err error
	var count int
	uuid := request.ID
	tx := GetTransaction(ctx)
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

	err = db.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}

//CreateVirtualDNS handle a Create API
// nolint
func (db *DB) CreateVirtualDNS(
	ctx context.Context,
	request *models.CreateVirtualDNSRequest) (*models.CreateVirtualDNSResponse, error) {
	model := request.VirtualDNS
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createVirtualDNS(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_DNS",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateVirtualDNSResponse{
		VirtualDNS: request.VirtualDNS,
	}, nil
}

//UpdateVirtualDNS handles a Update request.
func (db *DB) UpdateVirtualDNS(
	ctx context.Context,
	request *models.UpdateVirtualDNSRequest) (*models.UpdateVirtualDNSResponse, error) {
	model := request.VirtualDNS
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateVirtualDNS(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_DNS",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateVirtualDNSResponse{
		VirtualDNS: model,
	}, nil
}

//DeleteVirtualDNS delete a resource.
func (db *DB) DeleteVirtualDNS(ctx context.Context, request *models.DeleteVirtualDNSRequest) (*models.DeleteVirtualDNSResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteVirtualDNS(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteVirtualDNSResponse{
		ID: request.ID,
	}, nil
}

//GetVirtualDNS a Get request.
func (db *DB) GetVirtualDNS(ctx context.Context, request *models.GetVirtualDNSRequest) (response *models.GetVirtualDNSResponse, err error) {
	spec := &models.ListSpec{
		Limit:  1,
		Detail: true,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListVirtualDNSRequest{
		Spec: spec,
	}
	var result *models.ListVirtualDNSResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listVirtualDNS(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.VirtualDNSs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetVirtualDNSResponse{
		VirtualDNS: result.VirtualDNSs[0],
	}
	return response, nil
}

//ListVirtualDNS handles a List service Request.
// nolint
func (db *DB) ListVirtualDNS(
	ctx context.Context,
	request *models.ListVirtualDNSRequest) (response *models.ListVirtualDNSResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listVirtualDNS(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
