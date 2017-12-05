package db

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"strings"
)

const insertVPNGroupQuery = "insert into `vpn_group` (`type`,`display_name`,`key_value_pair`,`uuid`,`provisioning_start_time`,`provisioning_log`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`fq_name`,`provisioning_progress`,`provisioning_progress_stage`,`provisioning_state`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVPNGroupQuery = "update `vpn_group` set `type` = ?,`display_name` = ?,`key_value_pair` = ?,`uuid` = ?,`provisioning_start_time` = ?,`provisioning_log` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`fq_name` = ?,`provisioning_progress` = ?,`provisioning_progress_stage` = ?,`provisioning_state` = ?;"
const deleteVPNGroupQuery = "delete from `vpn_group` where uuid = ?"
const listVPNGroupQuery = "select `vpn_group`.`type`,`vpn_group`.`display_name`,`vpn_group`.`key_value_pair`,`vpn_group`.`uuid`,`vpn_group`.`provisioning_start_time`,`vpn_group`.`provisioning_log`,`vpn_group`.`user_visible`,`vpn_group`.`last_modified`,`vpn_group`.`group`,`vpn_group`.`group_access`,`vpn_group`.`owner`,`vpn_group`.`owner_access`,`vpn_group`.`other_access`,`vpn_group`.`enable`,`vpn_group`.`description`,`vpn_group`.`created`,`vpn_group`.`creator`,`vpn_group`.`perms2_owner`,`vpn_group`.`perms2_owner_access`,`vpn_group`.`global_access`,`vpn_group`.`share`,`vpn_group`.`fq_name`,`vpn_group`.`provisioning_progress`,`vpn_group`.`provisioning_progress_stage`,`vpn_group`.`provisioning_state` from `vpn_group`"
const showVPNGroupQuery = "select `vpn_group`.`type`,`vpn_group`.`display_name`,`vpn_group`.`key_value_pair`,`vpn_group`.`uuid`,`vpn_group`.`provisioning_start_time`,`vpn_group`.`provisioning_log`,`vpn_group`.`user_visible`,`vpn_group`.`last_modified`,`vpn_group`.`group`,`vpn_group`.`group_access`,`vpn_group`.`owner`,`vpn_group`.`owner_access`,`vpn_group`.`other_access`,`vpn_group`.`enable`,`vpn_group`.`description`,`vpn_group`.`created`,`vpn_group`.`creator`,`vpn_group`.`perms2_owner`,`vpn_group`.`perms2_owner_access`,`vpn_group`.`global_access`,`vpn_group`.`share`,`vpn_group`.`fq_name`,`vpn_group`.`provisioning_progress`,`vpn_group`.`provisioning_progress_stage`,`vpn_group`.`provisioning_state` from `vpn_group` where uuid = ?"

const insertVPNGroupLocationQuery = "insert into `ref_vpn_group_location` (`from`, `to` ) values (?, ?);"

func CreateVPNGroup(tx *sql.Tx, model *models.VPNGroup) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVPNGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.Type),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningLog),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		utils.MustJSON(model.FQName),
		int(model.ProvisioningProgress),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningState))

	stmtLocationRef, err := tx.Prepare(insertVPNGroupLocationQuery)
	if err != nil {
		return err
	}
	defer stmtLocationRef.Close()
	for _, ref := range model.LocationRefs {
		_, err = stmtLocationRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanVPNGroup(rows *sql.Rows) (*models.VPNGroup, error) {
	m := models.MakeVPNGroup()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.Type,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.UUID,
		&m.ProvisioningStartTime,
		&m.ProvisioningLog,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&jsonFQName,
		&m.ProvisioningProgress,
		&m.ProvisioningProgressStage,
		&m.ProvisioningState); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildVPNGroupWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["type"]; ok {
		results = append(results, "type = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListVPNGroup(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.VPNGroup, error) {
	result := models.MakeVPNGroupSlice()
	whereQuery, values := buildVPNGroupWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listVPNGroupQuery)
	query.WriteRune(' ')
	query.WriteString(whereQuery)
	query.WriteRune(' ')
	query.WriteString(pagenationQuery)
	rows, err = tx.Query(query.String(), values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanVPNGroup(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowVPNGroup(tx *sql.Tx, uuid string) (*models.VPNGroup, error) {
	rows, err := tx.Query(showVPNGroupQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanVPNGroup(rows)
	}
	return nil, nil
}

func UpdateVPNGroup(tx *sql.Tx, uuid string, model *models.VPNGroup) error {
	return nil
}

func DeleteVPNGroup(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVPNGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
