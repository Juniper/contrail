package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	dbPkg "github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestVirtualNetwork(t *testing.T) {
	t.Parallel()
	model := models.MakeVirtualNetwork()
	model.UUID = "vn_uuid"
	db := testDB

	policy := models.MakeNetworkPolicy()
	policy.UUID = "policy_uuid"

	model.NetworkPolicyRefs = append(model.NetworkPolicyRefs,
		&models.VirtualNetworkNetworkPolicyRef{
			UUID: policy.UUID,
			Attr: models.MakeVirtualNetworkPolicyType(),
		})
	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		err := dbPkg.CreateNetworkPolicy(tx, policy)
		if err != nil {
			return err
		}
		return dbPkg.CreateVirtualNetwork(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		list, err := dbPkg.ListVirtualNetwork(tx, &common.ListSpec{Limit: 1})
		if err != nil {
			return err
		}
		if len(list) != 1 {
			return fmt.Errorf("expected one element")
		}
		if len(list[0].NetworkPolicyRefs) != 1 {
			return fmt.Errorf("can't get reference")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		model, err := dbPkg.ShowVirtualNetwork(tx, model.UUID)
		if err != nil {
			return err
		}
		if model == nil || model.UUID != "vn_uuid" {
			return fmt.Errorf("uuid mismatch")
		}
		return nil
	})
	if err != nil {
		t.Fatal("show failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		err := dbPkg.DeleteVirtualNetwork(tx, model.UUID)
		if err != nil {
			return err
		}
		return dbPkg.DeleteNetworkPolicy(tx, policy.UUID)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := dbPkg.ListVirtualNetwork(tx, &common.ListSpec{Limit: 1})
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
