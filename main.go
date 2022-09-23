package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"webapp/app"
)

func main() {
	fmt.Println("initializing...")

	// manually declare os.Args
	var cfg app.Config
	flag.IntVar(&cfg.Port, "port", 3000, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "development", "Application environment (development|production)")
	flag.Parse()

	// choose where and how to send logs
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// our html templates
	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		return
	}

	// our webapp
	app := &app.Application{
		Config:   cfg,
		Logger:   logger,
		Template: tmpl,
	}
	app.SetupRouter()

	// custom server config
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.Handler(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	logger.Println("Starting server...")

	err = srv.ListenAndServe()
	if err != nil {
		logger.Println(err)
		return
	}
}
