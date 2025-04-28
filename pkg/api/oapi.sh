#!/bin/sh

cd ../..
go run cmd/app/main.go openapi
mkdir -p pkg/sdk
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest -generate "types,client" -package sdk doc/openapi.json | tail -n +2 > pkg/sdk/sdk.go
bunx openapi-typescript doc/openapi.json --alphabetize -o src/schema.d.ts
