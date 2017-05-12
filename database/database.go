package database

import (
	"fmt"

	"database/sql"

	"github.com/anemiq/anemiq/config"
)

type Database struct {
	db *sql.DB
}

func (self *Database) Tables(tablesNames []string) ([]*Table, error) {
	tables := make([]*Table, len(tablesNames))
	for i, tableName := range tablesNames {
		table, err := self.table(tableName)
		if err != nil {
			return nil, err
		}
		tables[i] = table
	}
	return tables, nil
}

func Open(conn config.DatabaseConn) (*Database, error) {
	db, err := sql.Open("mysql", buildDataSourceName(conn))
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func (self *Database) Close() {
	self.db.Close()
}

func (self *Database) table(name string) (*Table, error) {
	return newTable(self, name)
}

func buildDataSourceName(conn config.DatabaseConn) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conn.User, conn.Pass, conn.Host, conn.Port, conn.Name)
}
