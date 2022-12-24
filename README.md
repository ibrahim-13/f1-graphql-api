# f1-api

A tiny server which exposes GraphQL API for Formula 1 race schedule.
Data is retrieved from Linked Metadata.

## Development
This package usage (gqlgen)[https://gqlgen.com/getting-started/] for GraphQL

 - Create the project skeleton

```sh
go run github.com/99designs/gqlgen init
go mod tidy
```

 - Run server

 ```sh
 go run server.go
 ```

 - Run code generation

 ```sh
 go generate ./..
 ```

`resolver.go` contains the following comment at the top, which generates the code -
>//go:generate go run github.com/99designs/gqlgen generate