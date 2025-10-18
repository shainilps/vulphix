package main

import (
	"fmt"
	"os"

	"github.com/codeshaine/vulphix/cmd/build"
	"github.com/codeshaine/vulphix/cmd/preview"
)

var help = `vulphix

Ready to generate your doc sites

Usage:
build -  for building the project
preview - for previewing the build
`

func main() {
	if len(os.Args) == 1 {
		fmt.Println(help)
		return
	}
	cmd := os.Args[1]
	if cmd == "help" {
		fmt.Println(help)
	}
	if cmd == "build" {
		fmt.Println("building your project...!!")
		os.Exit(build.Build())
		return
	}

	if cmd == "preview" {
		os.Exit(preview.PreviewBuild())
		return
	}
	fmt.Println("unknown command!!")

}
