package storage

import (
	"math/big"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

type Execution struct {
	ShardID          int64
	NamespaceID      []byte
	WorkflowID       string
	RunID            []byte
	NextEventID      *big.Int
	LastWriteVersion *big.Int
	Data             []byte
	DataEncoding     string
	State            []byte
	StateEncoding    string
	DBRecordVersion  *big.Int
}

func (e *Execution) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{e.GetShardIDBytes(), e.GetNamespaceIDBytes(), e.GetWorkflowIDBytes(), e.GetRunIDBytes()})
}

func (e *Execution) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_namespace_id_workflow_id_run_id_idx", e.GetShardIDBytes(), e.GetNamespaceIDBytes(), e.GetWorkflowIDBytes(), e.GetRunIDBytes()})
}

func (e *Execution) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (e *Execution) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(e.ShardID))
	return buff
}

func (e *Execution) SetShardID(data []byte) {
	e.ShardID = int64(BytesOrder.Uint64(data))
}

func (e *Execution) GetNamespaceIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"namespace_id"})
}

func (e *Execution) GetNamespaceIDBytes() []byte {
	return e.NamespaceID
}

func (e *Execution) SetNamespaceID(data []byte) {
	e.NamespaceID = data
}

func (e *Execution) GetWorkflowIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"workflow_id"})
}

func (e *Execution) GetWorkflowIDBytes() []byte {
	return []byte(e.WorkflowID)
}

func (e *Execution) SetWorkflowID(data []byte) {
	e.WorkflowID = string(data)
}

func (e *Execution) GetRunIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"run_id"})
}

func (e *Execution) GetRunIDBytes() []byte {
	return e.RunID
}

func (e *Execution) SetRunID(data []byte) {
	e.RunID = data
}

func (e *Execution) GetNextEventIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"next_event_id"})
}

func (e *Execution) GetNextEventIDBytes() []byte {
	return e.NextEventID.Bytes()
}

func (e *Execution) SetNextEventID(data []byte) {
	e.NextEventID = big.NewInt(0).SetBytes(data)
}

func (e *Execution) GetLastWriteVersionKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"last_write_version"})
}

func (e *Execution) GetLastWriteVersionBytes() []byte {
	return e.LastWriteVersion.Bytes()
}

func (e *Execution) SetLastWriteVersion(data []byte) {
	e.LastWriteVersion = big.NewInt(0).SetBytes(data)
}

func (e *Execution) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (e *Execution) GetDataBytes() []byte {
	return e.Data
}

func (e *Execution) SetData(data []byte) {
	e.Data = data
}

func (e *Execution) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (e *Execution) GetDataEncodingBytes() []byte {
	return []byte(e.DataEncoding)
}

func (e *Execution) SetDataEncoding(data []byte) {
	e.DataEncoding = string(data)
}

func (e *Execution) GetStateKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"state"})
}

func (e *Execution) GetStateBytes() []byte {
	return e.State
}

func (e *Execution) SetState(data []byte) {
	e.State = data
}

func (e *Execution) GetStateEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"state_encoding"})
}

func (e *Execution) GetStateEncodingBytes() []byte {
	return []byte(e.StateEncoding)
}

func (e *Execution) SetStateEncoding(data []byte) {
	e.StateEncoding = string(data)
}

func (e *Execution) GetDBRecordVersionKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"db_record_version"})
}

func (e *Execution) GetDBRecordVersionBytes() []byte {
	return e.DBRecordVersion.Bytes()
}

func (e *Execution) SetDBRecordVersion(data []byte) {
	e.DBRecordVersion = big.NewInt(0).SetBytes(data)
}

func (e *Execution) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// Executions table row data
		// Note: ShardID, NamespaceID, WorkflowID and RunID are part of the primary key of namespaces
		// and as such are present in the prefix of all keys below. Storing a
		// key/value pair for those is not strictly necessary as it wastes
		// space duplicating data.
		tx.Set(e.GetNextEventIDKey(ss), e.GetNextEventIDBytes())
		tx.Set(e.GetLastWriteVersionKey(ss), e.GetLastWriteVersionBytes())
		tx.Set(e.GetDataKey(ss), e.GetDataBytes())
		tx.Set(e.GetDataEncodingKey(ss), e.GetDataEncodingBytes())
		tx.Set(e.GetStateKey(ss), e.GetStateBytes())
		tx.Set(e.GetStateEncodingKey(ss), e.GetStateEncodingBytes())
		tx.Set(e.GetDBRecordVersionKey(ss), e.GetDBRecordVersionBytes())

		// Indexes
		tx.Set(e.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type ExecutionRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewExecutionRepository(tr fdb.Transactor, ss subspace.Subspace) *ExecutionRepository {
	return &ExecutionRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *ExecutionRepository) Find(shardID int64, namespaceID []byte, workflowID string, runID []byte) (*Execution, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &Execution{
			ShardID:     shardID,
			NamespaceID: namespaceID,
			WorkflowID:  workflowID,
			RunID:       runID,
		}

		e.SetNextEventID(rtx.Get(e.GetNextEventIDKey(r.ss)).MustGet())
		e.SetLastWriteVersion(rtx.Get(e.GetLastWriteVersionKey(r.ss)).MustGet())
		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())
		e.SetState(rtx.Get(e.GetStateKey(r.ss)).MustGet())
		e.SetStateEncoding(rtx.Get(e.GetStateEncodingKey(r.ss)).MustGet())
		e.SetDBRecordVersion(rtx.Get(e.GetDBRecordVersionKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*Execution), err
}

// CREATE TABLE current_executions(
//
//	shard_id INTEGER NOT NULL,
//	namespace_id BYTEA NOT NULL,
//	workflow_id VARCHAR(255) NOT NULL,
//	--
//	run_id BYTEA NOT NULL,
//	create_request_id VARCHAR(255) NOT NULL,
//	state INTEGER NOT NULL,
//	status INTEGER NOT NULL,
//	start_version BIGINT NOT NULL DEFAULT 0,
//	last_write_version BIGINT NOT NULL,
//	PRIMARY KEY (shard_id, namespace_id, workflow_id)
//
// );
type CurrentExecution struct {
	ShardID          int64
	NamespaceID      []byte
	WorkflowID       string
	RunID            []byte
	CreateRequestID  string
	State            int64
	Status           int64
	StartVersion     *big.Int
	LastWriteVersion *big.Int
}

func (e *CurrentExecution) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{e.GetShardIDBytes(), e.GetNamespaceIDBytes(), e.GetWorkflowIDBytes()})
}

