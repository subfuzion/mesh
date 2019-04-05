package fsutil_test

import (
	"fmt"
	"net/http"

	"github.com/subfuzion/meshdemo/pkg/fsutil"
)

// This example shows the basic usage of the package: wrap an instance
// of an http.FileSystem.
func Example_basic() {

	// Create an http.FileSystem
	fs := http.Dir("assets")

	// Wrap it using an fsutil factory function
	assets := fsutil.NewAssets(fs)

	// Read in the full contents of a file
	path := "sample.txt"
	bytes, err := assets.ReadAll(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", bytes)

}
