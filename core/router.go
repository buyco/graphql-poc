package core

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/buyco/keel/pkg/app"
	go_graphql_poc "github.com/defgenx/go-graphql-poc"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	return &Router{
		mux.NewRouter()}
}

func (r *Router) SetRoutes() {
	r.setHandlers()
	r.setGraphQLRoutes()
}

func (r *Router) setHandlers() {
	r.Router.NotFoundHandler = http.HandlerFunc(
		app.ErrorJSONHandler(
			&app.ErrorApiResponse{
				Error: &app.Error{
					Message: http.StatusText(http.StatusNotFound),
					Code:    "GQ-001",
				},
				HttpResponse: &app.HttpResponse{Message: http.StatusText(http.StatusNotFound), Code: http.StatusNotFound},
			},
		),
	)
}
func (r *Router) setGraphQLRoutes() {
	r.Router.HandleFunc("/health", healthCheck).Name("HealthCheck").Methods("GET")
	r.Router.HandleFunc("/", handler.Playground("GraphQL playground", "/query"))
	r.Router.HandleFunc("/query", handler.GraphQL(go_graphql_poc.NewExecutableSchema(go_graphql_poc.Config{Resolvers: &go_graphql_poc.Resolver{}})))
}
