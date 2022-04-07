package leafGoMongo

import (
	"context"
	"fmt"
	leafNoSql "github.com/paulusrobin/leaf-utilities/database/nosql/nosql"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"reflect"
	"time"
)

type (
	implementation struct {
		client       *mgo.Client
		database     leafNoSql.Database
		logger       leafLogger.Logger
		databaseName string
	}
	decoder struct {
	}
)

func (d decoder) DecodeValue(dCtx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Kind() != reflect.String {
		return fmt.Errorf("bad type or not settable")
	}
	var str string
	var err error
	switch vr.Type() {
	case bsontype.String:
		if str, err = vr.ReadString(); err != nil {
			return err
		}
	case bsontype.Null: // THIS IS THE MISSING PIECE TO HANDLE NULL!
		if err = vr.ReadNull(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot decode %v into a string type", vr.Type())
	}

	val.SetString(str)
	return nil
}

func New(opts ...Option) (leafNoSql.Mongo, error) {
	option := defaultOptions()
	for _, opt := range opts {
		opt.Apply(&option)
	}

	if option.uri == "" {
		return nil, fmt.Errorf("uri is required")
	}

	if option.databaseName == "" {
		return nil, fmt.Errorf("database name is required")
	}

	clientOptions := options.Client().
		ApplyURI(option.uri).
		SetRegistry(bson.NewRegistryBuilder().RegisterDecoder(reflect.TypeOf(""), decoder{}).Build())

	for _, clientOpt := range option.mongoOptions {
		clientOptions = clientOpt(clientOptions)
	}

	ctx := context.Background()
	client, err := mgo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %+v", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping mongo : %+v", err)
	}

	database := NewDatabase(client.Database(option.databaseName))

	return &implementation{client, database, option.logger, option.databaseName}, nil
}

func (i *implementation) Ping(ctx context.Context) error {
	ctxData, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	return i.client.Ping(ctxData, readpref.Primary())
}

func (i *implementation) Client() *mgo.Client {
	return i.client
}

func (i *implementation) Close(ctx context.Context) error {
	return i.client.Disconnect(ctx)
}

func (i *implementation) DB() leafNoSql.Database {
	return i.database
}

func (i *implementation) FindAll(ctx context.Context, collection string, filter interface{}, results interface{}, options ...*options.FindOptions) error {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "FindAll",
		collectionName:     collection,
		parameterizedQuery: "find all where $1",
		queryParameters:    []interface{}{filter},
	})
	defer span.Finish()

	rs, err := i.database.Collection(collection).Find(ctx, filter, options...)

	if err != nil {
		return fmt.Errorf("failed to find all with context : %+v", err)
	}

	if err := rs.All(ctx, results); err != nil {
		return fmt.Errorf("failed to decode all : %+v", err)
	}

	return nil
}

func (i *implementation) FindOne(ctx context.Context, collection string, filter, object interface{}, options ...*options.FindOneOptions) error {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "FindOne",
		collectionName:     collection,
		parameterizedQuery: "find one where $1",
		queryParameters:    []interface{}{filter},
	})
	defer span.Finish()

	sr := i.database.Collection(collection).FindOne(ctx, filter, options...)

	if err := sr.Err(); err != nil {
		return fmt.Errorf("FindOne failed: %+v", err)
	}

	if err := sr.Decode(object); err != nil {
		return fmt.Errorf("FindOne decode failed: %+v", err)
	}

	return nil
}

func (i *implementation) Find(ctx context.Context, collection string, filter interface{}, callback leafNoSql.FindCallback, options ...*options.FindOptions) error {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "Find",
		collectionName:     collection,
		parameterizedQuery: "find where $1",
		queryParameters:    []interface{}{filter},
	})
	defer span.Finish()

	coll, err := i.database.Collection(collection).Find(ctx, filter, options...)
	if err != nil {
		return err
	}

	cursor, err := NewCursor(coll)

	defer func() {
		if cursor == nil {
			return
		}

		if err := cursor.Close(ctx); err != nil {
			i.logger.StandardLogger().Errorf("failed to close cursor %s", err)
		}
	}()

	if err != nil {
		return callback(nil, err)
	} else {
		return callback(cursor, nil)
	}
}

