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
	test.AssertEqual(t, conf.Conn.Host, "localhost")
	test.AssertEqual(t, conf.Conn.Port, "3306")
	test.AssertEqual(t, conf.Conn.Database, "mydb")
	test.AssertEqual(t, conf.Conn.User, "anemiq")
	test.AssertEqual(t, conf.Conn.Pass, "1234")
}
