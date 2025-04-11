package sql

import "embed"

//go:embed *.sql **/*.sql
var SQL embed.FS
