package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertLoadbalancerPoolQuery = "insert into `loadbalancer_pool` (`key_value_pair`,`loadbalancer_pool_provider`,`display_name`,`annotations_key_value_pair`,`uuid`,`admin_state`,`persistence_cookie_name`,`status_description`,`loadbalancer_method`,`status`,`protocol`,`subnet_id`,`session_persistence`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerPoolQuery = "update `loadbalancer_pool` set `key_value_pair` = ?,`loadbalancer_pool_provider` = ?,`display_name` = ?,`annotations_key_value_pair` = ?,`uuid` = ?,`admin_state` = ?,`persistence_cookie_name` = ?,`status_description` = ?,`loadbalancer_method` = ?,`status` = ?,`protocol` = ?,`subnet_id` = ?,`session_persistence` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?;"
const deleteLoadbalancerPoolQuery = "delete from `loadbalancer_pool` where uuid = ?"

// LoadbalancerPoolFields is db columns for LoadbalancerPool
var LoadbalancerPoolFields = []string{
	"key_value_pair",
	"loadbalancer_pool_provider",
	"display_name",
	"annotations_key_value_pair",
	"uuid",
	"admin_state",
	"persistence_cookie_name",
	"status_description",
	"loadbalancer_method",
	"status",
	"protocol",
	"subnet_id",
	"session_persistence",
	"fq_name",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"other_access",
	"group",
	"group_access",
	"owner",
	"owner_access",
	"enable",
	"description",
	"perms2_owner_access",
	"global_access",
	"share",
	"perms2_owner",
}

// LoadbalancerPoolRefFields is db reference fields for LoadbalancerPool
var LoadbalancerPoolRefFields = map[string][]string{

	"service_appliance_set": {
	// <common.Schema Value>

	},

	"virtual_machine_interface": {
	// <common.Schema Value>

	},

	"loadbalancer_listener": {
	// <common.Schema Value>

	},

	"service_instance": {
	// <common.Schema Value>

	},

	"loadbalancer_healthmonitor": {
	// <common.Schema Value>

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
	_, err = stmt.Exec(common.MustJSON(model.LoadbalancerPoolCustomAttributes.KeyValuePair),
		string(model.LoadbalancerPoolProvider),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		bool(model.LoadbalancerPoolProperties.AdminState),
		string(model.LoadbalancerPoolProperties.PersistenceCookieName),
		string(model.LoadbalancerPoolProperties.StatusDescription),
		string(model.LoadbalancerPoolProperties.LoadbalancerMethod),
		string(model.LoadbalancerPoolProperties.Status),
		string(model.LoadbalancerPoolProperties.Protocol),
		string(model.LoadbalancerPoolProperties.SubnetID),
		string(model.LoadbalancerPoolProperties.SessionPersistence),
		common.MustJSON(model.FQName),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner))
	if err != nil {
		return errors.Wrap(err, "create failed")
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

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanLoadbalancerPool(values map[string]interface{}) (*models.LoadbalancerPool, error) {
	m := models.MakeLoadbalancerPool()

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.LoadbalancerPoolCustomAttributes.KeyValuePair)

	}

	if value, ok := values["loadbalancer_pool_provider"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerPoolProvider = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["annotations_key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["admin_state"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.LoadbalancerPoolProperties.AdminState = castedValue

	}

	if value, ok := values["persistence_cookie_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerPoolProperties.PersistenceCookieName = castedValue

	}

	if value, ok := values["status_description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerPoolProperties.StatusDescription = castedValue

	}

	if value, ok := values["loadbalancer_method"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerPoolProperties.LoadbalancerMethod = models.LoadbalancerMethodType(castedValue)

	}

	if value, ok := values["status"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerPoolProperties.Status = castedValue

	}

	if value, ok := values["protocol"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerPoolProperties.Protocol = models.LoadbalancerProtocolType(castedValue)

	}

	if value, ok := values["subnet_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerPoolProperties.SubnetID = models.UuidStringType(castedValue)

	}

	if value, ok := values["session_persistence"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerPoolProperties.SessionPersistence = models.SessionPersistenceType(castedValue)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

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

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

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

	if value, ok := values["perms2_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["ref_service_appliance_set"]; ok {
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
			referenceModel := &models.LoadbalancerPoolServiceApplianceSetRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ServiceApplianceSetRefs = append(m.ServiceApplianceSetRefs, referenceModel)

		}
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
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.LoadbalancerPoolVirtualMachineInterfaceRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_loadbalancer_listener"]; ok {
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
			referenceModel := &models.LoadbalancerPoolLoadbalancerListenerRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.LoadbalancerListenerRefs = append(m.LoadbalancerListenerRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_instance"]; ok {
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
			referenceModel := &models.LoadbalancerPoolServiceInstanceRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_loadbalancer_healthmonitor"]; ok {
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
			referenceModel := &models.LoadbalancerPoolLoadbalancerHealthmonitorRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.LoadbalancerHealthmonitorRefs = append(m.LoadbalancerHealthmonitorRefs, referenceModel)

		}
	}

	return m, nil
}

// ListLoadbalancerPool lists LoadbalancerPool with list spec.
func ListLoadbalancerPool(tx *sql.Tx, spec *common.ListSpec) ([]*models.LoadbalancerPool, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "loadbalancer_pool"
	spec.Fields = LoadbalancerPoolFields
	spec.RefFields = LoadbalancerPoolRefFields
	result := models.MakeLoadbalancerPoolSlice()
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
	list, err := ListLoadbalancerPool(tx, &common.ListSpec{
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
