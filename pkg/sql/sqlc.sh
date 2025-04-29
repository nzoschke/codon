#!/bin/sh

rm -rf q
sqlc generate
go fmt ./q
