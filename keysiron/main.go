package main

import (
	"flag"
	"keysiron/config"
)

func main() {
	config.ArgumentVar.Parse()
	config.KeysironConfigVar.Init()
	config.KeysironConfigVar.Parse(config.ArgumentVar.ConfigFile)
	type commandFunc func()
	commands := map[string]commandFunc{
		"run":     runServer,
		"migrate": migrate,
		"usage":   flag.Usage,
	}
	function, ok := commands[config.ArgumentVar.Command]
	if !ok {
		function = commands["usage"]
	}
	function()
}
