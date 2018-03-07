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

// LoadbalancerPoolFields is db columns for LoadbalancerPool
var LoadbalancerPoolFields = []string{
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"loadbalancer_pool_provider",
	"subnet_id",
	"status_description",
	"status",
	"session_persistence",
	"protocol",
	"persistence_cookie_name",
	"loadbalancer_method",
	"admin_state",
	"key_value_pair",
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

// LoadbalancerPoolRefFields is db reference fields for LoadbalancerPool
var LoadbalancerPoolRefFields = map[string][]string{

	"loadbalancer_healthmonitor": []string{
	// <schema.Schema Value>

	},

	"service_appliance_set": []string{
	// <schema.Schema Value>

	},

	"virtual_machine_interface": []string{
	// <schema.Schema Value>

	},

	"loadbalancer_listener": []string{
	// <schema.Schema Value>

	},

	"service_instance": []string{
	// <schema.Schema Value>

	},
}

// LoadbalancerPoolBackRefFields is db back reference fields for LoadbalancerPool
var LoadbalancerPoolBackRefFields = map[string][]string{

	"loadbalancer_member": []string{
		"uuid",
		"share",
		"owner_access",
		"owner",
		"global_access",
		"parent_uuid",
		"parent_type",
		"weight",
		"status_description",
		"status",
		"protocol_port",
		"admin_state",
		"address",
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

// LoadbalancerPoolParentTypes is possible parents for LoadbalancerPool
var LoadbalancerPoolParents = []string{

	"project",
}

// CreateLoadbalancerPool inserts LoadbalancerPool to DB
// nolint
func (db *DB) createLoadbalancerPool(
	ctx context.Context,
	request *models.CreateLoadbalancerPoolRequest) error {
	qb := db.queryBuilders["loadbalancer_pool"]
	tx := GetTransaction(ctx)
	model := request.LoadbalancerPool
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		string(model.GetLoadbalancerPoolProvider()),
		string(model.GetLoadbalancerPoolProperties().GetSubnetID()),
		string(model.GetLoadbalancerPoolProperties().GetStatusDescription()),
		string(model.GetLoadbalancerPoolProperties().GetStatus()),
		string(model.GetLoadbalancerPoolProperties().GetSessionPersistence()),
		string(model.GetLoadbalancerPoolProperties().GetProtocol()),
		string(model.GetLoadbalancerPoolProperties().GetPersistenceCookieName()),
		string(model.GetLoadbalancerPoolProperties().GetLoadbalancerMethod()),
		bool(model.GetLoadbalancerPoolProperties().GetAdminState()),
		common.MustJSON(model.GetLoadbalancerPoolCustomAttributes().GetKeyValuePair()),
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

	for _, ref := range model.ServiceApplianceSetRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("service_appliance_set"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceApplianceSetRefs create failed")
		}
	}

	for _, ref := range model.VirtualMachineInterfaceRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("virtual_machine_interface"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	for _, ref := range model.LoadbalancerListenerRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("loadbalancer_listener"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "LoadbalancerListenerRefs create failed")
		}
	}

	for _, ref := range model.ServiceInstanceRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("service_instance"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs create failed")
		}
	}

	for _, ref := range model.LoadbalancerHealthmonitorRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("loadbalancer_healthmonitor"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "LoadbalancerHealthmonitorRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "loadbalancer_pool",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "loadbalancer_pool", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanLoadbalancerPool(values map[string]interface{}) (*models.LoadbalancerPool, error) {
	m := models.MakeLoadbalancerPool()

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

	if value, ok := values["loadbalancer_pool_provider"]; ok {

		m.LoadbalancerPoolProvider = common.InterfaceToString(value)

	}

	if value, ok := values["subnet_id"]; ok {

		m.LoadbalancerPoolProperties.SubnetID = common.InterfaceToString(value)

	}

	if value, ok := values["status_description"]; ok {

		m.LoadbalancerPoolProperties.StatusDescription = common.InterfaceToString(value)

	}

	if value, ok := values["status"]; ok {

		m.LoadbalancerPoolProperties.Status = common.InterfaceToString(value)

	}

	if value, ok := values["session_persistence"]; ok {

		m.LoadbalancerPoolProperties.SessionPersistence = common.InterfaceToString(value)

	}

	if value, ok := values["protocol"]; ok {

		m.LoadbalancerPoolProperties.Protocol = common.InterfaceToString(value)

	}

	if value, ok := values["persistence_cookie_name"]; ok {

		m.LoadbalancerPoolProperties.PersistenceCookieName = common.InterfaceToString(value)

	}

	if value, ok := values["loadbalancer_method"]; ok {

		m.LoadbalancerPoolProperties.LoadbalancerMethod = common.InterfaceToString(value)

	}

	if value, ok := values["admin_state"]; ok {

		m.LoadbalancerPoolProperties.AdminState = common.InterfaceToBool(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.LoadbalancerPoolCustomAttributes.KeyValuePair)

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
			referenceModel := &models.LoadbalancerPoolServiceApplianceSetRef{}
			referenceModel.UUID = uuid
			m.ServiceApplianceSetRefs = append(m.ServiceApplianceSetRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
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
			referenceModel := &models.LoadbalancerPoolVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_loadbalancer_listener"]; ok {
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
			referenceModel := &models.LoadbalancerPoolLoadbalancerListenerRef{}
			referenceModel.UUID = uuid
			m.LoadbalancerListenerRefs = append(m.LoadbalancerListenerRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_instance"]; ok {
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
			referenceModel := &models.LoadbalancerPoolServiceInstanceRef{}
			referenceModel.UUID = uuid
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_loadbalancer_healthmonitor"]; ok {
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
			referenceModel := &models.LoadbalancerPoolLoadbalancerHealthmonitorRef{}
			referenceModel.UUID = uuid
			m.LoadbalancerHealthmonitorRefs = append(m.LoadbalancerHealthmonitorRefs, referenceModel)

		}
	}

	if value, ok := values["backref_loadbalancer_member"]; ok {
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
			childModel := models.MakeLoadbalancerMember()
			m.LoadbalancerMembers = append(m.LoadbalancerMembers, childModel)

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

			if propertyValue, ok := childResourceMap["weight"]; ok && propertyValue != nil {

				childModel.LoadbalancerMemberProperties.Weight = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["status_description"]; ok && propertyValue != nil {

				childModel.LoadbalancerMemberProperties.StatusDescription = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["status"]; ok && propertyValue != nil {

				childModel.LoadbalancerMemberProperties.Status = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["protocol_port"]; ok && propertyValue != nil {

				childModel.LoadbalancerMemberProperties.ProtocolPort = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["admin_state"]; ok && propertyValue != nil {

				childModel.LoadbalancerMemberProperties.AdminState = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["address"]; ok && propertyValue != nil {

				childModel.LoadbalancerMemberProperties.Address = common.InterfaceToString(propertyValue)

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

// ListLoadbalancerPool lists LoadbalancerPool with list spec.
func (db *DB) listLoadbalancerPool(ctx context.Context, request *models.ListLoadbalancerPoolRequest) (response *models.ListLoadbalancerPoolResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["loadbalancer_pool"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.LoadbalancerPool{}

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
		m, err := scanLoadbalancerPool(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListLoadbalancerPoolResponse{
		LoadbalancerPools: result,
	}
	return response, nil
}

// UpdateLoadbalancerPool updates a resource
func (db *DB) updateLoadbalancerPool(
	ctx context.Context,
	request *models.UpdateLoadbalancerPoolRequest,
) error {
	//TODO
	return nil
}

// DeleteLoadbalancerPool deletes a resource
func (db *DB) deleteLoadbalancerPool(
	ctx context.Context,
	request *models.DeleteLoadbalancerPoolRequest) error {
	qb := db.queryBuilders["loadbalancer_pool"]

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

//CreateLoadbalancerPool handle a Create API
// nolint
func (db *DB) CreateLoadbalancerPool(
	ctx context.Context,
	request *models.CreateLoadbalancerPoolRequest) (*models.CreateLoadbalancerPoolResponse, error) {
	model := request.LoadbalancerPool
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createLoadbalancerPool(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_pool",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateLoadbalancerPoolResponse{
		LoadbalancerPool: request.LoadbalancerPool,
	}, nil
}

//UpdateLoadbalancerPool handles a Update request.
func (db *DB) UpdateLoadbalancerPool(
	ctx context.Context,
	request *models.UpdateLoadbalancerPoolRequest) (*models.UpdateLoadbalancerPoolResponse, error) {
	model := request.LoadbalancerPool
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateLoadbalancerPool(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_pool",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateLoadbalancerPoolResponse{
		LoadbalancerPool: model,
	}, nil
}

//DeleteLoadbalancerPool delete a resource.
func (db *DB) DeleteLoadbalancerPool(ctx context.Context, request *models.DeleteLoadbalancerPoolRequest) (*models.DeleteLoadbalancerPoolResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteLoadbalancerPool(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteLoadbalancerPoolResponse{
		ID: request.ID,
	}, nil
}

//GetLoadbalancerPool a Get request.
func (db *DB) GetLoadbalancerPool(ctx context.Context, request *models.GetLoadbalancerPoolRequest) (response *models.GetLoadbalancerPoolResponse, err error) {
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
	listRequest := &models.ListLoadbalancerPoolRequest{
		Spec: spec,
	}
	var result *models.ListLoadbalancerPoolResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listLoadbalancerPool(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.LoadbalancerPools) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetLoadbalancerPoolResponse{
		LoadbalancerPool: result.LoadbalancerPools[0],
	}
	return response, nil
}

//ListLoadbalancerPool handles a List service Request.
// nolint
func (db *DB) ListLoadbalancerPool(
	ctx context.Context,
	request *models.ListLoadbalancerPoolRequest) (response *models.ListLoadbalancerPoolResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listLoadbalancerPool(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
