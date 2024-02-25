package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Kiranraj27/go/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load the env")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("DB URL is not found in the environment")
	}

	queries := database.New(conn)
	if err != nil {
		log.Fatal("DB URL is not found in the environment")
	}

	apiCg := apiConfig{
		DB: queries,
	}

	router := chi.NewRouter()

	router.Use((cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/users", apiCg.handlerCreateUser)
	v1Router.Get("/users", apiCg.middlewareAuth(apiCg.handlerGetUser))
	v1Router.Post("/feeds", apiCg.middlewareAuth(apiCg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCg.handlerGetFeeds)
	v1Router.Post("/feed_follows", apiCg.middlewareAuth(apiCg.handlerCreateFeedFollows))
	v1Router.Get("/feed_follows", apiCg.middlewareAuth(apiCg.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{id}", apiCg.middlewareAuth(apiCg.handlerDeleteFeedFollows))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler:           router,
		Addr:              ":" + portString,
		ReadHeaderTimeout: time.Duration(5) * time.Second,
	}

	log.Printf("Server starting on PORT %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)
}
