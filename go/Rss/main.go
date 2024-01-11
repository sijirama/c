package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
	//"github.com/go-chi/cors"
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
	//r.Use(middleware.Logger)

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
