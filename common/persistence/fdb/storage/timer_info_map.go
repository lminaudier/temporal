package storage

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// CREATE TABLE timer_info_maps (
//
//	shard_id INTEGER NOT NULL,
//	namespace_id BYTEA NOT NULL,
//	workflow_id VARCHAR(255) NOT NULL,
//	run_id BYTEA NOT NULL,
//	timer_id VARCHAR(255) NOT NULL,
//
// --
//
//	data BYTEA NOT NULL,
//	data_encoding VARCHAR(16),
//	PRIMARY KEY (shard_id, namespace_id, workflow_id, run_id, timer_id)
//
// );
type TimerInfoMap struct {
	ShardID      int64
	NamespaceID  []byte
	WorkflowID   string
	RunID        []byte
	TimerID      string
	Data         []byte
	DataEncoding string
}

func (m *TimerInfoMap) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{m.GetShardIDBytes(), m.GetNamespaceIDBytes(), m.GetWorkflowIDBytes(), m.GetRunIDBytes(), m.GetTimerIDBytes()})
}

func (m *TimerInfoMap) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_namespace_id_workflow_id_run_id_timer_id_idx", m.GetShardIDBytes(), m.GetNamespaceIDBytes(), m.GetWorkflowIDBytes(), m.GetRunIDBytes(), m.GetTimerIDBytes()})
}

func (m *TimerInfoMap) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (m *TimerInfoMap) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(m.ShardID))
	return buff
}

func (m *TimerInfoMap) SetShardID(data []byte) {
	m.ShardID = int64(BytesOrder.Uint64(data))
}

func (m *TimerInfoMap) GetWorkflowIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"workflow_id"})
}

func (m *TimerInfoMap) GetWorkflowIDBytes() []byte {
	return []byte(m.WorkflowID)
}

func (m *TimerInfoMap) SetWorkflowID(data []byte) {
	m.WorkflowID = string(data)
}

func (m *TimerInfoMap) GetNamespaceIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"namespace_id"})
}

func (m *TimerInfoMap) GetNamespaceIDBytes() []byte {
	return m.NamespaceID
}

func (m *TimerInfoMap) SetNamespaceID(data []byte) {
	m.NamespaceID = data
}

func (m *TimerInfoMap) GetRunIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"run_id"})
}

func (m *TimerInfoMap) GetRunIDBytes() []byte {
	return m.RunID
}

func (m *TimerInfoMap) SetRunID(data []byte) {
	m.RunID = data
}

func (m *TimerInfoMap) GetTimerIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"timer_id"})
}

func (m *TimerInfoMap) GetTimerIDBytes() []byte {
	return []byte(m.TimerID)
}

func (m *TimerInfoMap) SetTimerID(data []byte) {
	m.TimerID = string(m.TimerID)
}

func (m *TimerInfoMap) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (m *TimerInfoMap) GetDataBytes() []byte {
	return m.Data
}

func (m *TimerInfoMap) SetData(data []byte) {
	m.Data = data
}

func (m *TimerInfoMap) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (m *TimerInfoMap) GetDataEncodingBytes() []byte {
	return []byte(m.DataEncoding)
}

func (m *TimerInfoMap) SetDataEncoding(data []byte) {
	m.DataEncoding = string(data)
}

func (m *TimerInfoMap) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// TimerInfoMap table row data
		// Note: ShardID, NamespaceID, WorkflowID, RunID and TimerID are part of the
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

type TimerInfoMapRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewTimerInfoMapRepository(tr fdb.Transactor, ss subspace.Subspace) *TimerInfoMapRepository {
	return &TimerInfoMapRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *TimerInfoMapRepository) Find(shardID int64, namespaceID []byte, workflowID string, runID []byte, timerID string) (*TimerInfoMap, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &TimerInfoMap{
			ShardID:     shardID,
			NamespaceID: namespaceID,
			WorkflowID:  workflowID,
			RunID:       runID,
			TimerID:     timerID,
		}

		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*TimerInfoMap), err
}
