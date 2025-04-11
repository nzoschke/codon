package sql

import "embed"

//go:embed **/*.sql
var SQL embed.FS

//go:generate rm -rf q
//go:generate sqlc generate
//go:generate go fmt ./q
