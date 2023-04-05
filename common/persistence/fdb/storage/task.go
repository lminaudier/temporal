package storage

import (
	"math/big"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

type Task struct {
	RangeHash    *big.Int
	TaskQueueID  []byte
	TaskID       *big.Int
	Data         []byte
	DataEncoding string
}

func (e *Task) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{e.GetRangeHashBytes(), e.GetTaskQueueIDBytes(), e.GetTaskIDBytes()})
}

func (e *Task) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"range_hash_task_queue_id_task_id_idx", e.GetRangeHashBytes(), e.GetTaskQueueIDBytes(), e.GetTaskIDBytes()})
}

func (e *Task) GetRangeHashKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"range_hash"})
}

func (e *Task) GetRangeHashBytes() []byte {
	return e.RangeHash.Bytes()
}

func (e *Task) SetRangeHash(data []byte) {
	e.RangeHash = big.NewInt(0).SetBytes(data)
}

func (e *Task) GetTaskQueueIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"task_queue_id"})
}

func (e *Task) GetTaskQueueIDBytes() []byte {
	return e.TaskQueueID
}

func (e *Task) SetTaskQueueID(data []byte) {
	e.TaskQueueID = data
}

func (e *Task) GetTaskIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"task_id"})
}

func (e *Task) GetTaskIDBytes() []byte {
	return e.TaskID.Bytes()
}

func (e *Task) SetTaskID(data []byte) {
	e.TaskID = big.NewInt(0).SetBytes(data)
}

func (e *Task) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (e *Task) GetDataBytes() []byte {
	return e.Data
}

func (e *Task) SetData(data []byte) {
	e.Data = data
}

func (e *Task) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (e *Task) GetDataEncodingBytes() []byte {
	return []byte(e.DataEncoding)
}

func (e *Task) SetDataEncoding(data []byte) {
	e.DataEncoding = string(data)
}

func (e *Task) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// Tasks table row data
		// Note: RangeHash, TaskQueueID and TaskId are part of the
		// primary key of tasks and as such are present in the prefix of
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

type TaskRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewTaskRepository(tr fdb.Transactor, ss subspace.Subspace) *TaskRepository {
	return &TaskRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *TaskRepository) Find(rangeHash *big.Int, taskQueueID []byte, taskID *big.Int) (*Task, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &Task{
			RangeHash:   rangeHash,
			TaskQueueID: taskQueueID,
			TaskID:      taskID,
		}

		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*Task), err
}
