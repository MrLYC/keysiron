package main

import (
	"github.com/go-martini/martini"
	"keysiron/config"
	"keysiron/storage"
	"os"
	"strconv"
)

func runServer() {
	os.Setenv("PORT", strconv.Itoa(config.KeysironConfigVar.Server.Port))

	m := martini.Classic()
	routesConfig(m)
	m.Run()
}

func migrate() {
	connection, err := storage.Connect()
	if err != nil {
		panic(err)
	}
	connection.CreateTable(&storage.User{})
}
