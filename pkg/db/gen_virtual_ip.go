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

// VirtualIPFields is db columns for VirtualIP
var VirtualIPFields = []string{
	"subnet_id",
	"status_description",
	"status",
	"protocol_port",
	"protocol",
	"persistence_type",
	"persistence_cookie_name",
	"connection_limit",
	"admin_state",
	"address",
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

// VirtualIPRefFields is db reference fields for VirtualIP
var VirtualIPRefFields = map[string][]string{

	"virtual_machine_interface": []string{
	// <schema.Schema Value>

	},

	"loadbalancer_pool": []string{
	// <schema.Schema Value>

	},
}

// VirtualIPBackRefFields is db back reference fields for VirtualIP
var VirtualIPBackRefFields = map[string][]string{}

// VirtualIPParentTypes is possible parents for VirtualIP
var VirtualIPParents = []string{

	"project",
}

// CreateVirtualIP inserts VirtualIP to DB
// nolint
func (db *DB) createVirtualIP(
	ctx context.Context,
	request *models.CreateVirtualIPRequest) error {
	qb := db.queryBuilders["virtual_ip"]
	tx := GetTransaction(ctx)
	model := request.VirtualIP
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetVirtualIPProperties().GetSubnetID()),
		string(model.GetVirtualIPProperties().GetStatusDescription()),
		string(model.GetVirtualIPProperties().GetStatus()),
		int(model.GetVirtualIPProperties().GetProtocolPort()),
		string(model.GetVirtualIPProperties().GetProtocol()),
		string(model.GetVirtualIPProperties().GetPersistenceType()),
		string(model.GetVirtualIPProperties().GetPersistenceCookieName()),
		int(model.GetVirtualIPProperties().GetConnectionLimit()),
		bool(model.GetVirtualIPProperties().GetAdminState()),
		string(model.GetVirtualIPProperties().GetAddress()),
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

	for _, ref := range model.LoadbalancerPoolRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("loadbalancer_pool"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "LoadbalancerPoolRefs create failed")
		}
	}

	for _, ref := range model.VirtualMachineInterfaceRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("virtual_machine_interface"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "virtual_ip",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "virtual_ip", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanVirtualIP(values map[string]interface{}) (*models.VirtualIP, error) {
	m := models.MakeVirtualIP()

	if value, ok := values["subnet_id"]; ok {

		m.VirtualIPProperties.SubnetID = common.InterfaceToString(value)

	}

	if value, ok := values["status_description"]; ok {

		m.VirtualIPProperties.StatusDescription = common.InterfaceToString(value)

	}

	if value, ok := values["status"]; ok {

		m.VirtualIPProperties.Status = common.InterfaceToString(value)

	}

	if value, ok := values["protocol_port"]; ok {

		m.VirtualIPProperties.ProtocolPort = common.InterfaceToInt64(value)

	}

	if value, ok := values["protocol"]; ok {

		m.VirtualIPProperties.Protocol = common.InterfaceToString(value)

	}

	if value, ok := values["persistence_type"]; ok {

		m.VirtualIPProperties.PersistenceType = common.InterfaceToString(value)

	}

	if value, ok := values["persistence_cookie_name"]; ok {

		m.VirtualIPProperties.PersistenceCookieName = common.InterfaceToString(value)

	}

	if value, ok := values["connection_limit"]; ok {

		m.VirtualIPProperties.ConnectionLimit = common.InterfaceToInt64(value)

	}

	if value, ok := values["admin_state"]; ok {

		m.VirtualIPProperties.AdminState = common.InterfaceToBool(value)

	}

	if value, ok := values["address"]; ok {

		m.VirtualIPProperties.Address = common.InterfaceToString(value)

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

	if value, ok := values["ref_loadbalancer_pool"]; ok {
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
			referenceModel := &models.VirtualIPLoadbalancerPoolRef{}
			referenceModel.UUID = uuid
			m.LoadbalancerPoolRefs = append(m.LoadbalancerPoolRefs, referenceModel)

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
			referenceModel := &models.VirtualIPVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	return m, nil
}

// ListVirtualIP lists VirtualIP with list spec.
func (db *DB) listVirtualIP(ctx context.Context, request *models.ListVirtualIPRequest) (response *models.ListVirtualIPResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["virtual_ip"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.VirtualIP{}

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
		m, err := scanVirtualIP(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListVirtualIPResponse{
		VirtualIPs: result,
	}
	return response, nil
}

// UpdateVirtualIP updates a resource
func (db *DB) updateVirtualIP(
	ctx context.Context,
	request *models.UpdateVirtualIPRequest,
) error {
	//TODO
	return nil
}

// DeleteVirtualIP deletes a resource
func (db *DB) deleteVirtualIP(
	ctx context.Context,
	request *models.DeleteVirtualIPRequest) error {
	qb := db.queryBuilders["virtual_ip"]

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

//CreateVirtualIP handle a Create API
// nolint
func (db *DB) CreateVirtualIP(
	ctx context.Context,
	request *models.CreateVirtualIPRequest) (*models.CreateVirtualIPResponse, error) {
	model := request.VirtualIP
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createVirtualIP(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_ip",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateVirtualIPResponse{
		VirtualIP: request.VirtualIP,
	}, nil
}

//UpdateVirtualIP handles a Update request.
func (db *DB) UpdateVirtualIP(
	ctx context.Context,
	request *models.UpdateVirtualIPRequest) (*models.UpdateVirtualIPResponse, error) {
	model := request.VirtualIP
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateVirtualIP(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_ip",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateVirtualIPResponse{
		VirtualIP: model,
	}, nil
}

//DeleteVirtualIP delete a resource.
func (db *DB) DeleteVirtualIP(ctx context.Context, request *models.DeleteVirtualIPRequest) (*models.DeleteVirtualIPResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteVirtualIP(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteVirtualIPResponse{
		ID: request.ID,
	}, nil
}

//GetVirtualIP a Get request.
func (db *DB) GetVirtualIP(ctx context.Context, request *models.GetVirtualIPRequest) (response *models.GetVirtualIPResponse, err error) {
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
	listRequest := &models.ListVirtualIPRequest{
		Spec: spec,
	}
	var result *models.ListVirtualIPResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listVirtualIP(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.VirtualIPs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetVirtualIPResponse{
		VirtualIP: result.VirtualIPs[0],
	}
	return response, nil
}

//ListVirtualIP handles a List service Request.
// nolint
func (db *DB) ListVirtualIP(
	ctx context.Context,
	request *models.ListVirtualIPRequest) (response *models.ListVirtualIPResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listVirtualIP(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
