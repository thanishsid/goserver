package assets

import (
	"embed"
)

//go:embed mail-templates/*.html
//go:embed sql/migrations
//go:embed casbin*
var Files embed.FS
