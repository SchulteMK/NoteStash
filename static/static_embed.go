//go:build embed

package frontend

import (
	"embed"
	"io/fs"
)

//go:embed all:build
var frontend embed.FS

func init() {
	var err error
	Frontend, err = fs.Sub(frontend, "static")
	if err != nil {
		panic(err)
	}
}
