package database

import (
	"fmt"

	"database/sql"

	"github.com/anemiq/anemiq/config"
)

type Database struct {
	db *sql.DB
}

func (self *Database) Table(name string) (*Table, error) {
	return newTable(self, name)
}

func Open(conn config.Conn) (*Database, error) {
	db, err := sql.Open("mysql", buildDataSourceName(conn))
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func (self *Database) Close() {
	self.db.Close()
}

func buildDataSourceName(conn config.Conn) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conn.User, conn.Pass, conn.Host, conn.Port, conn.Database)
}
