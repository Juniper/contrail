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

// ContrailConfigNodeFields is db columns for ContrailConfigNode
var ContrailConfigNodeFields = []string{
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
	"configuration_version",
	"key_value_pair",
}

// ContrailConfigNodeRefFields is db reference fields for ContrailConfigNode
var ContrailConfigNodeRefFields = map[string][]string{

	"node": []string{
	// <schema.Schema Value>

	},
}

// ContrailConfigNodeBackRefFields is db back reference fields for ContrailConfigNode
var ContrailConfigNodeBackRefFields = map[string][]string{}

// ContrailConfigNodeParentTypes is possible parents for ContrailConfigNode
var ContrailConfigNodeParents = []string{

	"contrail_cluster",
}

// CreateContrailConfigNode inserts ContrailConfigNode to DB
// nolint
func (db *DB) createContrailConfigNode(
	ctx context.Context,
	request *models.CreateContrailConfigNodeRequest) error {
	qb := db.queryBuilders["contrail_config_node"]
	tx := GetTransaction(ctx)
	model := request.ContrailConfigNode
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
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
		int(model.GetConfigurationVersion()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		log.WithFields(log.Fields{
			"model": model,
			"err":   err}).Debug("create failed")
		return errors.Wrap(err, "create failed")
	}

	for _, ref := range model.NodeRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("node"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "NodeRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "contrail_config_node",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "contrail_config_node", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanContrailConfigNode(values map[string]interface{}) (*models.ContrailConfigNode, error) {
	m := models.MakeContrailConfigNode()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

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

	if value, ok := values["configuration_version"]; ok {

		m.ConfigurationVersion = common.InterfaceToInt64(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_node"]; ok {
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
			referenceModel := &models.ContrailConfigNodeNodeRef{}
			referenceModel.UUID = uuid
			m.NodeRefs = append(m.NodeRefs, referenceModel)

		}
	}

	return m, nil
}

// ListContrailConfigNode lists ContrailConfigNode with list spec.
func (db *DB) listContrailConfigNode(ctx context.Context, request *models.ListContrailConfigNodeRequest) (response *models.ListContrailConfigNodeResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["contrail_config_node"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.ContrailConfigNode{}

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
		m, err := scanContrailConfigNode(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListContrailConfigNodeResponse{
		ContrailConfigNodes: result,
	}
	return response, nil
}

// UpdateContrailConfigNode updates a resource
func (db *DB) updateContrailConfigNode(
	ctx context.Context,
	request *models.UpdateContrailConfigNodeRequest,
) error {
	//TODO
	return nil
}

// DeleteContrailConfigNode deletes a resource
func (db *DB) deleteContrailConfigNode(
	ctx context.Context,
	request *models.DeleteContrailConfigNodeRequest) error {
	qb := db.queryBuilders["contrail_config_node"]

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

//CreateContrailConfigNode handle a Create API
// nolint
func (db *DB) CreateContrailConfigNode(
	ctx context.Context,
	request *models.CreateContrailConfigNodeRequest) (*models.CreateContrailConfigNodeResponse, error) {
	model := request.ContrailConfigNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createContrailConfigNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_config_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateContrailConfigNodeResponse{
		ContrailConfigNode: request.ContrailConfigNode,
	}, nil
}

//UpdateContrailConfigNode handles a Update request.
func (db *DB) UpdateContrailConfigNode(
	ctx context.Context,
	request *models.UpdateContrailConfigNodeRequest) (*models.UpdateContrailConfigNodeResponse, error) {
	model := request.ContrailConfigNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateContrailConfigNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_config_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateContrailConfigNodeResponse{
		ContrailConfigNode: model,
	}, nil
}

//DeleteContrailConfigNode delete a resource.
func (db *DB) DeleteContrailConfigNode(ctx context.Context, request *models.DeleteContrailConfigNodeRequest) (*models.DeleteContrailConfigNodeResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteContrailConfigNode(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteContrailConfigNodeResponse{
		ID: request.ID,
	}, nil
}

//GetContrailConfigNode a Get request.
func (db *DB) GetContrailConfigNode(ctx context.Context, request *models.GetContrailConfigNodeRequest) (response *models.GetContrailConfigNodeResponse, err error) {
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
	listRequest := &models.ListContrailConfigNodeRequest{
		Spec: spec,
	}
	var result *models.ListContrailConfigNodeResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listContrailConfigNode(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ContrailConfigNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetContrailConfigNodeResponse{
		ContrailConfigNode: result.ContrailConfigNodes[0],
	}
	return response, nil
}

//ListContrailConfigNode handles a List service Request.
// nolint
func (db *DB) ListContrailConfigNode(
	ctx context.Context,
	request *models.ListContrailConfigNodeRequest) (response *models.ListContrailConfigNodeResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listContrailConfigNode(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
