package main

import (
	"fmt"

	"github.com/subfuzion/mesh"
	"github.com/subfuzion/mesh/dotenv"
	"github.com/subfuzion/mesh/pkg/fsutil"
	"github.com/subfuzion/mesh/pkg/pretty"
)

func main() {
	fmt.Println("getting assets...")

	assets := fsutil.NewAssets(mesh.Assets)

	bytes, err := assets.ReadAll("template.env")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", bytes)

	t, err := dotenv.Unmarshal(string(bytes))
	if err != nil {
		fmt.Println(err)
	}
	pretty.Println(t)

	env, err := dotenv.Read()
	if err != nil {
		fmt.Println(err)
	}
	pretty.Println(env)
}
