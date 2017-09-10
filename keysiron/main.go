package main

import (
	"flag"
	"github.com/go-martini/martini"
	"keysiron/views"
	"os"
	"strconv"
)

// Argument : the type of ArgumentVar
type Argument struct {
	ConfigFile string
}

// Parse configutation from file
func (a *Argument) Parse() {
	flag.StringVar(&a.ConfigFile, "c", "keysiron.yaml", "YAML configutation")
	flag.Parse()
}

func routesConfig(m *martini.ClassicMartini) {
	m.Get("/", views.Index)
}

// ArgumentVar : command start argument
var ArgumentVar Argument

// KeysironConfigVar : keysiron configutation
var KeysironConfigVar KeysironConfig

func main() {
	ArgumentVar.Parse()
	KeysironConfigVar.Parse(ArgumentVar.ConfigFile)

	os.Setenv("PORT", strconv.Itoa(KeysironConfigVar.Server.Port))

	m := martini.Classic()
	routesConfig(m)
	m.Run()
}
