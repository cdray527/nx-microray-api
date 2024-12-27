package main

import (
	"log"
	"net/http"
	"nx-microray-api/graph"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "5200"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	if (os.Getenv("GO_ENV") == "development") {
		http.Handle("/api/product/playground", playground.Handler("Product API playground", "/api/product/playground"))
	}
	http.HandleFunc("/api/product/query", func(w http.ResponseWriter, r *http.Request) {
        // Allow CORS for all origins
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle preflight request
        if r.Method == http.MethodOptions {
            return
        }

        // Call the GraphQL server handler
        srv.ServeHTTP(w, r)
    })

	log.Printf("Product API Service started, access port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
