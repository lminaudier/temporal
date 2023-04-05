package storage

import (
	"math/big"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// CREATE TABLE signal_info_maps (
//
//	shard_id INTEGER NOT NULL,
//	namespace_id BYTEA NOT NULL,
//	workflow_id VARCHAR(255) NOT NULL,
//	run_id BYTEA NOT NULL,
//	initiated_id BIGINT NOT NULL,
//
// --
//
//	data BYTEA NOT NULL,
//	data_encoding VARCHAR(16),
//	PRIMARY KEY (shard_id, namespace_id, workflow_id, run_id, initiated_id)
//
// );
type SignalInfoMap struct {
	ShardID      int64
	NamespaceID  []byte
	WorkflowID   string
	RunID        []byte
	InitiatedID  *big.Int
	Data         []byte
	DataEncoding string
}

func (m *SignalInfoMap) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{m.GetShardIDBytes(), m.GetNamespaceIDBytes(), m.GetWorkflowIDBytes(), m.GetRunIDBytes(), m.GetScheduleIDBytes()})
}

func (m *SignalInfoMap) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_namespace_id_workflow_id_run_id_schedule_id_idx", m.GetShardIDBytes(), m.GetNamespaceIDBytes(), m.GetWorkflowIDBytes(), m.GetRunIDBytes(), m.GetScheduleIDBytes()})
}

func (m *SignalInfoMap) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (m *SignalInfoMap) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(m.ShardID))
	return buff
}

func (m *SignalInfoMap) SetShardID(data []byte) {
	m.ShardID = int64(BytesOrder.Uint64(data))
}

func (m *SignalInfoMap) GetWorkflowIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"workflow_id"})
}

func (m *SignalInfoMap) GetWorkflowIDBytes() []byte {
	return []byte(m.WorkflowID)
}

func (m *SignalInfoMap) SetWorkflowID(data []byte) {
	m.WorkflowID = string(data)
}

func (m *SignalInfoMap) GetNamespaceIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"namespace_id"})
}

func (m *SignalInfoMap) GetNamespaceIDBytes() []byte {
	return m.NamespaceID
}

func (m *SignalInfoMap) SetNamespaceID(data []byte) {
	m.NamespaceID = data
}

func (m *SignalInfoMap) GetRunIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"run_id"})
}

func (m *SignalInfoMap) GetRunIDBytes() []byte {
	return m.RunID
}

func (m *SignalInfoMap) SetRunID(data []byte) {
	m.RunID = data
}

func (m *SignalInfoMap) GetScheduleIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"initiated_id"})
}

func (m *SignalInfoMap) GetScheduleIDBytes() []byte {
	return m.InitiatedID.Bytes()
}

func (m *SignalInfoMap) SetScheduleID(data []byte) {
	m.InitiatedID = big.NewInt(0).SetBytes(data)
}

func (m *SignalInfoMap) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (m *SignalInfoMap) GetDataBytes() []byte {
	return m.Data
}

func (m *SignalInfoMap) SetData(data []byte) {
	m.Data = data
}

func (m *SignalInfoMap) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (m *SignalInfoMap) GetDataEncodingBytes() []byte {
	return []byte(m.DataEncoding)
}

func (m *SignalInfoMap) SetDataEncoding(data []byte) {
	m.DataEncoding = string(data)
}

func (m *SignalInfoMap) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// SignalInfoMap table row data
		// Note: ShardID, NamespaceID, WorkflowID, RunID and ScheduleID are part of the
		// primary key of namespaces and as such are present in the prefix of
		// all keys below. Storing a key/value pair for those is not strictly
		// necessary as it wastes space duplicating data.
		tx.Set(m.GetDataKey(ss), m.GetDataBytes())
		tx.Set(m.GetDataEncodingKey(ss), m.GetDataEncodingBytes())

		// Indexes
		tx.Set(m.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type SignalInfoMapRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewSignalInfoMapRepository(tr fdb.Transactor, ss subspace.Subspace) *SignalInfoMapRepository {
	return &SignalInfoMapRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *SignalInfoMapRepository) Find(shardID int64, namespaceID []byte, workflowID string, runID []byte, initiatedID *big.Int) (*SignalInfoMap, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &SignalInfoMap{
			ShardID:     shardID,
			NamespaceID: namespaceID,
			WorkflowID:  workflowID,
			RunID:       runID,
			InitiatedID: initiatedID,
		}

		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*SignalInfoMap), err
}
