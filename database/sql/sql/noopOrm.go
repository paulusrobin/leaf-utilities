package leafSql

import (
	"context"
	"database/sql"
	"fmt"
	leafModel "github.com/paulusrobin/leaf-utilities/common/model"
	leafTime "github.com/paulusrobin/leaf-utilities/time"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type noopORM struct {}

func (n noopORM) Session(config *gorm.Session) ORM { return n }

func (n noopORM) WithContext(ctx context.Context) ORM { return n }

func (n noopORM) Debug() ORM { return n }

func (n noopORM) Set(key string, value interface{}) ORM { return n }

func (n noopORM) Get(key string) (interface{}, bool) {
	return nil, false
}

func (n noopORM) InstanceSet(key string, value interface{}) ORM { return n }

func (n noopORM) InstanceGet(key string) (interface{}, bool) {
	return nil, false
}

func (n noopORM) AddError(err error) error {
	return nil
}

func (n noopORM) DB() (*sql.DB, error) {
	return nil, fmt.Errorf("no db connection")
}

func (n noopORM) SetupJoinTable(model interface{}, field string, joinTable interface{}) error {
	return nil
}

func (n noopORM) Use(plugin gorm.Plugin) error {
	return nil
}

func (n noopORM) SkipDefaultTransaction() bool {
	return false
}

func (n noopORM) NamingStrategy() schema.Namer {
	return nil
}

func (n noopORM) FullSaveAssociations() bool {
	return false
}

func (n noopORM) Logger() logger.Interface {
	return nil
}

func (n noopORM) NowFunc() func() time.Time {
	return leafTime.Now
}

func (n noopORM) DryRun() bool {
	return false
}

func (n noopORM) PrepareStmt() bool {
	return false
}

func (n noopORM) DisableAutomaticPing() bool {
	return false
}

func (n noopORM) DisableForeignKeyConstraintWhenMigrating() bool {
	return false
}

func (n noopORM) AllowGlobalUpdate() bool {
	return false
}

func (n noopORM) ClauseBuilders() map[string]clause.ClauseBuilder {
	return make(map[string]clause.ClauseBuilder)
}

func (n noopORM) ConnPool() gorm.ConnPool {
	return nil
}

func (n noopORM) Plugins() map[string]gorm.Plugin {
	return make(map[string]gorm.Plugin)
}

func (n noopORM) Model(value interface{}) ORM { return n }

func (n noopORM) Clauses(conds ...clause.Expression) ORM { return n }

func (n noopORM) Table(name string, args ...interface{}) ORM { return n }

func (n noopORM) Distinct(args ...interface{}) ORM { return n }

func (n noopORM) Select(query interface{}, args ...interface{}) ORM { return n }

func (n noopORM) Omit(columns ...string) ORM { return n }

func (n noopORM) Where(query interface{}, args ...interface{}) ORM { return n }

func (n noopORM) Not(query interface{}, args ...interface{}) ORM { return n }

func (n noopORM) Or(query interface{}, args ...interface{}) ORM { return n }

func (n noopORM) Joins(query string, args ...interface{}) ORM { return n }

func (n noopORM) Group(name string) ORM { return n }

func (n noopORM) Having(query interface{}, args ...interface{}) ORM { return n }

func (n noopORM) Order(value interface{}) ORM { return n }

func (n noopORM) Limit(limit int) ORM { return n }

func (n noopORM) Offset(offset int) ORM { return n }

func (n noopORM) Scopes(funcs ...func(db ORM) ORM) ORM { return n }

func (n noopORM) Preload(query string, args ...interface{}) ORM { return n }

func (n noopORM) Attrs(attrs ...interface{}) ORM { return n }

func (n noopORM) Assign(attrs ...interface{}) ORM { return n }

func (n noopORM) Unscoped() ORM { return n }

func (n noopORM) Raw(sql string, values ...interface{}) ORM { return n }

func (n noopORM) Create(ctx context.Context, value interface{}) ORM { return n }

func (n noopORM) Save(ctx context.Context, value interface{}) ORM { return n }

func (n noopORM) First(ctx context.Context, dest interface{}, conds ...interface{}) ORM { return n }

func (n noopORM) Take(ctx context.Context, dest interface{}, conds ...interface{}) ORM { return n }

func (n noopORM) Last(ctx context.Context, dest interface{}, conds ...interface{}) ORM { return n }

func (n noopORM) Find(ctx context.Context, dest interface{}, conds ...interface{}) ORM { return n }

func (n noopORM) FindInBatches(ctx context.Context, dest interface{}, batchSize int, fc func(tx ORM, batch int) error) ORM { return n }

func (n noopORM) FirstOrInit(ctx context.Context, dest interface{}, conds ...interface{}) ORM { return n }

func (n noopORM) FirstOrCreate(ctx context.Context, dest interface{}, conds ...interface{}) ORM { return n }

func (n noopORM) Update(ctx context.Context, column string, value interface{}) ORM { return n }

func (n noopORM) Updates(ctx context.Context, values interface{}) ORM { return n }

func (n noopORM) UpdateColumn(ctx context.Context, column string, value interface{}) ORM { return n }

func (n noopORM) UpdateColumns(ctx context.Context, values interface{}) ORM { return n }

func (n noopORM) Delete(ctx context.Context, value interface{}, conds ...interface{}) ORM { return n }

func (n noopORM) Count(ctx context.Context, count *int64) ORM { return n }

func (n noopORM) Row(ctx context.Context) *sql.Row {
	return &sql.Row{}
}

func (n noopORM) Rows(ctx context.Context) (*sql.Rows, error) {
	return &sql.Rows{}, nil
}

func (n noopORM) Scan(ctx context.Context, dest interface{}) ORM { return n }

func (n noopORM) Pluck(ctx context.Context, column string, dest interface{}) ORM { return n }

func (n noopORM) ScanRows(ctx context.Context, rows *sql.Rows, dest interface{}) error {
	return nil
}

func (n noopORM) Transaction(ctx context.Context, fc func(ORM) error, opts ...*sql.TxOptions) error {
	return nil
}

func (n noopORM) Begin(ctx context.Context, opts ...*sql.TxOptions) ORM { return n }

func (n noopORM) Commit(ctx context.Context) ORM { return n }

func (n noopORM) Rollback(ctx context.Context) ORM { return n }

func (n noopORM) SavePoint(ctx context.Context, name string) ORM { return n }

func (n noopORM) RollbackTo(ctx context.Context, name string) ORM { return n }

func (n noopORM) Exec(ctx context.Context, sql string, values ...interface{}) ORM { return n }

func (n noopORM) BulkInsert(ctx context.Context, batches int, data ...SqlQueryable) error {
	return nil
}

func (n noopORM) PaginateData(ctx context.Context, items interface{}, options PaginateOptions) (leafModel.BasePagingResponse, error) {
	return leafModel.BasePagingResponse{}, nil
}

func (n noopORM) SimplePaginateData(ctx context.Context, items interface{}, options PaginateOptions) (leafModel.BaseSimplePagingResponse, error) {
	return leafModel.BaseSimplePagingResponse{}, nil
}

func (n noopORM) Ping(ctx context.Context) error {
	return nil
}

func (n noopORM) Gorm() *gorm.DB {
	return nil
}

func (n noopORM) Dialector() gorm.Dialector {
	return nil
}

func (n noopORM) Migrator() gorm.Migrator {
	return nil
}

func (n noopORM) Error() error {
	return nil
}

func (n noopORM) RowsAffected() int64 {
	return 0
}

func (n noopORM) AutoMigrate(dst ...interface{}) error {
	return nil
}

func (n noopORM) Association(column string) *gorm.Association {
	return nil
}

func (n noopORM) Statement() *gorm.Statement {
	return nil
}

func NoopORM() ORM { return &noopORM{} }
