package main

import (
	"lightdoc/config"
	"lightdoc/route"
	"os"
)

func main() {


	if len(os.Args) > 1 {
		config.Root = os.Args[1]
	}
	if len(os.Args) > 2 {
		config.Dist = os.Args[2]
	}
	err := route.Init()
	if err != nil {
		panic(err)
	}
}
