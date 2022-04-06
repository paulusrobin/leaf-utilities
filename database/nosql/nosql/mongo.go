package leafNoSql

import (
	"context"
	leafModel "github.com/paulusrobin/leaf-utilities/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	PaginateOptions struct {
		// Paging
		Page  int
		Limit int
		Sort  []string

		// Database
		Collection string
		Filter     interface{}

		// Field Mapping
		FieldMap      map[string]string
		MapOrDefault  bool
		RequestFilter interface{}
	}

	FindCallback func(Cursor, error) error

	Mongo interface {
		FindOne(ctx context.Context, collection string, filter, object interface{}, options ...*options.FindOneOptions) error
		FindAll(ctx context.Context, collection string, filter interface{}, results interface{}, options ...*options.FindOptions) error
		Find(ctx context.Context, collection string, filter interface{}, callback FindCallback, options ...*options.FindOptions) error
		FindOneAndDelete(ctx context.Context, collection string, filter interface{}, options ...*options.FindOneAndDeleteOptions) error
		FindOneAndUpdate(ctx context.Context, collection string, filter, object interface{}, options ...*options.FindOneAndUpdateOptions) error
		Insert(ctx context.Context, collection string, object interface{}, options ...*options.InsertOneOptions) (*primitive.ObjectID, error)
		InsertMany(ctx context.Context, collection string, documents []interface{}, options ...*options.InsertManyOptions) ([]primitive.ObjectID, error)
		Update(ctx context.Context, collection string, filter, object interface{}, options ...*options.UpdateOptions) error
		UpdateMany(ctx context.Context, collection string, filter, object interface{}, options ...*options.UpdateOptions) error
		DeleteMany(ctx context.Context, collection string, filter interface{}, options ...*options.DeleteOptions) error
		Delete(ctx context.Context, collection string, filter interface{}, options ...*options.DeleteOptions) error
		BulkDocument(ctx context.Context, collection string, data []mgo.WriteModel) error
		Count(ctx context.Context, collection string, opts ...*options.CountOptions) (int64, error)
		CountWithFilter(ctx context.Context, collection string, filter interface{}, opts ...*options.CountOptions) (int64, error)
		Distinct(ctx context.Context, collection, field string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error)

		Aggregate(ctx context.Context, collection string, pipeline interface{}, callback FindCallback, options ...*options.AggregateOptions) error
		Paginate(ctx context.Context, items interface{}, options PaginateOptions) (leafModel.PagingResponse, error)

		Indexes(collection string) IndexView
		Client() *mgo.Client
		DB() Database
		Ping(ctx context.Context) error
		Close(ctx context.Context) error
	}
)
