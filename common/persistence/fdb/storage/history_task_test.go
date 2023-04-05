package storage

import (
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

func TestHistoryImmediateTaskSaveAndLoad(t *testing.T) {
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

	expected := &HistoryImmediateTask{
		ShardID:      42,
		CategoryID:   51,
		TaskID:       big.NewInt(18),
		Data:         []byte("some-data"),
		DataEncoding: "some-data-encoding",
	}
	err = expected.Save(db, ss)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewHistoryImmediateTaskRepository(db, ss)
	actual, err := repo.Find(42, 51, big.NewInt(18))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("didn't load expected serialized data from the database\n\texpected: %#v\n\tgot:      %#v", expected, actual)
	}
}

func TestHistoryScheduledTaskSaveAndLoad(t *testing.T) {
	fdb.MustAPIVersion(710)
	db := fdb.MustOpenDefault()
	dir, err := directory.CreateOrOpen(db, []string{"test"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	ss := dir.Sub("history_scheduled_tasks")
	defer func() {
		db.Transact(func(tx fdb.Transaction) (interface{}, error) {
			tx.ClearRange(ss)

			return nil, nil
		})
	}()

	now := time.Now()
	expected := &HistoryScheduledTask{
		ShardID:             42,
		CategoryID:          51,
		VisibilityTimestamp: now,
		TaskID:              big.NewInt(18),
		Data:                []byte("some-data"),
		DataEncoding:        "some-data-encoding",
	}
	err = expected.Save(db, ss)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewHistoryScheduledTaskRepository(db, ss)
	actual, err := repo.Find(42, 51, now, big.NewInt(18))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("didn't load expected serialized data from the database\n\texpected: %#v\n\tgot:      %#v", expected, actual)
	}
}
