package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mladenovic-13/rss-aggregator/internal/database"
)

type Context struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	if port == "" {
		log.Fatal("PORT is not found in environment variables")
	}
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in environment variables")
	}

	dbConnection, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Can not connect to database: ", err)
	}

	ctx := Context{
		DB: database.New(dbConnection),
	}

	router := chi.NewRouter()

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}
	router.Use(cors.Handler(corsOptions))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)

	v1Router.Post("/users", ctx.handlerCreateUser)
	v1Router.Get("/users", ctx.middlewareAuth(ctx.handlerGetUser))

	v1Router.Post("/feeds", ctx.middlewareAuth(ctx.handlerCreateFeed))
	v1Router.Get("/feeds", ctx.handlerGetFeeds)
	v1Router.Post("/feed_follows", ctx.middlewareAuth(ctx.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", ctx.middlewareAuth(ctx.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowId}", ctx.middlewareAuth(ctx.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port: %s", port)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
