package main

import (
	"lightdoc/config"
	"lightdoc/route"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		config.Pubip = os.Args[1]

	}
	if len(os.Args) > 2 {
		config.Root = os.Args[2]
	}
	if len(os.Args) > 3 {
		config.Dist = os.Args[3]
	}
	err := route.Init()
	if err != nil {
		panic(err)
	}
}
