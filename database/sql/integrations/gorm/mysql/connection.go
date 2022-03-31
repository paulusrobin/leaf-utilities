package leafGormMySql

import (
	"fmt"
	"strings"
)

type DbConnection struct {
	Address  []string
	Username string
	Password string
	DbName   string
}

func (db DbConnection) URI() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		db.Username, db.Password, strings.Join(db.Address, ","), db.DbName)
}
