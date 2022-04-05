package leafSql

import (
	"context"
	"database/sql"
	leafModel "github.com/paulusrobin/leaf-utilities/common/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type (
	SqlQueryable interface {
		TableName() string
		Fields() []string
		SqlQuery() (string, []interface{})
	}
	PaginateOptions struct {
		// Paging
		Page  int
		Limit int
		Sort  []string

		// Mapping
		FieldMap     map[string]string
		MapOrDefault bool

		// Filter
		Filter interface{}
	}
)

type (
	Config interface {
		// GORM perform single create, update, delete operations in transactions by default to ensure database data integrity
		// You can disable it by setting `SkipDefaultTransaction` to true
		SkipDefaultTransaction() bool
		// NamingStrategy tables, columns naming strategy
		NamingStrategy() schema.Namer
		// FullSaveAssociations full save associations
		FullSaveAssociations() bool
		// Logger
		Logger() logger.Interface
		// NowFunc the function to be used when creating a new timestamp
		NowFunc() func() time.Time
		// DryRun generate sql without execute
		DryRun() bool
		// PrepareStmt executes the given query in cached statement
		PrepareStmt() bool
		// DisableAutomaticPing
		DisableAutomaticPing() bool
		// DisableForeignKeyConstraintWhenMigrating
		DisableForeignKeyConstraintWhenMigrating() bool
		// AllowGlobalUpdate allow global update
		AllowGlobalUpdate() bool
		// ClauseBuilders clause builder
		ClauseBuilders() map[string]clause.ClauseBuilder
		// ConnPool db conn pool
		ConnPool() gorm.ConnPool
		// Plugins registered plugins
		Plugins() map[string]gorm.Plugin
	}

	API interface {
		// Session create new db session
		Session(config *gorm.Session) ORM

		// WithContext change current instance db's context to ctx
		WithContext(ctx context.Context) ORM

		// Debug start debug mode
		Debug() ORM

		// Set store value with key into current db instance's context
		Set(key string, value interface{}) ORM

		// Get get value with key from current db instance's context
		Get(key string) (interface{}, bool)

		// InstanceSet store value with key into current db instance's context
		InstanceSet(key string, value interface{}) ORM

		// InstanceGet get value with key from current db instance's context
		InstanceGet(key string) (interface{}, bool)

		// AddError add error to db
		AddError(err error) error

		// DB returns `*sql.DB`
		DB() (*sql.DB, error)

		SetupJoinTable(model interface{}, field string, joinTable interface{}) error

		Use(plugin gorm.Plugin) (err error)
	}

	ChainableAPI interface {
		Model(value interface{}) ORM

		// Clauses Add clauses
		Clauses(conds ...clause.Expression) ORM

		// Table specify the table you would like to run db operations
		Table(name string, args ...interface{}) ORM

		// Distinct specify distinct fields that you want querying
		Distinct(args ...interface{}) ORM

		// Select specify fields that you want when querying, creating, updating
		Select(query interface{}, args ...interface{}) ORM

		// Omit specify fields that you want to ignore when creating, updating and querying
		Omit(columns ...string) ORM

		// Where add conditions
		Where(query interface{}, args ...interface{}) ORM

		// Not add NOT conditions
		Not(query interface{}, args ...interface{}) ORM

		// Or add OR conditions
		Or(query interface{}, args ...interface{}) ORM

		// Joins specify Joins conditions
		//     db.Joins("Account").Find(&user)
		//     db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
		Joins(query string, args ...interface{}) ORM

		// Group specify the group method on the find
		Group(name string) ORM

		// Having specify HAVING conditions for GROUP BY
		Having(query interface{}, args ...interface{}) ORM

		// Order specify order when retrieve records from database
		//     db.Order("name DESC")
		//     db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: true})
		Order(value interface{}) ORM

		// Limit specify the number of records to be retrieved
		Limit(limit int) ORM

		// Offset specify the number of records to skip before starting to return the records
		Offset(offset int) ORM

		// Scopes pass current database connection to arguments `func(DB) DB`, which could be used to add conditions dynamically
		//     func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
		//         return db.Where("amount > ?", 1000)
		//     }
		//
		//     func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
		//         return func (db *gorm.DB) *gorm.DB {
		//             return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
		//         }
		//     }
		//
		//     db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
		Scopes(funcs ...func(db ORM) ORM) ORM

		// Preload preload associations with given conditions
		//    db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
		Preload(query string, args ...interface{}) ORM

		Attrs(attrs ...interface{}) ORM

		Assign(attrs ...interface{}) ORM

		Unscoped() ORM

		Raw(sql string, values ...interface{}) ORM
	}

	FinisherAPI interface {
		// Create insert the value into database
		Create(ctx context.Context, value interface{}) ORM

		// Save update value in database, if the value doesn't have primary key, will insert it
		Save(ctx context.Context,value interface{}) ORM

		// First find first record that match given conditions, order by primary key
		First(ctx context.Context, dest interface{}, conds ...interface{}) ORM

		// Take return a record that match given conditions, the order will depend on the database implementation
		Take(ctx context.Context, dest interface{}, conds ...interface{}) ORM

		// Last find last record that match given conditions, order by primary key
		Last(ctx context.Context, dest interface{}, conds ...interface{}) ORM

		// Find find records that match given conditions
		Find(ctx context.Context, dest interface{}, conds ...interface{}) ORM

		// FindInBatches find records in batches
		FindInBatches(ctx context.Context, dest interface{}, batchSize int, fc func(tx ORM, batch int) error) ORM

		FirstOrInit(ctx context.Context, dest interface{}, conds ...interface{}) ORM

		FirstOrCreate(ctx context.Context, dest interface{}, conds ...interface{}) ORM

		// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
		Update(ctx context.Context, column string, value interface{}) ORM

		// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
		Updates(ctx context.Context, values interface{}) ORM

		UpdateColumn(ctx context.Context, column string, value interface{}) ORM

		UpdateColumns(ctx context.Context, values interface{}) ORM

		// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
		Delete(ctx context.Context, value interface{}, conds ...interface{}) ORM

		Count(ctx context.Context, count *int64) ORM

		Row(ctx context.Context) *sql.Row

		Rows(ctx context.Context) (*sql.Rows, error)

		// Scan scan value to a struct
		Scan(ctx context.Context, dest interface{}) ORM

		// Pluck used to query single column from a model as a map
		//     var ages []int64
		//     db.Find(&users).Pluck("age", &ages)
		Pluck(ctx context.Context, column string, dest interface{}) ORM

		ScanRows(ctx context.Context, rows *sql.Rows, dest interface{}) error

		// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
		Transaction(ctx context.Context, fc func(ORM) error, opts ...*sql.TxOptions) (err error)

		// Begin begins a transaction
		Begin(ctx context.Context, opts ...*sql.TxOptions) ORM

		// Commit commit a transaction
		Commit(ctx context.Context) ORM

		// Rollback rollback a transaction
		Rollback(ctx context.Context) ORM

		SavePoint(ctx context.Context, name string) ORM

		RollbackTo(ctx context.Context, name string) ORM

		// Exec execute raw sql
		Exec(ctx context.Context, sql string, values ...interface{}) ORM

		BulkInsert(ctx context.Context, batches int, data ...SqlQueryable) error
	}

	PaginationAPI interface {
		PaginateData(ctx context.Context, items interface{}, options PaginateOptions) (leafModel.BasePagingResponse, error)
		SimplePaginateData(ctx context.Context, items interface{}, options PaginateOptions) (leafModel.BaseSimplePagingResponse, error)
	}

	ORM interface {
		API
		Config
		ChainableAPI
		FinisherAPI
		PaginationAPI

		Ping(ctx context.Context) error
		Gorm() *gorm.DB
		Dialector() gorm.Dialector
		Migrator() gorm.Migrator

		Error() error
		RowsAffected() int64

		AutoMigrate(dst ...interface{}) error
		Association(column string) *gorm.Association
		Statement() *gorm.Statement
	}
)
