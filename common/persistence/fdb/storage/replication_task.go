package storage

import (
	"math/big"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// CREATE TABLE replication_tasks (
//
//	shard_id INTEGER NOT NULL,
//	task_id BIGINT NOT NULL,
//	--
//	data BYTEA NOT NULL,
//	data_encoding VARCHAR(16) NOT NULL,
//	PRIMARY KEY (shard_id, task_id)
//
// );
type ReplicationTask struct {
	ShardID      int64
	TaskID       *big.Int
	Data         []byte
	DataEncoding string
}

func (t *ReplicationTask) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{t.GetShardIDBytes(), t.GetTaskIDBytes()})
}

func (t *ReplicationTask) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_task_id_idx", t.GetShardIDBytes(), t.GetTaskIDBytes()})
}

func (t *ReplicationTask) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (t *ReplicationTask) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(t.ShardID))
	return buff
}

func (t *ReplicationTask) SetShardID(data []byte) {
	t.ShardID = int64(BytesOrder.Uint64(data))
}

func (t *ReplicationTask) GetTaskIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"task_id"})
}

func (t *ReplicationTask) GetTaskIDBytes() []byte {
	return t.TaskID.Bytes()
}

func (t *ReplicationTask) SetTaskID(data []byte) {
	t.TaskID = big.NewInt(0).SetBytes(data)
}

func (t *ReplicationTask) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (t *ReplicationTask) GetDataBytes() []byte {
	return t.Data
}

func (t *ReplicationTask) SetData(data []byte) {
	t.Data = data
}

func (t *ReplicationTask) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (t *ReplicationTask) GetDataEncodingBytes() []byte {
	return []byte(t.DataEncoding)
}

func (t *ReplicationTask) SetDataEncoding(data []byte) {
	t.DataEncoding = string(data)
}

func (t *ReplicationTask) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// ReplicationTasks table row data
		// Note: ShardID and TaskID are part of the
		// primary key of transfer_tasks and as such are present in the prefix of
		// all keys below. Storing a key/value pair for those is not strictly
		// necessary as it wastes space duplicating data.
		tx.Set(t.GetDataKey(ss), t.GetDataBytes())
		tx.Set(t.GetDataEncodingKey(ss), t.GetDataEncodingBytes())

		// Indexes
		tx.Set(t.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type ReplicationTaskRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewReplicationTaskRepository(tr fdb.Transactor, ss subspace.Subspace) *ReplicationTaskRepository {
	return &ReplicationTaskRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *ReplicationTaskRepository) Find(shardID int64, taskID *big.Int) (*ReplicationTask, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &ReplicationTask{
			ShardID: shardID,
			TaskID:  taskID,
		}

		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*ReplicationTask), err
}

// CREATE TABLE replication_tasks_dlq (
//
//	source_cluster_name VARCHAR(255) NOT NULL,
//	shard_id INTEGER NOT NULL,
//	task_id BIGINT NOT NULL,
//	--
//	data BYTEA NOT NULL,
//	data_encoding VARCHAR(16) NOT NULL,
//	PRIMARY KEY (source_cluster_name, shard_id, task_id)
//
// );
type ReplicationTaskSQL struct {
	SourceClusterName string
	ShardID           int64
	TaskID            *big.Int
	Data              []byte
	DataEncoding      string
}

func (t *ReplicationTaskSQL) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{t.GetShardIDBytes(), t.GetTaskIDBytes()})
}

func (t *ReplicationTaskSQL) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"source_cluster_name_shard_id_task_id_idx", t.GetSourceClusterNameBytes(), t.GetShardIDBytes(), t.GetTaskIDBytes()})
}

func (t *ReplicationTaskSQL) GetSourceClusterNameKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"source_cluster_name"})
}

func (t *ReplicationTaskSQL) GetSourceClusterNameBytes() []byte {
	return []byte(t.DataEncoding)
}

func (t *ReplicationTaskSQL) SetSourceClusterName(data []byte) {
	t.DataEncoding = string(data)
}

func (t *ReplicationTaskSQL) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (t *ReplicationTaskSQL) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(t.ShardID))
	return buff
}

func (t *ReplicationTaskSQL) SetShardID(data []byte) {
	t.ShardID = int64(BytesOrder.Uint64(data))
}

func (t *ReplicationTaskSQL) GetTaskIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"task_id"})
}

func (t *ReplicationTaskSQL) GetTaskIDBytes() []byte {
	return t.TaskID.Bytes()
}

func (t *ReplicationTaskSQL) SetTaskID(data []byte) {
	t.TaskID = big.NewInt(0).SetBytes(data)
}

func (t *ReplicationTaskSQL) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (t *ReplicationTaskSQL) GetDataBytes() []byte {
	return t.Data
}

func (t *ReplicationTaskSQL) SetData(data []byte) {
	t.Data = data
}

func (t *ReplicationTaskSQL) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (t *ReplicationTaskSQL) GetDataEncodingBytes() []byte {
	return []byte(t.DataEncoding)
}

func (t *ReplicationTaskSQL) SetDataEncoding(data []byte) {
	t.DataEncoding = string(data)
}

func (t *ReplicationTaskSQL) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// ReplicationTaskSQLs table row data
		// Note: ShardID and TaskID are part of the
		// primary key of transfer_tasks and as such are present in the prefix of
		// all keys below. Storing a key/value pair for those is not strictly
		// necessary as it wastes space duplicating data.
		tx.Set(t.GetDataKey(ss), t.GetDataBytes())
		tx.Set(t.GetDataEncodingKey(ss), t.GetDataEncodingBytes())

		// Indexes
		tx.Set(t.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type ReplicationTaskSQLRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewReplicationTaskSQLRepository(tr fdb.Transactor, ss subspace.Subspace) *ReplicationTaskSQLRepository {
	return &ReplicationTaskSQLRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *ReplicationTaskSQLRepository) Find(sourceClusterName string, shardID int64, taskID *big.Int) (*ReplicationTaskSQL, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &ReplicationTaskSQL{
			SourceClusterName: sourceClusterName,
			ShardID:           shardID,
			TaskID:            taskID,
		}

		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*ReplicationTaskSQL), err
}
