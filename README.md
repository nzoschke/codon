# CODON

Codon is a minimal Go and TypeScript toolkit that makes web app packaging and
deploys trivial.

Static sites are built with Bun and can be deployed as a directory.

Dynamic sites are built with Go + SQLite and deployed as a single binary.

The goals are to increase productivity building and hosting web apps, while
reducing dependencies, build steps, config files, and other layers of cruft.

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
bunx serve build/dist
```
