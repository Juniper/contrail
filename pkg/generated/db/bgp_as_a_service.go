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

const insertBGPAsAServiceQuery = "insert into `bgp_as_a_service` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`bgpaas_suppress_route_advertisement`,`bgpaas_shared`,`bgpaas_session_attributes`,`bgpaas_ipv4_mapped_ipv6_nexthop`,`bgpaas_ip_address`,`autonomous_system`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteBGPAsAServiceQuery = "delete from `bgp_as_a_service` where uuid = ?"

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

const insertBGPAsAServiceVirtualMachineInterfaceQuery = "insert into `ref_bgp_as_a_service_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertBGPAsAServiceServiceHealthCheckQuery = "insert into `ref_bgp_as_a_service_service_health_check` (`from`, `to` ) values (?, ?);"

// CreateBGPAsAService inserts BGPAsAService to DB
func CreateBGPAsAService(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateBGPAsAServiceRequest) error {
	model := request.BGPAsAService
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBGPAsAServiceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertBGPAsAServiceQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
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

	stmtServiceHealthCheckRef, err := tx.Prepare(insertBGPAsAServiceServiceHealthCheckQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceHealthCheckRefs create statement failed")
	}
	defer stmtServiceHealthCheckRef.Close()
	for _, ref := range model.ServiceHealthCheckRefs {

		_, err = stmtServiceHealthCheckRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceHealthCheckRefs create failed")
		}
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertBGPAsAServiceVirtualMachineInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs create statement failed")
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {

		_, err = stmtVirtualMachineInterfaceRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "bgp_as_a_service",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "bgp_as_a_service", model.UUID, model.GetPerms2().GetShare())
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

		m.UUID = schema.InterfaceToString(value)

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

	if value, ok := values["bgpaas_suppress_route_advertisement"]; ok {

		m.BgpaasSuppressRouteAdvertisement = schema.InterfaceToBool(value)

	}

	if value, ok := values["bgpaas_shared"]; ok {

		m.BgpaasShared = schema.InterfaceToBool(value)

	}

	if value, ok := values["bgpaas_session_attributes"]; ok {

		m.BgpaasSessionAttributes = schema.InterfaceToString(value)

	}

	if value, ok := values["bgpaas_ipv4_mapped_ipv6_nexthop"]; ok {

		m.BgpaasIpv4MappedIpv6Nexthop = schema.InterfaceToBool(value)

	}

	if value, ok := values["bgpaas_ip_address"]; ok {

		m.BgpaasIPAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["autonomous_system"]; ok {

		m.AutonomousSystem = schema.InterfaceToInt64(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
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
			referenceModel := &models.BGPAsAServiceVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_health_check"]; ok {
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
			referenceModel := &models.BGPAsAServiceServiceHealthCheckRef{}
			referenceModel.UUID = uuid
			m.ServiceHealthCheckRefs = append(m.ServiceHealthCheckRefs, referenceModel)

		}
	}

	return m, nil
}

// ListBGPAsAService lists BGPAsAService with list spec.
func ListBGPAsAService(ctx context.Context, tx *sql.Tx, request *models.ListBGPAsAServiceRequest) (response *models.ListBGPAsAServiceResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "bgp_as_a_service"
	qb.Fields = BGPAsAServiceFields
	qb.RefFields = BGPAsAServiceRefFields
	qb.BackRefFields = BGPAsAServiceBackRefFields
	result := []*models.BGPAsAService{}

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
func UpdateBGPAsAService(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateBGPAsAServiceRequest,
) error {
	//TODO
	return nil
}

// DeleteBGPAsAService deletes a resource
func DeleteBGPAsAService(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteBGPAsAServiceRequest) error {
	deleteQuery := deleteBGPAsAServiceQuery
	selectQuery := "select count(uuid) from bgp_as_a_service where uuid = ?"
	var err error
	var count int
	uuid := request.ID
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
