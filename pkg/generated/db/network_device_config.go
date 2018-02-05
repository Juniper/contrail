package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertNetworkDeviceConfigQuery = "insert into `network_device_config` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteNetworkDeviceConfigQuery = "delete from `network_device_config` where uuid = ?"

// NetworkDeviceConfigFields is db columns for NetworkDeviceConfig
var NetworkDeviceConfigFields = []string{
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

// NetworkDeviceConfigRefFields is db reference fields for NetworkDeviceConfig
var NetworkDeviceConfigRefFields = map[string][]string{

	"physical_router": {
	// <common.Schema Value>

	},
}

// NetworkDeviceConfigBackRefFields is db back reference fields for NetworkDeviceConfig
var NetworkDeviceConfigBackRefFields = map[string][]string{}

// NetworkDeviceConfigParentTypes is possible parents for NetworkDeviceConfig
var NetworkDeviceConfigParents = []string{}

const insertNetworkDeviceConfigPhysicalRouterQuery = "insert into `ref_network_device_config_physical_router` (`from`, `to` ) values (?, ?);"

// CreateNetworkDeviceConfig inserts NetworkDeviceConfig to DB
func CreateNetworkDeviceConfig(tx *sql.Tx, model *models.NetworkDeviceConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNetworkDeviceConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertNetworkDeviceConfigQuery,
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

	stmtPhysicalRouterRef, err := tx.Prepare(insertNetworkDeviceConfigPhysicalRouterQuery)
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

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "network_device_config",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "network_device_config", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanNetworkDeviceConfig(values map[string]interface{}) (*models.NetworkDeviceConfig, error) {
	m := models.MakeNetworkDeviceConfig()

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
			referenceModel := &models.NetworkDeviceConfigPhysicalRouterRef{}
			referenceModel.UUID = uuid
			m.PhysicalRouterRefs = append(m.PhysicalRouterRefs, referenceModel)

		}
	}

	return m, nil
}

// ListNetworkDeviceConfig lists NetworkDeviceConfig with list spec.
func ListNetworkDeviceConfig(tx *sql.Tx, spec *common.ListSpec) ([]*models.NetworkDeviceConfig, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "network_device_config"
	spec.Fields = NetworkDeviceConfigFields
	spec.RefFields = NetworkDeviceConfigRefFields
	spec.BackRefFields = NetworkDeviceConfigBackRefFields
	result := models.MakeNetworkDeviceConfigSlice()

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
		m, err := scanNetworkDeviceConfig(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateNetworkDeviceConfig updates a resource
func UpdateNetworkDeviceConfig(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	// Prepare statement for updating data
	var updateNetworkDeviceConfigQuery = "update `network_device_config` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateNetworkDeviceConfigQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateNetworkDeviceConfigQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateNetworkDeviceConfigQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateNetworkDeviceConfigQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateNetworkDeviceConfigQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateNetworkDeviceConfigQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateNetworkDeviceConfigQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateNetworkDeviceConfigQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateNetworkDeviceConfigQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateNetworkDeviceConfigQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateNetworkDeviceConfigQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateNetworkDeviceConfigQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateNetworkDeviceConfigQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateNetworkDeviceConfigQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateNetworkDeviceConfigQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateNetworkDeviceConfigQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateNetworkDeviceConfigQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateNetworkDeviceConfigQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateNetworkDeviceConfigQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateNetworkDeviceConfigQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNetworkDeviceConfigQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateNetworkDeviceConfigQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateNetworkDeviceConfigQuery += ","
	}

	updateNetworkDeviceConfigQuery =
		updateNetworkDeviceConfigQuery[:len(updateNetworkDeviceConfigQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateNetworkDeviceConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateNetworkDeviceConfigQuery,
	}).Debug("update query")
	_, err = stmt.Exec(updatedValues...)
	if err != nil {
		return errors.Wrap(err, "update failed")
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
				refQuery = "insert into `ref_network_device_config_physical_router` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_network_device_config_physical_router` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref PhysicalRouterRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_network_device_config_physical_router` where `from` = ? AND `to`= ?;"
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

	share, ok := common.GetValueByPath(model, ".Perms2.Share", ".")
	if ok {
		err = common.UpdateSharing(tx, "network_device_config", string(uuid), share.([]interface{}))
		if err != nil {
			return err
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("updated")
	return err
}

// DeleteNetworkDeviceConfig deletes a resource
func DeleteNetworkDeviceConfig(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteNetworkDeviceConfigQuery
	selectQuery := "select count(uuid) from network_device_config where uuid = ?"
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
