package storage

import (
	"math/big"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// CREATE TABLE timer_tasks (
//
//	shard_id INTEGER NOT NULL,
//	visibility_timestamp TIMESTAMP NOT NULL,
//	task_id BIGINT NOT NULL,
//	--
//	data BYTEA NOT NULL,
//	data_encoding VARCHAR(16) NOT NULL,
//	PRIMARY KEY (shard_id, visibility_timestamp, task_id)
//
// );
type TimerTask struct {
	ShardID             int64
	VisibilityTimestamp time.Time
	TaskID              *big.Int
	Data                []byte
	DataEncoding        string
}

func (t *TimerTask) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{t.GetShardIDBytes(), t.GetVisibilityTimestampBytes(), t.GetTaskIDBytes()})
}

func (t *TimerTask) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_visibility_timestamp_task_id_idx", t.GetShardIDBytes(), t.GetVisibilityTimestampBytes(), t.GetTaskIDBytes()})
}

func (t *TimerTask) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (t *TimerTask) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(t.ShardID))
	return buff
}

func (t *TimerTask) SetShardID(data []byte) {
	t.ShardID = int64(BytesOrder.Uint64(data))
}

func (t *TimerTask) GetVisibilityTimestampKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"visibility_timestamp"})
}

func (t *TimerTask) GetVisibilityTimestampBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(t.VisibilityTimestamp.UnixNano()))
	return buff
}

func (t *TimerTask) SetVisibilityTimestamp(data []byte) {
	t.VisibilityTimestamp = time.Unix(0, int64(BytesOrder.Uint64(data)))
}

func (t *TimerTask) GetTaskIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"task_id"})
}

func (t *TimerTask) GetTaskIDBytes() []byte {
	return t.TaskID.Bytes()
}

func (t *TimerTask) SetTaskID(data []byte) {
	t.TaskID = big.NewInt(0).SetBytes(data)
}

func (t *TimerTask) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (t *TimerTask) GetDataBytes() []byte {
	return t.Data
}

func (t *TimerTask) SetData(data []byte) {
	t.Data = data
}

func (t *TimerTask) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (t *TimerTask) GetDataEncodingBytes() []byte {
	return []byte(t.DataEncoding)
}

func (t *TimerTask) SetDataEncoding(data []byte) {
	t.DataEncoding = string(data)
}

func (t *TimerTask) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// TimerTasks table row data
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

type TimerTaskRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewTimerTaskRepository(tr fdb.Transactor, ss subspace.Subspace) *TimerTaskRepository {
	return &TimerTaskRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *TimerTaskRepository) Find(shardID int64, visibilityTimestamp time.Time, taskID *big.Int) (*TimerTask, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &TimerTask{
			ShardID:             shardID,
			VisibilityTimestamp: visibilityTimestamp,
			TaskID:              taskID,
		}

		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*TimerTask), err
}
