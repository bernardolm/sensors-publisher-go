package queuer

import (
	"context"
)

type Queuer interface {
	Add(_ context.Context, key string, value []byte) error
	// ListPending(_ context.Context) ([][]byte, error)
	// MarkAsSent(_ context.Context, key string, when time.Time) error
}
