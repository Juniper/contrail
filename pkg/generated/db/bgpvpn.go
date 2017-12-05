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

const insertBGPVPNQuery = "insert into `bgpvpn` (`display_name`,`route_target`,`export_route_target_list_route_target`,`bgpvpn_type`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`fq_name`,`import_route_target_list_route_target`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBGPVPNQuery = "update `bgpvpn` set `display_name` = ?,`route_target` = ?,`export_route_target_list_route_target` = ?,`bgpvpn_type` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`fq_name` = ?,`import_route_target_list_route_target` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`uuid` = ?;"
const deleteBGPVPNQuery = "delete from `bgpvpn` where uuid = ?"
const listBGPVPNQuery = "select `bgpvpn`.`display_name`,`bgpvpn`.`route_target`,`bgpvpn`.`export_route_target_list_route_target`,`bgpvpn`.`bgpvpn_type`,`bgpvpn`.`description`,`bgpvpn`.`created`,`bgpvpn`.`creator`,`bgpvpn`.`user_visible`,`bgpvpn`.`last_modified`,`bgpvpn`.`owner_access`,`bgpvpn`.`other_access`,`bgpvpn`.`group`,`bgpvpn`.`group_access`,`bgpvpn`.`owner`,`bgpvpn`.`enable`,`bgpvpn`.`fq_name`,`bgpvpn`.`import_route_target_list_route_target`,`bgpvpn`.`key_value_pair`,`bgpvpn`.`perms2_owner_access`,`bgpvpn`.`global_access`,`bgpvpn`.`share`,`bgpvpn`.`perms2_owner`,`bgpvpn`.`uuid` from `bgpvpn`"
const showBGPVPNQuery = "select `bgpvpn`.`display_name`,`bgpvpn`.`route_target`,`bgpvpn`.`export_route_target_list_route_target`,`bgpvpn`.`bgpvpn_type`,`bgpvpn`.`description`,`bgpvpn`.`created`,`bgpvpn`.`creator`,`bgpvpn`.`user_visible`,`bgpvpn`.`last_modified`,`bgpvpn`.`owner_access`,`bgpvpn`.`other_access`,`bgpvpn`.`group`,`bgpvpn`.`group_access`,`bgpvpn`.`owner`,`bgpvpn`.`enable`,`bgpvpn`.`fq_name`,`bgpvpn`.`import_route_target_list_route_target`,`bgpvpn`.`key_value_pair`,`bgpvpn`.`perms2_owner_access`,`bgpvpn`.`global_access`,`bgpvpn`.`share`,`bgpvpn`.`perms2_owner`,`bgpvpn`.`uuid` from `bgpvpn` where uuid = ?"

func CreateBGPVPN(tx *sql.Tx, model *models.BGPVPN) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBGPVPNQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		utils.MustJSON(model.RouteTargetList.RouteTarget),
		utils.MustJSON(model.ExportRouteTargetList.RouteTarget),
		string(model.BGPVPNType),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		utils.MustJSON(model.FQName),
		utils.MustJSON(model.ImportRouteTargetList.RouteTarget),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID))

	return err
}

func scanBGPVPN(rows *sql.Rows) (*models.BGPVPN, error) {
	m := models.MakeBGPVPN()

	var jsonRouteTargetListRouteTarget string

	var jsonExportRouteTargetListRouteTarget string

	var jsonFQName string

	var jsonImportRouteTargetListRouteTarget string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.DisplayName,
		&jsonRouteTargetListRouteTarget,
		&jsonExportRouteTargetListRouteTarget,
		&m.BGPVPNType,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Enable,
		&jsonFQName,
		&jsonImportRouteTargetListRouteTarget,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonRouteTargetListRouteTarget), &m.RouteTargetList.RouteTarget)

	json.Unmarshal([]byte(jsonExportRouteTargetListRouteTarget), &m.ExportRouteTargetList.RouteTarget)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonImportRouteTargetListRouteTarget), &m.ImportRouteTargetList.RouteTarget)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildBGPVPNWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["bgpvpn_type"]; ok {
		results = append(results, "bgpvpn_type = ?")
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

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListBGPVPN(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.BGPVPN, error) {
	result := models.MakeBGPVPNSlice()
	whereQuery, values := buildBGPVPNWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listBGPVPNQuery)
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
		m, _ := scanBGPVPN(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowBGPVPN(tx *sql.Tx, uuid string) (*models.BGPVPN, error) {
	rows, err := tx.Query(showBGPVPNQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanBGPVPN(rows)
	}
	return nil, nil
}

func UpdateBGPVPN(tx *sql.Tx, uuid string, model *models.BGPVPN) error {
	return nil
}

func DeleteBGPVPN(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteBGPVPNQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
