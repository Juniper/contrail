package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertLoadbalancerQuery = "insert into `loadbalancer` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`loadbalancer_provider`,`vip_subnet_id`,`vip_address`,`status`,`provisioning_status`,`operating_status`,`admin_state`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteLoadbalancerQuery = "delete from `loadbalancer` where uuid = ?"

// LoadbalancerFields is db columns for Loadbalancer
var LoadbalancerFields = []string{
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"loadbalancer_provider",
	"vip_subnet_id",
	"vip_address",
	"status",
	"provisioning_status",
	"operating_status",
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

// LoadbalancerRefFields is db reference fields for Loadbalancer
var LoadbalancerRefFields = map[string][]string{

	"service_appliance_set": {
	// <common.Schema Value>

	},

	"virtual_machine_interface": {
	// <common.Schema Value>

	},

	"service_instance": {
	// <common.Schema Value>

	},
}

// LoadbalancerBackRefFields is db back reference fields for Loadbalancer
var LoadbalancerBackRefFields = map[string][]string{}

// LoadbalancerParentTypes is possible parents for Loadbalancer
var LoadbalancerParents = []string{

	"project",
}

const insertLoadbalancerServiceApplianceSetQuery = "insert into `ref_loadbalancer_service_appliance_set` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerVirtualMachineInterfaceQuery = "insert into `ref_loadbalancer_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerServiceInstanceQuery = "insert into `ref_loadbalancer_service_instance` (`from`, `to` ) values (?, ?);"

// CreateLoadbalancer inserts Loadbalancer to DB
func CreateLoadbalancer(tx *sql.Tx, model *models.Loadbalancer) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertLoadbalancerQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.UUID),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		string(model.LoadbalancerProvider),
		string(model.LoadbalancerProperties.VipSubnetID),
		string(model.LoadbalancerProperties.VipAddress),
		string(model.LoadbalancerProperties.Status),
		string(model.LoadbalancerProperties.ProvisioningStatus),
		string(model.LoadbalancerProperties.OperatingStatus),
		bool(model.LoadbalancerProperties.AdminState),
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

	stmtServiceInstanceRef, err := tx.Prepare(insertLoadbalancerServiceInstanceQuery)
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

	stmtServiceApplianceSetRef, err := tx.Prepare(insertLoadbalancerServiceApplianceSetQuery)
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

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertLoadbalancerVirtualMachineInterfaceQuery)
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

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "loadbalancer",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "loadbalancer", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanLoadbalancer(values map[string]interface{}) (*models.Loadbalancer, error) {
	m := models.MakeLoadbalancer()

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

	if value, ok := values["loadbalancer_provider"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerProvider = castedValue

	}

	if value, ok := values["vip_subnet_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerProperties.VipSubnetID = models.UuidStringType(castedValue)

	}

	if value, ok := values["vip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerProperties.VipAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["status"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerProperties.Status = castedValue

	}

	if value, ok := values["provisioning_status"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerProperties.ProvisioningStatus = castedValue

	}

	if value, ok := values["operating_status"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LoadbalancerProperties.OperatingStatus = castedValue

	}

	if value, ok := values["admin_state"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.LoadbalancerProperties.AdminState = castedValue

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
			referenceModel := &models.LoadbalancerVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

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
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.LoadbalancerServiceInstanceRef{}
			referenceModel.UUID = uuid
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

		}
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
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.LoadbalancerServiceApplianceSetRef{}
			referenceModel.UUID = uuid
			m.ServiceApplianceSetRefs = append(m.ServiceApplianceSetRefs, referenceModel)

		}
	}

	return m, nil
}

// ListLoadbalancer lists Loadbalancer with list spec.
func ListLoadbalancer(tx *sql.Tx, spec *common.ListSpec) ([]*models.Loadbalancer, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "loadbalancer"
	spec.Fields = LoadbalancerFields
	spec.RefFields = LoadbalancerRefFields
	spec.BackRefFields = LoadbalancerBackRefFields
	result := models.MakeLoadbalancerSlice()

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
		m, err := scanLoadbalancer(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateLoadbalancer updates a resource
func UpdateLoadbalancer(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	// Prepare statement for updating data
	var updateLoadbalancerQuery = "update `loadbalancer` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateLoadbalancerQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateLoadbalancerQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateLoadbalancerQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateLoadbalancerQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateLoadbalancerQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateLoadbalancerQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateLoadbalancerQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".LoadbalancerProvider", "."); ok {
		updateLoadbalancerQuery += "`loadbalancer_provider` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".LoadbalancerProperties.VipSubnetID", "."); ok {
		updateLoadbalancerQuery += "`vip_subnet_id` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".LoadbalancerProperties.VipAddress", "."); ok {
		updateLoadbalancerQuery += "`vip_address` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".LoadbalancerProperties.Status", "."); ok {
		updateLoadbalancerQuery += "`status` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".LoadbalancerProperties.ProvisioningStatus", "."); ok {
		updateLoadbalancerQuery += "`provisioning_status` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".LoadbalancerProperties.OperatingStatus", "."); ok {
		updateLoadbalancerQuery += "`operating_status` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".LoadbalancerProperties.AdminState", "."); ok {
		updateLoadbalancerQuery += "`admin_state` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateLoadbalancerQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateLoadbalancerQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateLoadbalancerQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateLoadbalancerQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateLoadbalancerQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateLoadbalancerQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateLoadbalancerQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateLoadbalancerQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateLoadbalancerQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateLoadbalancerQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateLoadbalancerQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateLoadbalancerQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateLoadbalancerQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateLoadbalancerQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateLoadbalancerQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateLoadbalancerQuery += ","
	}

	updateLoadbalancerQuery =
		updateLoadbalancerQuery[:len(updateLoadbalancerQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateLoadbalancerQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateLoadbalancerQuery,
	}).Debug("update query")
	_, err = stmt.Exec(updatedValues...)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}

	if value, ok := common.GetValueByPath(model, "ServiceApplianceSetRefs", "."); ok {
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
				refQuery = "insert into `ref_loadbalancer_service_appliance_set` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_loadbalancer_service_appliance_set` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref ServiceApplianceSetRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_loadbalancer_service_appliance_set` where `from` = ? AND `to`= ?;"
				refValues = refValues[len(refValues)-2:]
			default:
				return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
			}
			stmt, err := tx.Prepare(refQuery)
			if err != nil {
				return errors.Wrap(err, "preparing ServiceApplianceSetRefs update statement failed")
			}
			_, err = stmt.Exec(refValues...)
			if err != nil {
				return errors.Wrap(err, "ServiceApplianceSetRefs update failed")
			}
		}
	}

	if value, ok := common.GetValueByPath(model, "VirtualMachineInterfaceRefs", "."); ok {
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
				refQuery = "insert into `ref_loadbalancer_virtual_machine_interface` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_loadbalancer_virtual_machine_interface` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref VirtualMachineInterfaceRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_loadbalancer_virtual_machine_interface` where `from` = ? AND `to`= ?;"
				refValues = refValues[len(refValues)-2:]
			default:
				return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
			}
			stmt, err := tx.Prepare(refQuery)
			if err != nil {
				return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs update statement failed")
			}
			_, err = stmt.Exec(refValues...)
			if err != nil {
				return errors.Wrap(err, "VirtualMachineInterfaceRefs update failed")
			}
		}
	}

	if value, ok := common.GetValueByPath(model, "ServiceInstanceRefs", "."); ok {
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
				refQuery = "insert into `ref_loadbalancer_service_instance` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_loadbalancer_service_instance` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref ServiceInstanceRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_loadbalancer_service_instance` where `from` = ? AND `to`= ?;"
				refValues = refValues[len(refValues)-2:]
			default:
				return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
			}
			stmt, err := tx.Prepare(refQuery)
			if err != nil {
				return errors.Wrap(err, "preparing ServiceInstanceRefs update statement failed")
			}
			_, err = stmt.Exec(refValues...)
			if err != nil {
				return errors.Wrap(err, "ServiceInstanceRefs update failed")
			}
		}
	}

	share, ok := common.GetValueByPath(model, ".Perms2.Share", ".")
	if ok {
		err = common.UpdateSharing(tx, "loadbalancer", string(uuid), share.([]interface{}))
		if err != nil {
			return err
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("updated")
	return err
}

// DeleteLoadbalancer deletes a resource
func DeleteLoadbalancer(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteLoadbalancerQuery
	selectQuery := "select count(uuid) from loadbalancer where uuid = ?"
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
