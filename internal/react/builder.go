package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	fmt.Println("Start bundling javascript")
	minifyArg := os.Args[1]
	watchFlag := os.Args[2]
	fmt.Println(minifyArg, watchFlag)

	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"./internal/react/index.tsx"},
		Bundle:      true,
		Outdir:      "./public/assets",
		Write:       true,
		//	MinifyWhitespace:  true,
		//	MinifyIdentifiers: true,
		//	MinifySyntax:      true,
	})

	if len(result.Errors) != 0 {
		json.NewEncoder(os.Stdout).Encode(result)
		os.Exit(1)
	}
}
