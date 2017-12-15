package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertLoadbalancerListenerQuery = "insert into `loadbalancer_listener` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`sni_containers`,`protocol_port`,`protocol`,`default_tls_container`,`connection_limit`,`admin_state`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerListenerQuery = "update `loadbalancer_listener` set `uuid` = ?,`share` = ?,`owner_access` = ?,`owner` = ?,`global_access` = ?,`parent_uuid` = ?,`parent_type` = ?,`sni_containers` = ?,`protocol_port` = ?,`protocol` = ?,`default_tls_container` = ?,`connection_limit` = ?,`admin_state` = ?,`user_visible` = ?,`permissions_owner_access` = ?,`permissions_owner` = ?,`other_access` = ?,`group_access` = ?,`group` = ?,`last_modified` = ?,`enable` = ?,`description` = ?,`creator` = ?,`created` = ?,`fq_name` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteLoadbalancerListenerQuery = "delete from `loadbalancer_listener` where uuid = ?"

// LoadbalancerListenerFields is db columns for LoadbalancerListener
var LoadbalancerListenerFields = []string{
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"sni_containers",
	"protocol_port",
	"protocol",
	"default_tls_container",
	"connection_limit",
	"admin_state",
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

// LoadbalancerListenerRefFields is db reference fields for LoadbalancerListener
var LoadbalancerListenerRefFields = map[string][]string{

	"loadbalancer": {
	// <common.Schema Value>

	},
}

// LoadbalancerListenerBackRefFields is db back reference fields for LoadbalancerListener
var LoadbalancerListenerBackRefFields = map[string][]string{}

const insertLoadbalancerListenerLoadbalancerQuery = "insert into `ref_loadbalancer_listener_loadbalancer` (`from`, `to` ) values (?, ?);"

// CreateLoadbalancerListener inserts LoadbalancerListener to DB
func CreateLoadbalancerListener(tx *sql.Tx, model *models.LoadbalancerListener) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerListenerQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertLoadbalancerListenerQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.UUID),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		common.MustJSON(model.LoadbalancerListenerProperties.SniContainers),
		int(model.LoadbalancerListenerProperties.ProtocolPort),
		string(model.LoadbalancerListenerProperties.Protocol),
		string(model.LoadbalancerListenerProperties.DefaultTLSContainer),
		int(model.LoadbalancerListenerProperties.ConnectionLimit),
		bool(model.LoadbalancerListenerProperties.AdminState),
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

	stmtLoadbalancerRef, err := tx.Prepare(insertLoadbalancerListenerLoadbalancerQuery)
	if err != nil {
		return errors.Wrap(err, "preparing LoadbalancerRefs create statement failed")
	}
	defer stmtLoadbalancerRef.Close()
	for _, ref := range model.LoadbalancerRefs {

		_, err = stmtLoadbalancerRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "LoadbalancerRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanLoadbalancerListener(values map[string]interface{}) (*models.LoadbalancerListener, error) {
	m := models.MakeLoadbalancerListener()

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

	if value, ok := values["sni_containers"]; ok {

		json.Unmarshal(value.([]byte), &m.LoadbalancerListenerProperties.SniContainers)

	}

	if value, ok := values["protocol_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.LoadbalancerListenerProperties.ProtocolPort = castedValue

	}

	if value, ok := values["protocol"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerListenerProperties.Protocol = models.LoadbalancerProtocolType(castedValue)

	}

	if value, ok := values["default_tls_container"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerListenerProperties.DefaultTLSContainer = castedValue

	}

	if value, ok := values["connection_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.LoadbalancerListenerProperties.ConnectionLimit = castedValue

	}

	if value, ok := values["admin_state"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.LoadbalancerListenerProperties.AdminState = castedValue

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

	if value, ok := values["ref_loadbalancer"]; ok {
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
			referenceModel := &models.LoadbalancerListenerLoadbalancerRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.LoadbalancerRefs = append(m.LoadbalancerRefs, referenceModel)

		}
	}

	return m, nil
}

// ListLoadbalancerListener lists LoadbalancerListener with list spec.
func ListLoadbalancerListener(tx *sql.Tx, spec *common.ListSpec) ([]*models.LoadbalancerListener, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "loadbalancer_listener"
	spec.Fields = LoadbalancerListenerFields
	spec.RefFields = LoadbalancerListenerRefFields
	spec.BackRefFields = LoadbalancerListenerBackRefFields
	result := models.MakeLoadbalancerListenerSlice()
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
		m, err := scanLoadbalancerListener(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateLoadbalancerListener updates a resource
func UpdateLoadbalancerListener(tx *sql.Tx, uuid string, model *models.LoadbalancerListener) error {
	//TODO(nati) support update
	return nil
}

// DeleteLoadbalancerListener deletes a resource
func DeleteLoadbalancerListener(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	query := deleteLoadbalancerListenerQuery
	var err error

	if auth.IsAdmin() {
		_, err = tx.Exec(query, uuid)
	} else {
		query += " and owner = ?"
		_, err = tx.Exec(query, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}
