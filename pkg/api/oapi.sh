#!/bin/sh

cd ../..
go run cmd/app/main.go openapi
bunx openapi-typescript doc/openapi.json --alphabetize -o src/schema.d.ts
deno fmt doc/openapi.json src/schema.d.ts
