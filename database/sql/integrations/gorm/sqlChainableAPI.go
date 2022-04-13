package leafGorm

import (
	leafSql "github.com/paulusrobin/leaf-utilities/database/sql/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (i *Impl) Model(value interface{}) leafSql.ORM {
	db := i.GormDB.Model(value)
	return i.newImpl(db)
}

// Clauses Add clauses
func (i *Impl) Clauses(conds ...clause.Expression) leafSql.ORM {
	db := i.GormDB.Clauses(conds...)
	return i.newImpl(db)
}

// Table specify the table you would like to run db operations
func (i *Impl) Table(name string, args ...interface{}) leafSql.ORM {
	db := i.GormDB.Table(name, args...)
	return i.newImpl(db)
}

// Distinct specify distinct fields that you want querying
func (i *Impl) Distinct(args ...interface{}) leafSql.ORM {
	db := i.GormDB.Distinct(args...)
	return i.newImpl(db)
}

// Select specify fields that you want when querying, creating, updating
func (i *Impl) Select(query interface{}, args ...interface{}) leafSql.ORM {
	db := i.GormDB.Select(query, args...)
	return i.newImpl(db)
}

// Omit specify fields that you want to ignore when creating, updating and querying
func (i *Impl) Omit(columns ...string) leafSql.ORM {
	db := i.GormDB.Omit(columns...)
	return i.newImpl(db)
}

// Where add conditions
func (i *Impl) Where(query interface{}, args ...interface{}) leafSql.ORM {
	db := i.GormDB.Where(query, args...)
	return i.newImpl(db)
}

// Not add NOT conditions
func (i *Impl) Not(query interface{}, args ...interface{}) leafSql.ORM {
	db := i.GormDB.Not(query, args...)
	return i.newImpl(db)
}

// Or add OR conditions
func (i *Impl) Or(query interface{}, args ...interface{}) leafSql.ORM {
	db := i.GormDB.Or(query, args...)
	return i.newImpl(db)
}

// Joins specify Joins conditions
//     db.Joins("Account").Find(&user)
//     db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
func (i *Impl) Joins(query string, args ...interface{}) leafSql.ORM {
	db := i.GormDB.Joins(query, args...)
	return i.newImpl(db)
}

// Group specify the group method on the find
func (i *Impl) Group(name string) leafSql.ORM {
	db := i.GormDB.Group(name)
	return i.newImpl(db)
}

// Having specify HAVING conditions for GROUP BY
func (i *Impl) Having(query interface{}, args ...interface{}) leafSql.ORM {
	db := i.GormDB.Having(query, args...)
	return i.newImpl(db)
}

// Order specify order when retrieve records from database
//     db.Order("name DESC")
//     db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: true})
func (i *Impl) Order(value interface{}) leafSql.ORM {
	db := i.GormDB.Order(value)
	return i.newImpl(db)
}

// Limit specify the number of records to be retrieved
func (i *Impl) Limit(limit int) leafSql.ORM {
	db := i.GormDB.Limit(limit)
	return i.newImpl(db)
}

// Offset specify the number of records to skip before starting to return the records
func (i *Impl) Offset(offset int) leafSql.ORM {
	db := i.GormDB.Offset(offset)
	return i.newImpl(db)
}

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
func (i *Impl) Scopes(funcs ...func(db leafSql.ORM) leafSql.ORM) leafSql.ORM {
	lfuncs := len(funcs)
	nfuncs := make([]func(db *gorm.DB) *gorm.DB, lfuncs)
	for idx := 0; idx < lfuncs; idx++ {
		nfuncs[idx] = func(db *gorm.DB) *gorm.DB {
			return funcs[idx](i).Gorm()
		}
	}

	db := i.GormDB.Scopes(nfuncs...)
	return i.newImpl(db)
}

// Preload preload associations with given conditions
//    db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
func (i *Impl) Preload(query string, args ...interface{}) leafSql.ORM {
	db := i.GormDB.Preload(query, args...)
	return i.newImpl(db)
}

func (i *Impl) Attrs(attrs ...interface{}) leafSql.ORM {
	db := i.GormDB.Attrs(attrs)
	return i.newImpl(db)
}

func (i *Impl) Assign(attrs ...interface{}) leafSql.ORM {
	db := i.GormDB.Assign(attrs...)
	return i.newImpl(db)
}

func (i *Impl) Unscoped() leafSql.ORM {
	db := i.GormDB.Unscoped()
	return i.newImpl(db)
}

func (i *Impl) Raw(sql string, values ...interface{}) leafSql.ORM {
	db := i.GormDB.Raw(sql, values...)
	return i.newImpl(db)
}
