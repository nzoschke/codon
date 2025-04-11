#!/bin/sh

rm -rf q
sqlc generate
rm sqlc*.log q/crud_*.go
