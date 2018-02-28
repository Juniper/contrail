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

const insertVPNGroupQuery = "insert into `vpn_group` (`uuid`,`type`,`provisioning_state`,`provisioning_start_time`,`provisioning_progress_stage`,`provisioning_progress`,`provisioning_log`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteVPNGroupQuery = "delete from `vpn_group` where uuid = ?"

// VPNGroupFields is db columns for VPNGroup
var VPNGroupFields = []string{
	"uuid",
	"type",
	"provisioning_state",
	"provisioning_start_time",
	"provisioning_progress_stage",
	"provisioning_progress",
	"provisioning_log",
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

// VPNGroupRefFields is db reference fields for VPNGroup
var VPNGroupRefFields = map[string][]string{

	"location": []string{
	// <schema.Schema Value>

	},
}

// VPNGroupBackRefFields is db back reference fields for VPNGroup
var VPNGroupBackRefFields = map[string][]string{}

// VPNGroupParentTypes is possible parents for VPNGroup
var VPNGroupParents = []string{}

const insertVPNGroupLocationQuery = "insert into `ref_vpn_group_location` (`from`, `to` ) values (?, ?);"

// CreateVPNGroup inserts VPNGroup to DB
// nolint
func (db *DB) createVPNGroup(
	ctx context.Context,
	request *models.CreateVPNGroupRequest) error {
	tx := GetTransaction(ctx)
	model := request.VPNGroup
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVPNGroupQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertVPNGroupQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		string(model.GetType()),
		string(model.GetProvisioningState()),
		string(model.GetProvisioningStartTime()),
		string(model.GetProvisioningProgressStage()),
		int(model.GetProvisioningProgress()),
		string(model.GetProvisioningLog()),
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

	stmtLocationRef, err := tx.Prepare(insertVPNGroupLocationQuery)
	if err != nil {
		return errors.Wrap(err, "preparing LocationRefs create statement failed")
	}
	defer stmtLocationRef.Close()
	for _, ref := range model.LocationRefs {

		_, err = stmtLocationRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "LocationRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "vpn_group",
		FQName: model.FQName,
	}
	err = CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = CreateSharing(tx, "vpn_group", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

// nolint
func scanVPNGroup(values map[string]interface{}) (*models.VPNGroup, error) {
	m := models.MakeVPNGroup()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["type"]; ok {

		m.Type = common.InterfaceToString(value)

	}

	if value, ok := values["provisioning_state"]; ok {

		m.ProvisioningState = common.InterfaceToString(value)

	}

	if value, ok := values["provisioning_start_time"]; ok {

		m.ProvisioningStartTime = common.InterfaceToString(value)

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		m.ProvisioningProgressStage = common.InterfaceToString(value)

	}

	if value, ok := values["provisioning_progress"]; ok {

		m.ProvisioningProgress = common.InterfaceToInt64(value)

	}

	if value, ok := values["provisioning_log"]; ok {

		m.ProvisioningLog = common.InterfaceToString(value)

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

	if value, ok := values["ref_location"]; ok {
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
			referenceModel := &models.VPNGroupLocationRef{}
			referenceModel.UUID = uuid
			m.LocationRefs = append(m.LocationRefs, referenceModel)

		}
	}

	return m, nil
}

// ListVPNGroup lists VPNGroup with list spec.
// nolint
func (db *DB) listVPNGroup(ctx context.Context, request *models.ListVPNGroupRequest) (response *models.ListVPNGroupResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)
	qb := &ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "vpn_group"
	qb.Fields = VPNGroupFields
	qb.RefFields = VPNGroupRefFields
	qb.BackRefFields = VPNGroupBackRefFields
	result := []*models.VPNGroup{}

	if spec.ParentFQName != nil {
		parentMetaData, err := GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = models.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
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
		m, err := scanVPNGroup(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListVPNGroupResponse{
		VPNGroups: result,
	}
	return response, nil
}

// UpdateVPNGroup updates a resource
// nolint
func (db *DB) updateVPNGroup(
	ctx context.Context,
	request *models.UpdateVPNGroupRequest,
) error {
	//TODO
	return nil
}

// DeleteVPNGroup deletes a resource
// nolint
func (db *DB) deleteVPNGroup(
	ctx context.Context,
	request *models.DeleteVPNGroupRequest) error {
	deleteQuery := deleteVPNGroupQuery
	selectQuery := "select count(uuid) from vpn_group where uuid = ?"
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

	err = DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}

//CreateVPNGroup handle a Create API
// nolint
func (db *DB) CreateVPNGroup(
	ctx context.Context,
	request *models.CreateVPNGroupRequest) (*models.CreateVPNGroupResponse, error) {
	model := request.VPNGroup
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createVPNGroup(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "vpn_group",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateVPNGroupResponse{
		VPNGroup: request.VPNGroup,
	}, nil
}

//UpdateVPNGroup handles a Update request.
func (db *DB) UpdateVPNGroup(
	ctx context.Context,
	request *models.UpdateVPNGroupRequest) (*models.UpdateVPNGroupResponse, error) {
	model := request.VPNGroup
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateVPNGroup(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "vpn_group",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateVPNGroupResponse{
		VPNGroup: model,
	}, nil
}

//DeleteVPNGroup delete a resource.
func (db *DB) DeleteVPNGroup(ctx context.Context, request *models.DeleteVPNGroupRequest) (*models.DeleteVPNGroupResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteVPNGroup(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteVPNGroupResponse{
		ID: request.ID,
	}, nil
}

//GetVPNGroup a Get request.
// nolint
func (db *DB) GetVPNGroup(ctx context.Context, request *models.GetVPNGroupRequest) (response *models.GetVPNGroupResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListVPNGroupRequest{
		Spec: spec,
	}
	var result *models.ListVPNGroupResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listVPNGroup(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.VPNGroups) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetVPNGroupResponse{
		VPNGroup: result.VPNGroups[0],
	}
	return response, nil
}

//ListVPNGroup handles a List service Request.
// nolint
func (db *DB) ListVPNGroup(
	ctx context.Context,
	request *models.ListVPNGroupRequest) (response *models.ListVPNGroupResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listVPNGroup(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
