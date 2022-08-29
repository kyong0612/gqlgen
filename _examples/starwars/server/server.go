package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/kyong0612/gqlgen/_examples/starwars"
	"github.com/kyong0612/gqlgen/_examples/starwars/generated"
	"github.com/kyong0612/gqlgen/graphql"
	"github.com/kyong0612/gqlgen/graphql/handler"
	"github.com/kyong0612/gqlgen/graphql/playground"
)

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(starwars.NewResolver()))
	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		rc := graphql.GetFieldContext(ctx)
		fmt.Println("Entered", rc.Object, rc.Field.Name)
		res, err = next(ctx)
		fmt.Println("Left", rc.Object, rc.Field.Name, "=>", res, err)
		return res, err
	})

	http.Handle("/", playground.Handler("Starwars", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
