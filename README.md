# Codon

Codon is a minimal Go and Bun TypeScript toolkit that makes deployment trivial.

Static sites are built with Bun and deployed as a directory.

Dynamic sites are built with Go and deployed as a single binary.

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
