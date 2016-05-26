package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/pressly/chi"
	"github.com/unrolled/render"
	"golang.org/x/net/context"

	"github.com/sogko/gosg-graphql-go-demo/server/schema"
)

var R *render.Render

// don't judge me
var IP = "127.0.0.1"
var PORT = "3000"

func init() {

	R = render.New(render.Options{
		Directory:     "views",
		IsDevelopment: true,
		Extensions:    []string{".html"},
	})
}

func serveGraphQL(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// helper to parse request query
	opts := handler.NewRequestOptions(r)

	// execute graphql query
	params := graphql.Params{
		Schema:         schema.Root,
		RequestString:  opts.Query,
		VariableValues: opts.Variables,
		OperationName:  opts.OperationName,
		Context:        ctx,
	}
	result := graphql.Do(params)

	// render result
	R.JSON(w, http.StatusOK, result)
}

func main() {
	r := chi.NewRouter()

	// GraphQL Endpoint
	r.Handle("/graphql", serveGraphQL)

	// Serves static files for client v2
	r.FileServer("/v2", http.Dir("public-v2"))

	// Serves static files for client v1
	r.FileServer("/", http.Dir("public"))

	bind := fmt.Sprintf("%s:%s", IP, PORT)
	log.Println("Starting server at", bind)

	http.ListenAndServe(bind, r)
}
