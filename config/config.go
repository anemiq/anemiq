package config

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   Server
	Database DatabaseConn
	Tables   Tables
}

type Server struct {
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

//Read configuration from yaml file. First command-line argument
//indicates the file path. Default config file is ./anemiq.yaml
func Read() (*Config, error) {
	args := os.Args[1:]
	var filePath string
	if len(args) > 0 {
		filePath = args[0]
	} else {
		filePath = "./anemiq.yaml"
	}
	return read(filePath)
}

func read(filePath string) (*Config, error) {
	filename, _ := filepath.Abs(filePath)
	configFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("Config file not found " + filePath)
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, errors.New("error reading config file " + filePath)
	}
	return &config, nil
}
