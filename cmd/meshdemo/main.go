package main

import (
	"fmt"

	"github.com/subfuzion/meshdemo"
	"github.com/subfuzion/meshdemo/pkg/fsutil"
)

func main() {
	fmt.Println("getting assets...")

	assets := fsutil.NewAssets(meshdemo.Assets)

	bytes, err := assets.ReadAll("template.env")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", bytes)
}
