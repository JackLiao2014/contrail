package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestBridgeDomain(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "bridge_domain")
	defer func() {
		common.ClearTable(db, "bridge_domain")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeBridgeDomain()
	model.UUID = "bridge_domain_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "bridge_domain_dummy"}

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateBridgeDomain(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListBridgeDomain(tx, &common.ListSpec{Limit: 1})
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
		return DeleteBridgeDomain(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListBridgeDomain(tx, &common.ListSpec{Limit: 1})
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
