package leafGoMongo

import (
	"context"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
	leafNoSql "github.com/paulusrobin/leaf-utilities/database/nosql/nosql"
	"go.mongodb.org/mongo-driver/bson"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	databaseImplementation struct {
		database *mgo.Database
	}
)

func NewDatabase(database *mgo.Database) leafNoSql.Database {
	return &databaseImplementation{database: database}
}

func (d *databaseImplementation) Drop(ctx context.Context) error {
	return d.database.Drop(ctx)
}

func (d *databaseImplementation) CreateCollection(ctx context.Context, name string) error {
	return d.database.CreateCollection(ctx, name)
}

func (d *databaseImplementation) Collection(name string, opts ...*options.CollectionOptions) leafNoSql.Collection {
	return NewCollection(d.database.Collection(name, opts...))
}

func (d *databaseImplementation) HasCollection(ctx context.Context, name string) bool {
	collections, err := d.ListCollectionNames(ctx)
	if err != nil {
		return false
	}

	if len(collections) < 1 {
		return false
	}

	if idx := leafFunctions.IndexString(collections, name); idx == -1 {
		return false
	}
	return true
}

func (d *databaseImplementation) ListCollectionNames(ctx context.Context) ([]string, error) {
	return d.database.ListCollectionNames(ctx, bson.D{})
}
