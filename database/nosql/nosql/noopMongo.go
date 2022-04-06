package leafNoSql

import (
	"context"
	leafModel "github.com/paulusrobin/leaf-utilities/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	noopMongo struct{}
)

func (n noopMongo) FindOne(ctx context.Context, collection string, filter, object interface{}, options ...*options.FindOneOptions) error {
	return nil
}

func (n noopMongo) FindAll(ctx context.Context, collection string, filter interface{}, results interface{}, options ...*options.FindOptions) error {
	return nil
}

func (n noopMongo) Find(ctx context.Context, collection string, filter interface{}, callback FindCallback, options ...*options.FindOptions) error {
	return nil
}

func (n noopMongo) FindOneAndDelete(ctx context.Context, collection string, filter interface{}, options ...*options.FindOneAndDeleteOptions) error {
	return nil
}

func (n noopMongo) FindOneAndUpdate(ctx context.Context, collection string, filter, object interface{}, options ...*options.FindOneAndUpdateOptions) error {
	return nil
}

func (n noopMongo) Insert(ctx context.Context, collection string, object interface{}, options ...*options.InsertOneOptions) (*primitive.ObjectID, error) {
	return nil, nil
}

func (n noopMongo) InsertMany(ctx context.Context, collection string, documents []interface{}, options ...*options.InsertManyOptions) ([]primitive.ObjectID, error) {
	return nil, nil
}

func (n noopMongo) Update(ctx context.Context, collection string, filter, object interface{}, options ...*options.UpdateOptions) error {
	return nil
}

func (n noopMongo) UpdateMany(ctx context.Context, collection string, filter, object interface{}, options ...*options.UpdateOptions) error {
	return nil
}

func (n noopMongo) DeleteMany(ctx context.Context, collection string, filter interface{}, options ...*options.DeleteOptions) error {
	return nil
}

func (n noopMongo) Delete(ctx context.Context, collection string, filter interface{}, options ...*options.DeleteOptions) error {
	return nil
}

func (n noopMongo) BulkDocument(ctx context.Context, collection string, data []mgo.WriteModel) error {
	return nil
}

func (n noopMongo) Count(ctx context.Context, collection string, opts ...*options.CountOptions) (int64, error) {
	return 0, nil
}

func (n noopMongo) CountWithFilter(ctx context.Context, collection string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return 0, nil
}

func (n noopMongo) Distinct(ctx context.Context, collection, field string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (n noopMongo) Aggregate(ctx context.Context, collection string, pipeline interface{}, callback FindCallback, options ...*options.AggregateOptions) error {
	return nil
}

func (n noopMongo) Paginate(ctx context.Context, items interface{}, options PaginateOptions) (leafModel.PagingResponse, error) {
	return leafModel.PagingResponse{}, nil
}

func (n noopMongo) Indexes(collection string) IndexView {
	return NoopIndexView()
}

func (n noopMongo) Client() *mgo.Client { return nil }

func (n noopMongo) DB() Database {
	return NoopDatabase()
}

func (n noopMongo) Ping(ctx context.Context) error { return nil }

func (n noopMongo) Close(ctx context.Context) error { return nil }

func NoopMongo() Mongo {
	return &noopMongo{}
}
