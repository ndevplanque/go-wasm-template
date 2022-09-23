package app

import (
	"html/template"
	"log"

	"github.com/julienschmidt/httprouter"
)

const version = "1.0.0"

// Config contient les méta données du server
type Config struct {
	Env  string
	Port int
}

type Application struct {
	Config   Config
	Logger   *log.Logger
	Template *template.Template
	Router   *httprouter.Router
}