func (i *implementation) FindOneAndDelete(ctx context.Context, collection string, filter interface{}, options ...*options.FindOneAndDeleteOptions) error {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "FindOneAndDelete",
		collectionName:     collection,
		parameterizedQuery: "find where $1, delete",
		queryParameters:    []interface{}{filter},
	})
	defer span.Finish()

	sr := i.database.Collection(collection).FindOneAndDelete(ctx, filter, options...)

	if err := sr.Err(); err != nil {
		return fmt.Errorf("FindOneAndDeleteWithContext failed: %+v", err)
	}

	return nil
}

func (i *implementation) FindOneAndUpdate(ctx context.Context, collection string, filter, object interface{}, options ...*options.FindOneAndUpdateOptions) error {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "FindOneAndUpdate",
		collectionName:     collection,
		parameterizedQuery: "find where $1, set $2",
		queryParameters:    []interface{}{filter, object},
	})
	defer span.Finish()

	sr := i.database.Collection(collection).FindOneAndUpdate(ctx, filter, object, options...)

	if err := sr.Err(); err != nil {
		return fmt.Errorf("FindOneAndUpdateWithContext failed: %+v", err)
	}

	if err := sr.Decode(&object); err != nil {
		return fmt.Errorf("FindOneAndUpdate decode failed: %+v", err)
	}

	return nil
}

func (i *implementation) Insert(ctx context.Context, collection string, object interface{}, options ...*options.InsertOneOptions) (*primitive.ObjectID, error) {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "Insert",
		collectionName:     collection,
		parameterizedQuery: "insert $1",
		queryParameters:    []interface{}{object},
	})
	defer span.Finish()

	ir, err := i.database.Collection(collection).InsertOne(ctx, object, options...)

	if err != nil {
		return nil, fmt.Errorf("InsertOneWithContext failed: %+v", err)
	}

	id, ok := ir.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, fmt.Errorf("InsertWithContext failed to cast ObjectID")
	}

	return &id, nil
}

func (i *implementation) InsertMany(ctx context.Context, collection string, documents []interface{}, options ...*options.InsertManyOptions) ([]primitive.ObjectID, error) {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "InsertMany",
		collectionName:     collection,
		parameterizedQuery: "insert $1",
		queryParameters:    []interface{}{documents},
	})
	defer span.Finish()

	ir, err := i.database.Collection(collection).InsertMany(ctx, documents, options...)

	if err != nil {
		return nil, fmt.Errorf("InsertManyWithContext failed: %+v", err)
	}

	ids := make([]primitive.ObjectID, 0)

	for _, id := range ir.InsertedIDs {
		i, ok := id.(primitive.ObjectID)

		if !ok {
			err = fmt.Errorf("InsertWithContext failed to cast ObjectID %s", i)
			break
		}

		ids = append(ids, i)
	}

	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (i *implementation) Update(ctx context.Context, collection string, filter, object interface{}, options ...*options.UpdateOptions) error {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "Update",
		collectionName:     collection,
		parameterizedQuery: "update where $1 set $2",
		queryParameters:    []interface{}{filter, object},
	})
	defer span.Finish()

	if _, err := i.database.Collection(collection).UpdateOne(ctx, filter, object, options...); err != nil {
		return fmt.Errorf("UpdateWithContext failed: %+v", err)
	}

	return nil
}

func (i *implementation) UpdateMany(ctx context.Context, collection string, filter, object interface{}, options ...*options.UpdateOptions) error {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "UpdateMany",
		collectionName:     collection,
		parameterizedQuery: "update many where $1 set $2",
		queryParameters:    []interface{}{filter, object},
	})
	defer span.Finish()

	if _, err := i.database.Collection(collection).UpdateMany(ctx, filter, object, options...); err != nil {
		return fmt.Errorf("UpdateManyWithContext failed: %+v", err)
	}

	return nil
}

