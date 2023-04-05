package storage

import (
	"math/big"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

type Shard struct {
	ShardID      int64
	RangeID      *big.Int
	Data         []byte
	DataEncoding string
}

func (s *Shard) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{s.GetShardIDBytes()})
}

func (s *Shard) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_idx", s.GetShardIDBytes()})
}

func (s *Shard) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(s.ShardID))
	return buff
}

func (s *Shard) SetShardID(data []byte) {
	s.ShardID = int64(BytesOrder.Uint64(data))
}

func (s *Shard) GetRangeIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return s.PrimaryKeyKey(ss).Pack(tuple.Tuple{"range_id"})
}

func (s *Shard) GetRangeIDBytes() []byte {
	return s.RangeID.Bytes()
}

func (s *Shard) SetRangeID(data []byte) {
	s.RangeID = big.NewInt(0).SetBytes(data)
}

func (s *Shard) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return s.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (s *Shard) GetDataBytes() []byte {
	return s.Data
}

func (s *Shard) SetData(data []byte) {
	s.Data = data
}

func (s *Shard) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return s.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (s *Shard) GetDataEncodingBytes() []byte {
	return []byte(s.DataEncoding)
}

func (s *Shard) SetDataEncoding(data []byte) {
	s.DataEncoding = string(data)
}

func (s *Shard) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// Namespaces table row data
		// Note: SharID is part of the primary key of shards and as such are
		// present in the prefix of all keys below. Storing a key/value pair
		// for those is not strictly necessary as it wastes space duplicating
		// data.
		tx.Set(s.GetRangeIDKey(ss), s.GetRangeIDBytes())
		tx.Set(s.GetDataKey(ss), s.GetDataBytes())
		tx.Set(s.GetDataEncodingKey(ss), s.GetDataEncodingBytes())

		// Indexes
		tx.Set(s.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type ShardRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewShardRepository(tr fdb.Transactor, ss subspace.Subspace) *ShardRepository {
	return &ShardRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *ShardRepository) Find(shardID int64) (*Shard, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		ns := &Shard{
			ShardID: shardID,
		}

		ns.SetRangeID(rtx.Get(ns.GetRangeIDKey(r.ss)).MustGet())
		ns.SetData(rtx.Get(ns.GetDataKey(r.ss)).MustGet())
		ns.SetDataEncoding(rtx.Get(ns.GetDataEncodingKey(r.ss)).MustGet())

		return ns, nil
	})

	return res.(*Shard), err
}

func (r *ShardRepository) Save(in *Shard) error {
	return in.Save(r.tr, r.ss)
}
