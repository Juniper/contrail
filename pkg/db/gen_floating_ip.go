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

// FloatingIPFields is db columns for FloatingIP
var FloatingIPFields = []string{
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
	"floating_ip_traffic_direction",
	"port_mappings",
	"floating_ip_port_mappings_enable",
	"floating_ip_is_virtual_ip",
	"floating_ip_fixed_ip_address",
	"floating_ip_address_family",
	"floating_ip_address",
	"display_name",
	"key_value_pair",
}

// FloatingIPRefFields is db reference fields for FloatingIP
var FloatingIPRefFields = map[string][]string{

	"project": []string{
	// <schema.Schema Value>

	},

	"virtual_machine_interface": []string{
	// <schema.Schema Value>

	},
}

// FloatingIPBackRefFields is db back reference fields for FloatingIP
var FloatingIPBackRefFields = map[string][]string{}

// FloatingIPParentTypes is possible parents for FloatingIP
var FloatingIPParents = []string{

	"instance_ip",

	"floating_ip_pool",
}

// CreateFloatingIP inserts FloatingIP to DB
// nolint
func (db *DB) createFloatingIP(
	ctx context.Context,
	request *models.CreateFloatingIPRequest) error {
	qb := db.queryBuilders["floating_ip"]
	tx := GetTransaction(ctx)
	model := request.FloatingIP
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
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
		string(model.GetFloatingIPTrafficDirection()),
		common.MustJSON(model.GetFloatingIPPortMappings().GetPortMappings()),
		bool(model.GetFloatingIPPortMappingsEnable()),
		bool(model.GetFloatingIPIsVirtualIP()),
		string(model.GetFloatingIPFixedIPAddress()),
		string(model.GetFloatingIPAddressFamily()),
		string(model.GetFloatingIPAddress()),
		string(model.GetDisplayName()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	for _, ref := range model.ProjectRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("project"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ProjectRefs create failed")
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
		Type:   "floating_ip",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "floating_ip", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanFloatingIP(values map[string]interface{}) (*models.FloatingIP, error) {
	m := models.MakeFloatingIP()

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

	if value, ok := values["floating_ip_traffic_direction"]; ok {

		m.FloatingIPTrafficDirection = common.InterfaceToString(value)

	}

	if value, ok := values["port_mappings"]; ok {

		json.Unmarshal(value.([]byte), &m.FloatingIPPortMappings.PortMappings)

	}

	if value, ok := values["floating_ip_port_mappings_enable"]; ok {

		m.FloatingIPPortMappingsEnable = common.InterfaceToBool(value)

	}

	if value, ok := values["floating_ip_is_virtual_ip"]; ok {

		m.FloatingIPIsVirtualIP = common.InterfaceToBool(value)

	}

	if value, ok := values["floating_ip_fixed_ip_address"]; ok {

		m.FloatingIPFixedIPAddress = common.InterfaceToString(value)

	}

	if value, ok := values["floating_ip_address_family"]; ok {

		m.FloatingIPAddressFamily = common.InterfaceToString(value)

	}

	if value, ok := values["floating_ip_address"]; ok {

		m.FloatingIPAddress = common.InterfaceToString(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_project"]; ok {
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
			referenceModel := &models.FloatingIPProjectRef{}
			referenceModel.UUID = uuid
			m.ProjectRefs = append(m.ProjectRefs, referenceModel)

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
			referenceModel := &models.FloatingIPVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	return m, nil
}

// ListFloatingIP lists FloatingIP with list spec.
func (db *DB) listFloatingIP(ctx context.Context, request *models.ListFloatingIPRequest) (response *models.ListFloatingIPResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["floating_ip"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.FloatingIP{}

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
		m, err := scanFloatingIP(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListFloatingIPResponse{
		FloatingIPs: result,
	}
	return response, nil
}

// UpdateFloatingIP updates a resource
func (db *DB) updateFloatingIP(
	ctx context.Context,
	request *models.UpdateFloatingIPRequest,
) error {
	//TODO
	return nil
}

// DeleteFloatingIP deletes a resource
func (db *DB) deleteFloatingIP(
	ctx context.Context,
	request *models.DeleteFloatingIPRequest) error {
	qb := db.queryBuilders["floating_ip"]

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

//CreateFloatingIP handle a Create API
// nolint
func (db *DB) CreateFloatingIP(
	ctx context.Context,
	request *models.CreateFloatingIPRequest) (*models.CreateFloatingIPResponse, error) {
	model := request.FloatingIP
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createFloatingIP(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "floating_ip",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateFloatingIPResponse{
		FloatingIP: request.FloatingIP,
	}, nil
}

//UpdateFloatingIP handles a Update request.
func (db *DB) UpdateFloatingIP(
	ctx context.Context,
	request *models.UpdateFloatingIPRequest) (*models.UpdateFloatingIPResponse, error) {
	model := request.FloatingIP
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateFloatingIP(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "floating_ip",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateFloatingIPResponse{
		FloatingIP: model,
	}, nil
}

//DeleteFloatingIP delete a resource.
func (db *DB) DeleteFloatingIP(ctx context.Context, request *models.DeleteFloatingIPRequest) (*models.DeleteFloatingIPResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteFloatingIP(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteFloatingIPResponse{
		ID: request.ID,
	}, nil
}

//GetFloatingIP a Get request.
func (db *DB) GetFloatingIP(ctx context.Context, request *models.GetFloatingIPRequest) (response *models.GetFloatingIPResponse, err error) {
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
	listRequest := &models.ListFloatingIPRequest{
		Spec: spec,
	}
	var result *models.ListFloatingIPResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listFloatingIP(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.FloatingIPs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetFloatingIPResponse{
		FloatingIP: result.FloatingIPs[0],
	}
	return response, nil
}

//ListFloatingIP handles a List service Request.
// nolint
func (db *DB) ListFloatingIP(
	ctx context.Context,
	request *models.ListFloatingIPRequest) (response *models.ListFloatingIPResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listFloatingIP(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
