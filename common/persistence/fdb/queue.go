package fdb

import (
	"context"
	"math/big"

	commonpb "go.temporal.io/api/common/v1"
	"go.temporal.io/server/common/persistence"
	"go.temporal.io/server/common/persistence/fdb/storage"
)

type Queue struct {
	QueueType persistence.QueueType
	store     *storage.Storage
}

var _ persistence.Queue = (*Queue)(nil)

func (q *Queue) Close() {}

func (q *Queue) Init(ctx context.Context, blob *commonpb.DataBlob) error {
	qm, err := q.store.QueueMetadata.Find(int64(q.QueueType))
	if err != nil {
		return err
	}
	// FIXME: find a better way to check for not found error returned by Find above
	if qm.Data != nil {
		return nil
	}

	err = q.store.QueueMetadata.Save(&storage.QueueMetadata{
		QueueType:    int64(q.QueueType),
		Data:         blob.Data,
		DataEncoding: blob.EncodingType.String(),
		Version:      big.NewInt(0),
	})
	if err != nil {
		return err
	}

	return nil
}

func (q *Queue) EnqueueMessage(ctx context.Context, blob commonpb.DataBlob) error {
	panic("unimplemented")
}

func (q *Queue) ReadMessages(ctx context.Context, lastMessageID int64, maxCount int) ([]*persistence.QueueMessage, error) {
	panic("unimplemented")
}

func (q *Queue) DeleteMessagesBefore(ctx context.Context, messageID int64) error {
	panic("unimplemented")
}

func (q *Queue) UpdateAckLevel(ctx context.Context, metadata *persistence.InternalQueueMetadata) error {
	panic("unimplemented")
}

func (q *Queue) GetAckLevels(ctx context.Context) (*persistence.InternalQueueMetadata, error) {
	panic("unimplemented")
}

func (q *Queue) EnqueueMessageToDLQ(ctx context.Context, blob commonpb.DataBlob) (int64, error) {
	panic("unimplemented")
}

func (q *Queue) ReadMessagesFromDLQ(ctx context.Context, firstMessageID int64, lastMessageID int64, pageSize int, pageToken []byte) ([]*persistence.QueueMessage, []byte, error) {
	panic("unimplemented")
}

func (q *Queue) DeleteMessageFromDLQ(ctx context.Context, messageID int64) error {
	panic("unimplemented")
}

func (q *Queue) RangeDeleteMessagesFromDLQ(ctx context.Context, firstMessageID int64, lastMessageID int64) error {
	panic("unimplemented")
}

func (q *Queue) UpdateDLQAckLevel(ctx context.Context, metadata *persistence.InternalQueueMetadata) error {
	panic("unimplemented")
}

func (q *Queue) GetDLQAckLevels(ctx context.Context) (*persistence.InternalQueueMetadata, error) {
	panic("unimplemented")
}
