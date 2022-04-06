package leafNoSql

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type noopIndexView struct {
}

func (n noopIndexView) List(ctx context.Context, opts ...*options.ListIndexesOptions) (Cursor, error) {
	return NoopCursor(), nil
}

func (n noopIndexView) CreateMany(ctx context.Context, models []mgo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	return make([]string, 0), nil
}

func (n noopIndexView) CreateOne(ctx context.Context, model mgo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	return "", nil
}

func (n noopIndexView) DropOne(ctx context.Context, name string, opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	return bson.Raw{}, nil
}

func (n noopIndexView) DropAll(ctx context.Context, opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	return bson.Raw{}, nil
}

func NoopIndexView() IndexView {
	return &noopIndexView{}
}
