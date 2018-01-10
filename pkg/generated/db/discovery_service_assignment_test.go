package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestDiscoveryServiceAssignment(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "discovery_service_assignment")
	defer func() {
		common.ClearTable(db, "discovery_service_assignment")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeDiscoveryServiceAssignment()
	model.UUID = "discovery_service_assignment_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "discovery_service_assignment_dummy"}

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateDiscoveryServiceAssignment(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListDiscoveryServiceAssignment(tx, &common.ListSpec{Limit: 1})
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
		return DeleteDiscoveryServiceAssignment(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListDiscoveryServiceAssignment(tx, &common.ListSpec{Limit: 1})
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
