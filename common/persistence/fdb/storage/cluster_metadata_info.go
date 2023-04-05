package storage

import (
	"math/big"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// CREATE TABLE cluster_metadata_info (
//
//	metadata_partition        INTEGER NOT NULL,
//	cluster_name              VARCHAR(255) NOT NULL,
//	data                      BYTEA NOT NULL,
//	data_encoding             VARCHAR(16) NOT NULL,
//	version                   BIGINT NOT NULL,
//	PRIMARY KEY(metadata_partition, cluster_name)
//
// );
type ClusterMetadataInfo struct {
	MetadataPartition int64
	ClusterName       string
	Data              []byte
	DataEncoding      string
	Version           *big.Int
}

func (m *ClusterMetadataInfo) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{m.GetMetadataPartitionBytes(), m.GetClusterNameBytes()})
}

func ClusterMetadataInfoPrimaryKeyIndexPrefix(ss subspace.Subspace) fdb.Key {
	return ss.Pack(tuple.Tuple{"metadata_partition_cluster_name_idx"})
}

func (m *ClusterMetadataInfo) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"metadata_partition_cluster_name_idx", m.GetMetadataPartitionBytes(), m.GetClusterNameBytes()})
}

func (m *ClusterMetadataInfo) GetMetadataPartitionKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"metadata_partition"})
}

func (m *ClusterMetadataInfo) GetMetadataPartitionBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(m.MetadataPartition))
	return buff
}

func (m *ClusterMetadataInfo) SetMetadataPartition(data []byte) {
	m.MetadataPartition = int64(BytesOrder.Uint64(data))
}

func (m *ClusterMetadataInfo) GetClusterNameKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"cluster_name"})
}

func (m *ClusterMetadataInfo) GetClusterNameBytes() []byte {
	return []byte(m.ClusterName)
}

func (m *ClusterMetadataInfo) SetClusterName(data []byte) {
	m.ClusterName = string(data)
}

func (m *ClusterMetadataInfo) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (m *ClusterMetadataInfo) GetDataBytes() []byte {
	return m.Data
}

func (m *ClusterMetadataInfo) SetData(data []byte) {
	m.Data = data
}

func (m *ClusterMetadataInfo) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (m *ClusterMetadataInfo) GetDataEncodingBytes() []byte {
	return []byte(m.DataEncoding)
}

func (m *ClusterMetadataInfo) SetDataEncoding(data []byte) {
	m.DataEncoding = string(data)
}

func (m *ClusterMetadataInfo) GetVersionKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"version"})
}

func (m *ClusterMetadataInfo) GetVersionBytes() []byte {
	return m.Version.Bytes()
}

func (m *ClusterMetadataInfo) SetVersion(data []byte) {
	m.Version = big.NewInt(0).SetBytes(data)
}

func (m *ClusterMetadataInfo) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		tx.Set(m.GetDataKey(ss), m.GetDataBytes())
		tx.Set(m.GetDataEncodingKey(ss), m.GetDataEncodingBytes())
		tx.Set(m.GetVersionKey(ss), m.GetVersionBytes())

		// Indexes
		tx.Set(m.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type ClusterMetadataInfoRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewClusterMetadataInfoRepository(tr fdb.Transactor, ss subspace.Subspace) *ClusterMetadataInfoRepository {
	return &ClusterMetadataInfoRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *ClusterMetadataInfoRepository) Find(metadataPartiton int64, clusterName string) (*ClusterMetadataInfo, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		m := &ClusterMetadataInfo{
			MetadataPartition: metadataPartiton,
			ClusterName:       clusterName,
		}

		m.SetData(rtx.Get(m.GetDataKey(r.ss)).MustGet())
		m.SetDataEncoding(rtx.Get(m.GetDataEncodingKey(r.ss)).MustGet())
		m.SetVersion(rtx.Get(m.GetVersionKey(r.ss)).MustGet())

		return m, nil
	})

	return res.(*ClusterMetadataInfo), err
}

func (r *ClusterMetadataInfoRepository) List() ([]*ClusterMetadataInfo, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		result := make([]*ClusterMetadataInfo, 0)

		kr, err := fdb.PrefixRange(ClusterMetadataInfoPrimaryKeyIndexPrefix(r.ss))
		if err != nil {
			return nil, err
		}
		it := rtx.GetRange(kr, fdb.RangeOptions{}).Iterator()
		for it.Advance() {
			tuple, err := r.ss.Unpack(it.MustGet().Key)
			if err != nil {
				return nil, err
			}

			m := &ClusterMetadataInfo{}
			m.SetMetadataPartition(tuple[len(tuple)-2].([]byte))
			m.SetClusterName(tuple[len(tuple)-1].([]byte))

			m.SetData(rtx.Get(m.GetDataKey(r.ss)).MustGet())
			m.SetDataEncoding(rtx.Get(m.GetDataEncodingKey(r.ss)).MustGet())
			m.SetVersion(rtx.Get(m.GetVersionKey(r.ss)).MustGet())

			result = append(result, m)
		}

		return result, nil
	})

	return res.([]*ClusterMetadataInfo), err
}

func (r *ClusterMetadataInfoRepository) Save(in *ClusterMetadataInfo) error {
	return in.Save(r.tr, r.ss)
}
