#!/bin/sh

rm -rf q
sqlc generate
go fmt ./q

tygo generate
deno fmt models