func (i *implementation) DeleteMany(ctx context.Context, collection string, filter interface{}, options ...*options.DeleteOptions) error {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "DeleteMany",
		collectionName:     collection,
		parameterizedQuery: "delete many where $1",
		queryParameters:    []interface{}{filter},
	})
	defer span.Finish()

	if _, err := i.database.Collection(collection).DeleteMany(ctx, filter, options...); err != nil {
		return fmt.Errorf("DeleteManyWithContext failed: %+v", err)
	}

	return nil
}

func (i *implementation) Delete(ctx context.Context, collection string, filter interface{}, options ...*options.DeleteOptions) error {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "Delete",
		collectionName:     collection,
		parameterizedQuery: "delete one where $1",
		queryParameters:    []interface{}{filter},
	})
	defer span.Finish()

	if _, err := i.database.Collection(collection).DeleteOne(ctx, filter, options...); err != nil {
		return fmt.Errorf("DeleteWithContext failed: %+v", err)
	}

	return nil
}

func (i *implementation) countWithFilter(ctx context.Context, collection string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	coll := i.database.Collection(collection)
	total, err := coll.CountDocuments(ctx, filter, opts...)

	if err != nil {
		return 0, fmt.Errorf("count collection %s failed: %+v", collection, err)
	}

	return total, nil
}

func (i *implementation) CountWithFilter(ctx context.Context, collection string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "CountWithFilter",
		collectionName:     collection,
		parameterizedQuery: "count where $1",
		queryParameters:    []interface{}{filter},
	})
	defer span.Finish()
	return i.countWithFilter(ctx, collection, filter, opts...)
}

func (i *implementation) Distinct(ctx context.Context, collection, field string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "Distinct",
		collectionName:     collection,
		parameterizedQuery: "distinct",
		queryParameters:    []interface{}{field, filter},
	})
	defer span.Finish()

	object, err := i.database.Collection(collection).Distinct(ctx, field, filter, opts...)
	if err != nil {
		return nil, fmt.Errorf("Distinct failed: %+v", err)
	}

	return object, nil
}

func (i *implementation) Count(ctx context.Context, collection string, opts ...*options.CountOptions) (int64, error) {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "Count",
		collectionName:     collection,
		parameterizedQuery: "count",
		queryParameters:    []interface{}{},
	})
	defer span.Finish()
	return i.countWithFilter(ctx, collection, bson.D{}, opts...)
}

func (i *implementation) Indexes(collection string) leafNoSql.IndexView {
	return i.database.Collection(collection).Indexes()
}

func (i *implementation) BulkDocument(ctx context.Context, collection string, data []mgo.WriteModel) error {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "BulkDocument",
		collectionName:     collection,
		parameterizedQuery: "bulk document $1",
		queryParameters:    []interface{}{data},
	})
	defer span.Finish()

	_, err := i.database.Collection(collection).BulkWrite(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *implementation) Aggregate(ctx context.Context, collection string, pipeline interface{}, callback leafNoSql.FindCallback, options ...*options.AggregateOptions) error {
	span := startDataStoreSpan(&ctx, dataStoreParam{
		databaseName:       i.databaseName,
		operationName:      "Aggregate",
		collectionName:     collection,
		parameterizedQuery: "aggregate $1",
		queryParameters:    []interface{}{pipeline},
	})
	defer span.Finish()

	coll, err := i.database.Collection(collection).Aggregate(ctx, pipeline, options...)
	if err != nil {
		return err
	}

	cursor, err := NewCursor(coll)

	defer func() {
		if cursor == nil {
			return
		}

		if err := cursor.Close(ctx); err != nil {
			i.logger.StandardLogger().Errorf("failed to close cursor %s", err)
		}
	}()

	if err != nil {
		return callback(nil, err)
	} else {
		return callback(cursor, nil)
	}
}
