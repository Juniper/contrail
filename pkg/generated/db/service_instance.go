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

const insertServiceInstanceQuery = "insert into `service_instance` (`uuid`,`virtual_router_id`,`max_instances`,`auto_scale`,`right_virtual_network`,`right_ip_address`,`management_virtual_network`,`left_virtual_network`,`left_ip_address`,`interface_list`,`ha_mode`,`availability_zone`,`auto_policy`,`key_value_pair`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`annotations_key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteServiceInstanceQuery = "delete from `service_instance` where uuid = ?"

// ServiceInstanceFields is db columns for ServiceInstance
var ServiceInstanceFields = []string{
	"uuid",
	"virtual_router_id",
	"max_instances",
	"auto_scale",
	"right_virtual_network",
	"right_ip_address",
	"management_virtual_network",
	"left_virtual_network",
	"left_ip_address",
	"interface_list",
	"ha_mode",
	"availability_zone",
	"auto_policy",
	"key_value_pair",
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

// ServiceInstanceRefFields is db reference fields for ServiceInstance
var ServiceInstanceRefFields = map[string][]string{

	"service_template": []string{
	// <schema.Schema Value>

	},

	"instance_ip": []string{
		// <schema.Schema Value>
		"interface_type",
	},
}

// ServiceInstanceBackRefFields is db back reference fields for ServiceInstance
var ServiceInstanceBackRefFields = map[string][]string{

	"port_tuple": []string{
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
	},
}

// ServiceInstanceParentTypes is possible parents for ServiceInstance
var ServiceInstanceParents = []string{

	"project",
}

const insertServiceInstanceServiceTemplateQuery = "insert into `ref_service_instance_service_template` (`from`, `to` ) values (?, ?);"

const insertServiceInstanceInstanceIPQuery = "insert into `ref_service_instance_instance_ip` (`from`, `to` ,`interface_type`) values (?, ?,?);"

// CreateServiceInstance inserts ServiceInstance to DB
func (db *DB) createServiceInstance(
	ctx context.Context,
	request *models.CreateServiceInstanceRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.ServiceInstance
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertServiceInstanceQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		string(model.GetServiceInstanceProperties().GetVirtualRouterID()),
		int(model.GetServiceInstanceProperties().GetScaleOut().GetMaxInstances()),
		bool(model.GetServiceInstanceProperties().GetScaleOut().GetAutoScale()),
		string(model.GetServiceInstanceProperties().GetRightVirtualNetwork()),
		string(model.GetServiceInstanceProperties().GetRightIPAddress()),
		string(model.GetServiceInstanceProperties().GetManagementVirtualNetwork()),
		string(model.GetServiceInstanceProperties().GetLeftVirtualNetwork()),
		string(model.GetServiceInstanceProperties().GetLeftIPAddress()),
		common.MustJSON(model.GetServiceInstanceProperties().GetInterfaceList()),
		string(model.GetServiceInstanceProperties().GetHaMode()),
		string(model.GetServiceInstanceProperties().GetAvailabilityZone()),
		bool(model.GetServiceInstanceProperties().GetAutoPolicy()),
		common.MustJSON(model.GetServiceInstanceBindings().GetKeyValuePair()),
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

	stmtServiceTemplateRef, err := tx.Prepare(insertServiceInstanceServiceTemplateQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceTemplateRefs create statement failed")
	}
	defer stmtServiceTemplateRef.Close()
	for _, ref := range model.ServiceTemplateRefs {

		_, err = stmtServiceTemplateRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceTemplateRefs create failed")
		}
	}

	stmtInstanceIPRef, err := tx.Prepare(insertServiceInstanceInstanceIPQuery)
	if err != nil {
		return errors.Wrap(err, "preparing InstanceIPRefs create statement failed")
	}
	defer stmtInstanceIPRef.Close()
	for _, ref := range model.InstanceIPRefs {

		if ref.Attr == nil {
			ref.Attr = &models.ServiceInterfaceTag{}
		}

		_, err = stmtInstanceIPRef.ExecContext(ctx, model.UUID, ref.UUID, string(ref.Attr.GetInterfaceType()))
		if err != nil {
			return errors.Wrap(err, "InstanceIPRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "service_instance",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "service_instance", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanServiceInstance(values map[string]interface{}) (*models.ServiceInstance, error) {
	m := models.MakeServiceInstance()

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["virtual_router_id"]; ok {

		m.ServiceInstanceProperties.VirtualRouterID = schema.InterfaceToString(value)

	}

	if value, ok := values["max_instances"]; ok {

		m.ServiceInstanceProperties.ScaleOut.MaxInstances = schema.InterfaceToInt64(value)

	}

	if value, ok := values["auto_scale"]; ok {

		m.ServiceInstanceProperties.ScaleOut.AutoScale = schema.InterfaceToBool(value)

	}

	if value, ok := values["right_virtual_network"]; ok {

		m.ServiceInstanceProperties.RightVirtualNetwork = schema.InterfaceToString(value)

	}

	if value, ok := values["right_ip_address"]; ok {

		m.ServiceInstanceProperties.RightIPAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["management_virtual_network"]; ok {

		m.ServiceInstanceProperties.ManagementVirtualNetwork = schema.InterfaceToString(value)

	}

	if value, ok := values["left_virtual_network"]; ok {

		m.ServiceInstanceProperties.LeftVirtualNetwork = schema.InterfaceToString(value)

	}

	if value, ok := values["left_ip_address"]; ok {

		m.ServiceInstanceProperties.LeftIPAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["interface_list"]; ok {

		json.Unmarshal(value.([]byte), &m.ServiceInstanceProperties.InterfaceList)

	}

	if value, ok := values["ha_mode"]; ok {

		m.ServiceInstanceProperties.HaMode = schema.InterfaceToString(value)

	}

	if value, ok := values["availability_zone"]; ok {

		m.ServiceInstanceProperties.AvailabilityZone = schema.InterfaceToString(value)

	}

	if value, ok := values["auto_policy"]; ok {

		m.ServiceInstanceProperties.AutoPolicy = schema.InterfaceToBool(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.ServiceInstanceBindings.KeyValuePair)

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

	if value, ok := values["ref_service_template"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.ServiceInstanceServiceTemplateRef{}
			referenceModel.UUID = uuid
			m.ServiceTemplateRefs = append(m.ServiceTemplateRefs, referenceModel)

		}
	}

	if value, ok := values["ref_instance_ip"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.ServiceInstanceInstanceIPRef{}
			referenceModel.UUID = uuid
			m.InstanceIPRefs = append(m.InstanceIPRefs, referenceModel)

			attr := models.MakeServiceInterfaceTag()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["backref_port_tuple"]; ok {
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
			childModel := models.MakePortTuple()
			m.PortTuples = append(m.PortTuples, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = schema.InterfaceToString(propertyValue)

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

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListServiceInstance lists ServiceInstance with list spec.
func (db *DB) listServiceInstance(ctx context.Context, request *models.ListServiceInstanceRequest) (response *models.ListServiceInstanceResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "service_instance"
	qb.Fields = ServiceInstanceFields
	qb.RefFields = ServiceInstanceRefFields
	qb.BackRefFields = ServiceInstanceBackRefFields
	result := []*models.ServiceInstance{}

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
		m, err := scanServiceInstance(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListServiceInstanceResponse{
		ServiceInstances: result,
	}
	return response, nil
}

// UpdateServiceInstance updates a resource
func (db *DB) updateServiceInstance(
	ctx context.Context,
	request *models.UpdateServiceInstanceRequest,
) error {
	//TODO
	return nil
}

// DeleteServiceInstance deletes a resource
func (db *DB) deleteServiceInstance(
	ctx context.Context,
	request *models.DeleteServiceInstanceRequest) error {
	deleteQuery := deleteServiceInstanceQuery
	selectQuery := "select count(uuid) from service_instance where uuid = ?"
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

//CreateServiceInstance handle a Create API
func (db *DB) CreateServiceInstance(
	ctx context.Context,
	request *models.CreateServiceInstanceRequest) (*models.CreateServiceInstanceResponse, error) {
	model := request.ServiceInstance
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createServiceInstance(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_instance",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceInstanceResponse{
		ServiceInstance: request.ServiceInstance,
	}, nil
}

//UpdateServiceInstance handles a Update request.
func (db *DB) UpdateServiceInstance(
	ctx context.Context,
	request *models.UpdateServiceInstanceRequest) (*models.UpdateServiceInstanceResponse, error) {
	model := request.ServiceInstance
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateServiceInstance(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_instance",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceInstanceResponse{
		ServiceInstance: model,
	}, nil
}

//DeleteServiceInstance delete a resource.
func (db *DB) DeleteServiceInstance(ctx context.Context, request *models.DeleteServiceInstanceRequest) (*models.DeleteServiceInstanceResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteServiceInstance(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceInstanceResponse{
		ID: request.ID,
	}, nil
}

//GetServiceInstance a Get request.
func (db *DB) GetServiceInstance(ctx context.Context, request *models.GetServiceInstanceRequest) (response *models.GetServiceInstanceResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListServiceInstanceRequest{
		Spec: spec,
	}
	var result *models.ListServiceInstanceResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listServiceInstance(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceInstances) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceInstanceResponse{
		ServiceInstance: result.ServiceInstances[0],
	}
	return response, nil
}

//ListServiceInstance handles a List service Request.
func (db *DB) ListServiceInstance(
	ctx context.Context,
	request *models.ListServiceInstanceRequest) (response *models.ListServiceInstanceResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listServiceInstance(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
