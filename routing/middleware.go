package routing

import (
	"fmt"
	"net/http"
	// "github.com/justinas/nosurf"
)

func setHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")

		fmt.Println("hit the page")
		return
	})
}
