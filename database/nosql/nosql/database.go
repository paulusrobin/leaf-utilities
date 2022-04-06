package leafNoSql

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Database interface {
		Drop(ctx context.Context) error
		CreateCollection(ctx context.Context, name string) error
		HasCollection(ctx context.Context, name string) bool
		ListCollectionNames(ctx context.Context) ([]string, error)
		Collection(name string, opts ...*options.CollectionOptions) Collection
	}
)
