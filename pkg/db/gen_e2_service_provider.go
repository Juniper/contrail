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

// E2ServiceProviderFields is db columns for E2ServiceProvider
var E2ServiceProviderFields = []string{
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
	"e2_service_provider_promiscuous",
	"display_name",
	"configuration_version",
	"key_value_pair",
}

// E2ServiceProviderRefFields is db reference fields for E2ServiceProvider
var E2ServiceProviderRefFields = map[string][]string{

	"physical_router": []string{
	// <schema.Schema Value>

	},

	"peering_policy": []string{
	// <schema.Schema Value>

	},
}

// E2ServiceProviderBackRefFields is db back reference fields for E2ServiceProvider
var E2ServiceProviderBackRefFields = map[string][]string{}

// E2ServiceProviderParentTypes is possible parents for E2ServiceProvider
var E2ServiceProviderParents = []string{}

// CreateE2ServiceProvider inserts E2ServiceProvider to DB
// nolint
func (db *DB) createE2ServiceProvider(
	ctx context.Context,
	request *models.CreateE2ServiceProviderRequest) error {
	qb := db.queryBuilders["e2_service_provider"]
	tx := GetTransaction(ctx)
	model := request.E2ServiceProvider
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
		bool(model.GetE2ServiceProviderPromiscuous()),
		string(model.GetDisplayName()),
		int(model.GetConfigurationVersion()),
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

	for _, ref := range model.PeeringPolicyRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("peering_policy"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "PeeringPolicyRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "e2_service_provider",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "e2_service_provider", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanE2ServiceProvider(values map[string]interface{}) (*models.E2ServiceProvider, error) {
	m := models.MakeE2ServiceProvider()

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

	if value, ok := values["e2_service_provider_promiscuous"]; ok {

		m.E2ServiceProviderPromiscuous = common.InterfaceToBool(value)

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
			referenceModel := &models.E2ServiceProviderPhysicalRouterRef{}
			referenceModel.UUID = uuid
			m.PhysicalRouterRefs = append(m.PhysicalRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_peering_policy"]; ok {
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
			referenceModel := &models.E2ServiceProviderPeeringPolicyRef{}
			referenceModel.UUID = uuid
			m.PeeringPolicyRefs = append(m.PeeringPolicyRefs, referenceModel)

		}
	}

	return m, nil
}

// ListE2ServiceProvider lists E2ServiceProvider with list spec.
func (db *DB) listE2ServiceProvider(ctx context.Context, request *models.ListE2ServiceProviderRequest) (response *models.ListE2ServiceProviderResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["e2_service_provider"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.E2ServiceProvider{}

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
		m, err := scanE2ServiceProvider(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListE2ServiceProviderResponse{
		E2ServiceProviders: result,
	}
	return response, nil
}

// UpdateE2ServiceProvider updates a resource
func (db *DB) updateE2ServiceProvider(
	ctx context.Context,
	request *models.UpdateE2ServiceProviderRequest,
) error {
	//TODO
	return nil
}

// DeleteE2ServiceProvider deletes a resource
func (db *DB) deleteE2ServiceProvider(
	ctx context.Context,
	request *models.DeleteE2ServiceProviderRequest) error {
	qb := db.queryBuilders["e2_service_provider"]

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

//CreateE2ServiceProvider handle a Create API
// nolint
func (db *DB) CreateE2ServiceProvider(
	ctx context.Context,
	request *models.CreateE2ServiceProviderRequest) (*models.CreateE2ServiceProviderResponse, error) {
	model := request.E2ServiceProvider
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createE2ServiceProvider(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "e2_service_provider",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateE2ServiceProviderResponse{
		E2ServiceProvider: request.E2ServiceProvider,
	}, nil
}

//UpdateE2ServiceProvider handles a Update request.
func (db *DB) UpdateE2ServiceProvider(
	ctx context.Context,
	request *models.UpdateE2ServiceProviderRequest) (*models.UpdateE2ServiceProviderResponse, error) {
	model := request.E2ServiceProvider
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateE2ServiceProvider(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "e2_service_provider",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateE2ServiceProviderResponse{
		E2ServiceProvider: model,
	}, nil
}

//DeleteE2ServiceProvider delete a resource.
func (db *DB) DeleteE2ServiceProvider(ctx context.Context, request *models.DeleteE2ServiceProviderRequest) (*models.DeleteE2ServiceProviderResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteE2ServiceProvider(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteE2ServiceProviderResponse{
		ID: request.ID,
	}, nil
}

//GetE2ServiceProvider a Get request.
func (db *DB) GetE2ServiceProvider(ctx context.Context, request *models.GetE2ServiceProviderRequest) (response *models.GetE2ServiceProviderResponse, err error) {
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
	listRequest := &models.ListE2ServiceProviderRequest{
		Spec: spec,
	}
	var result *models.ListE2ServiceProviderResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listE2ServiceProvider(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.E2ServiceProviders) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetE2ServiceProviderResponse{
		E2ServiceProvider: result.E2ServiceProviders[0],
	}
	return response, nil
}

//ListE2ServiceProvider handles a List service Request.
// nolint
func (db *DB) ListE2ServiceProvider(
	ctx context.Context,
	request *models.ListE2ServiceProviderRequest) (response *models.ListE2ServiceProviderResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listE2ServiceProvider(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
