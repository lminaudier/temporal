package storage

import (
	"math/big"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

type BufferedEvent struct {
	ShardID     int64
	NamespaceID []byte
	WorkflowID  string
	RunID       []byte
	// TODO: check fdb versionstamp feature https://apple.github.io/foundationdb/data-modeling.html#versionstamps
	ID           *big.Int
	Data         []byte
	DataEncoding string
}

func (e *BufferedEvent) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{e.GetShardIDBytes(), e.GetNamespaceIDBytes(), e.GetWorkflowIDBytes(), e.GetRunIDBytes(), e.GetIDBytes()})
}

func (e *BufferedEvent) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_namespace_id_workflow_id_run_id_id_idx", e.GetShardIDBytes(), e.GetNamespaceIDBytes(), e.GetWorkflowIDBytes(), e.GetRunIDBytes(), e.GetIDBytes()})
}

func (e *BufferedEvent) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (e *BufferedEvent) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(e.ShardID))
	return buff
}

func (e *BufferedEvent) SetShardID(data []byte) {
	e.ShardID = int64(BytesOrder.Uint64(data))
}

func (e *BufferedEvent) GetNamespaceIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"namespace_id"})
}

func (e *BufferedEvent) GetNamespaceIDBytes() []byte {
	return e.NamespaceID
}

func (e *BufferedEvent) SetNamespaceID(data []byte) {
	e.NamespaceID = data
}

func (e *BufferedEvent) GetWorkflowIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"workflow_id"})
}

func (e *BufferedEvent) GetWorkflowIDBytes() []byte {
	return []byte(e.WorkflowID)
}

func (e *BufferedEvent) SetWorkflowID(data []byte) {
	e.WorkflowID = string(data)
}

func (e *BufferedEvent) GetRunIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"run_id"})
}

func (e *BufferedEvent) GetRunIDBytes() []byte {
	return e.RunID
}

func (e *BufferedEvent) SetRunID(data []byte) {
	e.RunID = data
}

func (e *BufferedEvent) GetIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"id"})
}

func (e *BufferedEvent) GetIDBytes() []byte {
	return e.ID.Bytes()
}

func (e *BufferedEvent) SetID(data []byte) {
	e.ID = big.NewInt(0).SetBytes(data)
}

func (e *BufferedEvent) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (e *BufferedEvent) GetDataBytes() []byte {
	return e.Data
}

func (e *BufferedEvent) SetData(data []byte) {
	e.Data = data
}

func (e *BufferedEvent) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (e *BufferedEvent) GetDataEncodingBytes() []byte {
	return []byte(e.DataEncoding)
}

func (e *BufferedEvent) SetDataEncoding(data []byte) {
	e.DataEncoding = string(data)
}

func (e *BufferedEvent) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// BufferedEvents table row data
		// Note: ShardID, NamespaceID, WorkflowID, RunID and ID are part of the
		// primary key of namespaces and as such are present in the prefix of
		// all keys below. Storing a key/value pair for those is not strictly
		// necessary as it wastes space duplicating data.
		tx.Set(e.GetDataKey(ss), e.GetDataBytes())
		tx.Set(e.GetDataEncodingKey(ss), e.GetDataEncodingBytes())

		// Indexes
		tx.Set(e.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type BufferedEventRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewBufferedEventRepository(tr fdb.Transactor, ss subspace.Subspace) *BufferedEventRepository {
	return &BufferedEventRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *BufferedEventRepository) Find(shardID int64, namespaceID []byte, workflowID string, runID []byte, id *big.Int) (*BufferedEvent, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &BufferedEvent{
			ShardID:     shardID,
			NamespaceID: namespaceID,
			WorkflowID:  workflowID,
			RunID:       runID,
			ID:          id,
		}

		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*BufferedEvent), err
}
