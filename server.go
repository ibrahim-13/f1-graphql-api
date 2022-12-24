package main

import (
	"f1api/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const (
	__defaultIpAddr = "localhost"
	__defaultPort   = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = __defaultPort
	}
	ipAddr := os.Getenv("IPADDR")
	if ipAddr == "" {
		ipAddr = __defaultIpAddr
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("GraphQL playground : http://%s:%s/", ipAddr, port)
	log.Fatal(http.ListenAndServe(ipAddr+":"+port, nil))
}
