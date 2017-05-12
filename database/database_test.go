package database

import (
	"testing"

	"github.com/anemiq/anemiq/config"
	"github.com/anemiq/anemiq/test"
)

func TestBuildDataSourceName(t *testing.T) {
	conn := config.DatabaseConn{
		Host: "localhost",
		Port: "3306",
		Name: "mydb",
		User: "user",
		Pass: "pass"}
	dataSource := buildDataSourceName(conn)
	test.AssertEqual(t, dataSource, "user:pass@tcp(localhost:3306)/mydb")
}
