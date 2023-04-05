package fdb

import (
	"context"

	"go.temporal.io/server/common/persistence"
)

type TaskStore struct{}

var _ persistence.TaskStore = (*TaskStore)(nil)

func (s *TaskStore) Close() {}

func (s *TaskStore) GetName() string {
	panic("unimplemented")
}

func (s *TaskStore) CreateTaskQueue(ctx context.Context, request *persistence.InternalCreateTaskQueueRequest) error {
	panic("unimplemented")
}

func (s *TaskStore) GetTaskQueue(ctx context.Context, request *persistence.InternalGetTaskQueueRequest) (*persistence.InternalGetTaskQueueResponse, error) {
	panic("unimplemented")
}

func (s *TaskStore) UpdateTaskQueue(ctx context.Context, request *persistence.InternalUpdateTaskQueueRequest) (*persistence.UpdateTaskQueueResponse, error) {
	panic("unimplemented")
}

func (s *TaskStore) ListTaskQueue(ctx context.Context, request *persistence.ListTaskQueueRequest) (*persistence.InternalListTaskQueueResponse, error) {
	panic("unimplemented")
}

func (s *TaskStore) DeleteTaskQueue(ctx context.Context, request *persistence.DeleteTaskQueueRequest) error {
	panic("unimplemented")
}

func (s *TaskStore) CreateTasks(ctx context.Context, request *persistence.InternalCreateTasksRequest) (*persistence.CreateTasksResponse, error) {
	panic("unimplemented")
}

func (s *TaskStore) GetTasks(ctx context.Context, request *persistence.GetTasksRequest) (*persistence.InternalGetTasksResponse, error) {
	panic("unimplemented")
}

func (s *TaskStore) CompleteTask(ctx context.Context, request *persistence.CompleteTaskRequest) error {
	panic("unimplemented")
}

func (s *TaskStore) CompleteTasksLessThan(ctx context.Context, request *persistence.CompleteTasksLessThanRequest) (int, error) {
	panic("unimplemented")
}
