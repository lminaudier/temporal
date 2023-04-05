package fdb

import (
	"context"
	"math/big"

	"go.temporal.io/server/common/persistence"
	"go.temporal.io/server/common/persistence/fdb/storage"
)

type ShardStore struct {
	store *storage.Storage
}

var _ persistence.ShardStore = (*ShardStore)(nil)

func (s *ShardStore) Close() {}

func (s *ShardStore) GetName() string {
	panic("unimplemented")
}

func (s *ShardStore) GetClusterName() string {
	panic("unimplemented")
}

func (s *ShardStore) GetOrCreateShard(ctx context.Context, request *persistence.InternalGetOrCreateShardRequest) (*persistence.InternalGetOrCreateShardResponse, error) {
	// Get section
	existing, err := s.store.Shard.Find(int64(request.ShardID))
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return &persistence.InternalGetOrCreateShardResponse{
			ShardInfo: persistence.NewDataBlob(existing.Data, existing.DataEncoding),
		}, nil
	}

	// Create section
	rangeID, shardInfo, err := request.CreateShardInfo()
	if err != nil {
		return nil, err
	}

	err = s.store.Shard.Save(&storage.Shard{
		ShardID:      int64(request.ShardID),
		RangeID:      big.NewInt(rangeID),
		Data:         shardInfo.Data,
		DataEncoding: shardInfo.EncodingType.String(),
	})
	if err != nil {
		return nil, err
	}

	return &persistence.InternalGetOrCreateShardResponse{
		ShardInfo: shardInfo,
	}, nil
}

func (s *ShardStore) UpdateShard(ctx context.Context, request *persistence.InternalUpdateShardRequest) error {
	panic("unimplemented")
}

func (s *ShardStore) AssertShardOwnership(ctx context.Context, request *persistence.AssertShardOwnershipRequest) error {
	panic("unimplemented")
}
