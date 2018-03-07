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

// ServiceApplianceFields is db columns for ServiceAppliance
var ServiceApplianceFields = []string{
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
	"configuration_version",
	"annotations_key_value_pair",
}

// ServiceApplianceRefFields is db reference fields for ServiceAppliance
var ServiceApplianceRefFields = map[string][]string{

	"physical_interface": []string{
		// <schema.Schema Value>
		"interface_type",
	},
}

// ServiceApplianceBackRefFields is db back reference fields for ServiceAppliance
var ServiceApplianceBackRefFields = map[string][]string{}

// ServiceApplianceParentTypes is possible parents for ServiceAppliance
var ServiceApplianceParents = []string{

	"service_appliance_set",
}

// CreateServiceAppliance inserts ServiceAppliance to DB
// nolint
func (db *DB) createServiceAppliance(
	ctx context.Context,
	request *models.CreateServiceApplianceRequest) error {
	qb := db.queryBuilders["service_appliance"]
	tx := GetTransaction(ctx)
	model := request.ServiceAppliance
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
		string(model.GetServiceApplianceUserCredentials().GetUsername()),
		string(model.GetServiceApplianceUserCredentials().GetPassword()),
		common.MustJSON(model.GetServiceApplianceProperties().GetKeyValuePair()),
		string(model.GetServiceApplianceIPAddress()),
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
		log.WithFields(log.Fields{
			"model": model,
			"err":   err}).Debug("create failed")
		return errors.Wrap(err, "create failed")
	}

	for _, ref := range model.PhysicalInterfaceRefs {

		if ref.Attr == nil {
			ref.Attr = &models.ServiceApplianceInterfaceType{}
		}

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("physical_interface"), model.UUID, ref.UUID, string(ref.Attr.GetInterfaceType()))
		if err != nil {
			return errors.Wrap(err, "PhysicalInterfaceRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "service_appliance",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "service_appliance", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanServiceAppliance(values map[string]interface{}) (*models.ServiceAppliance, error) {
	m := models.MakeServiceAppliance()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["username"]; ok {

		m.ServiceApplianceUserCredentials.Username = common.InterfaceToString(value)

	}

	if value, ok := values["password"]; ok {

		m.ServiceApplianceUserCredentials.Password = common.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.ServiceApplianceProperties.KeyValuePair)

	}

	if value, ok := values["service_appliance_ip_address"]; ok {

		m.ServiceApplianceIPAddress = common.InterfaceToString(value)

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

	if value, ok := values["annotations_key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_physical_interface"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.ServiceAppliancePhysicalInterfaceRef{}
			referenceModel.UUID = uuid
			m.PhysicalInterfaceRefs = append(m.PhysicalInterfaceRefs, referenceModel)

			attr := models.MakeServiceApplianceInterfaceType()
			referenceModel.Attr = attr

		}
	}

	return m, nil
}

// ListServiceAppliance lists ServiceAppliance with list spec.
func (db *DB) listServiceAppliance(ctx context.Context, request *models.ListServiceApplianceRequest) (response *models.ListServiceApplianceResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["service_appliance"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.ServiceAppliance{}

	if spec.ParentFQName != nil {
		parentMetaData, err := db.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = models.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}
	query, columns, values := qb.ListQuery(auth, spec)
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
		m, err := scanServiceAppliance(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListServiceApplianceResponse{
		ServiceAppliances: result,
	}
	return response, nil
}

// UpdateServiceAppliance updates a resource
func (db *DB) updateServiceAppliance(
	ctx context.Context,
	request *models.UpdateServiceApplianceRequest,
) error {
	//TODO
	return nil
}

// DeleteServiceAppliance deletes a resource
func (db *DB) deleteServiceAppliance(
	ctx context.Context,
	request *models.DeleteServiceApplianceRequest) error {
	qb := db.queryBuilders["service_appliance"]

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

//CreateServiceAppliance handle a Create API
// nolint
func (db *DB) CreateServiceAppliance(
	ctx context.Context,
	request *models.CreateServiceApplianceRequest) (*models.CreateServiceApplianceResponse, error) {
	model := request.ServiceAppliance
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createServiceAppliance(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_appliance",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceApplianceResponse{
		ServiceAppliance: request.ServiceAppliance,
	}, nil
}

//UpdateServiceAppliance handles a Update request.
func (db *DB) UpdateServiceAppliance(
	ctx context.Context,
	request *models.UpdateServiceApplianceRequest) (*models.UpdateServiceApplianceResponse, error) {
	model := request.ServiceAppliance
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateServiceAppliance(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_appliance",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceApplianceResponse{
		ServiceAppliance: model,
	}, nil
}

//DeleteServiceAppliance delete a resource.
func (db *DB) DeleteServiceAppliance(ctx context.Context, request *models.DeleteServiceApplianceRequest) (*models.DeleteServiceApplianceResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteServiceAppliance(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceApplianceResponse{
		ID: request.ID,
	}, nil
}

//GetServiceAppliance a Get request.
func (db *DB) GetServiceAppliance(ctx context.Context, request *models.GetServiceApplianceRequest) (response *models.GetServiceApplianceResponse, err error) {
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
	listRequest := &models.ListServiceApplianceRequest{
		Spec: spec,
	}
	var result *models.ListServiceApplianceResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listServiceAppliance(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceAppliances) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceApplianceResponse{
		ServiceAppliance: result.ServiceAppliances[0],
	}
	return response, nil
}

//ListServiceAppliance handles a List service Request.
// nolint
func (db *DB) ListServiceAppliance(
	ctx context.Context,
	request *models.ListServiceApplianceRequest) (response *models.ListServiceApplianceResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listServiceAppliance(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
