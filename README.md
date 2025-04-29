# Codon

Codon is a minimal Go, SQLite and TypeScript toolkit that makes web app
packaging and deploys trivial.

Static sites are built with Bun and can be deployed as a directory.

Dynamic apps are built with Go + SQLite and deployed as a single binary.

The goals are to increase productivity building and hosting web apps, with
modern tools that reduce dependencies, build steps, config files, and other
layers of cruft.

## Quick start

```bash
brew install deno go oven-sh/bun/bun sqlc
go install github.com/nzoschke/sqlc-gen-zz@latest
go install github.com/mitranim/gow@latest
go install github.com/gzuidhof/tygo@latest

go generate ./...
go test -v ./...
go build -o app cmd/app/main.go

./app
open http://localhost:1234
```

```bash
curl -d '{"email":"a@example.com","name":"Ann"}' -H "Content-Type: application/json" http://localhost:1234/api/contacts

{"created_at":"2025-04-13T18:39:04Z","email":"a@example.com","id":1,"info":{"age":0},"name":"Ann","phone":"","updated_at":"2025-04-13T18:39:04Z"}
```

## Development

### Go

```bash
LEVEL=debug gow -v run cmd/app/main.go -dev
```

To test:

```bash
go test -v ./...
bun test
```

To build:

````bash
go generate ./...
go build -o app cmd/app/main.go
./app

### TS

```bash
cp src/scripts/pre-commit .git/hooks
bun install
bun run build
bunx serve build/dist
```

## Philosophy

- All things equal the shorter solution is better
- Reduce dependencies

This leads us to:

- Bun and its build tools
- Deno and its fmt tool (until https://github.com/oven-sh/bun/issues/2246🤞)
- Go and its tools and stdlib
- Fuego and its Open API spec and client gen
- Svelte and its compiler (and without SvelteKit)
- SQLite and its file and memory backed databases
- Tailwind

References:

- https://dev.to/danielgtaylor/reducing-go-dependencies-dec
- https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
````
