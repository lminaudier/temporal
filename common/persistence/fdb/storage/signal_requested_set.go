package storage

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// CREATE TABLE signals_requested_sets (
//
//	shard_id INTEGER NOT NULL,
//	namespace_id BYTEA NOT NULL,
//	workflow_id VARCHAR(255) NOT NULL,
//	run_id BYTEA NOT NULL,
//	signal_id VARCHAR(255) NOT NULL,
//	--
//	PRIMARY KEY (shard_id, namespace_id, workflow_id, run_id, signal_id)
//
// );
type SignalRequestedSet struct {
	ShardID     int64
	NamespaceID []byte
	WorkflowID  string
	RunID       []byte
	SignalID    string
}

func (m *SignalRequestedSet) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{m.GetShardIDBytes(), m.GetNamespaceIDBytes(), m.GetWorkflowIDBytes(), m.GetRunIDBytes(), m.GetSignalIDBytes()})
}

func (m *SignalRequestedSet) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_namespace_id_workflow_id_run_id_signal_id_idx", m.GetShardIDBytes(), m.GetNamespaceIDBytes(), m.GetWorkflowIDBytes(), m.GetRunIDBytes(), m.GetSignalIDBytes()})
}

func (m *SignalRequestedSet) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (m *SignalRequestedSet) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(m.ShardID))
	return buff
}

func (m *SignalRequestedSet) SetShardID(data []byte) {
	m.ShardID = int64(BytesOrder.Uint64(data))
}

func (m *SignalRequestedSet) GetWorkflowIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"workflow_id"})
}

func (m *SignalRequestedSet) GetWorkflowIDBytes() []byte {
	return []byte(m.WorkflowID)
}

func (m *SignalRequestedSet) SetWorkflowID(data []byte) {
	m.WorkflowID = string(data)
}

func (m *SignalRequestedSet) GetNamespaceIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"namespace_id"})
}

func (m *SignalRequestedSet) GetNamespaceIDBytes() []byte {
	return m.NamespaceID
}

func (m *SignalRequestedSet) SetNamespaceID(data []byte) {
	m.NamespaceID = data
}

func (m *SignalRequestedSet) GetRunIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"run_id"})
}

func (m *SignalRequestedSet) GetRunIDBytes() []byte {
	return m.RunID
}

func (m *SignalRequestedSet) SetRunID(data []byte) {
	m.RunID = data
}

func (m *SignalRequestedSet) GetSignalIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"signal_id"})
}

func (m *SignalRequestedSet) GetSignalIDBytes() []byte {
	return []byte(m.SignalID)
}

func (m *SignalRequestedSet) SetSignalID(data []byte) {
	m.SignalID = string(m.SignalID)
}

func (m *SignalRequestedSet) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (m *SignalRequestedSet) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// SignalRequestedSet table is here only to guarantee unicity
		// TODO: expand why we only create the index

		// Indexes
		tx.Set(m.PrimaryKeyIndexKey(ss), m.PrimaryKeyIndexKey(ss).FDBKey())

		return nil, nil
	})

	return nil
}

type SignalRequestedSetRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewSignalRequestedSetRepository(tr fdb.Transactor, ss subspace.Subspace) *SignalRequestedSetRepository {
	return &SignalRequestedSetRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *SignalRequestedSetRepository) Find(shardID int64, namespaceID []byte, workflowID string, runID []byte, signalID string) (*SignalRequestedSet, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &SignalRequestedSet{
			ShardID:     shardID,
			NamespaceID: namespaceID,
			WorkflowID:  workflowID,
			RunID:       runID,
			SignalID:    signalID,
		}

		// We check if the key is indexed
		_, err := rtx.Get(e.PrimaryKeyIndexKey(r.ss)).Get()
		if err != nil {
			return nil, err
		}

		return e, nil
	})

	return res.(*SignalRequestedSet), err
}
