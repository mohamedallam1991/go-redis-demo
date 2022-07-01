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

const port = ":8081"

var app config.AppConfig

func main() {
	app.InProduction = false
	app.UseCache = false
	// handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("starting the app %v", port)
	portDock := fmt.Sprintf(":%s", os.Getenv("PORT"))

	srv := &http.Server{
		Addr:    portDock,
		Handler: routing.Routes(&app),
	}
	err := srv.ListenAndServe()

	checking.Checking(err, "fatal in main of the package")
}

// type ApiResponse struct {
// 	Cache bool                   `json:"cache"`
// 	Data  []TheNominatimResponse `json:"data"`
// }

// type TheNominatimResponse struct {
// 	PlaceID     int      `json:"place_id"`
// 	Licence     string   `json:"licence"`
// 	OsmType     string   `json:"osm_type"`
// 	OsmID       int      `json:"osm_id"`
// 	Boundingbox []string `json:"boundingbox"`
// 	Lat         string   `json:"lat"`
// 	Lon         string   `json:"lon"`
// 	DisplayName string   `json:"display_name"`
// 	Class       string   `json:"class"`
// 	Type        string   `json:"type"`
// 	Importance  float64  `json:"importance"`
// 	Icon        string   `json:"icon"`
// }
