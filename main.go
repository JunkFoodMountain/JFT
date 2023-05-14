package main

import (
	"JFT/app"
	"JFT/contract"
	"JFT/internal/repository"
	"JFT/internal/transport/gql/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vrischmann/envconfig"
	"log"
	"net/http"
)

func main() {
	repoConfig := repository.Config{}
	err := envconfig.Init(&repoConfig)
	killIf(err)
	db, err := setupDB(repoConfig)
	killIf(err)

	app := app.New()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.New(&app)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func setupDB(config repository.Config) (*contract.NFTRepository, error) {
	db, err := repository.OpenDB(config)
	if err != nil {
		return nil, err
	}

	err = repository.MigrateDB(db)

	return db, err
}

func killIf(err error) {
	if err != nil {
		panic(err)
	}
}
