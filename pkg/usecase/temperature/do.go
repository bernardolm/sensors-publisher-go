package temperature

import "context"

func (uc usecase) Do(ctx context.Context) error {
	return uc.collect(ctx)
}
