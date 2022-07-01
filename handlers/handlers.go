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
	resp["message"] = "Status its okay babe"
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

	// redisAddress := fmt.Sprintf("%s:6379", os.Getenv("REDIS_URL"))

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     redisAddress,
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })
	// var connection resources.API
	// connection.Cache = rdb

	raa := connect()
	// repo := resources.NewRepo(rdb)
	repo := resources.NewRepo(raa)
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

func connect() *redis.Client {
	var opts *redis.Options

	if os.Getenv("LOCAL") == "true" {
		redisuri := os.Getenv("REDIS_URL")
		fmt.Println("redisuri :", redisuri)

		redisAddress := fmt.Sprintf("%s:6379", redisuri)
		fmt.Println("redisAddress :", redisAddress)

		// fmt.Println("Redis address :", redisAddress)
		// fmt.Println("redis address :", redis:6379)

		opts = &redis.Options{
			Addr:     redisAddress,
			Password: "", // no password set
			DB:       0,  // use default DB
		}
		fmt.Println("builtOpts :", opts)

	} else {
		// redisUri := `redis://:p5d63c80679f27374749b8fdde15820fb74f7da276e7c6eb5e5ac6dae4cfb61c3@ec2-100-26-75-186.compute-1.amazonaws.com:16529:6379`
		// redisUri := `rediss://:p5d63c80679f27374749b8fdde15820fb74f7da276e7c6eb5e5ac6dae4cfb61c3@ec2-100-26-75-186.compute-1.amazonaws.com:16530`
		redisUri := `redis://:p5d63c80679f27374749b8fdde15820fb74f7da276e7c6eb5e5ac6dae4cfb61c3@ec2-100-26-75-186.compute-1.amazonaws.com:16529`
		// redis://:p5d63c80679f27374749b8fdde15820fb74f7da276e7c6eb5e5ac6dae4cfb61c3@ec2-100-26-75-186.compute-1.amazonaws.com:16529:6379:
		// fmt.Println("redisUri :", redisUri)
		// redisAddress := fmt.Sprintf("%s:6379", redisUri)
		// fmt.Println("redisAddress :", redisAddress)

		// redisAddress := fmt.Sprintf("%s:6379", redisUri)
		// fmt.Println("redisAddress :", redisAddress)
		//redis://<user>:<password>@<host>:<port>/<db_number>

		builtOpts, err := redis.ParseURL(redisUri)
		if err != nil {
			panic(err)
		}
		fmt.Println("builtOpts :", opts)

		// builtOpts, err := redis.ParseURL(os.Getenv("RZEDIS_URL"))
		// if err != nil {
		// 	panic(err)
		// }
		opts = builtOpts
	}

	rdb := redis.NewClient(opts)
	return rdb

	// return &resources.API{
	// 	cache: rdb,
	// }
}
