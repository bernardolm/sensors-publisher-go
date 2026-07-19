package bootstrap

import "context"

type dummy struct{}

func (dummy) Do(context.Context) error { return nil }
