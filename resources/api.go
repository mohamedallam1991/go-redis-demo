package resources

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mohamedallam1991/go-redis-demo/checking"
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

func (a *API) GetData(ctx context.Context, q string) ([]models.NominatimResponse, error) {
	fmt.Println("from the get data function")
	// value, err := a.Cache.Get(ctx, q).Result()
	// checking.Checking(err, "cant get the data from redis")

	// if err == redis.Nil {
	// we want call external data source

	resp := GetFromSource(q)
	data := make([]models.NominatimResponse, 0)

	_ = json.NewDecoder(resp.Body).Decode(&data)
	b, _ := json.Marshal(data)

	err := a.Cache.Set(ctx, q, bytes.NewBuffer(b).Bytes(), time.Second*15).Err()
	checking.Checking(err, "cant set the data")

	return data, nil
	// } else if err != nil {
	// 	fmt.Printf("error calling redis: %v\n", err)
	// 	return nil, err
	// } else {
	// 	log.Fatal("weird thing happening here")
	// 	data := make([]models.NominatimResponse, 0)

	// 	return data, nil

	// }
	// 	// cache hit
	// 	data := make([]models.NominatimResponse, 0)

	// 	// build response
	// 	err = json.Unmarshal(bytes.NewBufferString(value).Bytes(), &data)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	// return response
	// 	return data, nil
	// }
	// return nil, nil

}

func GetFromSource(qq string) (r *http.Response) {
	espaceQ := url.PathEscape(qq)
	address := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", espaceQ)
	resp, err := http.Get(address)
	if err != nil {
		log.Fatal(err, "error getting the damn data from api")
	}
	return resp
}

func (a *API) TryingCache(ctx context.Context, q string) ([]models.NominatimResponse, error) {
	fmt.Println("from the get trying function")
	value, err := a.Cache.Get(ctx, q).Result()
	// checking.Checking(err, "cant get the data from redis")

	if err != nil {
		fmt.Println("cant get the data from redis")
	}

	// cache hit
	data := make([]models.NominatimResponse, 0)

	// build response
	err = json.Unmarshal(bytes.NewBufferString(value).Bytes(), &data)
	// checking.Checking(err, "cant umarshell the data")

	if err != nil {
		fmt.Println("cant umarshell the data")
	}
	// return response
	return data, nil

}
