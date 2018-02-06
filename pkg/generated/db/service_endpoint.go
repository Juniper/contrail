package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertServiceEndpointQuery = "insert into `service_endpoint` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteServiceEndpointQuery = "delete from `service_endpoint` where uuid = ?"

// ServiceEndpointFields is db columns for ServiceEndpoint
var ServiceEndpointFields = []string{
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
	"key_value_pair",
}

// ServiceEndpointRefFields is db reference fields for ServiceEndpoint
var ServiceEndpointRefFields = map[string][]string{

	"service_connection_module": {
	// <common.Schema Value>

	},

	"physical_router": {
	// <common.Schema Value>

	},

	"service_object": {
	// <common.Schema Value>

	},
}

// ServiceEndpointBackRefFields is db back reference fields for ServiceEndpoint
var ServiceEndpointBackRefFields = map[string][]string{}

// ServiceEndpointParentTypes is possible parents for ServiceEndpoint
var ServiceEndpointParents = []string{}

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
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		bool(model.IDPerms.UserVisible),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OtherAccess),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Group),
		string(model.IDPerms.LastModified),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Creator),
		string(model.IDPerms.Created),
		common.MustJSON(model.FQName),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
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

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "service_endpoint",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "service_endpoint", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanServiceEndpoint(values map[string]interface{}) (*models.ServiceEndpoint, error) {
	m := models.MakeServiceEndpoint()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["parent_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentUUID = castedValue

	}

	if value, ok := values["parent_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentType = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

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
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.ServiceEndpointServiceConnectionModuleRef{}
			referenceModel.UUID = uuid
			m.ServiceConnectionModuleRefs = append(m.ServiceConnectionModuleRefs, referenceModel)

		}
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
			referenceModel := &models.ServiceEndpointPhysicalRouterRef{}
			referenceModel.UUID = uuid
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
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.ServiceEndpointServiceObjectRef{}
			referenceModel.UUID = uuid
			m.ServiceObjectRefs = append(m.ServiceObjectRefs, referenceModel)

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
	spec.BackRefFields = ServiceEndpointBackRefFields
	result := models.MakeServiceEndpointSlice()

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filter.AppendValues("parent_uuid", []string{parentMetaData.UUID})
	}

	query := spec.BuildQuery()
	columns := spec.Columns
	values := spec.Values
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
		m, err := scanServiceEndpoint(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateServiceEndpoint updates a resource
func UpdateServiceEndpoint(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	// Prepare statement for updating data
	var updateServiceEndpointQuery = "update `service_endpoint` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateServiceEndpointQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateServiceEndpointQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateServiceEndpointQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateServiceEndpointQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateServiceEndpointQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateServiceEndpointQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateServiceEndpointQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateServiceEndpointQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateServiceEndpointQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateServiceEndpointQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateServiceEndpointQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateServiceEndpointQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateServiceEndpointQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateServiceEndpointQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateServiceEndpointQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateServiceEndpointQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateServiceEndpointQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateServiceEndpointQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateServiceEndpointQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateServiceEndpointQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceEndpointQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateServiceEndpointQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateServiceEndpointQuery += ","
	}

	updateServiceEndpointQuery =
		updateServiceEndpointQuery[:len(updateServiceEndpointQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateServiceEndpointQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateServiceEndpointQuery,
	}).Debug("update query")
	_, err = stmt.Exec(updatedValues...)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}

	if value, ok := common.GetValueByPath(model, "ServiceConnectionModuleRefs", "."); ok {
		for _, ref := range value.([]interface{}) {
			refQuery := ""
			refValues := make([]interface{}, 0)
			refKeys := make([]string, 0)
			refUUID, ok := common.GetValueByPath(ref.(map[string]interface{}), "UUID", ".")
			if !ok {
				return errors.Wrap(err, "UUID is missing for referred resource. Failed to update Refs")
			}

			refValues = append(refValues, uuid)
			refValues = append(refValues, refUUID)
			operation, ok := common.GetValueByPath(ref.(map[string]interface{}), common.OPERATION, ".")
			switch operation {
			case common.ADD:
				refQuery = "insert into `ref_service_endpoint_service_connection_module` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_service_endpoint_service_connection_module` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref ServiceConnectionModuleRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_service_endpoint_service_connection_module` where `from` = ? AND `to`= ?;"
				refValues = refValues[len(refValues)-2:]
			default:
				return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
			}
			stmt, err := tx.Prepare(refQuery)
			if err != nil {
				return errors.Wrap(err, "preparing ServiceConnectionModuleRefs update statement failed")
			}
			_, err = stmt.Exec(refValues...)
			if err != nil {
				return errors.Wrap(err, "ServiceConnectionModuleRefs update failed")
			}
		}
	}

	if value, ok := common.GetValueByPath(model, "PhysicalRouterRefs", "."); ok {
		for _, ref := range value.([]interface{}) {
			refQuery := ""
			refValues := make([]interface{}, 0)
			refKeys := make([]string, 0)
			refUUID, ok := common.GetValueByPath(ref.(map[string]interface{}), "UUID", ".")
			if !ok {
				return errors.Wrap(err, "UUID is missing for referred resource. Failed to update Refs")
			}

			refValues = append(refValues, uuid)
			refValues = append(refValues, refUUID)
			operation, ok := common.GetValueByPath(ref.(map[string]interface{}), common.OPERATION, ".")
			switch operation {
			case common.ADD:
				refQuery = "insert into `ref_service_endpoint_physical_router` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_service_endpoint_physical_router` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref PhysicalRouterRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_service_endpoint_physical_router` where `from` = ? AND `to`= ?;"
				refValues = refValues[len(refValues)-2:]
			default:
				return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
			}
			stmt, err := tx.Prepare(refQuery)
			if err != nil {
				return errors.Wrap(err, "preparing PhysicalRouterRefs update statement failed")
			}
			_, err = stmt.Exec(refValues...)
			if err != nil {
				return errors.Wrap(err, "PhysicalRouterRefs update failed")
			}
		}
	}

	if value, ok := common.GetValueByPath(model, "ServiceObjectRefs", "."); ok {
		for _, ref := range value.([]interface{}) {
			refQuery := ""
			refValues := make([]interface{}, 0)
			refKeys := make([]string, 0)
			refUUID, ok := common.GetValueByPath(ref.(map[string]interface{}), "UUID", ".")
			if !ok {
				return errors.Wrap(err, "UUID is missing for referred resource. Failed to update Refs")
			}

			refValues = append(refValues, uuid)
			refValues = append(refValues, refUUID)
			operation, ok := common.GetValueByPath(ref.(map[string]interface{}), common.OPERATION, ".")
			switch operation {
			case common.ADD:
				refQuery = "insert into `ref_service_endpoint_service_object` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_service_endpoint_service_object` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref ServiceObjectRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_service_endpoint_service_object` where `from` = ? AND `to`= ?;"
				refValues = refValues[len(refValues)-2:]
			default:
				return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
			}
			stmt, err := tx.Prepare(refQuery)
			if err != nil {
				return errors.Wrap(err, "preparing ServiceObjectRefs update statement failed")
			}
			_, err = stmt.Exec(refValues...)
			if err != nil {
				return errors.Wrap(err, "ServiceObjectRefs update failed")
			}
		}
	}

	share, ok := common.GetValueByPath(model, ".Perms2.Share", ".")
	if ok {
		err = common.UpdateSharing(tx, "service_endpoint", string(uuid), share.([]interface{}))
		if err != nil {
			return err
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("updated")
	return err
}

// DeleteServiceEndpoint deletes a resource
func DeleteServiceEndpoint(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteServiceEndpointQuery
	selectQuery := "select count(uuid) from service_endpoint where uuid = ?"
	var err error
	var count int

	if auth.IsAdmin() {
		row := tx.QueryRow(selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.Exec(deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRow(selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.Exec(deleteQuery, uuid, auth.ProjectID())
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
