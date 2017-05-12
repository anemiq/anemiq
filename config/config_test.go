package config

import (
	"strings"
	"testing"

	"github.com/anemiq/anemiq/test"
)

func TestDefaultConfigFileIsMissing(t *testing.T) {
	_, err := read("")
	if !strings.Contains(err.Error(), "file not found") {
		t.Error()
	}
}

func TestMalformedConfigFile(t *testing.T) {
	_, err := read("./testdata/malformed_anemiq.yaml")
	if !strings.Contains(err.Error(), "error reading") {
		t.Error()
	}
}

func TestConfigFileIsReadenProperly(t *testing.T) {
	conf, _ := read("./testdata/anemiq.yaml")
	test.AssertEqual(t, conf.Database.Host, "localhost")
	test.AssertEqual(t, conf.Database.Port, "3306")
	test.AssertEqual(t, conf.Database.Name, "mydb")
	test.AssertEqual(t, conf.Database.User, "anemiq")
	test.AssertEqual(t, conf.Database.Pass, "1234")
}

func TestTablesAreReadenProperly(t *testing.T) {
	conf, _ := read("./testdata/anemiq.yaml")
	test.AssertEqual(t, conf.Tables[0], "users")
	test.AssertEqual(t, conf.Tables[1], "orders")
}

func TestNoTablesToExpose(t *testing.T) {
	conf, _ := read("./testdata/anemiq_no_tables.yaml")
	test.AssertEqual(t, len(conf.Tables), 0)
}
