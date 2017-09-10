package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// KeysironConfig : type of KeysironConfigVar
type KeysironConfig struct {
	Server struct {
		Port int `yaml:"port"`
	}
	Database struct {
		Type     string `yaml:"type"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}
}

// Init the config
func (c *KeysironConfig) Init() {
}

// Parse from path
func (c *KeysironConfig) Parse(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
}

// KeysironConfigVar : keysiron configutation
var KeysironConfigVar KeysironConfig
