package main

import (
	"BackEnd/graph"
	"BackEnd/graph/generated"
	"github.com/go-chi/chi/middleware"
	"github.com/go-redis/redis/v8"

	customMiddleware "BackEnd/middleware"
	"BackEnd/postgre"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v10"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

const defaultPort = "5000"

func main() {
	router := chi.NewRouter()

	db := postgre.New(&pg.Options{
		Addr: ":5432",
		User: "postgres",
		Password: "monster081",
		Database: "YouRJube",
	})

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	db.AddQueryHook(postgre.DBLogger{})

	defer db.Close()
	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "http://192.168.241.1:4200"},
		AllowedHeaders: []string{"Authorization","Content-Type","Origin"},
		AllowedMethods: []string{"GET","PUT","POST","DELETE","PATCH","OPTIONS"},
		AllowCredentials: true,
		Debug:            true,

	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	resolver := graph.Resolver{DB: db, RDB: rdb}

	router.Use(customMiddleware.AuthMiddleware(resolver.Query()))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db,RDB: rdb}}))

	queryHandler := srv
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", graph.DataLoaderMiddleware(db, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
