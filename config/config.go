package config

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConn
	Tables   Tables
}

type ServerConfig struct {
	Port string
}

type DatabaseConn struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

type Tables []string

func Read() (*Config, error) {
	configFilePath, _ := filepath.Abs("./anemiq.yml")
	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, errors.New("malformed anemiq.yml\n" + err.Error())
	}
	return &config, nil
}
