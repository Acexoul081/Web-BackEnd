package main

import (
	"BackEnd/graph"
	"BackEnd/graph/generated"
	customMiddleware "BackEnd/middleware"
	"BackEnd/postgre"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"time"
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mgdb, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://admin:admin@cluster0.nigor.gcp.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil { log.Fatal(err) }
	err = mgdb.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

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

	resolver := graph.Resolver{DB: db, RDB: rdb, MGDB: mgdb}

	router.Use(customMiddleware.AuthMiddleware(resolver.Query()))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db,RDB: rdb, MGDB: mgdb}}))

	queryHandler := srv
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", graph.DataLoaderMiddleware(db, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
