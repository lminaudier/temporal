package storage

import (
	"math/big"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// CREATE TABLE queue (
//
//	queue_type        INTEGER NOT NULL,
//	message_id        BIGINT NOT NULL,
//	message_payload   BYTEA NOT NULL,
//	message_encoding  VARCHAR(16) NOT NULL,
//	PRIMARY KEY(queue_type, message_id)
//
// );
type Queue struct {
	QueueType       int64
	MessageID       *big.Int
	MessagePayload  []byte
	MessageEncoding string
}

func (t *Queue) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{t.GetQueueTypeBytes(), t.GetMessageIDBytes()})
}

func (t *Queue) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"queue_type_message_id_idx", t.GetQueueTypeBytes(), t.GetMessageIDBytes()})
}

func (t *Queue) GetQueueTypeKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"queue_type"})
}

func (t *Queue) GetQueueTypeBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(t.QueueType))
	return buff
}

func (t *Queue) SetQueueType(data []byte) {
	t.QueueType = int64(BytesOrder.Uint64(data))
}

func (t *Queue) GetMessageIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"message_id"})
}

func (t *Queue) GetMessageIDBytes() []byte {
	return t.MessageID.Bytes()
}

func (t *Queue) SetMessageID(data []byte) {
	t.MessageID = big.NewInt(0).SetBytes(data)
}

func (t *Queue) GetMessagePayloadKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (t *Queue) GetMessagePayloadBytes() []byte {
	return t.MessagePayload
}

func (t *Queue) SetMessagePayload(data []byte) {
	t.MessagePayload = data
}

func (t *Queue) GetMessageEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (t *Queue) GetMessageEncodingBytes() []byte {
	return []byte(t.MessageEncoding)
}

func (t *Queue) SetMessageEncoding(data []byte) {
	t.MessageEncoding = string(data)
}

func (t *Queue) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		tx.Set(t.GetMessagePayloadKey(ss), t.GetMessagePayloadBytes())
		tx.Set(t.GetMessageEncodingKey(ss), t.GetMessageEncodingBytes())

		// Indexes
		tx.Set(t.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type QueueRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewQueueRepository(tr fdb.Transactor, ss subspace.Subspace) *QueueRepository {
	return &QueueRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *QueueRepository) Find(queueType int64, messageID *big.Int) (*Queue, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &Queue{
			QueueType: queueType,
			MessageID: messageID,
		}

		e.SetMessagePayload(rtx.Get(e.GetMessagePayloadKey(r.ss)).MustGet())
		e.SetMessageEncoding(rtx.Get(e.GetMessageEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*Queue), err
}

func (r *QueueRepository) Save(in *Queue) error {
	return in.Save(r.tr, r.ss)
}
