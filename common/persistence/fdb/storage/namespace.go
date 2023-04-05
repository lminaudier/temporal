package storage

import (
	"math/big"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

type Namespace struct {
	PartitionID         int64
	ID                  []byte
	Name                string
	NotificationVersion *big.Int
	Data                []byte
	DataEncoding        string
	IsGlobal            bool
}

func (n *Namespace) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{n.GetPartitionIDBytes(), n.GetIDBytes()})
}

func (n *Namespace) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"partition_id_and_id_idx", n.GetPartitionIDBytes(), n.GetIDBytes()})
}

func NamespacePrimaryKeyIndexPrefix(ss subspace.Subspace) fdb.Key {
	return ss.Pack(tuple.Tuple{"partition_id_and_id_idx"})
}

func (n *Namespace) GetPartitionIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return n.PrimaryKeyKey(ss).Pack(tuple.Tuple{"partition_id"})
}

func (n *Namespace) GetPartitionIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(n.PartitionID))
	return buff
}

func (n *Namespace) SetPartitionID(data []byte) {
	n.PartitionID = int64(BytesOrder.Uint64(data))
}

func (n *Namespace) GetIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return n.PrimaryKeyKey(ss).Pack(tuple.Tuple{"id"})
}

func (n *Namespace) GetIDBytes() []byte {
	return n.ID
}

func (n *Namespace) SetID(data []byte) {
	n.ID = data
}

func (n *Namespace) GetNameKey(ss subspace.Subspace) fdb.KeyConvertible {
	return n.PrimaryKeyKey(ss).Pack(tuple.Tuple{"name"})
}

func (n *Namespace) GetNameBytes() []byte {
	return []byte(n.Name)
}

func (n *Namespace) SetName(data []byte) {
	n.Name = string(data)
}

func (n *Namespace) GetNotificationVersionKey(ss subspace.Subspace) fdb.KeyConvertible {
	return n.PrimaryKeyKey(ss).Pack(tuple.Tuple{"notification_version"})
}

func (n *Namespace) GetNotificationVersionBytes() []byte {
	if n.NotificationVersion == nil {
		n.NotificationVersion = big.NewInt(0)
	}
	return n.NotificationVersion.Bytes()
}

func (n *Namespace) SetNotificationVersion(data []byte) {
	n.NotificationVersion = big.NewInt(0).SetBytes(data)
}

func (n *Namespace) GetDataKey(ss subspace.Subspace) fdb.KeyConvertible {
	return n.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data"})
}

func (n *Namespace) GetDataBytes() []byte {
	return n.Data
}

func (n *Namespace) SetData(data []byte) {
	n.Data = data
}

func (n *Namespace) GetDataEncodingKey(ss subspace.Subspace) fdb.KeyConvertible {
	return n.PrimaryKeyKey(ss).Pack(tuple.Tuple{"data_encoding"})
}

func (n *Namespace) GetDataEncodingBytes() []byte {
	return []byte(n.DataEncoding)
}

func (n *Namespace) SetDataEncoding(data []byte) {
	n.DataEncoding = string(data)
}

func (n *Namespace) GetIsGlobalKey(ss subspace.Subspace) fdb.KeyConvertible {
	return n.PrimaryKeyKey(ss).Pack(tuple.Tuple{"is_global"})
}

func (n *Namespace) GetIsGlobalBytes() []byte {
	if n.IsGlobal {
		return []byte{1}
	}
	return []byte{0}
}

func (n *Namespace) SetIsGlobal(data []byte) {
	// FIXME: is it a safe default? likely not
	if len(data) < 1 {
		n.IsGlobal = false
		return
	}
	n.IsGlobal = (data[0] == 1)
}

func (n *Namespace) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// Namespaces table row data
		// Note: PartitionID and ID are part of the primary key of namespaces
		// and as such are present in the prefix of all keys below. Storing a
		// key/value pair for those is not strictly necessary as it wastes
		// space duplicating data.
		tx.Set(n.GetNameKey(ss), n.GetNameBytes())
		tx.Set(n.GetNotificationVersionKey(ss), n.GetNotificationVersionBytes())
		tx.Set(n.GetDataKey(ss), n.GetDataBytes())
		tx.Set(n.GetDataEncodingKey(ss), n.GetDataEncodingBytes())
		tx.Set(n.GetIsGlobalKey(ss), n.GetIsGlobalBytes())

		// Indexes
		tx.Set(n.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type NamespaceRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewNamespaceRepository(tr fdb.Transactor, ss subspace.Subspace) *NamespaceRepository {
	return &NamespaceRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *NamespaceRepository) Find(partitionID int64, id []byte) (*Namespace, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		ns := &Namespace{
			PartitionID: partitionID,
			ID:          id,
		}

		ns.SetName(rtx.Get(ns.GetNameKey(r.ss)).MustGet())
		ns.SetNotificationVersion(rtx.Get(ns.GetNotificationVersionKey(r.ss)).MustGet())
		ns.SetData(rtx.Get(ns.GetDataKey(r.ss)).MustGet())
		ns.SetDataEncoding(rtx.Get(ns.GetDataEncodingKey(r.ss)).MustGet())
		ns.SetIsGlobal(rtx.Get(ns.GetIsGlobalKey(r.ss)).MustGet())

		return ns, nil
	})

	return res.(*Namespace), err
}

