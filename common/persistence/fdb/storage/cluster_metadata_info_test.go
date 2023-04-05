package storage

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

func TestClusterMetadataInfoSaveAndLoad(t *testing.T) {
	fdb.MustAPIVersion(710)
	db := fdb.MustOpenDefault()
	dir, err := directory.CreateOrOpen(db, []string{"test"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	ss := dir.Sub("cluster_metadata_info")
	defer func() {
		db.Transact(func(tx fdb.Transaction) (interface{}, error) {
			tx.ClearRange(ss)

			return nil, nil
		})
	}()

	expected := &ClusterMetadataInfo{
		MetadataPartition: 42,
		ClusterName:       "a-cluster-name",
		Data:              []byte("some-data"),
		DataEncoding:      "some-data-encoding",
		Version:           big.NewInt(18),
	}
	err = expected.Save(db, ss)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewClusterMetadataInfoRepository(db, ss)
	actual, err := repo.Find(42, "a-cluster-name")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("didn't load expected serialized data from the database\n\texpected: %#v\n\tgot:      %#v", expected, actual)
	}
}

func TestClusterMetadataInfoList(t *testing.T) {
	fdb.MustAPIVersion(710)
	db := fdb.MustOpenDefault()
	dir, err := directory.CreateOrOpen(db, []string{"test"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	ss := dir.Sub("cluster_metadata_info")
	defer func() {
		db.Transact(func(tx fdb.Transaction) (interface{}, error) {
			tx.ClearRange(ss)

			return nil, nil
		})
	}()

	expected1 := &ClusterMetadataInfo{
		MetadataPartition: 42,
		ClusterName:       "a-cluster-name",
		Data:              []byte("some-data"),
		DataEncoding:      "some-data-encoding",
		Version:           big.NewInt(18),
	}
	err = expected1.Save(db, ss)
	if err != nil {
		t.Fatal(err)
	}

	expected2 := &ClusterMetadataInfo{
		MetadataPartition: 101,
		ClusterName:       "another-cluster-name",
		Data:              []byte("other-data"),
		DataEncoding:      "some-other-encoding",
		Version:           big.NewInt(111),
	}
	err = expected2.Save(db, ss)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewClusterMetadataInfoRepository(db, ss)
	actual, err := repo.List()
	if err != nil {
		t.Fatal(err)
	}

	expected := []*ClusterMetadataInfo{expected1, expected2}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("didn't load expected serialized data from the database\n\texpected: %#v\n\tgot:      %#v", expected, actual)
	}
}
