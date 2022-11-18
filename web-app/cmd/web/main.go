package main

import (
	"fmt"
	"github.com/matheus1103/go-studies/pkg/config"
	"github.com/matheus1103/go-studies/pkg/handlers"
	"github.com/matheus1103/go-studies/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
