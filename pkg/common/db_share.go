package common

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Juniper/contrail/pkg/generated/models"
)

//CreateSharing creates sharing information in DB.
func CreateSharing(tx *sql.Tx, table string, uuid string, shares []*models.ShareType) error {
	for _, share := range shares {
		shareParts := strings.Split(share.Tenant, ":")
		shareType := "tenant"
		if shareParts[0] == "domain" {
			shareType = "tenant"
		}
		to := shareParts[0]
		_, err := tx.Exec(
			fmt.Sprintf("insert into `%s_share_%s`(`uuid`, `access`, `to`)",
				shareType, table), uuid, share.TenantAccess, to)
		if err != nil {
			return err
		}
	}
	return nil
}

//UpdateSharing updates sharing data for a object by UUID.
func UpdateSharing(tx *sql.Tx, table string, uuid string, shares []*models.ShareType) error {
	return nil
}
