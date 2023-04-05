package fdb

import (
	"context"
	"fmt"
	"math/big"

	"go.temporal.io/api/serviceerror"
	"go.temporal.io/server/common/persistence"
	"go.temporal.io/server/common/persistence/fdb/storage"
)

type ClusterMetadataStore struct {
	store *storage.Storage
}

var _ persistence.ClusterMetadataStore = (*ClusterMetadataStore)(nil)

func (s *ClusterMetadataStore) Close() {}

func (s *ClusterMetadataStore) GetName() string {
	panic("unimplemented")
}

func (s *ClusterMetadataStore) ListClusterMetadata(ctx context.Context, request *persistence.InternalListClusterMetadataRequest) (*persistence.InternalListClusterMetadataResponse, error) {
	items, err := s.store.ClusterMetadataInfo.List()
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, serviceerror.NewNotFound("cluster metadata list is empty")
	}

	var ms []*persistence.InternalGetClusterMetadataResponse
	for _, m := range items {
		ms = append(ms, &persistence.InternalGetClusterMetadataResponse{
			ClusterMetadata: persistence.NewDataBlob(m.Data, m.DataEncoding),
			Version:         m.Version.Int64(),
		})
	}

	return &persistence.InternalListClusterMetadataResponse{
		ClusterMetadata: ms,
	}, nil
}

const DefaultMetadataPartition int64 = 0

func (s *ClusterMetadataStore) GetClusterMetadata(ctx context.Context, request *persistence.InternalGetClusterMetadataRequest) (*persistence.InternalGetClusterMetadataResponse, error) {
	m, err := s.store.ClusterMetadataInfo.Find(DefaultMetadataPartition, request.ClusterName)
	if err != nil {
		return nil, serviceerror.NewNotFound(fmt.Sprintf("cluster metadata not found for cluster %q (metadata partition: %q)", request.ClusterName, DefaultMetadataPartition))
	}
	if m == nil {
		return nil, serviceerror.NewNotFound(fmt.Sprintf("cluster metadata not found for cluster %q (metadata partition: %q)", request.ClusterName, DefaultMetadataPartition))
	}
	// FIXME: implement a better way to check for existence of cluster metadata in the database
	if m.Data == nil {
		return nil, serviceerror.NewNotFound(fmt.Sprintf("cluster metadata not found for cluster %q (metadata partition: %q)", request.ClusterName, DefaultMetadataPartition))
	}

	return &persistence.InternalGetClusterMetadataResponse{
		ClusterMetadata: persistence.NewDataBlob(m.Data, m.DataEncoding),
		Version:         m.Version.Int64(),
	}, nil
}

func (s *ClusterMetadataStore) SaveClusterMetadata(ctx context.Context, request *persistence.InternalSaveClusterMetadataRequest) (bool, error) {
	err := s.store.ClusterMetadataInfo.Save(&storage.ClusterMetadataInfo{
		MetadataPartition: DefaultMetadataPartition,
		ClusterName:       request.ClusterName,
		Data:              request.ClusterMetadata.Data,
		DataEncoding:      request.ClusterMetadata.EncodingType.String(),
		Version:           big.NewInt(request.Version),
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *ClusterMetadataStore) DeleteClusterMetadata(ctx context.Context, request *persistence.InternalDeleteClusterMetadataRequest) error {
	panic("unimplemented")
}

// Membership APIs
func (s *ClusterMetadataStore) GetClusterMembers(ctx context.Context, request *persistence.GetClusterMembersRequest) (*persistence.GetClusterMembersResponse, error) {
	panic("unimplemented")
}

const DefaultMembershipPartition int64 = 0

func (s *ClusterMetadataStore) UpsertClusterMembership(ctx context.Context, request *persistence.UpsertClusterMembershipRequest) error {
	err := s.store.ClusterMembership.Save(&storage.ClusterMembership{
		MembershipPartition: DefaultMembershipPartition,
		HostID:              []byte(request.HostID),
		RPCAddress:          request.RPCAddress.String(),
		RPCPort:             int8(request.RPCPort),
		Role:                int8(request.Role),
		SessionStart:        request.SessionStart,
		RecordExpiry:        request.RecordExpiry,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *ClusterMetadataStore) PruneClusterMembership(ctx context.Context, request *persistence.PruneClusterMembershipRequest) error {
	panic("unimplemented")
}
