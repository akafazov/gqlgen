# GQLGEN - GraphQL library using `https://github.com/99designs/gqlgen`

This is an extensible version of `https://github.com/99designs/gqlgen` which provides some out-of-the-box GraphQL schema and logic, but still allowing the user to extend the schema and implement new queries, mutations and objects in the same server endpoint.

## Usage

Import the package
```bash
go get github.com/akafazov/gqlgen@latest
```

Start a HTTP service with a graph endpoint by using the provided resolver from this package
```go
import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/akafazov/gqlgen/graph"
)

func main() {
	c := graph.Config{Resolvers: &graph.Resolver{}}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
```

