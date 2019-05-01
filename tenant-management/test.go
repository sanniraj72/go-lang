package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("welcome"))
	})
	_ = http.ListenAndServe(":3000", r)
}
