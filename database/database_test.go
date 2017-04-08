package database

import (
	"testing"

	"github.com/anemiq/config"
	"github.com/anemiq/test"
)

func TestBuildDataSourceName(t *testing.T) {
	conn := config.Conn{
		Host:     "localhost",
		Port:     "3306",
		Database: "mydb",
		User:     "user",
		Pass:     "pass"}
	dataSource := buildDataSourceName(conn)
	test.AssertEqual(t, dataSource, "user:pass@tcp(localhost:3306)/mydb")
}
