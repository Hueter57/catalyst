package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hueter57/catalyst/backend/internal/domain"
	"github.com/hueter57/catalyst/backend/internal/graph"
	"github.com/hueter57/catalyst/backend/internal/repository"
	"github.com/hueter57/catalyst/backend/internal/resolver"
	"github.com/hueter57/catalyst/backend/internal/service"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	// Setup ent client
	client, err := domain.Connect()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()
	repo := repository.NewRepository(client)
	if err := repo.MigrateApply(ctx); err != nil {
		log.Fatalf("failed to apply migrations: %v", err)
	}

	service := service.New(repo)
	resolver := resolver.NewResolver(service)

	srv := newQueryHandler(resolver)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// TODO auth を実装する
	http.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func newQueryHandler(resolver *resolver.Resolver) *handler.Server {
	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
		// TODO Directives と Complexity を実装する
		// Directives: graph.NewDirective(),
		// Complexity: graph.ComplexityConfig(),
	}))
	// TODO ComplexityLimit と middleware を実装する

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return srv
}
