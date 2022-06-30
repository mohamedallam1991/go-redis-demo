package models

type ApiResponse struct {
	Cache bool                `json:"cache"`
	Data  []NominatimResponse `json:"data"`
}
