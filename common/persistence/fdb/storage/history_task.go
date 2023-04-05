package storage

import (
	"math/big"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// CREATE TABLE history_immediate_tasks(
//
//	shard_id INTEGER NOT NULL,
//	category_id INTEGER NOT NULL,
//	task_id BIGINT NOT NULL,
//	--
//	data BYTEA NOT NULL,
//	data_encoding VARCHAR(16) NOT NULL,
//	PRIMARY KEY (shard_id, category_id, task_id)
//
// );
type HistoryImmediateTask struct {
	ShardID      int64
	CategoryID   int64
	TaskID       *big.Int
	Data         []byte
	DataEncoding string
}

func (t *HistoryImmediateTask) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{t.GetShardIDBytes(), t.GetCategoryIDBytes(), t.GetTaskIDBytes()})
}

func (t *HistoryImmediateTask) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_category_id_task_id_idx", t.GetShardIDBytes(), t.GetCategoryIDBytes(), t.GetTaskIDBytes()})
}

func (t *HistoryImmediateTask) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (t *HistoryImmediateTask) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(t.ShardID))
	return buff
}

func (t *HistoryImmediateTask) SetShardID(data []byte) {
	t.ShardID = int64(BytesOrder.Uint64(data))
}

func (t *HistoryImmediateTask) GetCategoryIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"category_id"})
}

func (t *HistoryImmediateTask) GetCategoryIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(t.CategoryID))
	return buff
}

func (t *HistoryImmediateTask) SetCategoryID(data []byte) {
	t.CategoryID = int64(BytesOrder.Uint64(data))
}

func (t *HistoryImmediateTask) GetTaskIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"task_id"})
}

func (t *HistoryImmediateTask) GetTaskIDBytes() []byte {
	return t.TaskID.Bytes()
}

func (t *HistoryImmediateTask) SetTaskID(data []byte) {
	t.TaskID = big.NewInt(0).SetBytes(data)
}

func (t *HistoryImmediateTask) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (t *HistoryImmediateTask) GetDataBytes() []byte {
	return t.Data
}

func (t *HistoryImmediateTask) SetData(data []byte) {
	t.Data = data
}

func (t *HistoryImmediateTask) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (t *HistoryImmediateTask) GetDataEncodingBytes() []byte {
	return []byte(t.DataEncoding)
}

func (t *HistoryImmediateTask) SetDataEncoding(data []byte) {
	t.DataEncoding = string(data)
}

func (t *HistoryImmediateTask) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// HistoryImmediateTasks table row data
		// Note: ShardID, CategoryID and TaskID are part of the
		// primary key of history_immediate_tasks and as such are present in the prefix of
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

type HistoryImmediateTaskRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewHistoryImmediateTaskRepository(tr fdb.Transactor, ss subspace.Subspace) *HistoryImmediateTaskRepository {
	return &HistoryImmediateTaskRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *HistoryImmediateTaskRepository) Find(shardID int64, categoryID int64, taskID *big.Int) (*HistoryImmediateTask, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &HistoryImmediateTask{
			ShardID:    shardID,
			CategoryID: categoryID,
			TaskID:     taskID,
		}

		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*HistoryImmediateTask), err
}

type HistoryScheduledTask struct {
	ShardID             int64
	CategoryID          int64
	VisibilityTimestamp time.Time
	TaskID              *big.Int
	Data                []byte
	DataEncoding        string
}

func (t *HistoryScheduledTask) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{t.GetShardIDBytes(), t.GetCategoryIDBytes(), t.GetVisibilityTimestampBytes(), t.GetTaskIDBytes()})
}

func (t *HistoryScheduledTask) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_category_id_visibility_timestamp_task_id_idx", t.GetShardIDBytes(), t.GetCategoryIDBytes(), t.GetVisibilityTimestampBytes(), t.GetTaskIDBytes()})
}

func (t *HistoryScheduledTask) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (t *HistoryScheduledTask) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(t.ShardID))
	return buff
}

func (t *HistoryScheduledTask) SetShardID(data []byte) {
	t.ShardID = int64(BytesOrder.Uint64(data))
}

func (t *HistoryScheduledTask) GetCategoryIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"category_id"})
}

func (t *HistoryScheduledTask) GetCategoryIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(t.CategoryID))
	return buff
}

func (t *HistoryScheduledTask) SetCategoryID(data []byte) {
	t.CategoryID = int64(BytesOrder.Uint64(data))
}

func (t *HistoryScheduledTask) GetVisibilityTimestampKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"visibility_timestamp"})
}

func (t *HistoryScheduledTask) GetVisibilityTimestampBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(t.VisibilityTimestamp.UnixNano()))
	return buff
}

func (t *HistoryScheduledTask) SetVisibilityTimestamp(data []byte) {
	t.VisibilityTimestamp = time.Unix(0, int64(BytesOrder.Uint64(data)))
}

func (t *HistoryScheduledTask) GetTaskIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"task_id"})
}

func (t *HistoryScheduledTask) GetTaskIDBytes() []byte {
	return t.TaskID.Bytes()
}

func (t *HistoryScheduledTask) SetTaskID(data []byte) {
	t.TaskID = big.NewInt(0).SetBytes(data)
}

func (t *HistoryScheduledTask) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (t *HistoryScheduledTask) GetDataBytes() []byte {
	return t.Data
}

func (t *HistoryScheduledTask) SetData(data []byte) {
	t.Data = data
}

func (t *HistoryScheduledTask) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (t *HistoryScheduledTask) GetDataEncodingBytes() []byte {
	return []byte(t.DataEncoding)
}

func (t *HistoryScheduledTask) SetDataEncoding(data []byte) {
	t.DataEncoding = string(data)
}

func (t *HistoryScheduledTask) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// HistoryScheduledTasks table row data
		// Note: ShardID, CategoryID, VisibilityTimestamp and TaskID are part of the
		// primary key of history_scheduled_tasks and as such are present in the prefix of
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

type HistoryScheduledTaskRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewHistoryScheduledTaskRepository(tr fdb.Transactor, ss subspace.Subspace) *HistoryScheduledTaskRepository {
	return &HistoryScheduledTaskRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *HistoryScheduledTaskRepository) Find(shardID int64, categoryID int64, visibilityTimestamp time.Time, taskID *big.Int) (*HistoryScheduledTask, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &HistoryScheduledTask{
			ShardID:             shardID,
			CategoryID:          categoryID,
			VisibilityTimestamp: visibilityTimestamp,
			TaskID:              taskID,
		}

		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*HistoryScheduledTask), err
}
