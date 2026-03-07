package main
import (
	"flag"
	"fmt"
	"log"
	"time"

	"hotreload/internal/builder"
	"hotreload/internal/runner"
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
	proc := &runner.Process{}
	err = builder.RunBuild(*buildCmd)
if err != nil {
	log.Fatal("Initial build failed")
}

err = proc.Start(*execCmd)
if err != nil {
	log.Fatal("Failed to start server")
}
debounceDelay := 500 * time.Millisecond
var debounceTimer *time.Timer

	for {
		select {

		case event := <-w.Watcher.Events:

	if debounceTimer != nil {
		debounceTimer.Stop()
	}

	debounceTimer = time.AfterFunc(debounceDelay, func() {

		log.Println("Debounced file change:", event.Name)

		proc.Stop()

		err := builder.RunBuild(*buildCmd)
		if err != nil {
			log.Println("Build failed, not restarting server")
			return
		}

		err = proc.Start(*execCmd)
		if err != nil {
			log.Println("Failed to start server:", err)
		}

	})

		case err := <-w.Watcher.Errors:
			log.Println("Watcher error:", err)

		}
	}
}