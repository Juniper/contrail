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

const insertLoadbalancerQuery = "insert into `loadbalancer` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`loadbalancer_provider`,`vip_subnet_id`,`vip_address`,`status`,`provisioning_status`,`operating_status`,`admin_state`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteLoadbalancerQuery = "delete from `loadbalancer` where uuid = ?"

// LoadbalancerFields is db columns for Loadbalancer
var LoadbalancerFields = []string{
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"loadbalancer_provider",
	"vip_subnet_id",
	"vip_address",
	"status",
	"provisioning_status",
	"operating_status",
	"admin_state",
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

// LoadbalancerRefFields is db reference fields for Loadbalancer
var LoadbalancerRefFields = map[string][]string{

	"service_appliance_set": []string{
	// <schema.Schema Value>

	},

	"virtual_machine_interface": []string{
	// <schema.Schema Value>

	},

	"service_instance": []string{
	// <schema.Schema Value>

	},
}

// LoadbalancerBackRefFields is db back reference fields for Loadbalancer
var LoadbalancerBackRefFields = map[string][]string{}

// LoadbalancerParentTypes is possible parents for Loadbalancer
var LoadbalancerParents = []string{

	"project",
}

const insertLoadbalancerServiceApplianceSetQuery = "insert into `ref_loadbalancer_service_appliance_set` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerVirtualMachineInterfaceQuery = "insert into `ref_loadbalancer_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerServiceInstanceQuery = "insert into `ref_loadbalancer_service_instance` (`from`, `to` ) values (?, ?);"

// CreateLoadbalancer inserts Loadbalancer to DB
func (db *DB) createLoadbalancer(
	ctx context.Context,
	request *models.CreateLoadbalancerRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.Loadbalancer
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertLoadbalancerQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		string(model.GetLoadbalancerProvider()),
		string(model.GetLoadbalancerProperties().GetVipSubnetID()),
		string(model.GetLoadbalancerProperties().GetVipAddress()),
		string(model.GetLoadbalancerProperties().GetStatus()),
		string(model.GetLoadbalancerProperties().GetProvisioningStatus()),
		string(model.GetLoadbalancerProperties().GetOperatingStatus()),
		bool(model.GetLoadbalancerProperties().GetAdminState()),
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

	stmtServiceApplianceSetRef, err := tx.Prepare(insertLoadbalancerServiceApplianceSetQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceApplianceSetRefs create statement failed")
	}
	defer stmtServiceApplianceSetRef.Close()
	for _, ref := range model.ServiceApplianceSetRefs {

		_, err = stmtServiceApplianceSetRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceApplianceSetRefs create failed")
		}
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertLoadbalancerVirtualMachineInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs create statement failed")
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {

		_, err = stmtVirtualMachineInterfaceRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	stmtServiceInstanceRef, err := tx.Prepare(insertLoadbalancerServiceInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceInstanceRefs create statement failed")
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {

		_, err = stmtServiceInstanceRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "loadbalancer",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "loadbalancer", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanLoadbalancer(values map[string]interface{}) (*models.Loadbalancer, error) {
	m := models.MakeLoadbalancer()

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

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

	if value, ok := values["loadbalancer_provider"]; ok {

		m.LoadbalancerProvider = schema.InterfaceToString(value)

	}

	if value, ok := values["vip_subnet_id"]; ok {

		m.LoadbalancerProperties.VipSubnetID = schema.InterfaceToString(value)

	}

	if value, ok := values["vip_address"]; ok {

		m.LoadbalancerProperties.VipAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["status"]; ok {

		m.LoadbalancerProperties.Status = schema.InterfaceToString(value)

	}

	if value, ok := values["provisioning_status"]; ok {

		m.LoadbalancerProperties.ProvisioningStatus = schema.InterfaceToString(value)

	}

	if value, ok := values["operating_status"]; ok {

		m.LoadbalancerProperties.OperatingStatus = schema.InterfaceToString(value)

	}

	if value, ok := values["admin_state"]; ok {

		m.LoadbalancerProperties.AdminState = schema.InterfaceToBool(value)

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

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_service_appliance_set"]; ok {
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
			referenceModel := &models.LoadbalancerServiceApplianceSetRef{}
			referenceModel.UUID = uuid
			m.ServiceApplianceSetRefs = append(m.ServiceApplianceSetRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
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
			referenceModel := &models.LoadbalancerVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_instance"]; ok {
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
			referenceModel := &models.LoadbalancerServiceInstanceRef{}
			referenceModel.UUID = uuid
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

		}
	}

	return m, nil
}

// ListLoadbalancer lists Loadbalancer with list spec.
func (db *DB) listLoadbalancer(ctx context.Context, request *models.ListLoadbalancerRequest) (response *models.ListLoadbalancerResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "loadbalancer"
	qb.Fields = LoadbalancerFields
	qb.RefFields = LoadbalancerRefFields
	qb.BackRefFields = LoadbalancerBackRefFields
	result := []*models.Loadbalancer{}

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
		m, err := scanLoadbalancer(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListLoadbalancerResponse{
		Loadbalancers: result,
	}
	return response, nil
}

// UpdateLoadbalancer updates a resource
func (db *DB) updateLoadbalancer(
	ctx context.Context,
	request *models.UpdateLoadbalancerRequest,
) error {
	//TODO
	return nil
}

// DeleteLoadbalancer deletes a resource
func (db *DB) deleteLoadbalancer(
	ctx context.Context,
	request *models.DeleteLoadbalancerRequest) error {
	deleteQuery := deleteLoadbalancerQuery
	selectQuery := "select count(uuid) from loadbalancer where uuid = ?"
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

//CreateLoadbalancer handle a Create API
func (db *DB) CreateLoadbalancer(
	ctx context.Context,
	request *models.CreateLoadbalancerRequest) (*models.CreateLoadbalancerResponse, error) {
	model := request.Loadbalancer
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createLoadbalancer(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateLoadbalancerResponse{
		Loadbalancer: request.Loadbalancer,
	}, nil
}

//UpdateLoadbalancer handles a Update request.
func (db *DB) UpdateLoadbalancer(
	ctx context.Context,
	request *models.UpdateLoadbalancerRequest) (*models.UpdateLoadbalancerResponse, error) {
	model := request.Loadbalancer
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateLoadbalancer(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateLoadbalancerResponse{
		Loadbalancer: model,
	}, nil
}

//DeleteLoadbalancer delete a resource.
func (db *DB) DeleteLoadbalancer(ctx context.Context, request *models.DeleteLoadbalancerRequest) (*models.DeleteLoadbalancerResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteLoadbalancer(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteLoadbalancerResponse{
		ID: request.ID,
	}, nil
}

//GetLoadbalancer a Get request.
func (db *DB) GetLoadbalancer(ctx context.Context, request *models.GetLoadbalancerRequest) (response *models.GetLoadbalancerResponse, err error) {
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
	listRequest := &models.ListLoadbalancerRequest{
		Spec: spec,
	}
	var result *models.ListLoadbalancerResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listLoadbalancer(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Loadbalancers) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetLoadbalancerResponse{
		Loadbalancer: result.Loadbalancers[0],
	}
	return response, nil
}

//ListLoadbalancer handles a List service Request.
func (db *DB) ListLoadbalancer(
	ctx context.Context,
	request *models.ListLoadbalancerRequest) (response *models.ListLoadbalancerResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listLoadbalancer(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
