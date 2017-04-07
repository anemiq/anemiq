package config

import (
    "testing"
    "strings"
 )

func TestDefaultConfigFileIsMissing(t *testing.T) {
    _, err := Read("")
    if !strings.Contains(err.Error(), "file not found")  {
        t.Error()
    }
}

func TestMalformedConfigFile(t *testing.T) {
    _, err := Read("./testdata/malformed_anemiq.yaml")
    if !strings.Contains(err.Error(), "error reading")  {
        t.Error()
    }
}

func TestConfigFileIsReadenProperly(t *testing.T) {
    conf, _ := Read("./testdata/anemiq.yaml")
    assertEqual(t, conf.Conn.Host, "localhost")
    assertEqual(t, conf.Conn.Port, "3306")
    assertEqual(t, conf.Conn.User, "anemiq")
    assertEqual(t, conf.Conn.Pass, "1234")
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
    if a != b {
        t.Fatal()
    }
}