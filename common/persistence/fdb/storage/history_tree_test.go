package storage

import (
	"reflect"
	"testing"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

func TestHistoryTreeSaveAndLoad(t *testing.T) {
	fdb.MustAPIVersion(710)
	db := fdb.MustOpenDefault()
	dir, err := directory.CreateOrOpen(db, []string{"test"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	ss := dir.Sub("history_immediate_tasks")
	defer func() {
		db.Transact(func(tx fdb.Transaction) (interface{}, error) {
			tx.ClearRange(ss)

			return nil, nil
		})
	}()

	expected := &HistoryTree{
		ShardID:      42,
		TreeID:       []byte("a-tree-id"),
		BranchID:     []byte("a-branch-id"),
		Data:         []byte("some-data"),
		DataEncoding: "some-data-encoding",
	}
	err = expected.Save(db, ss)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewHistoryTreeRepository(db, ss)
	actual, err := repo.Find(42, []byte("a-tree-id"), []byte("a-branch-id"))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("didn't load expected serialized data from the database\n\texpected: %#v\n\tgot:      %#v", expected, actual)
	}
}
