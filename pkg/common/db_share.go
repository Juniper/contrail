package common

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"
)

//CreateSharing creates sharing information in DB.
func CreateSharing(tx *sql.Tx, table string, uuid string, shares []*models.ShareType) error {
	for _, share := range shares {
		err := createSharingEntry(tx, table, uuid, share.Tenant, int(share.TenantAccess))
		if err != nil {
			return err
		}
	}
	return nil
}

//UpdateSharing updates sharing data for a object by UUID.
func UpdateSharing(tx *sql.Tx, table string, uuid string, shares []interface{}) error {
	if len(shares) == 0 {
		return nil
	}
	_, err := tx.Exec(
		fmt.Sprintf("delete from `domain_share_%s` where `uuid` = ?;", table), uuid)
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		fmt.Sprintf("delete from `tenant_share_%s` where `uuid` = ?;", table), uuid)
	if err != nil {
		return err
	}
	for _, share := range shares {
		err = createSharingEntry(tx, table, uuid, InterfaceToString(share.(map[string]interface{})["tenant"]), InterfaceToInt(share.(map[string]interface{})["tenant_access"]))
		if err != nil {
			return err
		}
	}
	return nil
}

func createSharingEntry(tx *sql.Tx, table string, uuid string, tenant string, tenantAccess int) error {
	shareParts := strings.Split(tenant, ":")
	shareType := "domain"
	if len(shareParts) > 1 {
		shareType = "tenant"
	}
	resourceMetaData, err := GetMetaData(tx, "", shareParts)
	if err != nil {
		return errors.Wrap(err, "can't find resource")
	}
	to := resourceMetaData.UUID
	_, err = tx.Exec(
		fmt.Sprintf("insert into `%s_share_%s`(`uuid`, `access`, `to`) values (?,?,?);",
			shareType, table), uuid, tenantAccess, to)
	return err
}
