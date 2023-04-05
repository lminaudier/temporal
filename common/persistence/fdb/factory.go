package fdb

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"go.temporal.io/server/common/config"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/common/persistence"
	"go.temporal.io/server/common/persistence/client"
	"go.temporal.io/server/common/persistence/fdb/storage"
	"go.temporal.io/server/common/resolver"
)

type (
	AbstractFactory struct{}

	Factory struct {
		store *storage.Storage
	}
)

func NewAbstractFactory() client.AbstractDataStoreFactory {
	return &AbstractFactory{}
}

func (f *AbstractFactory) NewFactory(
	cfg config.CustomDatastoreConfig,
	r resolver.ServiceResolver,
	clusterName string,
	logger log.Logger,
	metricsHandler metrics.Handler,
) client.DataStoreFactory {
	fdb.MustAPIVersion(710)
	// TODO: use config to properly connect to fdb
	db := fdb.MustOpenDefault()

	dir, err := directory.CreateOrOpen(db, []string{"default"}, nil)
	if err != nil {
		// TODO: see if we can avoid panicing
		panic(err)
	}

	return &Factory{
		store: storage.NewStorage(db, dir),
	}
}

func (f *Factory) Close() {}

func (f *Factory) NewTaskStore() (persistence.TaskStore, error) {
	return &TaskStore{}, nil
}

func (f *Factory) NewShardStore() (persistence.ShardStore, error) {
	return &ShardStore{
		store: f.store,
	}, nil
}

func (f *Factory) NewMetadataStore() (persistence.MetadataStore, error) {
	return &MetadataStore{
		store: f.store,
	}, nil
}

func (f *Factory) NewExecutionStore() (persistence.ExecutionStore, error) {
	return &ExecutionStore{}, nil
}

func (f *Factory) NewQueue(queueType persistence.QueueType) (persistence.Queue, error) {
	return &Queue{
		QueueType: queueType,
		store:     f.store,
	}, nil
}

func (f *Factory) NewClusterMetadataStore() (persistence.ClusterMetadataStore, error) {
	return &ClusterMetadataStore{
		store: f.store,
	}, nil
}
