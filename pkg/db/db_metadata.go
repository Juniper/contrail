package db

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

//MetaData represents resource meta data.
type MetaData struct {
	UUID   string
	FQName []string
	Type   string
}

//FQNameToString returns string representation of FQName.
func FQNameToString(fqName []string) string {
	return strings.Join(fqName, ":")
}

//CreateMetaData creates fqname, uuid pair with type.
func CreateMetaData(tx *sql.Tx, metaData *MetaData) error {
	_, err := tx.Exec("insert into `metadata` (`uuid`,`type`,`fq_name`) values (?,?,?);",
		metaData.UUID, metaData.Type, FQNameToString(metaData.FQName))
	return errors.Wrap(err, "failed to create metadata")
}

//GetMetaData gets metadata from database.
func GetMetaData(tx *sql.Tx, uuid string, fqName []string) (*MetaData, error) {
	var query bytes.Buffer
	query.WriteString("select `uuid`,`type`,`fq_name` from `metadata` where ")
	var row *sql.Row

	log.Debug(fqName)
	if uuid != "" {
		query.WriteString("uuid = ?")
		row = tx.QueryRow(query.String(), uuid)
	} else if fqName != nil {
		query.WriteString("fq_name = ?")
		log.Debug(query.String())
		row = tx.QueryRow(query.String(), FQNameToString(fqName))
	} else {
		return nil, fmt.Errorf("uuid and fqName unspecified ")
	}
	metaData := &MetaData{}
	var fqNameString string
	err := row.Scan(&metaData.UUID, &metaData.Type, &fqNameString)
	metaData.FQName = models.ParseFQName(fqNameString)
	return metaData, errors.Wrap(err, "failed to get metadata")
}

//DeleteMetaData deltes metadata by uuid.
func DeleteMetaData(tx *sql.Tx, uuid string) error {
	_, err := tx.Exec("delete from `metadata` where uuid = ?", uuid)
	return errors.Wrap(err, "failed to delete metadata")
}
