package main

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

// KeysironConfigVar : keysiron configutation
var KeysironConfigVar KeysironConfig

func main() {
	ArgumentVar.Parse()
	KeysironConfigVar.Init()
	KeysironConfigVar.Parse(ArgumentVar.ConfigFile)
	type commandFunc func()
	commands := map[string]commandFunc{
		"run":   runServer,
		"usage": flag.Usage,
	}
	function, ok := commands[ArgumentVar.Command]
	if !ok {
		function = commands["usage"]
	}
	function()
}
