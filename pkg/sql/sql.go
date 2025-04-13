package sql

import "embed"

//go:embed **/*.sql
var SQL embed.FS

//go:generate ./sqlc.sh
