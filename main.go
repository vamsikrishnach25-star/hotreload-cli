package main

import (
	"flag"
	"fmt"
	"log"

	"hotreload/internal/watcher"
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

	w, err := watcher.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = w.WatchDir(*root)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {

		case event := <-w.Watcher.Events:
			log.Println("File changed:", event.Name)

		case err := <-w.Watcher.Errors:
			log.Println("Watcher error:", err)

		}
	}
}