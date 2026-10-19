package watcher
import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)
type Watcher struct {
	Watcher *fsnotify.Watcher
}

func NewWatcher() (*Watcher, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Watcher{
		Watcher: w,
	}, nil
}
func shouldIgnore(path string) bool {

	ignoreDirs := []string{
		".git",
		"node_modules",
		"bin",
		"build",
		"tmp",
	}

	for _, dir := range ignoreDirs {
		if filepath.Base(path) == dir {
			return true
		}
	}

	return false
}

func (w *Watcher) WatchDir(root string) error {

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {

		if shouldIgnore(path) {
			log.Println("Ignoring:", path)
			return filepath.SkipDir
			}

			log.Println("Watching:", path)
			return w.Watcher.Add(path)
		}

		return nil
	})

	return err
}

func (w *Watcher) Start() {

	for {
		select {

		case event := <-w.Watcher.Events:
			log.Println("File changed:", event.Name)

		case err := <-w.Watcher.Errors:
			log.Println("Watcher error:", err)

		}
	}
}