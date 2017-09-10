package main

import (
	"github.com/go-martini/martini"
	"keysiron/config"
	"os"
	"strconv"
)

func runServer() {
	os.Setenv("PORT", strconv.Itoa(config.KeysironConfigVar.Server.Port))

	m := martini.Classic()
	routesConfig(m)
	m.Run()
}
