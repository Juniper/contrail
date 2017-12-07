package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestServiceGroup(t *testing.T) {
	t.Parallel()
	model := models.MakeServiceGroup()
	model.UUID = "dummy_uuid"
	db := testDB

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceGroup(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListServiceGroup(tx, &common.ListSpec{Limit: 1})
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

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		model, err := ShowServiceGroup(tx, model.UUID)
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

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceGroup(tx, model.UUID)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListServiceGroup(tx, &common.ListSpec{Limit: 1})
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
