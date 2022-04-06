package leafNoSql

import (
	"context"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type noopCollection struct {
}

func (n noopCollection) Drop(ctx context.Context) error {
	return nil
}

func (n noopCollection) Indexes() IndexView {
	return NoopIndexView()
}

func (n noopCollection) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mgo.Cursor, error) {
	return &mgo.Cursor{}, nil
}

func (n noopCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mgo.Cursor, error) {
	return &mgo.Cursor{}, nil
}

func (n noopCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mgo.SingleResult {
	return &mgo.SingleResult{}
}

func (n noopCollection) BulkWrite(ctx context.Context, models []mgo.WriteModel, opts ...*options.BulkWriteOptions) (*mgo.BulkWriteResult, error) {
	return &mgo.BulkWriteResult{}, nil
}

func (n noopCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return 0, nil
}

func (n noopCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mgo.DeleteResult, error) {
	return &mgo.DeleteResult{}, nil
}

func (n noopCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mgo.DeleteResult, error) {
	return &mgo.DeleteResult{}, nil
}

func (n noopCollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mgo.UpdateResult, error) {
	return &mgo.UpdateResult{}, nil
}

func (n noopCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mgo.UpdateResult, error) {
	return &mgo.UpdateResult{}, nil
}

func (n noopCollection) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mgo.InsertManyResult, error) {
	return &mgo.InsertManyResult{}, nil
}

func (n noopCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mgo.InsertOneResult, error) {
	return &mgo.InsertOneResult{}, nil
}

func (n noopCollection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mgo.SingleResult {
	return &mgo.SingleResult{}
}

func (n noopCollection) FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mgo.SingleResult {
	return &mgo.SingleResult{}
}

func (n noopCollection) Distinct(ctx context.Context, field string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	return []interface{}{}, nil
}

func NoopCollection() Collection {
	return &noopCollection{}
}
