package leafNoSql

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	IndexView interface {
		List(ctx context.Context, opts ...*options.ListIndexesOptions) (Cursor, error)
		CreateMany(ctx context.Context, models []mgo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error)
		CreateOne(ctx context.Context, model mgo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error)
		DropOne(ctx context.Context, name string, opts ...*options.DropIndexesOptions) (bson.Raw, error)
		DropAll(ctx context.Context, opts ...*options.DropIndexesOptions) (bson.Raw, error)
	}
)
