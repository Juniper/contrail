package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertDiscoveryServiceAssignmentQuery = "insert into `discovery_service_assignment` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDiscoveryServiceAssignmentQuery = "update `discovery_service_assignment` set `uuid` = ?,`share` = ?,`owner_access` = ?,`owner` = ?,`global_access` = ?,`parent_uuid` = ?,`parent_type` = ?,`user_visible` = ?,`permissions_owner_access` = ?,`permissions_owner` = ?,`other_access` = ?,`group_access` = ?,`group` = ?,`last_modified` = ?,`enable` = ?,`description` = ?,`creator` = ?,`created` = ?,`fq_name` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteDiscoveryServiceAssignmentQuery = "delete from `discovery_service_assignment` where uuid = ?"

// DiscoveryServiceAssignmentFields is db columns for DiscoveryServiceAssignment
var DiscoveryServiceAssignmentFields = []string{
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

// DiscoveryServiceAssignmentRefFields is db reference fields for DiscoveryServiceAssignment
var DiscoveryServiceAssignmentRefFields = map[string][]string{}

// DiscoveryServiceAssignmentBackRefFields is db back reference fields for DiscoveryServiceAssignment
var DiscoveryServiceAssignmentBackRefFields = map[string][]string{

	"dsa_rule": {
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
		"subscriber",
		"ep_version",
		"ep_type",
		"ip_prefix_len",
		"ip_prefix",
		"ep_id",
		"display_name",
		"key_value_pair",
	},
}

// CreateDiscoveryServiceAssignment inserts DiscoveryServiceAssignment to DB
func CreateDiscoveryServiceAssignment(tx *sql.Tx, model *models.DiscoveryServiceAssignment) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDiscoveryServiceAssignmentQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertDiscoveryServiceAssignmentQuery,
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

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanDiscoveryServiceAssignment(values map[string]interface{}) (*models.DiscoveryServiceAssignment, error) {
	m := models.MakeDiscoveryServiceAssignment()

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

	if value, ok := values["backref_dsa_rule"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			if childResourceMap["uuid"] == "" {
				continue
			}
			childModel := models.MakeDsaRule()
			m.DsaRules = append(m.DsaRules, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.Perms2.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.GlobalAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentUUID = castedValue

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentType = castedValue

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.UserVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Group = castedValue

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.LastModified = castedValue

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.Enable = castedValue

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Description = castedValue

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Creator = castedValue

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Created = castedValue

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["subscriber"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.DsaRuleEntry.Subscriber)

			}

			if propertyValue, ok := childResourceMap["ep_version"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DsaRuleEntry.Publisher.EpVersion = castedValue

			}

			if propertyValue, ok := childResourceMap["ep_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DsaRuleEntry.Publisher.EpType = castedValue

			}

			if propertyValue, ok := childResourceMap["ip_prefix_len"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen = castedValue

			}

			if propertyValue, ok := childResourceMap["ip_prefix"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DsaRuleEntry.Publisher.EpPrefix.IPPrefix = castedValue

			}

			if propertyValue, ok := childResourceMap["ep_id"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DsaRuleEntry.Publisher.EpID = castedValue

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListDiscoveryServiceAssignment lists DiscoveryServiceAssignment with list spec.
func ListDiscoveryServiceAssignment(tx *sql.Tx, spec *common.ListSpec) ([]*models.DiscoveryServiceAssignment, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "discovery_service_assignment"
	spec.Fields = DiscoveryServiceAssignmentFields
	spec.RefFields = DiscoveryServiceAssignmentRefFields
	spec.BackRefFields = DiscoveryServiceAssignmentBackRefFields
	result := models.MakeDiscoveryServiceAssignmentSlice()
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
		m, err := scanDiscoveryServiceAssignment(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowDiscoveryServiceAssignment shows DiscoveryServiceAssignment resource
func ShowDiscoveryServiceAssignment(tx *sql.Tx, uuid string) (*models.DiscoveryServiceAssignment, error) {
	list, err := ListDiscoveryServiceAssignment(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateDiscoveryServiceAssignment updates a resource
func UpdateDiscoveryServiceAssignment(tx *sql.Tx, uuid string, model *models.DiscoveryServiceAssignment) error {
	//TODO(nati) support update
	return nil
}

// DeleteDiscoveryServiceAssignment deletes a resource
func DeleteDiscoveryServiceAssignment(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteDiscoveryServiceAssignmentQuery)
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
