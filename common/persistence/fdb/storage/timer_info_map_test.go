package storage

import (
	"reflect"
	"testing"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

func TestTimerInfoMapSaveAndLoad(t *testing.T) {
	fdb.MustAPIVersion(710)
	db := fdb.MustOpenDefault()
	dir, err := directory.CreateOrOpen(db, []string{"test"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	ss := dir.Sub("timer_info_maps")
	defer func() {
		db.Transact(func(tx fdb.Transaction) (interface{}, error) {
			tx.ClearRange(ss)

			return nil, nil
		})
	}()

	expected := &TimerInfoMap{
		ShardID:      42,
		NamespaceID:  []byte("a-namespace-id"),
		WorkflowID:   "a-workflow-id",
		RunID:        []byte("a-run-id"),
		TimerID:      "a-timer-id",
		Data:         []byte("some-data"),
		DataEncoding: "some-data-encoding",
	}
	err = expected.Save(db, ss)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewTimerInfoMapRepository(db, ss)
	actual, err := repo.Find(42, []byte("a-namespace-id"), "a-workflow-id", []byte("a-run-id"), "a-timer-id")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("didn't load expected serialized data from the database\n\texpected: %#v\n\tgot:      %#v", expected, actual)
	}
}
