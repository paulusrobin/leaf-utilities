package leafNoSql

import (
	"context"
)

type (
	Cursor interface {
		Next(context.Context) bool
		Close(ctx context.Context) error
		Decode(val interface{}) error
	}
)
