package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestContrailAnalyticsDatabaseNodeRole(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "contrail_analytics_database_node_role")
	defer func() {
		common.ClearTable(db, "contrail_analytics_database_node_role")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeContrailAnalyticsDatabaseNodeRole()
	model.UUID = "contrail_analytics_database_node_role_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "contrail_analytics_database_node_role_dummy"}

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateContrailAnalyticsDatabaseNodeRole(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListContrailAnalyticsDatabaseNodeRole(tx, &common.ListSpec{Limit: 1})
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
		return DeleteContrailAnalyticsDatabaseNodeRole(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListContrailAnalyticsDatabaseNodeRole(tx, &common.ListSpec{Limit: 1})
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
