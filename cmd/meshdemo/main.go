package main

import (
	"fmt"

	"github.com/subfuzion/meshdemo"
	"github.com/subfuzion/meshdemo/dotenv"
	"github.com/subfuzion/meshdemo/pkg/fsutil"
	"github.com/subfuzion/meshdemo/pkg/pretty"
)

func main() {
	fmt.Println("getting assets...")

	assets := fsutil.NewAssets(meshdemo.Assets)

	bytes, err := assets.ReadAll("template.env")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", bytes)

	env, err := dotenv.Read()
	if err != nil {
		fmt.Println(err)
	}
	pretty.Println(env)
}
