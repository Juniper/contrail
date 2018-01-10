package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestVirtualMachineInterface(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "virtual_machine_interface")
	defer func() {
		common.ClearTable(db, "virtual_machine_interface")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeVirtualMachineInterface()
	model.UUID = "virtual_machine_interface_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "virtual_machine_interface_dummy"}

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateVirtualMachineInterface(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListVirtualMachineInterface(tx, &common.ListSpec{Limit: 1})
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
		return DeleteVirtualMachineInterface(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListVirtualMachineInterface(tx, &common.ListSpec{Limit: 1})
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
