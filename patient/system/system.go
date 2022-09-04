package system

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Mysql string `yaml:"mysql"` //mysql url
	Addr  string `yaml:"addr"`  //server addr
}

var configuration *Config

func LoadConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	configuration = &config
	return nil
}

func GetConfig() *Config {
	return configuration
}
