package storage

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

func TestExecutionSaveAndLoad(t *testing.T) {
	fdb.MustAPIVersion(710)
	db := fdb.MustOpenDefault()
	dir, err := directory.CreateOrOpen(db, []string{"test"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	ss := dir.Sub("executions")
	defer func() {
		db.Transact(func(tx fdb.Transaction) (interface{}, error) {
			tx.ClearRange(ss)

			return nil, nil
		})
	}()

	expected := &Execution{
		ShardID:          42,
		NamespaceID:      []byte("a-namespace-id"),
		WorkflowID:       "a-workflow-id",
		RunID:            []byte("a-run-id"),
		NextEventID:      big.NewInt(51),
		LastWriteVersion: big.NewInt(12),
		Data:             []byte("some-data"),
		DataEncoding:     "some-data-encoding",
		State:            []byte("some-state"),
		StateEncoding:    "some-state-encoding",
		DBRecordVersion:  big.NewInt(99),
	}
	err = expected.Save(db, ss)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewExecutionRepository(db, ss)
	actual, err := repo.Find(42, []byte("a-namespace-id"), "a-workflow-id", []byte("a-run-id"))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("didn't load expected serialized data from the database\n\texpected: %#v\n\tgot:      %#v", expected, actual)
	}
}

func TestCurrentExecutionSaveAndLoad(t *testing.T) {
	fdb.MustAPIVersion(710)
	db := fdb.MustOpenDefault()
	dir, err := directory.CreateOrOpen(db, []string{"test"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	ss := dir.Sub("executions")
	defer func() {
		db.Transact(func(tx fdb.Transaction) (interface{}, error) {
			tx.ClearRange(ss)

			return nil, nil
		})
	}()

	expected := &CurrentExecution{
		ShardID:          42,
		NamespaceID:      []byte("a-namespace-id"),
		WorkflowID:       "a-workflow-id",
		RunID:            []byte("a-run-id"),
		CreateRequestID:  "a-create-request-id",
		State:            18,
		Status:           21,
		StartVersion:     big.NewInt(51),
		LastWriteVersion: big.NewInt(12),
	}
	err = expected.Save(db, ss)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewCurrentExecutionRepository(db, ss)
	actual, err := repo.Find(42, []byte("a-namespace-id"), "a-workflow-id")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("didn't load expected serialized data from the database\n\texpected: %#v\n\tgot:      %#v", expected, actual)
	}
}
