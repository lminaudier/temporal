package storage

import (
	"math/big"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// CREATE TABLE queue_metadata (
//
//	queue_type     INTEGER NOT NULL,
//	data BYTEA     NOT NULL,
//	data_encoding  VARCHAR(16) NOT NULL,
//	version        BIGINT NOT NULL,
//	PRIMARY KEY(queue_type)
//
// );
type QueueMetadata struct {
	QueueType    int64
	Data         []byte
	DataEncoding string
	Version      *big.Int
}

func (m *QueueMetadata) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{m.GetQueueTypeBytes()})
}

func (m *QueueMetadata) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"queue_type_idx", m.GetQueueTypeBytes()})
}

func (m *QueueMetadata) GetQueueTypeKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"queue_type"})
}

func (m *QueueMetadata) GetQueueTypeBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(m.QueueType))
	return buff
}

func (m *QueueMetadata) SetQueueType(data []byte) {
	m.QueueType = int64(BytesOrder.Uint64(data))
}

func (m *QueueMetadata) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (m *QueueMetadata) GetDataBytes() []byte {
	return m.Data
}

func (m *QueueMetadata) SetData(data []byte) {
	m.Data = data
}

func (m *QueueMetadata) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (m *QueueMetadata) GetDataEncodingBytes() []byte {
	return []byte(m.DataEncoding)
}

func (m *QueueMetadata) SetDataEncoding(data []byte) {
	m.DataEncoding = string(data)
}

func (m *QueueMetadata) GetVersionKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"version"})
}

func (m *QueueMetadata) GetVersionBytes() []byte {
	return m.Version.Bytes()
}

func (m *QueueMetadata) SetVersion(data []byte) {
	m.Version = big.NewInt(0).SetBytes(data)
}

func (m *QueueMetadata) Save(tr fdb.Transactor, ss subspace.Subspace) error {
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

type QueueMetadataRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewQueueMetadataRepository(tr fdb.Transactor, ss subspace.Subspace) *QueueMetadataRepository {
	return &QueueMetadataRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *QueueMetadataRepository) Find(queueType int64) (*QueueMetadata, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &QueueMetadata{
			QueueType: queueType,
		}

		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())
		e.SetVersion(rtx.Get(e.GetVersionKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*QueueMetadata), err
}

func (r *QueueMetadataRepository) Save(in *QueueMetadata) error {
	return in.Save(r.tr, r.ss)
}
