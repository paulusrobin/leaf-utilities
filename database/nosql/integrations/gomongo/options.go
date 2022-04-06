package leafGoMongo

import (
	leafLogrus "github.com/paulusrobin/leaf-utilities/logger/integrations/logrus"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Option interface {
		Apply(o *mongoOptions)
	}
	ClientOptionBuilderFunc func(*options.ClientOptions) *options.ClientOptions
	mongoOptions            struct {
		uri          string
		databaseName string
		logger       leafLogger.Logger
		mongoOptions []ClientOptionBuilderFunc
	}
)

func defaultOptions() mongoOptions {
	return mongoOptions{
		uri:          "",
		databaseName: "",
		logger:       leafLogrus.DefaultLog(),
		mongoOptions: make([]ClientOptionBuilderFunc, 0),
	}
}

type withURI string

func (w withURI) Apply(o *mongoOptions) {
	o.uri = string(w)
}

func WithURI(uri string) Option {
	return withURI(uri)
}

type withDatabaseName string

func (w withDatabaseName) Apply(o *mongoOptions) {
	o.databaseName = string(w)
}

func WithDatabaseName(databaseName string) Option {
	return withDatabaseName(databaseName)
}

type withLogger struct{ leafLogger.Logger }

func (w withLogger) Apply(o *mongoOptions) {
	o.logger = w.Logger
}

func WithLogger(logger leafLogger.Logger) Option {
	return withLogger{logger}
}

type withMongoOptions struct{ ClientOptionBuilderFunc }

func (w withMongoOptions) Apply(o *mongoOptions) {
	o.mongoOptions = append(o.mongoOptions, w.ClientOptionBuilderFunc)
}

func WithMongoOptions(fn ClientOptionBuilderFunc) Option {
	return withMongoOptions{fn}
}
