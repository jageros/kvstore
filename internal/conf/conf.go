package conf

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	Model      string `yaml:"model"`
	ListenAddr string `yaml:"listen_addr"`
	Mongo      struct {
		Addr     string `yaml:"addr"`
		DB       string `yaml:"db"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"mongo"`
}

func Parse(path string) *Config {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Read Config yaml file err: %v", err)
	}

	conf := new(Config)
	err = yaml.Unmarshal(yamlFile, &conf)

	if err != nil {
		log.Fatalf("Read Config yaml Unmarshal err: %v", err)
	}
	return conf
}
