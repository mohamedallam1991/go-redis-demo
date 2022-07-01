package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
	"github.com/mohamedallam1991/go-redis-demo/checking"
	"github.com/mohamedallam1991/go-redis-demo/config"
	"github.com/mohamedallam1991/go-redis-demo/models"
	"github.com/mohamedallam1991/go-redis-demo/resources"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// New Handlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("from the home function")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status its okay nigga"
	resp["content"] = "Here goes our content"
	jsonResp, err := json.Marshal(resp)
	checking.Checking(err, "error in json unmarshling and marshling")

	w.Write(jsonResp)
	return

	// render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// var connection config.AppConfig
// var connection resources.API

func (m *Repository) All(w http.ResponseWriter, r *http.Request) {

	q := chi.URLParam(r, "q")
	fmt.Println("city name:", q)

	redisAddress := fmt.Sprintf("%s:6379", os.Getenv("REDIS_URL"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var connection resources.API
	connection.Cache = rdb

	repo := resources.NewRepo(rdb)
	resources.NewConnection(repo)
	data, err := resources.TheAPI.GetData(r.Context(), q)
	checking.Checking(err, "cant get the data")

	resp := models.ApiResponse{
		Cache: false,
		Data:  data,
	}

	jsonResp, err := json.Marshal(resp)
	checking.Checking(err, "error in json unmarshling and marshling")

	w.Write(jsonResp)
}

func (m *Repository) Try(w http.ResponseWriter, r *http.Request) {

	q := chi.URLParam(r, "q")
	fmt.Println("city name:", q)

	redisAddress := fmt.Sprintf("%s:6379", os.Getenv("REDIS_URL"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var connection resources.API
	connection.Cache = rdb

	repo := resources.NewRepo(rdb)
	resources.NewConnection(repo)
	data, err := resources.TheAPI.TryingCache(r.Context(), q)
	checking.Checking(err, "cant get the data")

	resp := models.ApiResponse{
		Cache: false,
		Data:  data,
	}

	jsonResp, err := json.Marshal(resp)
	checking.Checking(err, "error in json unmarshling and marshling")

	w.Write(jsonResp)
}
