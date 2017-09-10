package config

import (
	"flag"
)

// Argument : the type of ArgumentVar
type Argument struct {
	ConfigFile string
	Command    string
}

// Parse configutation from file
func (a *Argument) Parse() {
	flag.StringVar(&a.ConfigFile, "c", "keysiron.yaml", "YAML configutation")
	flag.StringVar(&a.Command, "x", "run", "Command to run")
	flag.Parse()
}

// ArgumentVar : command start argument
var ArgumentVar Argument
