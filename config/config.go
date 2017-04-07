package config

import (
    "path/filepath"
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "errors"
 )

type Config struct {
    Conn Conn
}

type Conn struct {
    Host string
    Port string
    User string
    Pass string
}

func Read(filePath string) (*Config, error) {
    if filePath == "" {
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