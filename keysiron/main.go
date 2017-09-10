package main

import "flag"
import "os"
import "strconv"

import "github.com/go-martini/martini"

import "keysiron/views"

type argument struct {
	Port int
}

func (a *argument) Parse() {
	flag.IntVar(&a.Port, "p", 8000, "Port to run on")
	flag.Parse()
}

func routesConfig(m *martini.ClassicMartini) {
	m.Get("/", views.Index)
}

func main() {
	a := new(argument)
	a.Parse()
	os.Setenv("PORT", strconv.Itoa(a.Port))
	m := martini.Classic()
	routesConfig(m)
	m.Run()
}
