package leafGormPostgreSql

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
	return fmt.Sprintf(`postgres://%s:%s@%s/%s?sslmode=disable`,
		db.Username, db.Password, strings.Join(db.Address, ","), db.DbName)
}
