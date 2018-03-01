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

const insertKubernetesNodeQuery = "insert into `kubernetes_node` (`uuid`,`provisioning_state`,`provisioning_start_time`,`provisioning_progress_stage`,`provisioning_progress`,`provisioning_log`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteKubernetesNodeQuery = "delete from `kubernetes_node` where uuid = ?"

// KubernetesNodeFields is db columns for KubernetesNode
var KubernetesNodeFields = []string{
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

// KubernetesNodeRefFields is db reference fields for KubernetesNode
var KubernetesNodeRefFields = map[string][]string{

	"node": []string{
	// <schema.Schema Value>

	},
}

// KubernetesNodeBackRefFields is db back reference fields for KubernetesNode
var KubernetesNodeBackRefFields = map[string][]string{}

// KubernetesNodeParentTypes is possible parents for KubernetesNode
var KubernetesNodeParents = []string{

	"contrail_cluster",
}

const insertKubernetesNodeNodeQuery = "insert into `ref_kubernetes_node_node` (`from`, `to` ) values (?, ?);"

// CreateKubernetesNode inserts KubernetesNode to DB
func (db *DB) createKubernetesNode(
	ctx context.Context,
	request *models.CreateKubernetesNodeRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.KubernetesNode
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertKubernetesNodeQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertKubernetesNodeQuery,
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

	stmtNodeRef, err := tx.Prepare(insertKubernetesNodeNodeQuery)
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
		Type:   "kubernetes_node",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "kubernetes_node", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanKubernetesNode(values map[string]interface{}) (*models.KubernetesNode, error) {
	m := models.MakeKubernetesNode()

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
			referenceModel := &models.KubernetesNodeNodeRef{}
			referenceModel.UUID = uuid
			m.NodeRefs = append(m.NodeRefs, referenceModel)

		}
	}

	return m, nil
}

// ListKubernetesNode lists KubernetesNode with list spec.
func (db *DB) listKubernetesNode(ctx context.Context, request *models.ListKubernetesNodeRequest) (response *models.ListKubernetesNodeResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "kubernetes_node"
	qb.Fields = KubernetesNodeFields
	qb.RefFields = KubernetesNodeRefFields
	qb.BackRefFields = KubernetesNodeBackRefFields
	result := []*models.KubernetesNode{}

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
		m, err := scanKubernetesNode(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListKubernetesNodeResponse{
		KubernetesNodes: result,
	}
	return response, nil
}

// UpdateKubernetesNode updates a resource
func (db *DB) updateKubernetesNode(
	ctx context.Context,
	request *models.UpdateKubernetesNodeRequest,
) error {
	//TODO
	return nil
}

// DeleteKubernetesNode deletes a resource
func (db *DB) deleteKubernetesNode(
	ctx context.Context,
	request *models.DeleteKubernetesNodeRequest) error {
	deleteQuery := deleteKubernetesNodeQuery
	selectQuery := "select count(uuid) from kubernetes_node where uuid = ?"
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

//CreateKubernetesNode handle a Create API
func (db *DB) CreateKubernetesNode(
	ctx context.Context,
	request *models.CreateKubernetesNodeRequest) (*models.CreateKubernetesNodeResponse, error) {
	model := request.KubernetesNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createKubernetesNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "kubernetes_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateKubernetesNodeResponse{
		KubernetesNode: request.KubernetesNode,
	}, nil
}

//UpdateKubernetesNode handles a Update request.
func (db *DB) UpdateKubernetesNode(
	ctx context.Context,
	request *models.UpdateKubernetesNodeRequest) (*models.UpdateKubernetesNodeResponse, error) {
	model := request.KubernetesNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateKubernetesNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "kubernetes_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateKubernetesNodeResponse{
		KubernetesNode: model,
	}, nil
}

//DeleteKubernetesNode delete a resource.
func (db *DB) DeleteKubernetesNode(ctx context.Context, request *models.DeleteKubernetesNodeRequest) (*models.DeleteKubernetesNodeResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteKubernetesNode(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteKubernetesNodeResponse{
		ID: request.ID,
	}, nil
}

//GetKubernetesNode a Get request.
func (db *DB) GetKubernetesNode(ctx context.Context, request *models.GetKubernetesNodeRequest) (response *models.GetKubernetesNodeResponse, err error) {
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
	listRequest := &models.ListKubernetesNodeRequest{
		Spec: spec,
	}
	var result *models.ListKubernetesNodeResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listKubernetesNode(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.KubernetesNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetKubernetesNodeResponse{
		KubernetesNode: result.KubernetesNodes[0],
	}
	return response, nil
}

//ListKubernetesNode handles a List service Request.
func (db *DB) ListKubernetesNode(
	ctx context.Context,
	request *models.ListKubernetesNodeRequest) (response *models.ListKubernetesNodeResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listKubernetesNode(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
