package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertServiceApplianceSetQuery = "insert into `service_appliance_set` (`uuid`,`key_value_pair`,`service_appliance_ha_mode`,`service_appliance_driver`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`annotations_key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteServiceApplianceSetQuery = "delete from `service_appliance_set` where uuid = ?"

// ServiceApplianceSetFields is db columns for ServiceApplianceSet
var ServiceApplianceSetFields = []string{
	"uuid",
	"key_value_pair",
	"service_appliance_ha_mode",
	"service_appliance_driver",
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
	"annotations_key_value_pair",
}

// ServiceApplianceSetRefFields is db reference fields for ServiceApplianceSet
var ServiceApplianceSetRefFields = map[string][]string{}

// ServiceApplianceSetBackRefFields is db back reference fields for ServiceApplianceSet
var ServiceApplianceSetBackRefFields = map[string][]string{

	"service_appliance": []string{
		"uuid",
		"username",
		"password",
		"key_value_pair",
		"service_appliance_ip_address",
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
		"annotations_key_value_pair",
	},
}

// ServiceApplianceSetParentTypes is possible parents for ServiceApplianceSet
var ServiceApplianceSetParents = []string{

	"global_system_config",
}

// CreateServiceApplianceSet inserts ServiceApplianceSet to DB
func (db *DB) createServiceApplianceSet(
	ctx context.Context,
	request *models.CreateServiceApplianceSetRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.ServiceApplianceSet
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceApplianceSetQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertServiceApplianceSetQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		common.MustJSON(model.GetServiceApplianceSetProperties().GetKeyValuePair()),
		string(model.GetServiceApplianceHaMode()),
		string(model.GetServiceApplianceDriver()),
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
		Type:   "service_appliance_set",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "service_appliance_set", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanServiceApplianceSet(values map[string]interface{}) (*models.ServiceApplianceSet, error) {
	m := models.MakeServiceApplianceSet()

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.ServiceApplianceSetProperties.KeyValuePair)

	}

	if value, ok := values["service_appliance_ha_mode"]; ok {

		m.ServiceApplianceHaMode = schema.InterfaceToString(value)

	}

	if value, ok := values["service_appliance_driver"]; ok {

		m.ServiceApplianceDriver = schema.InterfaceToString(value)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		m.Perms2.OwnerAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["owner"]; ok {

		m.Perms2.Owner = schema.InterfaceToString(value)

	}

	if value, ok := values["global_access"]; ok {

		m.Perms2.GlobalAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = schema.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = schema.InterfaceToString(value)

	}

	if value, ok := values["user_visible"]; ok {

		m.IDPerms.UserVisible = schema.InterfaceToBool(value)

	}

	if value, ok := values["permissions_owner_access"]; ok {

		m.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["permissions_owner"]; ok {

		m.IDPerms.Permissions.Owner = schema.InterfaceToString(value)

	}

	if value, ok := values["other_access"]; ok {

		m.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["group_access"]; ok {

		m.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["group"]; ok {

		m.IDPerms.Permissions.Group = schema.InterfaceToString(value)

	}

	if value, ok := values["last_modified"]; ok {

		m.IDPerms.LastModified = schema.InterfaceToString(value)

	}

	if value, ok := values["enable"]; ok {

		m.IDPerms.Enable = schema.InterfaceToBool(value)

	}

	if value, ok := values["description"]; ok {

		m.IDPerms.Description = schema.InterfaceToString(value)

	}

	if value, ok := values["creator"]; ok {

		m.IDPerms.Creator = schema.InterfaceToString(value)

	}

	if value, ok := values["created"]; ok {

		m.IDPerms.Created = schema.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = schema.InterfaceToString(value)

	}

	if value, ok := values["annotations_key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["backref_service_appliance"]; ok {
		var childResources []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeServiceAppliance()
			m.ServiceAppliances = append(m.ServiceAppliances, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["username"]; ok && propertyValue != nil {

				childModel.ServiceApplianceUserCredentials.Username = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["password"]; ok && propertyValue != nil {

				childModel.ServiceApplianceUserCredentials.Password = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.ServiceApplianceProperties.KeyValuePair)

			}

			if propertyValue, ok := childResourceMap["service_appliance_ip_address"]; ok && propertyValue != nil {

				childModel.ServiceApplianceIPAddress = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["annotations_key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListServiceApplianceSet lists ServiceApplianceSet with list spec.
func (db *DB) listServiceApplianceSet(ctx context.Context, request *models.ListServiceApplianceSetRequest) (response *models.ListServiceApplianceSetResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "service_appliance_set"
	qb.Fields = ServiceApplianceSetFields
	qb.RefFields = ServiceApplianceSetRefFields
	qb.BackRefFields = ServiceApplianceSetBackRefFields
	result := []*models.ServiceApplianceSet{}

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
		m, err := scanServiceApplianceSet(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListServiceApplianceSetResponse{
		ServiceApplianceSets: result,
	}
	return response, nil
}

// UpdateServiceApplianceSet updates a resource
func (db *DB) updateServiceApplianceSet(
	ctx context.Context,
	request *models.UpdateServiceApplianceSetRequest,
) error {
	//TODO
	return nil
}

// DeleteServiceApplianceSet deletes a resource
func (db *DB) deleteServiceApplianceSet(
	ctx context.Context,
	request *models.DeleteServiceApplianceSetRequest) error {
	deleteQuery := deleteServiceApplianceSetQuery
	selectQuery := "select count(uuid) from service_appliance_set where uuid = ?"
	var err error
	var count int
	uuid := request.ID
	tx := common.GetTransaction(ctx)
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

//CreateServiceApplianceSet handle a Create API
func (db *DB) CreateServiceApplianceSet(
	ctx context.Context,
	request *models.CreateServiceApplianceSetRequest) (*models.CreateServiceApplianceSetResponse, error) {
	model := request.ServiceApplianceSet
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createServiceApplianceSet(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_appliance_set",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceApplianceSetResponse{
		ServiceApplianceSet: request.ServiceApplianceSet,
	}, nil
}

//UpdateServiceApplianceSet handles a Update request.
func (db *DB) UpdateServiceApplianceSet(
	ctx context.Context,
	request *models.UpdateServiceApplianceSetRequest) (*models.UpdateServiceApplianceSetResponse, error) {
	model := request.ServiceApplianceSet
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateServiceApplianceSet(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_appliance_set",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceApplianceSetResponse{
		ServiceApplianceSet: model,
	}, nil
}

//DeleteServiceApplianceSet delete a resource.
func (db *DB) DeleteServiceApplianceSet(ctx context.Context, request *models.DeleteServiceApplianceSetRequest) (*models.DeleteServiceApplianceSetResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteServiceApplianceSet(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceApplianceSetResponse{
		ID: request.ID,
	}, nil
}

//GetServiceApplianceSet a Get request.
func (db *DB) GetServiceApplianceSet(ctx context.Context, request *models.GetServiceApplianceSetRequest) (response *models.GetServiceApplianceSetResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListServiceApplianceSetRequest{
		Spec: spec,
	}
	var result *models.ListServiceApplianceSetResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listServiceApplianceSet(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceApplianceSets) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceApplianceSetResponse{
		ServiceApplianceSet: result.ServiceApplianceSets[0],
	}
	return response, nil
}

//ListServiceApplianceSet handles a List service Request.
func (db *DB) ListServiceApplianceSet(
	ctx context.Context,
	request *models.ListServiceApplianceSetRequest) (response *models.ListServiceApplianceSetResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listServiceApplianceSet(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
