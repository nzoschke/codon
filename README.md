# Codon

Codon is a minimal Go, SQLite and TypeScript toolkit that makes web app
packaging and deploys trivial.

Static sites are built with Bun and can be deployed as a directory.

Dynamic apps are built with Go + SQLite and deployed as a single binary.

The goals are to increase productivity building and hosting web apps, with
modern tools that reduce dependencies, build steps, config files, and other
layers of cruft.

## Go

```bash
go tool gow run cmd/app/main.go
```

## TS

```bash
bun install
bun run dev
```

To build:

```bash
bun run build

# static server
bunx serve build/dist

# app server
go build -o app cmd/app/main.go
./app
```

## Philosophy

- All things equal the shorter solution is better
- Reduce dependencies

This leads us to:

- Bun and its build tools
- Go and its tools and stdlib
- Svelte and its compiler (and without SvelteKit)
- SQLite and its file and memory backed databases
- Tailwind

References:

- https://dev.to/danielgtaylor/reducing-go-dependencies-dec
