package leafNoSql

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type noopDatabase struct {
}

func (n noopDatabase) Drop(ctx context.Context) error {
	return nil
}

func (n noopDatabase) CreateCollection(ctx context.Context, name string) error {
	return nil
}

func (n noopDatabase) HasCollection(ctx context.Context, name string) bool {
	return false
}

func (n noopDatabase) ListCollectionNames(ctx context.Context) ([]string, error) {
	return make([]string, 0), nil
}

func (n noopDatabase) Collection(name string, opts ...*options.CollectionOptions) Collection {
	return NoopCollection()
}

func NoopDatabase() Database {
	return &noopDatabase{}
}
