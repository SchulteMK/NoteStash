//go:build !embed

package static

import (
	"os"
)

func init() {
	Frontend = os.DirFS("./frontend/static")
}
