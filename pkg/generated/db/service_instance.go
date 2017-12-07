package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertServiceInstanceQuery = "insert into `service_instance` (`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`service_instance_bindings`,`right_ip_address`,`ha_mode`,`auto_scale`,`max_instances`,`virtual_router_id`,`left_virtual_network`,`right_virtual_network`,`availability_zone`,`left_ip_address`,`interface_list`,`auto_policy`,`management_virtual_network`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceInstanceQuery = "update `service_instance` set `enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`service_instance_bindings` = ?,`right_ip_address` = ?,`ha_mode` = ?,`auto_scale` = ?,`max_instances` = ?,`virtual_router_id` = ?,`left_virtual_network` = ?,`right_virtual_network` = ?,`availability_zone` = ?,`left_ip_address` = ?,`interface_list` = ?,`auto_policy` = ?,`management_virtual_network` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteServiceInstanceQuery = "delete from `service_instance` where uuid = ?"

// ServiceInstanceFields is db columns for ServiceInstance
var ServiceInstanceFields = []string{
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"service_instance_bindings",
	"right_ip_address",
	"ha_mode",
	"auto_scale",
	"max_instances",
	"virtual_router_id",
	"left_virtual_network",
	"right_virtual_network",
	"availability_zone",
	"left_ip_address",
	"interface_list",
	"auto_policy",
	"management_virtual_network",
	"display_name",
	"key_value_pair",
	"share",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"uuid",
	"fq_name",
}

// ServiceInstanceRefFields is db reference fields for ServiceInstance
var ServiceInstanceRefFields = map[string][]string{

	"service_template": {
	// <common.Schema Value>

	},

	"instance_ip": {
	// <common.Schema Value>

	},
}

const insertServiceInstanceServiceTemplateQuery = "insert into `ref_service_instance_service_template` (`from`, `to` ) values (?, ?);"

const insertServiceInstanceInstanceIPQuery = "insert into `ref_service_instance_instance_ip` (`from`, `to` ) values (?, ?);"

// CreateServiceInstance inserts ServiceInstance to DB
func CreateServiceInstance(tx *sql.Tx, model *models.ServiceInstance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertServiceInstanceQuery,
	}).Debug("create query")
	_, err = stmt.Exec(bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		common.MustJSON(model.ServiceInstanceBindings),
		string(model.ServiceInstanceProperties.RightIPAddress),
		string(model.ServiceInstanceProperties.HaMode),
		bool(model.ServiceInstanceProperties.ScaleOut.AutoScale),
		int(model.ServiceInstanceProperties.ScaleOut.MaxInstances),
		string(model.ServiceInstanceProperties.VirtualRouterID),
		string(model.ServiceInstanceProperties.LeftVirtualNetwork),
		string(model.ServiceInstanceProperties.RightVirtualNetwork),
		string(model.ServiceInstanceProperties.AvailabilityZone),
		string(model.ServiceInstanceProperties.LeftIPAddress),
		common.MustJSON(model.ServiceInstanceProperties.InterfaceList),
		bool(model.ServiceInstanceProperties.AutoPolicy),
		string(model.ServiceInstanceProperties.ManagementVirtualNetwork),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		common.MustJSON(model.FQName))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtServiceTemplateRef, err := tx.Prepare(insertServiceInstanceServiceTemplateQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceTemplateRefs create statement failed")
	}
	defer stmtServiceTemplateRef.Close()
	for _, ref := range model.ServiceTemplateRefs {

		_, err = stmtServiceTemplateRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceTemplateRefs create failed")
		}
	}

	stmtInstanceIPRef, err := tx.Prepare(insertServiceInstanceInstanceIPQuery)
	if err != nil {
		return errors.Wrap(err, "preparing InstanceIPRefs create statement failed")
	}
	defer stmtInstanceIPRef.Close()
	for _, ref := range model.InstanceIPRefs {

		if ref.Attr == nil {
			ref.Attr = models.MakeServiceInterfaceTag()
		}

		_, err = stmtInstanceIPRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "InstanceIPRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanServiceInstance(values map[string]interface{}) (*models.ServiceInstance, error) {
	m := models.MakeServiceInstance()

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

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

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

	if value, ok := values["service_instance_bindings"]; ok {

		json.Unmarshal(value.([]byte), &m.ServiceInstanceBindings)

	}

	if value, ok := values["right_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceInstanceProperties.RightIPAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["ha_mode"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceInstanceProperties.HaMode = models.AddressMode(castedValue)

	}

	if value, ok := values["auto_scale"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ServiceInstanceProperties.ScaleOut.AutoScale = castedValue

	}

	if value, ok := values["max_instances"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ServiceInstanceProperties.ScaleOut.MaxInstances = castedValue

	}

	if value, ok := values["virtual_router_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceInstanceProperties.VirtualRouterID = castedValue

	}

	if value, ok := values["left_virtual_network"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceInstanceProperties.LeftVirtualNetwork = castedValue

	}

	if value, ok := values["right_virtual_network"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceInstanceProperties.RightVirtualNetwork = castedValue

	}

	if value, ok := values["availability_zone"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceInstanceProperties.AvailabilityZone = castedValue

	}

	if value, ok := values["left_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceInstanceProperties.LeftIPAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["interface_list"]; ok {

		json.Unmarshal(value.([]byte), &m.ServiceInstanceProperties.InterfaceList)

	}

	if value, ok := values["auto_policy"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ServiceInstanceProperties.AutoPolicy = castedValue

	}

	if value, ok := values["management_virtual_network"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceInstanceProperties.ManagementVirtualNetwork = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

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

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["ref_service_template"]; ok {
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
			referenceModel := &models.ServiceInstanceServiceTemplateRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ServiceTemplateRefs = append(m.ServiceTemplateRefs, referenceModel)

		}
	}

	if value, ok := values["ref_instance_ip"]; ok {
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
			referenceModel := &models.ServiceInstanceInstanceIPRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.InstanceIPRefs = append(m.InstanceIPRefs, referenceModel)

			attr := models.MakeServiceInterfaceTag()
			referenceModel.Attr = attr

		}
	}

	return m, nil
}

// ListServiceInstance lists ServiceInstance with list spec.
func ListServiceInstance(tx *sql.Tx, spec *common.ListSpec) ([]*models.ServiceInstance, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "service_instance"
	spec.Fields = ServiceInstanceFields
	spec.RefFields = ServiceInstanceRefFields
	result := models.MakeServiceInstanceSlice()
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
		m, err := scanServiceInstance(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowServiceInstance shows ServiceInstance resource
func ShowServiceInstance(tx *sql.Tx, uuid string) (*models.ServiceInstance, error) {
	list, err := ListServiceInstance(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateServiceInstance updates a resource
func UpdateServiceInstance(tx *sql.Tx, uuid string, model *models.ServiceInstance) error {
	//TODO(nati) support update
	return nil
}

// DeleteServiceInstance deletes a resource
func DeleteServiceInstance(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceInstanceQuery)
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
