package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func Routes(router *chi.Mux) {
	router.Get("/ping", pong)
}
