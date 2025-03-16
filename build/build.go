package build

import "embed"

//go:embed all:dist
var Dist embed.FS
