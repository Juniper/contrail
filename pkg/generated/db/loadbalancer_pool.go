package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertLoadbalancerPoolQuery = "insert into `loadbalancer_pool` (`session_persistence`,`admin_state`,`persistence_cookie_name`,`status_description`,`loadbalancer_method`,`status`,`protocol`,`subnet_id`,`loadbalancer_pool_provider`,`display_name`,`uuid`,`fq_name`,`key_value_pair`,`annotations_key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerPoolQuery = "update `loadbalancer_pool` set `session_persistence` = ?,`admin_state` = ?,`persistence_cookie_name` = ?,`status_description` = ?,`loadbalancer_method` = ?,`status` = ?,`protocol` = ?,`subnet_id` = ?,`loadbalancer_pool_provider` = ?,`display_name` = ?,`uuid` = ?,`fq_name` = ?,`key_value_pair` = ?,`annotations_key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?;"
const deleteLoadbalancerPoolQuery = "delete from `loadbalancer_pool` where uuid = ?"

// LoadbalancerPoolFields is db columns for LoadbalancerPool
var LoadbalancerPoolFields = []string{
	"session_persistence",
	"admin_state",
	"persistence_cookie_name",
	"status_description",
	"loadbalancer_method",
	"status",
	"protocol",
	"subnet_id",
	"loadbalancer_pool_provider",
	"display_name",
	"uuid",
	"fq_name",
	"key_value_pair",
	"annotations_key_value_pair",
	"owner",
	"owner_access",
	"global_access",
	"share",
	"last_modified",
	"group_access",
	"permissions_owner",
	"permissions_owner_access",
	"other_access",
	"group",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
}

// LoadbalancerPoolRefFields is db reference fields for LoadbalancerPool
var LoadbalancerPoolRefFields = map[string][]string{

	"service_appliance_set": {
	// <utils.Schema Value>

	},

	"virtual_machine_interface": {
	// <utils.Schema Value>

	},

	"loadbalancer_listener": {
	// <utils.Schema Value>

	},

	"service_instance": {
	// <utils.Schema Value>

	},

	"loadbalancer_healthmonitor": {
	// <utils.Schema Value>

	},
}

const insertLoadbalancerPoolServiceApplianceSetQuery = "insert into `ref_loadbalancer_pool_service_appliance_set` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerPoolVirtualMachineInterfaceQuery = "insert into `ref_loadbalancer_pool_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerPoolLoadbalancerListenerQuery = "insert into `ref_loadbalancer_pool_loadbalancer_listener` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerPoolServiceInstanceQuery = "insert into `ref_loadbalancer_pool_service_instance` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerPoolLoadbalancerHealthmonitorQuery = "insert into `ref_loadbalancer_pool_loadbalancer_healthmonitor` (`from`, `to` ) values (?, ?);"

