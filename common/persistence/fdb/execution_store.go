package fdb

import (
	"context"

	"go.temporal.io/server/common/persistence"
)

type ExecutionStore struct{}

var _ persistence.ExecutionStore = (*ExecutionStore)(nil)

func (s *ExecutionStore) Close() {}

func (s *ExecutionStore) GetName() string {
	panic("unimplemented")
}

func (s *ExecutionStore) CreateWorkflowExecution(ctx context.Context, request *persistence.InternalCreateWorkflowExecutionRequest) (*persistence.InternalCreateWorkflowExecutionResponse, error) {
	panic("unimplemented")
}

func (s *ExecutionStore) UpdateWorkflowExecution(ctx context.Context, request *persistence.InternalUpdateWorkflowExecutionRequest) error {
	panic("unimplemented")
}

func (s *ExecutionStore) ConflictResolveWorkflowExecution(ctx context.Context, request *persistence.InternalConflictResolveWorkflowExecutionRequest) error {
	panic("unimplemented")
}

func (s *ExecutionStore) DeleteWorkflowExecution(ctx context.Context, request *persistence.DeleteWorkflowExecutionRequest) error {
	panic("unimplemented")
}

func (s *ExecutionStore) DeleteCurrentWorkflowExecution(ctx context.Context, request *persistence.DeleteCurrentWorkflowExecutionRequest) error {
	panic("unimplemented")
}

func (s *ExecutionStore) GetCurrentExecution(ctx context.Context, request *persistence.GetCurrentExecutionRequest) (*persistence.InternalGetCurrentExecutionResponse, error) {
	panic("unimplemented")
}

func (s *ExecutionStore) GetWorkflowExecution(ctx context.Context, request *persistence.GetWorkflowExecutionRequest) (*persistence.InternalGetWorkflowExecutionResponse, error) {
	panic("unimplemented")
}

func (s *ExecutionStore) SetWorkflowExecution(ctx context.Context, request *persistence.InternalSetWorkflowExecutionRequest) error {
	panic("unimplemented")
}

func (s *ExecutionStore) ListConcreteExecutions(ctx context.Context, request *persistence.ListConcreteExecutionsRequest) (*persistence.InternalListConcreteExecutionsResponse, error) {
	panic("unimplemented")
}

func (s *ExecutionStore) AddHistoryTasks(ctx context.Context, request *persistence.InternalAddHistoryTasksRequest) error {
	panic("unimplemented")
}

func (s *ExecutionStore) GetHistoryTask(ctx context.Context, request *persistence.GetHistoryTaskRequest) (*persistence.InternalGetHistoryTaskResponse, error) {
	panic("unimplemented")
}

func (s *ExecutionStore) GetHistoryTasks(ctx context.Context, request *persistence.GetHistoryTasksRequest) (*persistence.InternalGetHistoryTasksResponse, error) {
	panic("unimplemented")
}

func (s *ExecutionStore) CompleteHistoryTask(ctx context.Context, request *persistence.CompleteHistoryTaskRequest) error {
	panic("unimplemented")
}

func (s *ExecutionStore) RangeCompleteHistoryTasks(ctx context.Context, request *persistence.RangeCompleteHistoryTasksRequest) error {
	panic("unimplemented")
}

func (s *ExecutionStore) PutReplicationTaskToDLQ(ctx context.Context, request *persistence.PutReplicationTaskToDLQRequest) error {
	panic("unimplemented")
}

func (s *ExecutionStore) GetReplicationTasksFromDLQ(ctx context.Context, request *persistence.GetReplicationTasksFromDLQRequest) (*persistence.InternalGetReplicationTasksFromDLQResponse, error) {
	panic("unimplemented")
}

func (s *ExecutionStore) DeleteReplicationTaskFromDLQ(ctx context.Context, request *persistence.DeleteReplicationTaskFromDLQRequest) error {
	panic("unimplemented")
}

func (s *ExecutionStore) RangeDeleteReplicationTaskFromDLQ(ctx context.Context, request *persistence.RangeDeleteReplicationTaskFromDLQRequest) error {
	panic("unimplemented")
}

// The below are history V2 APIs
// V2 regards history events growing as a tree, decoupled from workflow concepts

// AppendHistoryNodes add a node to history node table
func (s *ExecutionStore) AppendHistoryNodes(ctx context.Context, request *persistence.InternalAppendHistoryNodesRequest) error {
	panic("unimplemented")
}

// DeleteHistoryNodes delete a node from history node table
func (s *ExecutionStore) DeleteHistoryNodes(ctx context.Context, request *persistence.InternalDeleteHistoryNodesRequest) error {
	panic("unimplemented")
}

// ParseHistoryBranchInfo parses the history branch for branch information
func (s *ExecutionStore) ParseHistoryBranchInfo(ctx context.Context, request *persistence.ParseHistoryBranchInfoRequest) (*persistence.ParseHistoryBranchInfoResponse, error) {
	panic("unimplemented")
}

// UpdateHistoryBranchInfo updates the history branch with branch information
func (s *ExecutionStore) UpdateHistoryBranchInfo(ctx context.Context, request *persistence.UpdateHistoryBranchInfoRequest) (*persistence.UpdateHistoryBranchInfoResponse, error) {
	panic("unimplemented")
}

// NewHistoryBranch initializes a new history branch
func (s *ExecutionStore) NewHistoryBranch(ctx context.Context, request *persistence.NewHistoryBranchRequest) (*persistence.NewHistoryBranchResponse, error) {
	panic("unimplemented")
}

// ReadHistoryBranch returns history node data for a branch
func (s *ExecutionStore) ReadHistoryBranch(ctx context.Context, request *persistence.InternalReadHistoryBranchRequest) (*persistence.InternalReadHistoryBranchResponse, error) {
	panic("unimplemented")
}

// ForkHistoryBranch forks a new branch from a old branch
func (s *ExecutionStore) ForkHistoryBranch(ctx context.Context, request *persistence.InternalForkHistoryBranchRequest) error {
	panic("unimplemented")
}

// DeleteHistoryBranch removes a branch
func (s *ExecutionStore) DeleteHistoryBranch(ctx context.Context, request *persistence.InternalDeleteHistoryBranchRequest) error {
	panic("unimplemented")
}

// GetHistoryTree returns all branch information of a tree
func (s *ExecutionStore) GetHistoryTree(ctx context.Context, request *persistence.GetHistoryTreeRequest) (*persistence.InternalGetHistoryTreeResponse, error) {
	panic("unimplemented")
}

// GetAllHistoryTreeBranches returns all branches of all trees.
// Note that branches may be skipped or duplicated across pages if there are branches created or deleted while
// paginating through results.
func (s *ExecutionStore) GetAllHistoryTreeBranches(ctx context.Context, request *persistence.GetAllHistoryTreeBranchesRequest) (*persistence.InternalGetAllHistoryTreeBranchesResponse, error) {
	panic("unimplemented")
}
