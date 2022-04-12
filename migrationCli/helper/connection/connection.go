package connection

import (
	"fmt"
	"github.com/pkg/errors"
)

const (
	MySQL   = "mysql"
	Postgre = "postgre"
	Mongo   = "mongo"
)

func CheckConnection(connections []string) error {
	for _, connection := range connections {
		if !IsValid(connection) {
			return errors.New(fmt.Sprintf("%s is not valid", connection))
		}
	}
	return nil
}

func IsValid(conn string) bool {
	return MySQL == conn || Postgre == conn || Mongo == conn
}

func IsMongo(conn string) bool {
	return Mongo == conn
}
