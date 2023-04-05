package storage

import (
	"math/big"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// -- history eventsV2: history_node stores history event data
// CREATE TABLE history_node (
//
//	shard_id       INTEGER NOT NULL,
//	tree_id        BYTEA NOT NULL,
//	branch_id      BYTEA NOT NULL,
//	node_id        BIGINT NOT NULL,
//	txn_id         BIGINT NOT NULL,
//	--
//	prev_txn_id    BIGINT NOT NULL DEFAULT 0,
//	data           BYTEA NOT NULL,
//	data_encoding  VARCHAR(16) NOT NULL,
//	PRIMARY KEY (shard_id, tree_id, branch_id, node_id, txn_id)
//
// );
type HistoryNode struct {
	ShardID      int64
	TreeID       []byte
	BranchID     []byte
	NodeID       *big.Int
	TxnID        *big.Int
	PrevTxnID    *big.Int
	Data         []byte
	DataEncoding string
}

func (n *HistoryNode) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{n.GetShardIDBytes(), n.GetTreeIDBytes(), n.GetBranchIDBytes(), n.GetNodeIDBytes(), n.GetTxnIDBytes()})
}

func (n *HistoryNode) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"shard_id_tree_id_branch_id_node_id_txn_id_idx", n.GetShardIDBytes(), n.GetTreeIDBytes(), n.GetBranchIDBytes(), n.GetNodeIDBytes(), n.GetTxnIDBytes()})
}

func (m *HistoryNode) GetShardIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"shard_id"})
}

func (m *HistoryNode) GetShardIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(m.ShardID))
	return buff
}

func (m *HistoryNode) SetShardID(data []byte) {
	m.ShardID = int64(BytesOrder.Uint64(data))
}

func (n *HistoryNode) GetTreeIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return n.PrimaryKeyKey(ss).Pack(tuple.Tuple{"tree_id"})
}

func (n *HistoryNode) GetTreeIDBytes() []byte {
	return n.TreeID
}

func (n *HistoryNode) SetTreeID(data []byte) {
	n.TreeID = data
}

func (n *HistoryNode) GetBranchdIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return n.PrimaryKeyKey(ss).Pack(tuple.Tuple{"branch_id"})
}

func (n *HistoryNode) GetBranchIDBytes() []byte {
	return n.BranchID
}

func (n *HistoryNode) SetBranchID(data []byte) {
	n.BranchID = data
}

func (n *HistoryNode) GetNodeIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return n.PrimaryKeyKey(ss).Pack(tuple.Tuple{"node_id"})
}

func (n *HistoryNode) GetNodeIDBytes() []byte {
	return n.NodeID.Bytes()
}

func (n *HistoryNode) SetNodeID(data []byte) {
	n.NodeID = big.NewInt(0).SetBytes(data)
}

func (n *HistoryNode) GetTxnIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return n.PrimaryKeyKey(ss).Pack(tuple.Tuple{"txn_id"})
}

func (n *HistoryNode) GetTxnIDBytes() []byte {
	return n.TxnID.Bytes()
}

func (n *HistoryNode) SetTxnID(data []byte) {
	n.TxnID = big.NewInt(0).SetBytes(data)
}

func (n *HistoryNode) GetPrevTxnIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return n.PrimaryKeyKey(ss).Pack(tuple.Tuple{"prev_txn_id"})
}

func (n *HistoryNode) GetPrevTxnIDBytes() []byte {
	return n.PrevTxnID.Bytes()
}

func (n *HistoryNode) SetPrevTxnID(data []byte) {
	n.PrevTxnID = big.NewInt(0).SetBytes(data)
}

func (n *HistoryNode) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return n.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (n *HistoryNode) GetDataBytes() []byte {
	return n.Data
}

func (n *HistoryNode) SetData(data []byte) {
	n.Data = data
}

func (n *HistoryNode) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return n.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (n *HistoryNode) GetDataEncodingBytes() []byte {
	return []byte(n.DataEncoding)
}

func (n *HistoryNode) SetDataEncoding(data []byte) {
	n.DataEncoding = string(data)
}

func (n *HistoryNode) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		tx.Set(n.GetPrevTxnIDKey(ss), n.GetPrevTxnIDBytes())
		tx.Set(n.GetDataKey(ss), n.GetDataBytes())
		tx.Set(n.GetDataEncodingKey(ss), n.GetDataEncodingBytes())

		// Indexes
		tx.Set(n.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type HistoryNodeRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewHistoryNodeRepository(tr fdb.Transactor, ss subspace.Subspace) *HistoryNodeRepository {
	return &HistoryNodeRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *HistoryNodeRepository) Find(shardID int64, treeID []byte, branchID []byte, nodeID *big.Int, txnID *big.Int) (*HistoryNode, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		e := &HistoryNode{
			ShardID:  shardID,
			TreeID:   treeID,
			BranchID: branchID,
			NodeID:   nodeID,
			TxnID:    txnID,
		}

		e.SetPrevTxnID(rtx.Get(e.GetPrevTxnIDKey(r.ss)).MustGet())
		e.SetData(rtx.Get(e.GetDataKey(r.ss)).MustGet())
		e.SetDataEncoding(rtx.Get(e.GetDataEncodingKey(r.ss)).MustGet())

		return e, nil
	})

	return res.(*HistoryNode), err
}
