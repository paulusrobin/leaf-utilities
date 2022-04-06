package leafGoMongo

import (
	"context"
	leafNoSql "github.com/paulusrobin/leaf-utilities/database/nosql/nosql"
	mgo "go.mongodb.org/mongo-driver/mongo"
)

type (
	cursorImplementation struct {
		cursor *mgo.Cursor
	}
)

func NewCursor(curr *mgo.Cursor) (leafNoSql.Cursor, error) {
	return &cursorImplementation{cursor: curr}, nil
}

func (c *cursorImplementation) Next(ctx context.Context) bool {
	return c.cursor.Next(ctx)
}

func (c *cursorImplementation) Close(ctx context.Context) error {
	return c.cursor.Close(ctx)
}

func (c *cursorImplementation) Decode(val interface{}) error {
	return c.cursor.Decode(val)
}