func (r *NamespaceRepository) List() ([]*Namespace, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		result := make([]*Namespace, 0)

		kr, err := fdb.PrefixRange(NamespacePrimaryKeyIndexPrefix(r.ss))
		if err != nil {
			return nil, err
		}
		it := rtx.GetRange(kr, fdb.RangeOptions{}).Iterator()
		for it.Advance() {
			tuple, err := r.ss.Unpack(it.MustGet().Key)
			if err != nil {
				return nil, err
			}

			m := &Namespace{}
			m.SetPartitionID(tuple[len(tuple)-2].([]byte))
			m.SetID(tuple[len(tuple)-1].([]byte))

			m.SetNotificationVersion(rtx.Get(m.GetNotificationVersionKey(r.ss)).MustGet())
			m.SetData(rtx.Get(m.GetDataKey(r.ss)).MustGet())
			m.SetDataEncoding(rtx.Get(m.GetDataEncodingKey(r.ss)).MustGet())
			m.SetIsGlobal(rtx.Get(m.GetIsGlobalKey(r.ss)).MustGet())

			result = append(result, m)
		}

		return result, nil
	})

	return res.([]*Namespace), err
}

func (r *NamespaceRepository) Save(ns *Namespace) error {
	return ns.Save(r.tr, r.ss)
}

type NamespaceMetadata struct {
	PartitionID         int64
	NotificationVersion *big.Int
}

func (m *NamespaceMetadata) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{m.GetPartitionIDBytes()})
}

func (m *NamespaceMetadata) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"partition_id_idx", m.GetPartitionIDBytes()})
}

func (m *NamespaceMetadata) GetPartitionIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"partition_id"})
}

func (m *NamespaceMetadata) GetPartitionIDBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(m.PartitionID))
	return buff
}

func (m *NamespaceMetadata) SetPartitionID(data []byte) {
	m.PartitionID = int64(BytesOrder.Uint64(data))
}

func (m *NamespaceMetadata) GetNotificationVersionKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"notification_version"})
}

func (m *NamespaceMetadata) GetNotificationVersionBytes() []byte {
	return m.NotificationVersion.Bytes()
}

func (m *NamespaceMetadata) SetNotificationVersion(data []byte) {
	m.NotificationVersion = big.NewInt(0).SetBytes(data)
}

func (m *NamespaceMetadata) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		// Namespaces table row data
		// Note: PartitionID is part of the primary key of namespace_metadata
		// and as such are present in the prefix of all keys below. Storing a
		// key/value pair for those is not strictly necessary as it wastes
		// space duplicating data.
		tx.Set(m.GetNotificationVersionKey(ss), m.GetNotificationVersionBytes())

		// Indexes
		tx.Set(m.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type NamespaceMetadataRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewNamespaceMetadataRepository(tr fdb.Transactor, ss subspace.Subspace) *NamespaceMetadataRepository {
	return &NamespaceMetadataRepository{
		tr: tr,
		ss: ss,
	}
}

// TODO: PostgreSQL schema is inserting a default namespace metadata row, this is replicating that.
// Check how we should integrate that call in the cluster bootstrap tooling
func (r *NamespaceMetadataRepository) Init() (*NamespaceMetadata, error) {
	m := &NamespaceMetadata{
		PartitionID:         45001,
		NotificationVersion: big.NewInt(1),
	}
	err := m.Save(r.tr, r.ss)
	if err != nil {
		return m, err
	}

	return m, nil
}

func (r *NamespaceMetadataRepository) Find(partitionID int64) (*NamespaceMetadata, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		ns := &NamespaceMetadata{
			PartitionID: partitionID,
		}

		ns.SetNotificationVersion(rtx.Get(ns.GetNotificationVersionKey(r.ss)).MustGet())

		return ns, nil
	})

	return res.(*NamespaceMetadata), err
}
