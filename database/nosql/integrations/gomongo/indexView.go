package leafGoMongo

import (
	"context"
	leafNoSql "github.com/paulusrobin/leaf-utilities/database/nosql/nosql"
	"go.mongodb.org/mongo-driver/bson"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	indexViewImplementation struct {
		indexView mgo.IndexView
	}
)

func NewIndexView(indexView mgo.IndexView) leafNoSql.IndexView {
	return &indexViewImplementation{indexView: indexView}
}

func (i *indexViewImplementation) List(
	ctx context.Context,
	opts ...*options.ListIndexesOptions,
) (leafNoSql.Cursor, error) {
	return i.indexView.List(ctx, opts...)
}

func (i *indexViewImplementation) CreateMany(
	ctx context.Context,
	models []mgo.IndexModel,
	opts ...*options.CreateIndexesOptions,
) ([]string, error) {
	return i.indexView.CreateMany(ctx, models, opts...)
}

func (i *indexViewImplementation) CreateOne(
	ctx context.Context,
	model mgo.IndexModel,
	opts ...*options.CreateIndexesOptions,
) (string, error) {
	return i.indexView.CreateOne(ctx, model, opts...)
}

func (i *indexViewImplementation) DropOne(
	ctx context.Context,
	name string,
	opts ...*options.DropIndexesOptions,
) (bson.Raw, error) {
	return i.indexView.DropOne(ctx, name, opts...)
}

func (i *indexViewImplementation) DropAll(
	ctx context.Context,
	opts ...*options.DropIndexesOptions,
) (bson.Raw, error) {
	return i.indexView.DropAll(ctx, opts...)
}
