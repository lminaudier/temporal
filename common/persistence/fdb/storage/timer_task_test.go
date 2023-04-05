package storage

import (
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

func TestTimerTaskSaveAndLoad(t *testing.T) {
	fdb.MustAPIVersion(710)
	db := fdb.MustOpenDefault()
	dir, err := directory.CreateOrOpen(db, []string{"test"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	ss := dir.Sub("timer_tasks")
	defer func() {
		db.Transact(func(tx fdb.Transaction) (interface{}, error) {
			tx.ClearRange(ss)

			return nil, nil
		})
	}()

	now := time.Now()
	expected := &TimerTask{
		ShardID:             42,
		VisibilityTimestamp: now,
		TaskID:              big.NewInt(18),
		Data:                []byte("some-data"),
		DataEncoding:        "some-data-encoding",
	}
	err = expected.Save(db, ss)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewTimerTaskRepository(db, ss)
	actual, err := repo.Find(42, now, big.NewInt(18))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("didn't load expected serialized data from the database\n\texpected: %#v\n\tgot:      %#v", expected, actual)
	}
}
