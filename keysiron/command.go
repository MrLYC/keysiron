package main

import (
	"github.com/go-martini/martini"
	"os"
	"strconv"
)

func runServer() {
	os.Setenv("PORT", strconv.Itoa(KeysironConfigVar.Server.Port))

	m := martini.Classic()
	routesConfig(m)
	m.Run()
}
