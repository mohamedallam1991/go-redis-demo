package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mohamedallam1991/go-redis-demo/checking"
	"github.com/mohamedallam1991/go-redis-demo/config"
	"github.com/mohamedallam1991/go-redis-demo/handlers"
	"github.com/mohamedallam1991/go-redis-demo/routing"
)

// const port = ":8081"

var app config.AppConfig

func main() {
	app.InProduction = false
	app.UseCache = false
	// handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("starting the app %v", os.Getenv("PORT"))
	portDock := fmt.Sprintf(":%s", os.Getenv("PORT"))

	srv := &http.Server{
		Addr:    portDock,
		Handler: routing.Routes(&app),
	}
	err := srv.ListenAndServe()

	checking.Checking(err, "fatal in main of the package")
}
