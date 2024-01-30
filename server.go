package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/IsaacDSC/fullcycle_catalog_ecommerce/external/config"
	"github.com/IsaacDSC/fullcycle_catalog_ecommerce/graph"
	"github.com/IsaacDSC/fullcycle_catalog_ecommerce/shared"
)

const defaultPort = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := config.DbAnaLyticsConn().Ping(); err != nil {
		panic(err)
	}

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					Repositories: shared.NewContainerRepository(),
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
