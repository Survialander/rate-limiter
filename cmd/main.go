package main

import (
	"net/http"

	"github.com/Survialander/rate-limitter/configs"
	"github.com/Survialander/rate-limitter/internal/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	configs.LoadConfig(".")
	router := chi.NewRouter()
	router.Use(middleware.RealIP)

	router.Use(middlewares.RateLimiter)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	http.ListenAndServe(":8080", router)
}
