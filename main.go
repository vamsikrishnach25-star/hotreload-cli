package main

import (
	"flag"
	"fmt"
)

func main() {

	root := flag.String("root", ".", "project root directory")
	buildCmd := flag.String("build", "", "build command")
	execCmd := flag.String("exec", "", "run command")

	flag.Parse()

	fmt.Println("HotReload Starting...")
	fmt.Println("Root:", *root)
	fmt.Println("Build Command:", *buildCmd)
	fmt.Println("Exec Command:", *execCmd)
}