package bootstrap

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/pkg/contract"
	// sqlitemeasurementrepository
	// "github.com/bernardolm/replications-publisher-go/pkg/adapter/repository/measurement/sqlite"
)

type ReplicationRepository struct {
	sqliteClient contract.SQLiteClient
}

func (ReplicationRepository) Do(context.Context) error { return nil }

func ProvideReplicationRepository(
	sqliteClient contract.SQLiteClient,
) contract.ReplicationRepository {

	// replicationRepository, err := sqlitereplicationrepository.New(ctx, sqliteClient)
	// if err != nil {
	// 	_ = sqliteClient.Close()
	// 	logger.Log.WithError(err).Fatal("cmd: failed to initialize SQLite replication repository")
	// }

	obj := ReplicationRepository{
		sqliteClient: sqliteClient,
	}
	return &obj
}

// --------------

type ReplicationWorker struct {
	logger                contract.Logger
	replicationRepository contract.ReplicationRepository
}

func (ReplicationWorker) Do(context.Context) error { return nil }

func ProvideReplicationWorker(logger contract.Logger, rep contract.Repository) contract.Worker {
	obj := ReplicationWorker{logger: logger, replicationRepository: rep}
	return &obj
}
