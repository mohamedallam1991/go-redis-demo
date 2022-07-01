package resources

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/go-redis/redis/v8"
	"github.com/mohamedallam1991/go-redis-demo/models"
)

// var Repo *Repository
var TheAPI *API

type API struct {
	Cache *redis.Client
}

func NewRepo(a *redis.Client) *API {
	return &API{
		Cache: a,
	}
}

func NewConnection(a *API) {
	TheAPI = a
}

func (a *API) GetData(q string) ([]models.NominatimResponse, error) {
	fmt.Println("from the data function")

	espaceQ := url.PathEscape(q)
	address := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", espaceQ)
	resp, err := http.Get(address)
	if err != nil {
		log.Fatal(err, "error getting the damn data")
	}

	data := make([]models.NominatimResponse, 0)

	_ = json.NewDecoder(resp.Body).Decode(&data)
	return data, nil
}
