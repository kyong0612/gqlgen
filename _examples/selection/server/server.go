package main

import (
	"log"
	"net/http"

	"github.com/kyong0612/gqlgen/_examples/selection"
	"github.com/kyong0612/gqlgen/graphql/handler"
	"github.com/kyong0612/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Selection Demo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
