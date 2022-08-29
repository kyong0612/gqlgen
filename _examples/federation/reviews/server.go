//go:generate go run ../../../testdata/gqlgen.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kyong0612/gqlgen/_examples/federation/reviews/graph"
	"github.com/kyong0612/gqlgen/_examples/federation/reviews/graph/generated"
	"github.com/kyong0612/gqlgen/graphql/handler"
	"github.com/kyong0612/gqlgen/graphql/handler/debug"
	"github.com/kyong0612/gqlgen/graphql/playground"
)

const defaultPort = "4003"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.Use(&debug.Tracer{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
