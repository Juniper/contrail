package db

import (
	"database/sql"
	"fmt"
	"testing"

	dbPkg "github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestOpenstackCluster(t *testing.T) {
	t.Parallel()
	model := models.MakeOpenstackCluster()
	model.UUID = "dummy_uuid"
	db := testServer.DB

	err := dbPkg.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateOpenstackCluster(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = dbPkg.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListOpenstackCluster(tx, &dbPkg.ListSpec{Limit: 1})
		if err != nil {
			return err
		}
		if len(models) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	err = dbPkg.DoInTransaction(db, func(tx *sql.Tx) error {
		model, err := ShowOpenstackCluster(tx, model.UUID)
		if err != nil {
			return err
		}
		if model == nil || model.UUID != "dummy_uuid" {
			return fmt.Errorf("show failed")
		}
		return nil
	})
	if err != nil {
		t.Fatal("show failed", err)
	}

	err = dbPkg.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteOpenstackCluster(tx, model.UUID)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = dbPkg.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListOpenstackCluster(tx, &dbPkg.ListSpec{Limit: 1})
		if err != nil {
			return err
		}
		if len(models) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