// CreateLoadbalancerPool inserts LoadbalancerPool to DB
func CreateLoadbalancerPool(tx *sql.Tx, model *models.LoadbalancerPool) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerPoolQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertLoadbalancerPoolQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.LoadbalancerPoolProperties.SessionPersistence),
		bool(model.LoadbalancerPoolProperties.AdminState),
		string(model.LoadbalancerPoolProperties.PersistenceCookieName),
		string(model.LoadbalancerPoolProperties.StatusDescription),
		string(model.LoadbalancerPoolProperties.LoadbalancerMethod),
		string(model.LoadbalancerPoolProperties.Status),
		string(model.LoadbalancerPoolProperties.Protocol),
		string(model.LoadbalancerPoolProperties.SubnetID),
		string(model.LoadbalancerPoolProvider),
		string(model.DisplayName),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		utils.MustJSON(model.LoadbalancerPoolCustomAttributes.KeyValuePair),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtServiceApplianceSetRef, err := tx.Prepare(insertLoadbalancerPoolServiceApplianceSetQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceApplianceSetRefs create statement failed")
	}
	defer stmtServiceApplianceSetRef.Close()
	for _, ref := range model.ServiceApplianceSetRefs {
		_, err = stmtServiceApplianceSetRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceApplianceSetRefs create failed")
		}
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertLoadbalancerPoolVirtualMachineInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs create statement failed")
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {
		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	stmtLoadbalancerListenerRef, err := tx.Prepare(insertLoadbalancerPoolLoadbalancerListenerQuery)
	if err != nil {
		return errors.Wrap(err, "preparing LoadbalancerListenerRefs create statement failed")
	}
	defer stmtLoadbalancerListenerRef.Close()
	for _, ref := range model.LoadbalancerListenerRefs {
		_, err = stmtLoadbalancerListenerRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "LoadbalancerListenerRefs create failed")
		}
	}

	stmtServiceInstanceRef, err := tx.Prepare(insertLoadbalancerPoolServiceInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceInstanceRefs create statement failed")
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {
		_, err = stmtServiceInstanceRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs create failed")
		}
	}

	stmtLoadbalancerHealthmonitorRef, err := tx.Prepare(insertLoadbalancerPoolLoadbalancerHealthmonitorQuery)
	if err != nil {
		return errors.Wrap(err, "preparing LoadbalancerHealthmonitorRefs create statement failed")
	}
	defer stmtLoadbalancerHealthmonitorRef.Close()
	for _, ref := range model.LoadbalancerHealthmonitorRefs {
		_, err = stmtLoadbalancerHealthmonitorRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "LoadbalancerHealthmonitorRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanLoadbalancerPool(values map[string]interface{}) (*models.LoadbalancerPool, error) {
	m := models.MakeLoadbalancerPool()

	if value, ok := values["session_persistence"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerPoolProperties.SessionPersistence = models.SessionPersistenceType(castedValue)

	}

	if value, ok := values["admin_state"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.LoadbalancerPoolProperties.AdminState = castedValue

	}

	if value, ok := values["persistence_cookie_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerPoolProperties.PersistenceCookieName = castedValue

	}

	if value, ok := values["status_description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerPoolProperties.StatusDescription = castedValue

	}

	if value, ok := values["loadbalancer_method"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerPoolProperties.LoadbalancerMethod = models.LoadbalancerMethodType(castedValue)

	}

	if value, ok := values["status"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerPoolProperties.Status = castedValue

	}

	if value, ok := values["protocol"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerPoolProperties.Protocol = models.LoadbalancerProtocolType(castedValue)

	}

	if value, ok := values["subnet_id"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerPoolProperties.SubnetID = models.UuidStringType(castedValue)

	}

	if value, ok := values["loadbalancer_pool_provider"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerPoolProvider = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.LoadbalancerPoolCustomAttributes.KeyValuePair)

	}

	if value, ok := values["annotations_key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["ref_service_appliance_set"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.LoadbalancerPoolServiceApplianceSetRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.ServiceApplianceSetRefs = append(m.ServiceApplianceSetRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.LoadbalancerPoolVirtualMachineInterfaceRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_loadbalancer_listener"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.LoadbalancerPoolLoadbalancerListenerRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.LoadbalancerListenerRefs = append(m.LoadbalancerListenerRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_instance"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.LoadbalancerPoolServiceInstanceRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_loadbalancer_healthmonitor"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.LoadbalancerPoolLoadbalancerHealthmonitorRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.LoadbalancerHealthmonitorRefs = append(m.LoadbalancerHealthmonitorRefs, referenceModel)

		}
	}

	return m, nil
}

// ListLoadbalancerPool lists LoadbalancerPool with list spec.
func ListLoadbalancerPool(tx *sql.Tx, spec *db.ListSpec) ([]*models.LoadbalancerPool, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "loadbalancer_pool"
	spec.Fields = LoadbalancerPoolFields
	spec.RefFields = LoadbalancerPoolRefFields
	result := models.MakeLoadbalancerPoolSlice()
	query, columns, values := db.BuildListQuery(spec)
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
		m, err := scanLoadbalancerPool(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowLoadbalancerPool shows LoadbalancerPool resource
func ShowLoadbalancerPool(tx *sql.Tx, uuid string) (*models.LoadbalancerPool, error) {
	list, err := ListLoadbalancerPool(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateLoadbalancerPool updates a resource
func UpdateLoadbalancerPool(tx *sql.Tx, uuid string, model *models.LoadbalancerPool) error {
	//TODO(nati) support update
	return nil
}

// DeleteLoadbalancerPool deletes a resource
func DeleteLoadbalancerPool(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLoadbalancerPoolQuery)
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
