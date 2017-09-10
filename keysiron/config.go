package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// KeysironConfig : type of KeysironConfigVar
type KeysironConfig struct {
	Server struct {
		Port int `yaml:"port"`
	}
	MySQL struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
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
