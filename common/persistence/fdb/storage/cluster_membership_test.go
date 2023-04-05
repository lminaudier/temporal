package storage

import (
	"reflect"
	"testing"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

func TestClusterMembershipSaveAndLoad(t *testing.T) {
	fdb.MustAPIVersion(710)
	db := fdb.MustOpenDefault()
	dir, err := directory.CreateOrOpen(db, []string{"test"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	ss := dir.Sub("cluster_memberships")
	defer func() {
		db.Transact(func(tx fdb.Transaction) (interface{}, error) {
			tx.ClearRange(ss)

			return nil, nil
		})
	}()

	expected := &ClusterMembership{
		MembershipPartition: 42,
		HostID:              []byte("a-host-id"),
		RPCAddress:          "rpc-address",
		RPCPort:             18,
		Role:                21,
		SessionStart:        time.Now().UTC(),
		LastHeartbeat:       time.Now().UTC().Add(-1 * time.Minute),
		RecordExpiry:        2 * time.Minute,
	}
	err = expected.Save(db, ss)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewClusterMembershipRepository(db, ss)
	actual, err := repo.Find(42, []byte("a-host-id"))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("didn't load expected serialized data from the database\n\texpected: %#v\n\tgot:      %#v", expected, actual)
	}
}
