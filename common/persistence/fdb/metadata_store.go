package fdb

import (
	"context"

	"go.temporal.io/api/serviceerror"
	"go.temporal.io/server/common/persistence"
	"go.temporal.io/server/common/persistence/fdb/storage"
)

type MetadataStore struct {
	store *storage.Storage
}

var _ persistence.MetadataStore = (*MetadataStore)(nil)

func (s *MetadataStore) Close() {}

func (s *MetadataStore) GetName() string {
	panic("unimplemented")
}

const DefaultPartitionID int64 = 54321

func (s *MetadataStore) CreateNamespace(ctx context.Context, request *persistence.InternalCreateNamespaceRequest) (*persistence.CreateNamespaceResponse, error) {
	err := s.store.Namespace.Save(&storage.Namespace{
		PartitionID:  DefaultPartitionID,
		ID:           []byte(request.ID),
		Name:         request.Name,
		Data:         request.Namespace.Data,
		DataEncoding: request.Namespace.EncodingType.String(),
		IsGlobal:     request.IsGlobal,
	})
	if err != nil {
		return nil, err
	}

	return &persistence.CreateNamespaceResponse{
		ID: request.ID,
	}, nil
}

func (s *MetadataStore) GetNamespace(ctx context.Context, request *persistence.GetNamespaceRequest) (*persistence.InternalGetNamespaceResponse, error) {
	ns, err := s.store.Namespace.Find(DefaultPartitionID, []byte(request.ID))
	if err != nil {
		return nil, err
	}
	if ns == nil {
		return nil, serviceerror.NewNamespaceNotFound(request.ID)
	}

	return &persistence.InternalGetNamespaceResponse{
		Namespace:           persistence.NewDataBlob(ns.Data, ns.DataEncoding),
		IsGlobal:            ns.IsGlobal,
		NotificationVersion: ns.NotificationVersion.Int64(),
	}, nil
}

func (s *MetadataStore) UpdateNamespace(ctx context.Context, request *persistence.InternalUpdateNamespaceRequest) error {
	panic("unimplemented")
}

func (s *MetadataStore) RenameNamespace(ctx context.Context, request *persistence.InternalRenameNamespaceRequest) error {
	panic("unimplemented")
}

func (s *MetadataStore) DeleteNamespace(ctx context.Context, request *persistence.DeleteNamespaceRequest) error {
	panic("unimplemented")
}

func (s *MetadataStore) DeleteNamespaceByName(ctx context.Context, request *persistence.DeleteNamespaceByNameRequest) error {
	panic("unimplemented")
}

func (s *MetadataStore) ListNamespaces(ctx context.Context, request *persistence.InternalListNamespacesRequest) (*persistence.InternalListNamespacesResponse, error) {
	items, err := s.store.Namespace.List()
	if err != nil {
		return nil, err
	}

	resps := make([]*persistence.InternalGetNamespaceResponse, 0)
	for _, n := range items {
		resps = append(resps, &persistence.InternalGetNamespaceResponse{
			Namespace:           persistence.NewDataBlob(n.Data, n.DataEncoding),
			IsGlobal:            n.IsGlobal,
			NotificationVersion: n.NotificationVersion.Int64(),
		})
	}

	return &persistence.InternalListNamespacesResponse{
		Namespaces: resps,
	}, nil
}

func (s *MetadataStore) GetMetadata(ctx context.Context) (*persistence.GetMetadataResponse, error) {
	panic("unimplemented")
}
