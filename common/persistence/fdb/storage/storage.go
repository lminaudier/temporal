package storage

import (
	"encoding/binary"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

// BytesOrder is the default binary byte order we use to encode keys and
// values in FDB
var BytesOrder = binary.BigEndian

type Storage struct {
	ActivityInfoMap       *ActivityInfoMapRepository
	BufferedEvent         *BufferedEventRepository
	ChildExecutionInfoMap *ChildExecutionInfoMapRepository
	ClusterMembership     *ClusterMembershipRepository
	ClusterMetadataInfo   *ClusterMetadataInfoRepository
	CurrentExecution      *CurrentExecutionRepository
	Execution             *ExecutionRepository
	HistoryImmediateTask  *HistoryImmediateTaskRepository
	HistoryNode           *HistoryNodeRepository
	HistoryScheduledTask  *HistoryScheduledTaskRepository
	HistoryTree           *HistoryTreeRepository
	Namespace             *NamespaceRepository
	NamespaceMetadata     *NamespaceMetadataRepository
	Queue                 *QueueRepository
	QueueMetadata         *QueueMetadataRepository
	ReplicationTask       *ReplicationTaskRepository
	ReplicationTaskSQL    *ReplicationTaskSQLRepository
	RequestCancelInfoMap  *RequestCancelInfoMapRepository
	Shard                 *ShardRepository
	SignalInfoMap         *SignalInfoMapRepository
	SignalRequestedSet    *SignalRequestedSetRepository
	Task                  *TaskRepository
	TaskQueue             *TaskQueueRepository
	TimerInfoMap          *TimerInfoMapRepository
	TimerTask             *TimerTaskRepository
	TransferTask          *TransferTaskRepository
	VisibilityTask        *VisibilityTaskRepository
}

func NewStorage(db fdb.Database, dir directory.DirectorySubspace) *Storage {
	return &Storage{
		ActivityInfoMap:       NewActivityInfoMapRepository(db, dir.Sub("activity_info_maps")),
		BufferedEvent:         NewBufferedEventRepository(db, dir.Sub("buffered_events")),
		ChildExecutionInfoMap: NewChildExecutionInfoMapRepository(db, dir.Sub("child_execution_info_maps")),
		ClusterMembership:     NewClusterMembershipRepository(db, dir.Sub("cluster_memberships")),
		ClusterMetadataInfo:   NewClusterMetadataInfoRepository(db, dir.Sub("cluster_metadata_info")),
		CurrentExecution:      NewCurrentExecutionRepository(db, dir.Sub("current_executions")),
		Execution:             NewExecutionRepository(db, dir.Sub("executions")),
		HistoryImmediateTask:  NewHistoryImmediateTaskRepository(db, dir.Sub("history_immediate_tasks")),
		HistoryNode:           NewHistoryNodeRepository(db, dir.Sub("history_nodes")),
		HistoryScheduledTask:  NewHistoryScheduledTaskRepository(db, dir.Sub("history_scheduled_tasks")),
		HistoryTree:           NewHistoryTreeRepository(db, dir.Sub("history_trees")),
		Namespace:             NewNamespaceRepository(db, dir.Sub("namespaces")),
		NamespaceMetadata:     NewNamespaceMetadataRepository(db, dir.Sub("namespaces_metadata")),
		Queue:                 NewQueueRepository(db, dir.Sub("queues")),
		QueueMetadata:         NewQueueMetadataRepository(db, dir.Sub("queues_metadata")),
		ReplicationTask:       NewReplicationTaskRepository(db, dir.Sub("replication_tasks")),
		ReplicationTaskSQL:    NewReplicationTaskSQLRepository(db, dir.Sub("replication_tasks_sql")),
		RequestCancelInfoMap:  NewRequestCancelInfoMapRepository(db, dir.Sub("request_cancel_info_maps")),
		Shard:                 NewShardRepository(db, dir.Sub("shards")),
		SignalInfoMap:         NewSignalInfoMapRepository(db, dir.Sub("signal_info_map")),
		SignalRequestedSet:    NewSignalRequestedSetRepository(db, dir.Sub("signal_requested_set")),
		Task:                  NewTaskRepository(db, dir.Sub("tasks")),
		TaskQueue:             NewTaskQueueRepository(db, dir.Sub("task_queues")),
		TimerInfoMap:          NewTimerInfoMapRepository(db, dir.Sub("timer_info_maps")),
		TimerTask:             NewTimerTaskRepository(db, dir.Sub("timer_tasks")),
		TransferTask:          NewTransferTaskRepository(db, dir.Sub("transfer_tasks")),
		VisibilityTask:        NewVisibilityTaskRepository(db, dir.Sub("visibility_tasks")),
	}
}
