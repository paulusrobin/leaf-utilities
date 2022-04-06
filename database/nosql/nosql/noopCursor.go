package leafNoSql

import (
	"context"
)

type noopCursor struct{}

func (n noopCursor) Next(ctx context.Context) bool {
	return false
}

func (n noopCursor) Close(ctx context.Context) error {
	return nil
}

func (n noopCursor) Decode(val interface{}) error {
	return nil
}

func NoopCursor() Cursor {
	return &noopCursor{}
}
