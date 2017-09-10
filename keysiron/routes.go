package main

import (
	"github.com/go-martini/martini"
	"keysiron/views"
)

func routesConfig(m *martini.ClassicMartini) {
	m.Get("/", views.Index)
}
