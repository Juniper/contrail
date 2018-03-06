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

// LogicalRouterFields is db columns for LogicalRouter
var LogicalRouterFields = []string{
	"vxlan_network_identifier",
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
	"route_target",
	"key_value_pair",
}

// LogicalRouterRefFields is db reference fields for LogicalRouter
var LogicalRouterRefFields = map[string][]string{

	"route_target": []string{
	// <schema.Schema Value>

	},

	"virtual_machine_interface": []string{
	// <schema.Schema Value>

	},

	"service_instance": []string{
	// <schema.Schema Value>

	},

	"route_table": []string{
	// <schema.Schema Value>

	},

	"virtual_network": []string{
	// <schema.Schema Value>

	},

	"physical_router": []string{
	// <schema.Schema Value>

	},

	"bgpvpn": []string{
	// <schema.Schema Value>

	},
}

// LogicalRouterBackRefFields is db back reference fields for LogicalRouter
var LogicalRouterBackRefFields = map[string][]string{}

// LogicalRouterParentTypes is possible parents for LogicalRouter
var LogicalRouterParents = []string{

	"project",
}

// CreateLogicalRouter inserts LogicalRouter to DB
// nolint
func (db *DB) createLogicalRouter(
	ctx context.Context,
	request *models.CreateLogicalRouterRequest) error {
	qb := db.queryBuilders["logical_router"]
	tx := GetTransaction(ctx)
	model := request.LogicalRouter
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetVxlanNetworkIdentifier()),
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
		common.MustJSON(model.GetConfiguredRouteTargetList().GetRouteTarget()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	for _, ref := range model.PhysicalRouterRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("physical_router"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "PhysicalRouterRefs create failed")
		}
	}

	for _, ref := range model.BGPVPNRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("bgpvpn"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "BGPVPNRefs create failed")
		}
	}

	for _, ref := range model.RouteTargetRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("route_target"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "RouteTargetRefs create failed")
		}
	}

	for _, ref := range model.VirtualMachineInterfaceRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("virtual_machine_interface"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	for _, ref := range model.ServiceInstanceRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("service_instance"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs create failed")
		}
	}

	for _, ref := range model.RouteTableRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("route_table"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "RouteTableRefs create failed")
		}
	}

	for _, ref := range model.VirtualNetworkRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("virtual_network"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "logical_router",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "logical_router", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanLogicalRouter(values map[string]interface{}) (*models.LogicalRouter, error) {
	m := models.MakeLogicalRouter()

	if value, ok := values["vxlan_network_identifier"]; ok {

		m.VxlanNetworkIdentifier = common.InterfaceToString(value)

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

	if value, ok := values["route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.ConfiguredRouteTargetList.RouteTarget)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_route_target"]; ok {
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
			referenceModel := &models.LogicalRouterRouteTargetRef{}
			referenceModel.UUID = uuid
			m.RouteTargetRefs = append(m.RouteTargetRefs, referenceModel)

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
			referenceModel := &models.LogicalRouterVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

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
			referenceModel := &models.LogicalRouterServiceInstanceRef{}
			referenceModel.UUID = uuid
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_route_table"]; ok {
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
			referenceModel := &models.LogicalRouterRouteTableRef{}
			referenceModel.UUID = uuid
			m.RouteTableRefs = append(m.RouteTableRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_network"]; ok {
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
			referenceModel := &models.LogicalRouterVirtualNetworkRef{}
			referenceModel.UUID = uuid
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
	}

	if value, ok := values["ref_physical_router"]; ok {
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
			referenceModel := &models.LogicalRouterPhysicalRouterRef{}
			referenceModel.UUID = uuid
			m.PhysicalRouterRefs = append(m.PhysicalRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_bgpvpn"]; ok {
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
			referenceModel := &models.LogicalRouterBGPVPNRef{}
			referenceModel.UUID = uuid
			m.BGPVPNRefs = append(m.BGPVPNRefs, referenceModel)

		}
	}

	return m, nil
}

// ListLogicalRouter lists LogicalRouter with list spec.
func (db *DB) listLogicalRouter(ctx context.Context, request *models.ListLogicalRouterRequest) (response *models.ListLogicalRouterResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["logical_router"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.LogicalRouter{}

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
		m, err := scanLogicalRouter(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListLogicalRouterResponse{
		LogicalRouters: result,
	}
	return response, nil
}

// UpdateLogicalRouter updates a resource
func (db *DB) updateLogicalRouter(
	ctx context.Context,
	request *models.UpdateLogicalRouterRequest,
) error {
	//TODO
	return nil
}

// DeleteLogicalRouter deletes a resource
func (db *DB) deleteLogicalRouter(
	ctx context.Context,
	request *models.DeleteLogicalRouterRequest) error {
	qb := db.queryBuilders["logical_router"]

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

//CreateLogicalRouter handle a Create API
// nolint
func (db *DB) CreateLogicalRouter(
	ctx context.Context,
	request *models.CreateLogicalRouterRequest) (*models.CreateLogicalRouterResponse, error) {
	model := request.LogicalRouter
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createLogicalRouter(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "logical_router",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateLogicalRouterResponse{
		LogicalRouter: request.LogicalRouter,
	}, nil
}

//UpdateLogicalRouter handles a Update request.
func (db *DB) UpdateLogicalRouter(
	ctx context.Context,
	request *models.UpdateLogicalRouterRequest) (*models.UpdateLogicalRouterResponse, error) {
	model := request.LogicalRouter
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateLogicalRouter(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "logical_router",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateLogicalRouterResponse{
		LogicalRouter: model,
	}, nil
}

//DeleteLogicalRouter delete a resource.
func (db *DB) DeleteLogicalRouter(ctx context.Context, request *models.DeleteLogicalRouterRequest) (*models.DeleteLogicalRouterResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteLogicalRouter(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteLogicalRouterResponse{
		ID: request.ID,
	}, nil
}

//GetLogicalRouter a Get request.
func (db *DB) GetLogicalRouter(ctx context.Context, request *models.GetLogicalRouterRequest) (response *models.GetLogicalRouterResponse, err error) {
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
	listRequest := &models.ListLogicalRouterRequest{
		Spec: spec,
	}
	var result *models.ListLogicalRouterResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listLogicalRouter(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.LogicalRouters) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetLogicalRouterResponse{
		LogicalRouter: result.LogicalRouters[0],
	}
	return response, nil
}

//ListLogicalRouter handles a List service Request.
// nolint
func (db *DB) ListLogicalRouter(
	ctx context.Context,
	request *models.ListLogicalRouterRequest) (response *models.ListLogicalRouterResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listLogicalRouter(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
