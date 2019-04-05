// +build ignore

package tools

// This package is not intended to be built, but its presence
// ensures that `go mod` will include the imported packages
// that will be used in `go generate` directives. Since these
// directives are in contained in comments, they won't normally
// be seen by `go mod`.
import (
	// used for embedding files
	_ "github.com/shurcooL/vfsgen"
)
