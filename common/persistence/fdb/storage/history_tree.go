package storage

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// -- history eventsV2: history_tree stores branch metadata
// CREATE TABLE history_tree (
//
//	shard_id       INTEGER NOT NULL,
//	tree_id        BYTEA NOT NULL,
//	branch_id      BYTEA NOT NULL,
//	--
//	data           BYTEA NOT NULL,
//	data_encoding  VARCHAR(16) NOT NULL,
//	PRIMARY KEY (shard_id, tree_id, branch_id)
//
// );
type HistoryTree struct {
	ShardID      int64
	TreeID       []byte
	BranchID     []byte
	Data         []byte
	DataEncoding string
}

func (t *HistoryTree) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{t.GetShardIDBytes(), t.GetTreeIDBytes(), t.GetBranchIDBytes()})
}

func (t *HistoryTree) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_tree_id_branch_id_idx", t.GetShardIDBytes(), t.GetTreeIDBytes(), t.GetBranchIDBytes()})
}

func (t *HistoryTree) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (t *HistoryTree) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(t.ShardID))
	return buff
}

func (t *HistoryTree) SetShardID(data []byte) {
	t.ShardID = int64(BytesOrder.Uint64(data))
}

func (t *HistoryTree) GetTreeIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"tree_id"})
}

func (t *HistoryTree) GetTreeIDBytes() []byte {
	return t.TreeID
}

func (t *HistoryTree) SetTreeID(data []byte) {
	t.TreeID = data
}

func (t *HistoryTree) GetBranchIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"tree_id"})
}

func (t *HistoryTree) GetBranchIDBytes() []byte {
	return t.BranchID
}

func (t *HistoryTree) SetBranchID(data []byte) {
	t.BranchID = data
}

func (t *HistoryTree) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (t *HistoryTree) GetDataBytes() []byte {
	return t.Data
}

func (t *HistoryTree) SetData(data []byte) {
	t.Data = data
}

func (t *HistoryTree) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return t.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (t *HistoryTree) GetDataEncodingBytes() []byte {
	return []byte(t.DataEncoding)
}

func (t *HistoryTree) SetDataEncoding(data []byte) {
	t.DataEncoding = string(data)
}

func (t *HistoryTree) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		tx.Set(t.GetDataKey(ss), t.GetDataBytes())
		tx.Set(t.GetDataEncodingKey(ss), t.GetDataEncodingBytes())

		// Indexes
		tx.Set(t.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type HistoryTreeRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewHistoryTreeRepository(tr fdb.Transactor, ss subspace.Subspace) *HistoryTreeRepository {
	return &HistoryTreeRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *HistoryTreeRepository) Find(shardID int64, treeID []byte, branchID []byte) (*HistoryTree, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &HistoryTree{
			ShardID:  shardID,
			TreeID:   treeID,
			BranchID: branchID,
		}

		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*HistoryTree), err
}
