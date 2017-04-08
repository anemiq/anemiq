package database

import (
	"fmt"

	"github.com/anemiq/config"
)

type Database interface {
	Table(name string) Table
}

type Table struct {
	Name string
	Cols []Column
}

type Column struct {
	Name string
	Type string
}

func Open(conn config.Conn) (*Database, error) {
	panic("not implemented yet")
}

func buildDataSourceName(conn config.Conn) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conn.User, conn.Pass, conn.Host, conn.Port, conn.Database)
}
