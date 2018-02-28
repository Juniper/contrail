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

// BGPAsAServiceFields is db columns for BGPAsAService
var BGPAsAServiceFields = []string{
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
	"bgpaas_suppress_route_advertisement",
	"bgpaas_shared",
	"bgpaas_session_attributes",
	"bgpaas_ipv4_mapped_ipv6_nexthop",
	"bgpaas_ip_address",
	"autonomous_system",
	"key_value_pair",
}

// BGPAsAServiceRefFields is db reference fields for BGPAsAService
var BGPAsAServiceRefFields = map[string][]string{

	"virtual_machine_interface": []string{
	// <schema.Schema Value>

	},

	"service_health_check": []string{
	// <schema.Schema Value>

	},
}

// BGPAsAServiceBackRefFields is db back reference fields for BGPAsAService
var BGPAsAServiceBackRefFields = map[string][]string{}

// BGPAsAServiceParentTypes is possible parents for BGPAsAService
var BGPAsAServiceParents = []string{

	"project",
}

// CreateBGPAsAService inserts BGPAsAService to DB
// nolint
func (db *DB) createBGPAsAService(
	ctx context.Context,
	request *models.CreateBGPAsAServiceRequest) error {
	qb := db.queryBuilders["bgp_as_a_service"]
	tx := GetTransaction(ctx)
	model := request.BGPAsAService
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
		string(model.GetDisplayName()),
		bool(model.GetBgpaasSuppressRouteAdvertisement()),
		bool(model.GetBgpaasShared()),
		string(model.GetBgpaasSessionAttributes()),
		bool(model.GetBgpaasIpv4MappedIpv6Nexthop()),
		string(model.GetBgpaasIPAddress()),
		int(model.GetAutonomousSystem()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	for _, ref := range model.VirtualMachineInterfaceRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("virtual_machine_interface"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	for _, ref := range model.ServiceHealthCheckRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("service_health_check"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceHealthCheckRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "bgp_as_a_service",
		FQName: model.FQName,
	}
	err = CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = CreateSharing(tx, "bgp_as_a_service", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanBGPAsAService(values map[string]interface{}) (*models.BGPAsAService, error) {
	m := models.MakeBGPAsAService()

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

	if value, ok := values["bgpaas_suppress_route_advertisement"]; ok {

		m.BgpaasSuppressRouteAdvertisement = common.InterfaceToBool(value)

	}

	if value, ok := values["bgpaas_shared"]; ok {

		m.BgpaasShared = common.InterfaceToBool(value)

	}

	if value, ok := values["bgpaas_session_attributes"]; ok {

		m.BgpaasSessionAttributes = common.InterfaceToString(value)

	}

	if value, ok := values["bgpaas_ipv4_mapped_ipv6_nexthop"]; ok {

		m.BgpaasIpv4MappedIpv6Nexthop = common.InterfaceToBool(value)

	}

	if value, ok := values["bgpaas_ip_address"]; ok {

		m.BgpaasIPAddress = common.InterfaceToString(value)

	}

	if value, ok := values["autonomous_system"]; ok {

		m.AutonomousSystem = common.InterfaceToInt64(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

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
			referenceModel := &models.BGPAsAServiceVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_health_check"]; ok {
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
			referenceModel := &models.BGPAsAServiceServiceHealthCheckRef{}
			referenceModel.UUID = uuid
			m.ServiceHealthCheckRefs = append(m.ServiceHealthCheckRefs, referenceModel)

		}
	}

	return m, nil
}

// ListBGPAsAService lists BGPAsAService with list spec.
func (db *DB) listBGPAsAService(ctx context.Context, request *models.ListBGPAsAServiceRequest) (response *models.ListBGPAsAServiceResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["bgp_as_a_service"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.BGPAsAService{}

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
		m, err := scanBGPAsAService(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListBGPAsAServiceResponse{
		BGPAsAServices: result,
	}
	return response, nil
}

// UpdateBGPAsAService updates a resource
func (db *DB) updateBGPAsAService(
	ctx context.Context,
	request *models.UpdateBGPAsAServiceRequest,
) error {
	//TODO
	return nil
}

// DeleteBGPAsAService deletes a resource
func (db *DB) deleteBGPAsAService(
	ctx context.Context,
	request *models.DeleteBGPAsAServiceRequest) error {
	qb := db.queryBuilders["bgp_as_a_service"]

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

//CreateBGPAsAService handle a Create API
// nolint
func (db *DB) CreateBGPAsAService(
	ctx context.Context,
	request *models.CreateBGPAsAServiceRequest) (*models.CreateBGPAsAServiceResponse, error) {
	model := request.BGPAsAService
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createBGPAsAService(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_as_a_service",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateBGPAsAServiceResponse{
		BGPAsAService: request.BGPAsAService,
	}, nil
}

//UpdateBGPAsAService handles a Update request.
func (db *DB) UpdateBGPAsAService(
	ctx context.Context,
	request *models.UpdateBGPAsAServiceRequest) (*models.UpdateBGPAsAServiceResponse, error) {
	model := request.BGPAsAService
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateBGPAsAService(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_as_a_service",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateBGPAsAServiceResponse{
		BGPAsAService: model,
	}, nil
}

//DeleteBGPAsAService delete a resource.
func (db *DB) DeleteBGPAsAService(ctx context.Context, request *models.DeleteBGPAsAServiceRequest) (*models.DeleteBGPAsAServiceResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteBGPAsAService(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteBGPAsAServiceResponse{
		ID: request.ID,
	}, nil
}

//GetBGPAsAService a Get request.
func (db *DB) GetBGPAsAService(ctx context.Context, request *models.GetBGPAsAServiceRequest) (response *models.GetBGPAsAServiceResponse, err error) {
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
	listRequest := &models.ListBGPAsAServiceRequest{
		Spec: spec,
	}
	var result *models.ListBGPAsAServiceResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listBGPAsAService(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.BGPAsAServices) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetBGPAsAServiceResponse{
		BGPAsAService: result.BGPAsAServices[0],
	}
	return response, nil
}

//ListBGPAsAService handles a List service Request.
// nolint
func (db *DB) ListBGPAsAService(
	ctx context.Context,
	request *models.ListBGPAsAServiceRequest) (response *models.ListBGPAsAServiceResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listBGPAsAService(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
