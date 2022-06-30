package resources

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/mohamedallam1991/go-redis-demo/models"
)

func GetData(q string) ([]models.NominatimResponse, error) {
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
