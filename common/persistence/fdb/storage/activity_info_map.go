package storage

import (
	"math/big"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// CREATE TABLE activity_info_maps (
// -- each row corresponds to one key of one map<string, ActivityInfo>
//
//	shard_id INTEGER NOT NULL,
//	namespace_id BYTEA NOT NULL,
//	workflow_id VARCHAR(255) NOT NULL,
//	run_id BYTEA NOT NULL,
//	schedule_id BIGINT NOT NULL,
//
// --
//
//	data BYTEA NOT NULL,
//	data_encoding VARCHAR(16),
//	PRIMARY KEY (shard_id, namespace_id, workflow_id, run_id, schedule_id)
//
// );
type ActivityInfoMap struct {
	ShardID      int64
	NamespaceID  []byte
	WorkflowID   string
	RunID        []byte
	ScheduleID   *big.Int
	Data         []byte
	DataEncoding string
}

func (m *ActivityInfoMap) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{m.GetShardIDBytes(), m.GetNamespaceIDBytes(), m.GetWorkflowIDBytes(), m.GetRunIDBytes(), m.GetScheduleIDBytes()})
}

func (m *ActivityInfoMap) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_namespace_id_workflow_id_run_id_schedule_id_idx", m.GetShardIDBytes(), m.GetNamespaceIDBytes(), m.GetWorkflowIDBytes(), m.GetRunIDBytes(), m.GetScheduleIDBytes()})
}

func (m *ActivityInfoMap) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (m *ActivityInfoMap) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(m.ShardID))
	return buff
}

func (m *ActivityInfoMap) SetShardID(data []byte) {
	m.ShardID = int64(BytesOrder.Uint64(data))
}

func (m *ActivityInfoMap) GetWorkflowIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"workflow_id"})
}

func (m *ActivityInfoMap) GetWorkflowIDBytes() []byte {
	return []byte(m.WorkflowID)
}

func (m *ActivityInfoMap) SetWorkflowID(data []byte) {
	m.WorkflowID = string(data)
}

func (m *ActivityInfoMap) GetNamespaceIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"namespace_id"})
}

func (m *ActivityInfoMap) GetNamespaceIDBytes() []byte {
	return m.NamespaceID
}

func (m *ActivityInfoMap) SetNamespaceID(data []byte) {
	m.NamespaceID = data
}

func (m *ActivityInfoMap) GetRunIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"run_id"})
}

func (m *ActivityInfoMap) GetRunIDBytes() []byte {
	return m.RunID
}

func (m *ActivityInfoMap) SetRunID(data []byte) {
	m.RunID = data
}

func (m *ActivityInfoMap) GetScheduleIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"schedule_id"})
}

func (m *ActivityInfoMap) GetScheduleIDBytes() []byte {
	return m.ScheduleID.Bytes()
}

func (m *ActivityInfoMap) SetScheduleID(data []byte) {
	m.ScheduleID = big.NewInt(0).SetBytes(data)
}

func (m *ActivityInfoMap) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (m *ActivityInfoMap) GetDataBytes() []byte {
	return m.Data
}

func (m *ActivityInfoMap) SetData(data []byte) {
	m.Data = data
}

func (m *ActivityInfoMap) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (m *ActivityInfoMap) GetDataEncodingBytes() []byte {
	return []byte(m.DataEncoding)
}

func (m *ActivityInfoMap) SetDataEncoding(data []byte) {
	m.DataEncoding = string(data)
}

func (m *ActivityInfoMap) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// ActivityInfoMap table row data
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

type ActivityInfoMapRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewActivityInfoMapRepository(tr fdb.Transactor, ss subspace.Subspace) *ActivityInfoMapRepository {
	return &ActivityInfoMapRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *ActivityInfoMapRepository) Find(shardID int64, namespaceID []byte, workflowID string, runID []byte, scheduleID *big.Int) (*ActivityInfoMap, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &ActivityInfoMap{
			ShardID:     shardID,
			NamespaceID: namespaceID,
			WorkflowID:  workflowID,
			RunID:       runID,
			ScheduleID:  scheduleID,
		}

		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*ActivityInfoMap), err
}
