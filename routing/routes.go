package routing

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mohamedallam1991/go-redis-demo/config"
	"github.com/mohamedallam1991/go-redis-demo/handlers"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	// mux.Use(setHeaders)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/all/{q}", handlers.Repo.All)
	// mux.Get("/numbers/{id}", handlers.Repo.Numbers)

	// fileServer := http.FileServer(http.Dir("./static/"))
	// mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
