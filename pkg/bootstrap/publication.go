package bootstrap

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/pkg/contract"
	// sqlitemeasurementrepository
	// "github.com/bernardolm/publications-publisher-go/pkg/adapter/repository/measurement/sqlite"
)

type PublicationRepository struct {
	sqliteClient contract.SQLiteClient
}

func (PublicationRepository) Do(context.Context) error { return nil }

func ProvidePublicationRepository(
	sqliteClient contract.SQLiteClient,
) contract.PublicationRepository {

	// publicationRepository, err := sqlitepublicationrepository.New(ctx, sqliteClient)
	// if err != nil {
	// 	_ = sqliteClient.Close()
	// 	logger.Log.WithError(err).Fatal("cmd: failed to initialize SQLite publication repository")
	// }

	obj := PublicationRepository{
		sqliteClient: sqliteClient,
	}
	return &obj
}

// --------------

type PublicationWorker struct {
	logger                contract.Logger
	publicationRepository contract.PublicationRepository
}

func (PublicationWorker) Do(context.Context) error { return nil }

func ProvidePublicationWorker(logger contract.Logger, rep contract.Repository) contract.Worker {
	obj := PublicationWorker{logger: logger, publicationRepository: rep}
	return &obj
}
