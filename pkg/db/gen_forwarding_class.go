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

// ForwardingClassFields is db columns for ForwardingClass
var ForwardingClassFields = []string{
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
	"forwarding_class_vlan_priority",
	"forwarding_class_mpls_exp",
	"forwarding_class_id",
	"forwarding_class_dscp",
	"display_name",
	"key_value_pair",
}

// ForwardingClassRefFields is db reference fields for ForwardingClass
var ForwardingClassRefFields = map[string][]string{

	"qos_queue": []string{
	// <schema.Schema Value>

	},
}

// ForwardingClassBackRefFields is db back reference fields for ForwardingClass
var ForwardingClassBackRefFields = map[string][]string{}

// ForwardingClassParentTypes is possible parents for ForwardingClass
var ForwardingClassParents = []string{

	"global_qos_config",
}

// CreateForwardingClass inserts ForwardingClass to DB
// nolint
func (db *DB) createForwardingClass(
	ctx context.Context,
	request *models.CreateForwardingClassRequest) error {
	qb := db.queryBuilders["forwarding_class"]
	tx := GetTransaction(ctx)
	model := request.ForwardingClass
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
		int(model.GetForwardingClassVlanPriority()),
		int(model.GetForwardingClassMPLSExp()),
		int(model.GetForwardingClassID()),
		int(model.GetForwardingClassDSCP()),
		string(model.GetDisplayName()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	for _, ref := range model.QosQueueRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("qos_queue"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "QosQueueRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "forwarding_class",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "forwarding_class", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanForwardingClass(values map[string]interface{}) (*models.ForwardingClass, error) {
	m := models.MakeForwardingClass()

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

	if value, ok := values["forwarding_class_vlan_priority"]; ok {

		m.ForwardingClassVlanPriority = common.InterfaceToInt64(value)

	}

	if value, ok := values["forwarding_class_mpls_exp"]; ok {

		m.ForwardingClassMPLSExp = common.InterfaceToInt64(value)

	}

	if value, ok := values["forwarding_class_id"]; ok {

		m.ForwardingClassID = common.InterfaceToInt64(value)

	}

	if value, ok := values["forwarding_class_dscp"]; ok {

		m.ForwardingClassDSCP = common.InterfaceToInt64(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_qos_queue"]; ok {
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
			referenceModel := &models.ForwardingClassQosQueueRef{}
			referenceModel.UUID = uuid
			m.QosQueueRefs = append(m.QosQueueRefs, referenceModel)

		}
	}

	return m, nil
}

// ListForwardingClass lists ForwardingClass with list spec.
func (db *DB) listForwardingClass(ctx context.Context, request *models.ListForwardingClassRequest) (response *models.ListForwardingClassResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["forwarding_class"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.ForwardingClass{}

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
		m, err := scanForwardingClass(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListForwardingClassResponse{
		ForwardingClasss: result,
	}
	return response, nil
}

// UpdateForwardingClass updates a resource
func (db *DB) updateForwardingClass(
	ctx context.Context,
	request *models.UpdateForwardingClassRequest,
) error {
	//TODO
	return nil
}

// DeleteForwardingClass deletes a resource
func (db *DB) deleteForwardingClass(
	ctx context.Context,
	request *models.DeleteForwardingClassRequest) error {
	qb := db.queryBuilders["forwarding_class"]

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

//CreateForwardingClass handle a Create API
// nolint
func (db *DB) CreateForwardingClass(
	ctx context.Context,
	request *models.CreateForwardingClassRequest) (*models.CreateForwardingClassResponse, error) {
	model := request.ForwardingClass
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createForwardingClass(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "forwarding_class",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateForwardingClassResponse{
		ForwardingClass: request.ForwardingClass,
	}, nil
}

//UpdateForwardingClass handles a Update request.
func (db *DB) UpdateForwardingClass(
	ctx context.Context,
	request *models.UpdateForwardingClassRequest) (*models.UpdateForwardingClassResponse, error) {
	model := request.ForwardingClass
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateForwardingClass(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "forwarding_class",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateForwardingClassResponse{
		ForwardingClass: model,
	}, nil
}

//DeleteForwardingClass delete a resource.
func (db *DB) DeleteForwardingClass(ctx context.Context, request *models.DeleteForwardingClassRequest) (*models.DeleteForwardingClassResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteForwardingClass(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteForwardingClassResponse{
		ID: request.ID,
	}, nil
}

//GetForwardingClass a Get request.
func (db *DB) GetForwardingClass(ctx context.Context, request *models.GetForwardingClassRequest) (response *models.GetForwardingClassResponse, err error) {
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
	listRequest := &models.ListForwardingClassRequest{
		Spec: spec,
	}
	var result *models.ListForwardingClassResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listForwardingClass(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ForwardingClasss) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetForwardingClassResponse{
		ForwardingClass: result.ForwardingClasss[0],
	}
	return response, nil
}

//ListForwardingClass handles a List service Request.
// nolint
func (db *DB) ListForwardingClass(
	ctx context.Context,
	request *models.ListForwardingClassRequest) (response *models.ListForwardingClassResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listForwardingClass(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
