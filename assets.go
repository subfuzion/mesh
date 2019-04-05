// +build dev

package meshdemo

import (
	"net/http"
)

// Assets contains project assets, such as template files.
// When building with the `dev` tag, access these files
// directly from the file system; otherwise the embedded
// assets generated using `vfsgen` will be accessed.
var Assets http.FileSystem = http.Dir("assets")
