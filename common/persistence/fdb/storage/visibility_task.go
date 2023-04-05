package storage

import (
	"math/big"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// CREATE TABLE visibility_tasks(
//
//	shard_id INTEGER NOT NULL,
//	task_id BIGINT NOT NULL,
//	--
//	data BYTEA NOT NULL,
//	data_encoding VARCHAR(16) NOT NULL,
//	PRIMARY KEY (shard_id, task_id)
//
// );
type VisibilityTask struct {
	ShardID      int64
	TaskID       *big.Int
	Data         []byte
	DataEncoding string
}

func (t *VisibilityTask) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{t.GetShardIDBytes(), t.GetTaskIDBytes()})
}

func (t *VisibilityTask) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_task_id_idx", t.GetShardIDBytes(), t.GetTaskIDBytes()})
}

func (t *VisibilityTask) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (t *VisibilityTask) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(t.ShardID))
	return buff
}

func (t *VisibilityTask) SetShardID(data []byte) {
	t.ShardID = int64(BytesOrder.Uint64(data))
}

func (t *VisibilityTask) GetTaskIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"task_id"})
}

func (t *VisibilityTask) GetTaskIDBytes() []byte {
	return t.TaskID.Bytes()
}

func (t *VisibilityTask) SetTaskID(data []byte) {
	t.TaskID = big.NewInt(0).SetBytes(data)
}

func (t *VisibilityTask) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (t *VisibilityTask) GetDataBytes() []byte {
	return t.Data
}

func (t *VisibilityTask) SetData(data []byte) {
	t.Data = data
}

func (t *VisibilityTask) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (t *VisibilityTask) GetDataEncodingBytes() []byte {
	return []byte(t.DataEncoding)
}

func (t *VisibilityTask) SetDataEncoding(data []byte) {
	t.DataEncoding = string(data)
}

func (t *VisibilityTask) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// VisibilityTasks table row data
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

type VisibilityTaskRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewVisibilityTaskRepository(tr fdb.Transactor, ss subspace.Subspace) *VisibilityTaskRepository {
	return &VisibilityTaskRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *VisibilityTaskRepository) Find(shardID int64, taskID *big.Int) (*VisibilityTask, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &VisibilityTask{
			ShardID: shardID,
			TaskID:  taskID,
		}

		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*VisibilityTask), err
}
