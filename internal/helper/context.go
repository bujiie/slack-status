package helper

import (
	"context"
)

func AddValues(ctx context.Context, kvPairs map[any]any) context.Context {
	localCtx := ctx
	for key, value := range kvPairs {
		localCtx = context.WithValue(localCtx, key, value)
	}
	return localCtx
}
