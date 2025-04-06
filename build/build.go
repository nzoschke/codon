package build

import "embed"

//go:embed all:dist
var Dist embed.FS

//go:generate bun install
//go:generate bun run build
