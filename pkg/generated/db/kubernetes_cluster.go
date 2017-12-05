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

const insertKubernetesClusterQuery = "insert into `kubernetes_cluster` (`kuberunetes_dashboard`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`fq_name`,`contrail_cluster_id`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateKubernetesClusterQuery = "update `kubernetes_cluster` set `kuberunetes_dashboard` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`contrail_cluster_id` = ?;"
const deleteKubernetesClusterQuery = "delete from `kubernetes_cluster` where uuid = ?"
const listKubernetesClusterQuery = "select `kubernetes_cluster`.`kuberunetes_dashboard`,`kubernetes_cluster`.`creator`,`kubernetes_cluster`.`user_visible`,`kubernetes_cluster`.`last_modified`,`kubernetes_cluster`.`owner_access`,`kubernetes_cluster`.`other_access`,`kubernetes_cluster`.`group`,`kubernetes_cluster`.`group_access`,`kubernetes_cluster`.`owner`,`kubernetes_cluster`.`enable`,`kubernetes_cluster`.`description`,`kubernetes_cluster`.`created`,`kubernetes_cluster`.`display_name`,`kubernetes_cluster`.`key_value_pair`,`kubernetes_cluster`.`share`,`kubernetes_cluster`.`perms2_owner`,`kubernetes_cluster`.`perms2_owner_access`,`kubernetes_cluster`.`global_access`,`kubernetes_cluster`.`uuid`,`kubernetes_cluster`.`fq_name`,`kubernetes_cluster`.`contrail_cluster_id` from `kubernetes_cluster`"
const showKubernetesClusterQuery = "select `kubernetes_cluster`.`kuberunetes_dashboard`,`kubernetes_cluster`.`creator`,`kubernetes_cluster`.`user_visible`,`kubernetes_cluster`.`last_modified`,`kubernetes_cluster`.`owner_access`,`kubernetes_cluster`.`other_access`,`kubernetes_cluster`.`group`,`kubernetes_cluster`.`group_access`,`kubernetes_cluster`.`owner`,`kubernetes_cluster`.`enable`,`kubernetes_cluster`.`description`,`kubernetes_cluster`.`created`,`kubernetes_cluster`.`display_name`,`kubernetes_cluster`.`key_value_pair`,`kubernetes_cluster`.`share`,`kubernetes_cluster`.`perms2_owner`,`kubernetes_cluster`.`perms2_owner_access`,`kubernetes_cluster`.`global_access`,`kubernetes_cluster`.`uuid`,`kubernetes_cluster`.`fq_name`,`kubernetes_cluster`.`contrail_cluster_id` from `kubernetes_cluster` where uuid = ?"

func CreateKubernetesCluster(tx *sql.Tx, model *models.KubernetesCluster) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertKubernetesClusterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.KuberunetesDashboard),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.ContrailClusterID))

	return err
}

func scanKubernetesCluster(rows *sql.Rows) (*models.KubernetesCluster, error) {
	m := models.MakeKubernetesCluster()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.KuberunetesDashboard,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID,
		&jsonFQName,
		&m.ContrailClusterID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildKubernetesClusterWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["kuberunetes_dashboard"]; ok {
		results = append(results, "kuberunetes_dashboard = ?")
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

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
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

	if value, ok := where["contrail_cluster_id"]; ok {
		results = append(results, "contrail_cluster_id = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListKubernetesCluster(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.KubernetesCluster, error) {
	result := models.MakeKubernetesClusterSlice()
	whereQuery, values := buildKubernetesClusterWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listKubernetesClusterQuery)
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
		m, _ := scanKubernetesCluster(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowKubernetesCluster(tx *sql.Tx, uuid string) (*models.KubernetesCluster, error) {
	rows, err := tx.Query(showKubernetesClusterQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanKubernetesCluster(rows)
	}
	return nil, nil
}

func UpdateKubernetesCluster(tx *sql.Tx, uuid string, model *models.KubernetesCluster) error {
	return nil
}

func DeleteKubernetesCluster(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteKubernetesClusterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
