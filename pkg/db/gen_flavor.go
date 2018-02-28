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

// FlavorFields is db columns for Flavor
var FlavorFields = []string{
	"vcpus",
	"uuid",
	"swap",
	"rxtx_factor",
	"ram",
	"property",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"name",
	"type",
	"rel",
	"href",
	"is_public",
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
	"id",
	"fq_name",
	"ephemeral",
	"display_name",
	"disk",
	"key_value_pair",
}

// FlavorRefFields is db reference fields for Flavor
var FlavorRefFields = map[string][]string{}

// FlavorBackRefFields is db back reference fields for Flavor
var FlavorBackRefFields = map[string][]string{}

// FlavorParentTypes is possible parents for Flavor
var FlavorParents = []string{}

// CreateFlavor inserts Flavor to DB
// nolint
func (db *DB) createFlavor(
	ctx context.Context,
	request *models.CreateFlavorRequest) error {
	qb := db.queryBuilders["flavor"]
	tx := GetTransaction(ctx)
	model := request.Flavor
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), int(model.GetVcpus()),
		string(model.GetUUID()),
		int(model.GetSwap()),
		int(model.GetRXTXFactor()),
		int(model.GetRAM()),
		string(model.GetProperty()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		string(model.GetName()),
		string(model.GetLinks().GetType()),
		string(model.GetLinks().GetRel()),
		string(model.GetLinks().GetHref()),
		bool(model.GetIsPublic()),
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
		string(model.GetID()),
		common.MustJSON(model.GetFQName()),
		int(model.GetEphemeral()),
		string(model.GetDisplayName()),
		int(model.GetDisk()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "flavor",
		FQName: model.FQName,
	}
	err = CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = CreateSharing(tx, "flavor", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanFlavor(values map[string]interface{}) (*models.Flavor, error) {
	m := models.MakeFlavor()

	if value, ok := values["vcpus"]; ok {

		m.Vcpus = common.InterfaceToInt64(value)

	}

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["swap"]; ok {

		m.Swap = common.InterfaceToInt64(value)

	}

	if value, ok := values["rxtx_factor"]; ok {

		m.RXTXFactor = common.InterfaceToInt64(value)

	}

	if value, ok := values["ram"]; ok {

		m.RAM = common.InterfaceToInt64(value)

	}

	if value, ok := values["property"]; ok {

		m.Property = common.InterfaceToString(value)

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

	if value, ok := values["name"]; ok {

		m.Name = common.InterfaceToString(value)

	}

	if value, ok := values["type"]; ok {

		m.Links.Type = common.InterfaceToString(value)

	}

	if value, ok := values["rel"]; ok {

		m.Links.Rel = common.InterfaceToString(value)

	}

	if value, ok := values["href"]; ok {

		m.Links.Href = common.InterfaceToString(value)

	}

	if value, ok := values["is_public"]; ok {

		m.IsPublic = common.InterfaceToBool(value)

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

	if value, ok := values["id"]; ok {

		m.ID = common.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["ephemeral"]; ok {

		m.Ephemeral = common.InterfaceToInt64(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["disk"]; ok {

		m.Disk = common.InterfaceToInt64(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListFlavor lists Flavor with list spec.
func (db *DB) listFlavor(ctx context.Context, request *models.ListFlavorRequest) (response *models.ListFlavorResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["flavor"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.Flavor{}

	if spec.ParentFQName != nil {
		parentMetaData, err := GetMetaData(tx, "", spec.ParentFQName)
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
		m, err := scanFlavor(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListFlavorResponse{
		Flavors: result,
	}
	return response, nil
}

// UpdateFlavor updates a resource
func (db *DB) updateFlavor(
	ctx context.Context,
	request *models.UpdateFlavorRequest,
) error {
	//TODO
	return nil
}

// DeleteFlavor deletes a resource
func (db *DB) deleteFlavor(
	ctx context.Context,
	request *models.DeleteFlavorRequest) error {
	qb := db.queryBuilders["flavor"]

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

	err = DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}

//CreateFlavor handle a Create API
// nolint
func (db *DB) CreateFlavor(
	ctx context.Context,
	request *models.CreateFlavorRequest) (*models.CreateFlavorResponse, error) {
	model := request.Flavor
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createFlavor(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "flavor",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateFlavorResponse{
		Flavor: request.Flavor,
	}, nil
}

//UpdateFlavor handles a Update request.
func (db *DB) UpdateFlavor(
	ctx context.Context,
	request *models.UpdateFlavorRequest) (*models.UpdateFlavorResponse, error) {
	model := request.Flavor
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateFlavor(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "flavor",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateFlavorResponse{
		Flavor: model,
	}, nil
}

//DeleteFlavor delete a resource.
func (db *DB) DeleteFlavor(ctx context.Context, request *models.DeleteFlavorRequest) (*models.DeleteFlavorResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteFlavor(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteFlavorResponse{
		ID: request.ID,
	}, nil
}

//GetFlavor a Get request.
func (db *DB) GetFlavor(ctx context.Context, request *models.GetFlavorRequest) (response *models.GetFlavorResponse, err error) {
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
	listRequest := &models.ListFlavorRequest{
		Spec: spec,
	}
	var result *models.ListFlavorResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listFlavor(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Flavors) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetFlavorResponse{
		Flavor: result.Flavors[0],
	}
	return response, nil
}

//ListFlavor handles a List service Request.
// nolint
func (db *DB) ListFlavor(
	ctx context.Context,
	request *models.ListFlavorRequest) (response *models.ListFlavorResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listFlavor(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
