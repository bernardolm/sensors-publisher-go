package publisher

import "context"

type Publisher interface {
	Publish(_ context.Context, content any) error
}
