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

// ServiceTemplateFields is db columns for ServiceTemplate
var ServiceTemplateFields = []string{
	"uuid",
	"vrouter_instance_type",
	"version",
	"service_virtualization_type",
	"service_type",
	"service_scaling",
	"service_mode",
	"ordered_interfaces",
	"interface_type",
	"instance_data",
	"image_name",
	"flavor",
	"availability_zone_enable",
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

// ServiceTemplateRefFields is db reference fields for ServiceTemplate
var ServiceTemplateRefFields = map[string][]string{

	"service_appliance_set": []string{
	// <schema.Schema Value>

	},
}

// ServiceTemplateBackRefFields is db back reference fields for ServiceTemplate
var ServiceTemplateBackRefFields = map[string][]string{}

// ServiceTemplateParentTypes is possible parents for ServiceTemplate
var ServiceTemplateParents = []string{

	"domain",
}

// CreateServiceTemplate inserts ServiceTemplate to DB
// nolint
func (db *DB) createServiceTemplate(
	ctx context.Context,
	request *models.CreateServiceTemplateRequest) error {
	qb := db.queryBuilders["service_template"]
	tx := GetTransaction(ctx)
	model := request.ServiceTemplate
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
		string(model.GetServiceTemplateProperties().GetVrouterInstanceType()),
		int(model.GetServiceTemplateProperties().GetVersion()),
		string(model.GetServiceTemplateProperties().GetServiceVirtualizationType()),
		string(model.GetServiceTemplateProperties().GetServiceType()),
		bool(model.GetServiceTemplateProperties().GetServiceScaling()),
		string(model.GetServiceTemplateProperties().GetServiceMode()),
		bool(model.GetServiceTemplateProperties().GetOrderedInterfaces()),
		common.MustJSON(model.GetServiceTemplateProperties().GetInterfaceType()),
		string(model.GetServiceTemplateProperties().GetInstanceData()),
		string(model.GetServiceTemplateProperties().GetImageName()),
		string(model.GetServiceTemplateProperties().GetFlavor()),
		bool(model.GetServiceTemplateProperties().GetAvailabilityZoneEnable()),
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

	for _, ref := range model.ServiceApplianceSetRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("service_appliance_set"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceApplianceSetRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "service_template",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "service_template", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanServiceTemplate(values map[string]interface{}) (*models.ServiceTemplate, error) {
	m := models.MakeServiceTemplate()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["vrouter_instance_type"]; ok {

		m.ServiceTemplateProperties.VrouterInstanceType = common.InterfaceToString(value)

	}

	if value, ok := values["version"]; ok {

		m.ServiceTemplateProperties.Version = common.InterfaceToInt64(value)

	}

	if value, ok := values["service_virtualization_type"]; ok {

		m.ServiceTemplateProperties.ServiceVirtualizationType = common.InterfaceToString(value)

	}

	if value, ok := values["service_type"]; ok {

		m.ServiceTemplateProperties.ServiceType = common.InterfaceToString(value)

	}

	if value, ok := values["service_scaling"]; ok {

		m.ServiceTemplateProperties.ServiceScaling = common.InterfaceToBool(value)

	}

	if value, ok := values["service_mode"]; ok {

		m.ServiceTemplateProperties.ServiceMode = common.InterfaceToString(value)

	}

	if value, ok := values["ordered_interfaces"]; ok {

		m.ServiceTemplateProperties.OrderedInterfaces = common.InterfaceToBool(value)

	}

	if value, ok := values["interface_type"]; ok {

		json.Unmarshal(value.([]byte), &m.ServiceTemplateProperties.InterfaceType)

	}

	if value, ok := values["instance_data"]; ok {

		m.ServiceTemplateProperties.InstanceData = common.InterfaceToString(value)

	}

	if value, ok := values["image_name"]; ok {

		m.ServiceTemplateProperties.ImageName = common.InterfaceToString(value)

	}

	if value, ok := values["flavor"]; ok {

		m.ServiceTemplateProperties.Flavor = common.InterfaceToString(value)

	}

	if value, ok := values["availability_zone_enable"]; ok {

		m.ServiceTemplateProperties.AvailabilityZoneEnable = common.InterfaceToBool(value)

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

	if value, ok := values["ref_service_appliance_set"]; ok {
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
			referenceModel := &models.ServiceTemplateServiceApplianceSetRef{}
			referenceModel.UUID = uuid
			m.ServiceApplianceSetRefs = append(m.ServiceApplianceSetRefs, referenceModel)

		}
	}

	return m, nil
}

// ListServiceTemplate lists ServiceTemplate with list spec.
func (db *DB) listServiceTemplate(ctx context.Context, request *models.ListServiceTemplateRequest) (response *models.ListServiceTemplateResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["service_template"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.ServiceTemplate{}

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
		m, err := scanServiceTemplate(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListServiceTemplateResponse{
		ServiceTemplates: result,
	}
	return response, nil
}

// UpdateServiceTemplate updates a resource
func (db *DB) updateServiceTemplate(
	ctx context.Context,
	request *models.UpdateServiceTemplateRequest,
) error {
	//TODO
	return nil
}

// DeleteServiceTemplate deletes a resource
func (db *DB) deleteServiceTemplate(
	ctx context.Context,
	request *models.DeleteServiceTemplateRequest) error {
	qb := db.queryBuilders["service_template"]

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

//CreateServiceTemplate handle a Create API
// nolint
func (db *DB) CreateServiceTemplate(
	ctx context.Context,
	request *models.CreateServiceTemplateRequest) (*models.CreateServiceTemplateResponse, error) {
	model := request.ServiceTemplate
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createServiceTemplate(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_template",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceTemplateResponse{
		ServiceTemplate: request.ServiceTemplate,
	}, nil
}

//UpdateServiceTemplate handles a Update request.
func (db *DB) UpdateServiceTemplate(
	ctx context.Context,
	request *models.UpdateServiceTemplateRequest) (*models.UpdateServiceTemplateResponse, error) {
	model := request.ServiceTemplate
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateServiceTemplate(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_template",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceTemplateResponse{
		ServiceTemplate: model,
	}, nil
}

//DeleteServiceTemplate delete a resource.
func (db *DB) DeleteServiceTemplate(ctx context.Context, request *models.DeleteServiceTemplateRequest) (*models.DeleteServiceTemplateResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteServiceTemplate(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceTemplateResponse{
		ID: request.ID,
	}, nil
}

//GetServiceTemplate a Get request.
func (db *DB) GetServiceTemplate(ctx context.Context, request *models.GetServiceTemplateRequest) (response *models.GetServiceTemplateResponse, err error) {
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
	listRequest := &models.ListServiceTemplateRequest{
		Spec: spec,
	}
	var result *models.ListServiceTemplateResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listServiceTemplate(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceTemplates) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceTemplateResponse{
		ServiceTemplate: result.ServiceTemplates[0],
	}
	return response, nil
}

//ListServiceTemplate handles a List service Request.
// nolint
func (db *DB) ListServiceTemplate(
	ctx context.Context,
	request *models.ListServiceTemplateRequest) (response *models.ListServiceTemplateResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listServiceTemplate(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
