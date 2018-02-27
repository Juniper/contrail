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

const insertContrailAnalyticsNodeQuery = "insert into `contrail_analytics_node` (`uuid`,`provisioning_state`,`provisioning_start_time`,`provisioning_progress_stage`,`provisioning_progress`,`provisioning_log`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteContrailAnalyticsNodeQuery = "delete from `contrail_analytics_node` where uuid = ?"

// ContrailAnalyticsNodeFields is db columns for ContrailAnalyticsNode
var ContrailAnalyticsNodeFields = []string{
	"uuid",
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

// ContrailAnalyticsNodeRefFields is db reference fields for ContrailAnalyticsNode
var ContrailAnalyticsNodeRefFields = map[string][]string{

	"node": []string{
	// <schema.Schema Value>

	},
}

// ContrailAnalyticsNodeBackRefFields is db back reference fields for ContrailAnalyticsNode
var ContrailAnalyticsNodeBackRefFields = map[string][]string{}

// ContrailAnalyticsNodeParentTypes is possible parents for ContrailAnalyticsNode
var ContrailAnalyticsNodeParents = []string{

	"contrail_cluster",
}

const insertContrailAnalyticsNodeNodeQuery = "insert into `ref_contrail_analytics_node_node` (`from`, `to` ) values (?, ?);"

// CreateContrailAnalyticsNode inserts ContrailAnalyticsNode to DB
func (db *DB) createContrailAnalyticsNode(
	ctx context.Context,
	request *models.CreateContrailAnalyticsNodeRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.ContrailAnalyticsNode
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailAnalyticsNodeQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertContrailAnalyticsNodeQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
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

	stmtNodeRef, err := tx.Prepare(insertContrailAnalyticsNodeNodeQuery)
	if err != nil {
		return errors.Wrap(err, "preparing NodeRefs create statement failed")
	}
	defer stmtNodeRef.Close()
	for _, ref := range model.NodeRefs {

		_, err = stmtNodeRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "NodeRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "contrail_analytics_node",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "contrail_analytics_node", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanContrailAnalyticsNode(values map[string]interface{}) (*models.ContrailAnalyticsNode, error) {
	m := models.MakeContrailAnalyticsNode()

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["provisioning_state"]; ok {

		m.ProvisioningState = schema.InterfaceToString(value)

	}

	if value, ok := values["provisioning_start_time"]; ok {

		m.ProvisioningStartTime = schema.InterfaceToString(value)

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		m.ProvisioningProgressStage = schema.InterfaceToString(value)

	}

	if value, ok := values["provisioning_progress"]; ok {

		m.ProvisioningProgress = schema.InterfaceToInt64(value)

	}

	if value, ok := values["provisioning_log"]; ok {

		m.ProvisioningLog = schema.InterfaceToString(value)

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

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_node"]; ok {
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
			referenceModel := &models.ContrailAnalyticsNodeNodeRef{}
			referenceModel.UUID = uuid
			m.NodeRefs = append(m.NodeRefs, referenceModel)

		}
	}

	return m, nil
}

// ListContrailAnalyticsNode lists ContrailAnalyticsNode with list spec.
func (db *DB) listContrailAnalyticsNode(ctx context.Context, request *models.ListContrailAnalyticsNodeRequest) (response *models.ListContrailAnalyticsNodeResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "contrail_analytics_node"
	qb.Fields = ContrailAnalyticsNodeFields
	qb.RefFields = ContrailAnalyticsNodeRefFields
	qb.BackRefFields = ContrailAnalyticsNodeBackRefFields
	result := []*models.ContrailAnalyticsNode{}

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
		m, err := scanContrailAnalyticsNode(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListContrailAnalyticsNodeResponse{
		ContrailAnalyticsNodes: result,
	}
	return response, nil
}

// UpdateContrailAnalyticsNode updates a resource
func (db *DB) updateContrailAnalyticsNode(
	ctx context.Context,
	request *models.UpdateContrailAnalyticsNodeRequest,
) error {
	//TODO
	return nil
}

// DeleteContrailAnalyticsNode deletes a resource
func (db *DB) deleteContrailAnalyticsNode(
	ctx context.Context,
	request *models.DeleteContrailAnalyticsNodeRequest) error {
	deleteQuery := deleteContrailAnalyticsNodeQuery
	selectQuery := "select count(uuid) from contrail_analytics_node where uuid = ?"
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

//CreateContrailAnalyticsNode handle a Create API
func (db *DB) CreateContrailAnalyticsNode(
	ctx context.Context,
	request *models.CreateContrailAnalyticsNodeRequest) (*models.CreateContrailAnalyticsNodeResponse, error) {
	model := request.ContrailAnalyticsNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createContrailAnalyticsNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateContrailAnalyticsNodeResponse{
		ContrailAnalyticsNode: request.ContrailAnalyticsNode,
	}, nil
}

//UpdateContrailAnalyticsNode handles a Update request.
func (db *DB) UpdateContrailAnalyticsNode(
	ctx context.Context,
	request *models.UpdateContrailAnalyticsNodeRequest) (*models.UpdateContrailAnalyticsNodeResponse, error) {
	model := request.ContrailAnalyticsNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateContrailAnalyticsNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateContrailAnalyticsNodeResponse{
		ContrailAnalyticsNode: model,
	}, nil
}

//DeleteContrailAnalyticsNode delete a resource.
func (db *DB) DeleteContrailAnalyticsNode(ctx context.Context, request *models.DeleteContrailAnalyticsNodeRequest) (*models.DeleteContrailAnalyticsNodeResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteContrailAnalyticsNode(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteContrailAnalyticsNodeResponse{
		ID: request.ID,
	}, nil
}

//GetContrailAnalyticsNode a Get request.
func (db *DB) GetContrailAnalyticsNode(ctx context.Context, request *models.GetContrailAnalyticsNodeRequest) (response *models.GetContrailAnalyticsNodeResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListContrailAnalyticsNodeRequest{
		Spec: spec,
	}
	var result *models.ListContrailAnalyticsNodeResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listContrailAnalyticsNode(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ContrailAnalyticsNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetContrailAnalyticsNodeResponse{
		ContrailAnalyticsNode: result.ContrailAnalyticsNodes[0],
	}
	return response, nil
}

//ListContrailAnalyticsNode handles a List service Request.
func (db *DB) ListContrailAnalyticsNode(
	ctx context.Context,
	request *models.ListContrailAnalyticsNodeRequest) (response *models.ListContrailAnalyticsNodeResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listContrailAnalyticsNode(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