func (e *CurrentExecution) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_namespace_id_workflow_id_idx", e.GetShardIDBytes(), e.GetNamespaceIDBytes(), e.GetWorkflowIDBytes()})
}

func (e *CurrentExecution) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (e *CurrentExecution) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(e.ShardID))
	return buff
}

func (e *CurrentExecution) SetShardID(data []byte) {
	e.ShardID = int64(BytesOrder.Uint64(data))
}

func (e *CurrentExecution) GetNamespaceIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"namespace_id"})
}

func (e *CurrentExecution) GetNamespaceIDBytes() []byte {
	return e.NamespaceID
}

func (e *CurrentExecution) SetNamespaceID(data []byte) {
	e.NamespaceID = data
}

func (e *CurrentExecution) GetWorkflowIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"workflow_id"})
}

func (e *CurrentExecution) GetWorkflowIDBytes() []byte {
	return []byte(e.WorkflowID)
}

func (e *CurrentExecution) SetWorkflowID(data []byte) {
	e.WorkflowID = string(data)
}

func (e *CurrentExecution) GetRunIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"run_id"})
}

func (e *CurrentExecution) GetRunIDBytes() []byte {
	return e.RunID
}

func (e *CurrentExecution) SetRunID(data []byte) {
	e.RunID = data
}

func (e *CurrentExecution) GetCreateRequestIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"create_request_id"})
}

func (e *CurrentExecution) GetCreateRequestIDBytes() []byte {
	return []byte(e.CreateRequestID)
}

func (e *CurrentExecution) SetCreateRequestID(data []byte) {
	e.CreateRequestID = string(data)
}

func (e *CurrentExecution) GetStateKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"state"})
}

func (e *CurrentExecution) GetStateBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(e.State))
	return buff
}

func (e *CurrentExecution) SetState(data []byte) {
	e.State = int64(BytesOrder.Uint64(data))
}

func (e *CurrentExecution) GetStatusKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"status"})
}

func (e *CurrentExecution) GetStatusBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(e.Status))
	return buff
}

func (e *CurrentExecution) SetStatus(data []byte) {
	e.Status = int64(BytesOrder.Uint64(data))
}

func (e *CurrentExecution) GetStartVersionKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"start_version"})
}

func (e *CurrentExecution) GetStartVersionBytes() []byte {
	return e.StartVersion.Bytes()
}

func (e *CurrentExecution) SetStartVersion(data []byte) {
	e.StartVersion = big.NewInt(0).SetBytes(data)
}

func (e *CurrentExecution) GetLastWriteVersionKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"last_write_version"})
}

func (e *CurrentExecution) GetLastWriteVersionBytes() []byte {
	return e.LastWriteVersion.Bytes()
}

func (e *CurrentExecution) SetLastWriteVersion(data []byte) {
	e.LastWriteVersion = big.NewInt(0).SetBytes(data)
}

func (e *CurrentExecution) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// CurrentExecutions table row data
		// Note: ShardID, NamespaceID and WorkflowID are part of the primary
		// key of namespaces and as such are present in the prefix of all keys
		// below. Storing a key/value pair for those is not strictly necessary
		// as it wastes space duplicating data.
		tx.Set(e.GetRunIDKey(ss), e.GetRunIDBytes())
		tx.Set(e.GetCreateRequestIDKey(ss), e.GetCreateRequestIDBytes())
		tx.Set(e.GetStateKey(ss), e.GetStateBytes())
		tx.Set(e.GetStatusKey(ss), e.GetStatusBytes())
		tx.Set(e.GetStartVersionKey(ss), e.GetStartVersionBytes())
		tx.Set(e.GetLastWriteVersionKey(ss), e.GetLastWriteVersionBytes())

		// Indexes
		tx.Set(e.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type CurrentExecutionRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewCurrentExecutionRepository(tr fdb.Transactor, ss subspace.Subspace) *CurrentExecutionRepository {
	return &CurrentExecutionRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *CurrentExecutionRepository) Find(shardID int64, namespaceID []byte, workflowID string) (*CurrentExecution, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &CurrentExecution{
			ShardID:     shardID,
			NamespaceID: namespaceID,
			WorkflowID:  workflowID,
		}

		e.SetRunID(rtx.Get(e.GetRunIDKey(r.ss)).MustGet())
		e.SetCreateRequestID(rtx.Get(e.GetCreateRequestIDKey(r.ss)).MustGet())
		e.SetState(rtx.Get(e.GetStateKey(r.ss)).MustGet())
		e.SetStatus(rtx.Get(e.GetStatusKey(r.ss)).MustGet())
		e.SetStartVersion(rtx.Get(e.GetStartVersionKey(r.ss)).MustGet())
		e.SetLastWriteVersion(rtx.Get(e.GetLastWriteVersionKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*CurrentExecution), err
}
