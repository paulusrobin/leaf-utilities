package leafGoMongo

import "context"

type dataStoreParam struct {
	databaseName       string
	operationName      string
	collectionName     string
	parameterizedQuery string
	queryParameters    []interface{}
}

func startDataStoreSpan(ctx *context.Context, param dataStoreParam) {

}
