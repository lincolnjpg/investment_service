package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/lincolnjpg/investment_service/cmd/graphql/gqlgen"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type graphQl struct {
	app  ports.Application
	port string
}

func (g graphQl) Run() {
	r, _ := g.app.(gqlgen.Resolver)
	srv := handler.NewDefaultServer(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: &r}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", g.port)
	log.Fatal(http.ListenAndServe(":"+g.port, nil))
}

func NewGraphQl(app ports.Application, port string) *graphQl {
	return &graphQl{
		app:  app,
		port: port,
	}
}
