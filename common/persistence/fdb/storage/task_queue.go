package storage

import (
	"math/big"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

type TaskQueue struct {
	RangeHash    *big.Int
	TaskQueueID  []byte
	RangeID      *big.Int
	Data         []byte
	DataEncoding string
}

func (e *TaskQueue) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{e.GetRangeHashBytes(), e.GetTaskQueueIDBytes()})
}

func (e *TaskQueue) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"range_hash_task_queue_id_task_id_idx", e.GetRangeHashBytes(), e.GetTaskQueueIDBytes()})
}

func (e *TaskQueue) GetRangeHashKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"range_hash"})
}

func (e *TaskQueue) GetRangeHashBytes() []byte {
	return e.RangeHash.Bytes()
}

func (e *TaskQueue) SetRangeHash(data []byte) {
	e.RangeHash = big.NewInt(0).SetBytes(data)
}

func (e *TaskQueue) GetTaskQueueIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"task_queue_id"})
}

func (e *TaskQueue) GetTaskQueueIDBytes() []byte {
	return e.TaskQueueID
}

func (e *TaskQueue) SetTaskQueueID(data []byte) {
	e.TaskQueueID = data
}

func (e *TaskQueue) GetRangeIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"task_id"})
}

func (e *TaskQueue) GetRangeIDBytes() []byte {
	return e.RangeID.Bytes()
}

func (e *TaskQueue) SetRangeID(data []byte) {
	e.RangeID = big.NewInt(0).SetBytes(data)
}

func (e *TaskQueue) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (e *TaskQueue) GetDataBytes() []byte {
	return e.Data
}

func (e *TaskQueue) SetData(data []byte) {
	e.Data = data
}

func (e *TaskQueue) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return e.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (e *TaskQueue) GetDataEncodingBytes() []byte {
	return []byte(e.DataEncoding)
}

func (e *TaskQueue) SetDataEncoding(data []byte) {
	e.DataEncoding = string(data)
}

func (e *TaskQueue) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// TaskQueues table row data
		// Note: RangeHash and TaskQueueID are part of the
		// primary key of task_queues and as such are present in the prefix of
		// all keys below. Storing a key/value pair for those is not strictly
		// necessary as it wastes space duplicating data.
		tx.Set(e.GetRangeIDKey(ss), e.GetRangeIDBytes())
		tx.Set(e.GetDataKey(ss), e.GetDataBytes())
		tx.Set(e.GetDataEncodingKey(ss), e.GetDataEncodingBytes())

		// Indexes
		tx.Set(e.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type TaskQueueRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewTaskQueueRepository(tr fdb.Transactor, ss subspace.Subspace) *TaskQueueRepository {
	return &TaskQueueRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *TaskQueueRepository) Find(rangeHash *big.Int, taskQueueID []byte) (*TaskQueue, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &TaskQueue{
			RangeHash:   rangeHash,
			TaskQueueID: taskQueueID,
		}

		e.SetRangeID(rtx.Get(e.GetRangeIDKey(r.ss)).MustGet())
		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*TaskQueue), err
}
