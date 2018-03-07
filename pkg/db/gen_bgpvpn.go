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

// BGPVPNFields is db columns for BGPVPN
var BGPVPNFields = []string{
	"uuid",
	"route_target",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"import_route_target_list_route_target",
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
	"export_route_target_list_route_target",
	"display_name",
	"configuration_version",
	"bgpvpn_type",
	"key_value_pair",
}

// BGPVPNRefFields is db reference fields for BGPVPN
var BGPVPNRefFields = map[string][]string{}

// BGPVPNBackRefFields is db back reference fields for BGPVPN
var BGPVPNBackRefFields = map[string][]string{}

// BGPVPNParentTypes is possible parents for BGPVPN
var BGPVPNParents = []string{

	"project",
}

// CreateBGPVPN inserts BGPVPN to DB
// nolint
func (db *DB) createBGPVPN(
	ctx context.Context,
	request *models.CreateBGPVPNRequest) error {
	qb := db.queryBuilders["bgpvpn"]
	tx := GetTransaction(ctx)
	model := request.BGPVPN
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
		common.MustJSON(model.GetRouteTargetList().GetRouteTarget()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		common.MustJSON(model.GetImportRouteTargetList().GetRouteTarget()),
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
		common.MustJSON(model.GetExportRouteTargetList().GetRouteTarget()),
		string(model.GetDisplayName()),
		int(model.GetConfigurationVersion()),
		string(model.GetBGPVPNType()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "bgpvpn",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "bgpvpn", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanBGPVPN(values map[string]interface{}) (*models.BGPVPN, error) {
	m := models.MakeBGPVPN()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.RouteTargetList.RouteTarget)

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

	if value, ok := values["import_route_target_list_route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.ImportRouteTargetList.RouteTarget)

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

	if value, ok := values["export_route_target_list_route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.ExportRouteTargetList.RouteTarget)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["configuration_version"]; ok {

		m.ConfigurationVersion = common.InterfaceToInt64(value)

	}

	if value, ok := values["bgpvpn_type"]; ok {

		m.BGPVPNType = common.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListBGPVPN lists BGPVPN with list spec.
func (db *DB) listBGPVPN(ctx context.Context, request *models.ListBGPVPNRequest) (response *models.ListBGPVPNResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["bgpvpn"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.BGPVPN{}

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
		m, err := scanBGPVPN(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListBGPVPNResponse{
		BGPVPNs: result,
	}
	return response, nil
}

// UpdateBGPVPN updates a resource
func (db *DB) updateBGPVPN(
	ctx context.Context,
	request *models.UpdateBGPVPNRequest,
) error {
	//TODO
	return nil
}

// DeleteBGPVPN deletes a resource
func (db *DB) deleteBGPVPN(
	ctx context.Context,
	request *models.DeleteBGPVPNRequest) error {
	qb := db.queryBuilders["bgpvpn"]

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

//CreateBGPVPN handle a Create API
// nolint
func (db *DB) CreateBGPVPN(
	ctx context.Context,
	request *models.CreateBGPVPNRequest) (*models.CreateBGPVPNResponse, error) {
	model := request.BGPVPN
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createBGPVPN(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgpvpn",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateBGPVPNResponse{
		BGPVPN: request.BGPVPN,
	}, nil
}

//UpdateBGPVPN handles a Update request.
func (db *DB) UpdateBGPVPN(
	ctx context.Context,
	request *models.UpdateBGPVPNRequest) (*models.UpdateBGPVPNResponse, error) {
	model := request.BGPVPN
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateBGPVPN(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgpvpn",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateBGPVPNResponse{
		BGPVPN: model,
	}, nil
}

//DeleteBGPVPN delete a resource.
func (db *DB) DeleteBGPVPN(ctx context.Context, request *models.DeleteBGPVPNRequest) (*models.DeleteBGPVPNResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteBGPVPN(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteBGPVPNResponse{
		ID: request.ID,
	}, nil
}

//GetBGPVPN a Get request.
func (db *DB) GetBGPVPN(ctx context.Context, request *models.GetBGPVPNRequest) (response *models.GetBGPVPNResponse, err error) {
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
	listRequest := &models.ListBGPVPNRequest{
		Spec: spec,
	}
	var result *models.ListBGPVPNResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listBGPVPN(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.BGPVPNs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetBGPVPNResponse{
		BGPVPN: result.BGPVPNs[0],
	}
	return response, nil
}

//ListBGPVPN handles a List service Request.
// nolint
func (db *DB) ListBGPVPN(
	ctx context.Context,
	request *models.ListBGPVPNRequest) (response *models.ListBGPVPNResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listBGPVPN(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
