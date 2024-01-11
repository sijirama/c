package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	//NOTE: load .env
	godotenv.Load(".env")

	//NOTE: get PORT
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT must be set")
	}

	//NOTE: chi router
	r := chi.NewRouter()

	//NOTE: Middlewares
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	//NOTE: V1 router
	v1Router := chi.NewRouter()

	//NOTE: V1 routes
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)

	//NOTE: mount v1 router
	r.Mount("/v1", v1Router)

	//NOTE: http server
	server := &http.Server{
		Handler: r,
		Addr:    ":" + PORT,
	}

	//NOTE: server listening
	log.Printf("Server is finally listening on %v", PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
