package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type health struct {
	Status   string   `json:"status"`
	Messages []string `json:"messages"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc(
		"/healthcheck",
		func(w http.ResponseWriter, r *http.Request) {
			h := health{
				Status:   "OK",
				Messages: []string{},
			}

			b, _ := json.Marshal(h)

			w.WriteHeader(http.StatusOK)
			w.Write(b)
		})

	s := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	s.ListenAndServe()

}
