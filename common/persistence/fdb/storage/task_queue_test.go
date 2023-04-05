package storage

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

func TestTaskQueueSaveAndLoad(t *testing.T) {
	fdb.MustAPIVersion(710)
	db := fdb.MustOpenDefault()
	dir, err := directory.CreateOrOpen(db, []string{"test"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	ss := dir.Sub("task_queues")
	defer func() {
		db.Transact(func(tx fdb.Transaction) (interface{}, error) {
			tx.ClearRange(ss)

			return nil, nil
		})
	}()

	expected := &TaskQueue{
		RangeHash:    big.NewInt(42),
		TaskQueueID:  []byte("a-task-queue-id"),
		RangeID:      big.NewInt(51),
		Data:         []byte("some-data"),
		DataEncoding: "an-encoding",
	}
	err = expected.Save(db, ss)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewTaskQueueRepository(db, ss)
	actual, err := repo.Find(big.NewInt(42), []byte("a-task-queue-id"))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("didn't load expected serialized data from the database\n\texpected: %#v\n\tgot:      %#v", expected, actual)
	}
}
