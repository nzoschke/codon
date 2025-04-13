#!/bin/sh

rm -rf q
sqlc generate
rm q/crud*.go
go fmt ./q

tygo generate
