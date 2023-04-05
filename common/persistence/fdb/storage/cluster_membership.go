package storage

import (
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// CREATE TABLE cluster_membership
// (
//
//	membership_partition INTEGER NOT NULL,
//	host_id              BYTEA NOT NULL,
//	rpc_address          VARCHAR(128) NOT NULL,
//	rpc_port             SMALLINT NOT NULL,
//	role                 SMALLINT NOT NULL,
//	session_start        TIMESTAMP DEFAULT '1970-01-01 00:00:01+00:00',
//	last_heartbeat       TIMESTAMP DEFAULT '1970-01-01 00:00:01+00:00',
//	record_expiry        TIMESTAMP DEFAULT '1970-01-01 00:00:01+00:00',
//	PRIMARY KEY (membership_partition, host_id)
//
// );
type ClusterMembership struct {
	MembershipPartition int64
	HostID              []byte
	RPCAddress          string
	RPCPort             int8
	Role                int8
	SessionStart        time.Time
	LastHeartbeat       time.Time
	RecordExpiry        time.Duration
}

func (m *ClusterMembership) PrimaryKeyKey(ss subspace.Subspace) subspace.Subspace {
	return ss.Sub(tuple.Tuple{m.GetMembershipPartitionBytes(), m.GetHostIDBytes()})
}

func (m *ClusterMembership) PrimaryKeyIndexKey(ss subspace.Subspace) fdb.KeyConvertible {
	return ss.Pack(tuple.Tuple{"membership_partition_host_id_idx", m.GetMembershipPartitionBytes(), m.GetHostIDBytes()})
}

func (m *ClusterMembership) GetMembershipPartitionKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"membership_partition"})
}

func (m *ClusterMembership) GetMembershipPartitionBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(m.MembershipPartition))
	return buff
}

func (m *ClusterMembership) SetMembershipPartition(data []byte) {
	m.MembershipPartition = int64(BytesOrder.Uint64(data))
}

func (m *ClusterMembership) GetHostIDKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"host_id"})
}

func (m *ClusterMembership) GetHostIDBytes() []byte {
	return m.HostID
}

func (m *ClusterMembership) SetHostID(data []byte) {
	m.HostID = data
}

func (m *ClusterMembership) GetRPCAddressKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"rpc_address"})
}

func (m *ClusterMembership) GetRPCAddressBytes() []byte {
	return []byte(m.RPCAddress)
}

func (m *ClusterMembership) SetRPCAddress(data []byte) {
	m.RPCAddress = string(data)
}

func (m *ClusterMembership) GetRPCPortKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"rpc_port"})
}

func (m *ClusterMembership) GetRPCPortBytes() []byte {
	buff := make([]byte, 2)
	BytesOrder.PutUint16(buff, uint16(m.RPCPort))
	return buff
}

func (m *ClusterMembership) SetRPCPort(data []byte) {
	m.RPCPort = int8(BytesOrder.Uint16(data))
}

func (m *ClusterMembership) GetRoleKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"role"})
}

func (m *ClusterMembership) GetRoleBytes() []byte {
	buff := make([]byte, 2)
	BytesOrder.PutUint16(buff, uint16(m.Role))
	return buff
}

func (m *ClusterMembership) SetRole(data []byte) {
	m.Role = int8(BytesOrder.Uint16(data))
}

func (m *ClusterMembership) GetSessionStartKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"session_start"})
}

func (m *ClusterMembership) GetSessionStartBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(m.SessionStart.UTC().UnixNano()))
	return buff
}

func (m *ClusterMembership) SetSessionStart(data []byte) {
	m.SessionStart = time.Unix(0, int64(BytesOrder.Uint64(data))).UTC()
}

func (m *ClusterMembership) GetLastHeartbeatKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"last_heartbeat"})
}

func (m *ClusterMembership) GetLastHeartbeatBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(m.LastHeartbeat.UTC().UnixNano()))
	return buff
}

func (m *ClusterMembership) SetLastHeartbeat(data []byte) {
	m.LastHeartbeat = time.Unix(0, int64(BytesOrder.Uint64(data))).UTC()
}

func (m *ClusterMembership) GetRecordExpiryKey(ss subspace.Subspace) fdb.KeyConvertible {
	return m.PrimaryKeyKey(ss).Pack(tuple.Tuple{"record_expiry"})
}

func (m *ClusterMembership) GetRecordExpiryBytes() []byte {
	buff := make([]byte, 8)
	BytesOrder.PutUint64(buff, uint64(m.RecordExpiry.Nanoseconds()))
	return buff
}

func (m *ClusterMembership) SetRecordExpiry(data []byte) {
	m.RecordExpiry = time.Duration(int64(BytesOrder.Uint64(data)))
}

func (m *ClusterMembership) Save(tr fdb.Transactor, ss subspace.Subspace) error {
	tr.Transact(func(tx fdb.Transaction) (interface{}, error) {
		tx.Set(m.GetRPCAddressKey(ss), m.GetRPCAddressBytes())
		tx.Set(m.GetRPCPortKey(ss), m.GetRPCPortBytes())
		tx.Set(m.GetRoleKey(ss), m.GetRoleBytes())
		tx.Set(m.GetSessionStartKey(ss), m.GetSessionStartBytes())
		tx.Set(m.GetLastHeartbeatKey(ss), m.GetLastHeartbeatBytes())
		tx.Set(m.GetRecordExpiryKey(ss), m.GetRecordExpiryBytes())

		// Indexes
		// TODO: create indexes
		// CREATE UNIQUE INDEX cm_idx_rolehost ON cluster_membership (role, host_id);
		// CREATE INDEX cm_idx_rolelasthb ON cluster_membership (role, last_heartbeat);
		// CREATE INDEX cm_idx_rpchost ON cluster_membership (rpc_address, role);
		// CREATE INDEX cm_idx_lasthb ON cluster_membership (last_heartbeat);
		// CREATE INDEX cm_idx_recordexpiry ON cluster_membership (record_expiry);
		tx.Set(m.PrimaryKeyIndexKey(ss), []byte{})

		return nil, nil
	})

	return nil
}

type ClusterMembershipRepository struct {
	tr fdb.Transactor
	ss subspace.Subspace
}

func NewClusterMembershipRepository(tr fdb.Transactor, ss subspace.Subspace) *ClusterMembershipRepository {
	return &ClusterMembershipRepository{
		tr: tr,
		ss: ss,
	}
}

func (r *ClusterMembershipRepository) Find(membershipPartition int64, hostID []byte) (*ClusterMembership, error) {
	res, err := r.tr.ReadTransact(func(rtx fdb.ReadTransaction) (interface{}, error) {
		m := &ClusterMembership{
			MembershipPartition: membershipPartition,
			HostID:              hostID,
		}

		m.SetRPCAddress(rtx.Get(m.GetRPCAddressKey(r.ss)).MustGet())
		m.SetRPCPort(rtx.Get(m.GetRPCPortKey(r.ss)).MustGet())
		m.SetRole(rtx.Get(m.GetRoleKey(r.ss)).MustGet())
		m.SetSessionStart(rtx.Get(m.GetSessionStartKey(r.ss)).MustGet())
		m.SetLastHeartbeat(rtx.Get(m.GetLastHeartbeatKey(r.ss)).MustGet())
		m.SetRecordExpiry(rtx.Get(m.GetRecordExpiryKey(r.ss)).MustGet())

		return m, nil
	})

	return res.(*ClusterMembership), err
}

func (r *ClusterMembershipRepository) Save(in *ClusterMembership) error {
	return in.Save(r.tr, r.ss)
}
