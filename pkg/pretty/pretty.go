package pretty

import (
	"encoding/json"
	"fmt"
	// 	"github.com/kr/pretty"
	// 	"github.com/davecgh/go-spew/spew"
)

// Println prints a pretty version to stdout
func Println(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
}
