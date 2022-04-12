package mongo

import (
	leafGoMongo "github.com/paulusrobin/leaf-utilities/database/nosql/integrations/gomongo"
	leafNoSql "github.com/paulusrobin/leaf-utilities/database/nosql/nosql"
	"github.com/paulusrobin/leaf-utilities/leafMigration/config"
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
	"sync"
)

var (
	mongoConnection leafNoSql.Mongo
	once            sync.Once
)

func GetMongoConnection() (leafNoSql.Mongo, error) {
	var err error
	once.Do(func() {
		mongoConnection, err = leafGoMongo.New(leafGoMongo.WithURI(config.GetConfig().MongoUri),
			leafGoMongo.WithDatabaseName(config.GetConfig().MongoDatabase),
			leafGoMongo.WithLogger(logger.GetLogger()))
	})
	if err != nil {
		return nil, err
	}
	return mongoConnection, nil
}
