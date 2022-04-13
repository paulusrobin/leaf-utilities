package leafGorm

import (
	"context"
	"database/sql"
	"fmt"
	leafSql "github.com/paulusrobin/leaf-utilities/database/sql/sql"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
	"strings"
)

// Create insert the value into database
func (i *Impl) Create(ctx context.Context, value interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "CREATE", i.Statement())
	db := i.GormDB.Create(value)
	span.Finish()
	return i.newImpl(db)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (i *Impl) Save(ctx context.Context, value interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "SAVE", i.Statement())
	db := i.GormDB.Save(value)
	span.Finish()
	return i.newImpl(db)
}

// First find first record that match given conditions, order by primary key
func (i *Impl) First(ctx context.Context, dest interface{}, conds ...interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "FIRST", i.Statement())
	db := i.GormDB.First(dest, conds...)
	span.Finish()
	return i.newImpl(db)
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (i *Impl) Take(ctx context.Context, dest interface{}, conds ...interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "TAKE", i.Statement())
	db := i.GormDB.Take(dest, conds...)
	span.Finish()
	return i.newImpl(db)
}

// Last find last record that match given conditions, order by primary key
func (i *Impl) Last(ctx context.Context, dest interface{}, conds ...interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "LAST", i.Statement())
	db := i.GormDB.Last(dest, conds...)
	span.Finish()
	return i.newImpl(db)
}

// Find find records that match given conditions
func (i *Impl) Find(ctx context.Context, dest interface{}, conds ...interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "FIND", i.Statement())
	db := i.GormDB.Find(dest, conds...)
	span.Finish()
	return i.newImpl(db)
}

// FindInBatches find records in batches
func (i *Impl) FindInBatches(ctx context.Context, dest interface{}, batchSize int, fc func(tx leafSql.ORM, batch int) error) leafSql.ORM {
	nfc := func(tx *gorm.DB, batch int) error {
		return fc(i, batchSize)
	}

	span := i.startDatastoreSegment(&ctx, "FIND_IN_BATCHES", i.Statement())
	db := i.GormDB.FindInBatches(dest, batchSize, nfc)
	span.Finish()
	return i.newImpl(db)
}

func (i *Impl) FirstOrInit(ctx context.Context, dest interface{}, conds ...interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "FIRST_OR_INIT", i.Statement())
	db := i.GormDB.FirstOrInit(dest, conds...)
	span.Finish()
	return i.newImpl(db)
}

func (i *Impl) FirstOrCreate(ctx context.Context, dest interface{}, conds ...interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "FIRST_OR_CREATE", i.Statement())
	db := i.GormDB.FirstOrCreate(dest, conds...)
	span.Finish()
	return i.newImpl(db)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (i *Impl) Update(ctx context.Context, column string, value interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "UPDATE", i.Statement())
	db := i.GormDB.Update(column, value)
	span.Finish()
	return i.newImpl(db)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (i *Impl) Updates(ctx context.Context, values interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "UPDATES", i.Statement())
	db := i.GormDB.Updates(values)
	span.Finish()
	return i.newImpl(db)
}

func (i *Impl) UpdateColumn(ctx context.Context, column string, value interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "UPDATE_COLUMN", i.Statement())
	db := i.GormDB.UpdateColumn(column, value)
	span.Finish()
	return i.newImpl(db)
}

func (i *Impl) UpdateColumns(ctx context.Context, values interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "UPDATE_COLUMNS", i.Statement())
	db := i.GormDB.UpdateColumns(values)
	span.Finish()
	return i.newImpl(db)
}

// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
func (i *Impl) Delete(ctx context.Context, value interface{}, conds ...interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "DELETE", i.Statement())
	db := i.GormDB.Delete(value, conds...)
	span.Finish()
	return i.newImpl(db)
}

func (i *Impl) Count(ctx context.Context, count *int64) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "COUNT", i.Statement())
	db := i.GormDB.Count(count)
	span.Finish()
	return i.newImpl(db)
}

func (i *Impl) Row(ctx context.Context) *sql.Row {
	return i.GormDB.Row()
}

func (i *Impl) Rows(ctx context.Context) (*sql.Rows, error) {
	sqlRows, err := i.GormDB.Rows()
	return sqlRows, err
}

// Scan scan value to a struct
func (i *Impl) Scan(ctx context.Context, dest interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "SCAN", i.Statement())
	db := i.GormDB.Scan(dest)
	span.Finish()
	return i.newImpl(db)
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Find(&users).Pluck("age", &ages)
func (i *Impl) Pluck(ctx context.Context, column string, dest interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "PLUCK", i.Statement())
	db := i.GormDB.Pluck(column, dest)
	span.Finish()
	return i.newImpl(db)
}

func (i *Impl) ScanRows(ctx context.Context, rows *sql.Rows, dest interface{}) error {
	return i.GormDB.ScanRows(rows, dest)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (i *Impl) Transaction(ctx context.Context, fc func(leafSql.ORM) error, opts ...*sql.TxOptions) error {
	nfc := func(db *gorm.DB) error {
		return fc(i)
	}
	return i.GormDB.Transaction(nfc, opts...)
}

// Begin begins a transaction
func (i *Impl) Begin(ctx context.Context, opts ...*sql.TxOptions) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "BEGIN", i.Statement())
	db := i.GormDB.Begin(opts...)
	span.Finish()
	return i.newImpl(db)
}

// Commit commit a transaction
func (i *Impl) Commit(ctx context.Context) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "COMMIT", i.Statement())
	db := i.GormDB.Commit()
	span.Finish()
	return i.newImpl(db)
}

// Rollback rollback a transaction
func (i *Impl) Rollback(ctx context.Context) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "ROLLBACK", i.Statement())
	db := i.GormDB.Rollback()
	span.Finish()
	return i.newImpl(db)
}

func (i *Impl) SavePoint(ctx context.Context, name string) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "SAVE_POINT", i.Statement())
	db := i.GormDB.SavePoint(name)
	span.Finish()
	return i.newImpl(db)
}

func (i *Impl) RollbackTo(ctx context.Context, name string) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "ROLLBACK_TO", i.Statement())
	db := i.GormDB.RollbackTo(name)
	span.Finish()
	return i.newImpl(db)
}

// Exec execute raw sql
func (i *Impl) Exec(ctx context.Context, sql string, values ...interface{}) leafSql.ORM {
	span := i.startDatastoreSegment(&ctx, "EXEC", i.Statement())
	db := i.GormDB.Exec(sql, values...)
	span.Finish()
	return i.newImpl(db)
}

func (i *Impl) BulkInsert(ctx context.Context, batches int, data ...leafSql.SqlQueryable) error {
	tx := i.Begin(ctx)
	chunkList := funk.Chunk(data, batches)
	for _, chunk := range chunkList.([][]leafSql.SqlQueryable) {
		tableName := ""
		valueFields := make([]string, 0)
		valueStrings := make([]string, 0)
		valueArgs := make([]interface{}, 0)

		for i, c := range chunk {
			if i == 0 {
				tableName = chunk[i].TableName()
				valueFields = c.Fields()
			}
			queryString, queryArgs := c.SqlQuery()
			valueStrings = append(valueStrings, queryString)
			valueArgs = append(valueArgs, queryArgs...)
		}

		if tableName == "" {
			tx.Rollback(ctx)
			return fmt.Errorf("table name is required")
		}

		stmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s",
			tableName,
			strings.Join(valueFields, ","),
			strings.Join(valueStrings, ","))
		err := tx.Exec(ctx, stmt, valueArgs...).Error()
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	return tx.Commit(ctx).Error()
}
