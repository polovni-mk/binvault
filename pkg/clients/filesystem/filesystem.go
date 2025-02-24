package filesystem

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

var FileQueue = make(chan string, 10)

func initQueue(dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			absPath, err := filepath.Abs(file.Name())
			if err != nil {
				log.Println("error getting absolute path:", err)
				continue
			}
			FileQueue <- absPath
		}
	}
}

func WatchFolder(dirPath string) {
	initQueue(dirPath)
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		fmt.Println("error creating folder watcher", err)
	}

	defer watcher.Close()
	log.Default().Println("Watching folder:", dirPath)
	watcher.Add(dirPath)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				filePath := event.Name
				fmt.Println("New file detected:", filePath)
				FileQueue <- filePath
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Watcher error:", err)
		}
	}
}
