graphql-init:
	go run github.com/99designs/gqlgen init

graphql-generate:
	go run github.com/99designs/gqlgen generate

graphql-server:
	go run cmd/graphql/server.go