package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertServiceEndpointQuery = "insert into `service_endpoint` (`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceEndpointQuery = "update `service_endpoint` set `uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteServiceEndpointQuery = "delete from `service_endpoint` where uuid = ?"

// ServiceEndpointFields is db columns for ServiceEndpoint
var ServiceEndpointFields = []string{
	"uuid",
	"fq_name",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"other_access",
	"group",
	"group_access",
	"owner",
	"owner_access",
	"display_name",
	"key_value_pair",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"share",
}

// ServiceEndpointRefFields is db reference fields for ServiceEndpoint
var ServiceEndpointRefFields = map[string][]string{

	"service_object": {
	// <common.Schema Value>

	},

	"service_connection_module": {
	// <common.Schema Value>

	},

	"physical_router": {
	// <common.Schema Value>

	},
}

const insertServiceEndpointServiceConnectionModuleQuery = "insert into `ref_service_endpoint_service_connection_module` (`from`, `to` ) values (?, ?);"

const insertServiceEndpointPhysicalRouterQuery = "insert into `ref_service_endpoint_physical_router` (`from`, `to` ) values (?, ?);"

const insertServiceEndpointServiceObjectQuery = "insert into `ref_service_endpoint_service_object` (`from`, `to` ) values (?, ?);"

// CreateServiceEndpoint inserts ServiceEndpoint to DB
func CreateServiceEndpoint(tx *sql.Tx, model *models.ServiceEndpoint) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceEndpointQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertServiceEndpointQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.UUID),
		common.MustJSON(model.FQName),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtServiceConnectionModuleRef, err := tx.Prepare(insertServiceEndpointServiceConnectionModuleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceConnectionModuleRefs create statement failed")
	}
	defer stmtServiceConnectionModuleRef.Close()
	for _, ref := range model.ServiceConnectionModuleRefs {

		_, err = stmtServiceConnectionModuleRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceConnectionModuleRefs create failed")
		}
	}

	stmtPhysicalRouterRef, err := tx.Prepare(insertServiceEndpointPhysicalRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing PhysicalRouterRefs create statement failed")
	}
	defer stmtPhysicalRouterRef.Close()
	for _, ref := range model.PhysicalRouterRefs {

		_, err = stmtPhysicalRouterRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "PhysicalRouterRefs create failed")
		}
	}

	stmtServiceObjectRef, err := tx.Prepare(insertServiceEndpointServiceObjectQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceObjectRefs create statement failed")
	}
	defer stmtServiceObjectRef.Close()
	for _, ref := range model.ServiceObjectRefs {

		_, err = stmtServiceObjectRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceObjectRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanServiceEndpoint(values map[string]interface{}) (*models.ServiceEndpoint, error) {
	m := models.MakeServiceEndpoint()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

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
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.ServiceEndpointPhysicalRouterRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.PhysicalRouterRefs = append(m.PhysicalRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_object"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.ServiceEndpointServiceObjectRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ServiceObjectRefs = append(m.ServiceObjectRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_connection_module"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.ServiceEndpointServiceConnectionModuleRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ServiceConnectionModuleRefs = append(m.ServiceConnectionModuleRefs, referenceModel)

		}
	}

	return m, nil
}

// ListServiceEndpoint lists ServiceEndpoint with list spec.
func ListServiceEndpoint(tx *sql.Tx, spec *common.ListSpec) ([]*models.ServiceEndpoint, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "service_endpoint"
	spec.Fields = ServiceEndpointFields
	spec.RefFields = ServiceEndpointRefFields
	result := models.MakeServiceEndpointSlice()
	query, columns, values := common.BuildListQuery(spec)
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.Query(query, values...)
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
		log.WithFields(log.Fields{
			"valuesMap": valuesMap,
		}).Debug("valueMap")
		m, err := scanServiceEndpoint(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowServiceEndpoint shows ServiceEndpoint resource
func ShowServiceEndpoint(tx *sql.Tx, uuid string) (*models.ServiceEndpoint, error) {
	list, err := ListServiceEndpoint(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateServiceEndpoint updates a resource
func UpdateServiceEndpoint(tx *sql.Tx, uuid string, model *models.ServiceEndpoint) error {
	//TODO(nati) support update
	return nil
}

// DeleteServiceEndpoint deletes a resource
func DeleteServiceEndpoint(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceEndpointQuery)
	if err != nil {
		return errors.Wrap(err, "preparing delete query failed")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	if err != nil {
		return errors.Wrap(err, "delete failed")
	}
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}
