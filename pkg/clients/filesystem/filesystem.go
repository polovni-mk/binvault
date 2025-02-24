package filesystem

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

var FileQueue = make(chan string, 10)

func initQueue() {
	// TODO: read existing files and schedule them for processing
}

func WatchFolder(folderPath string) {
	watcher, err := fsnotify.NewWatcher()

	initQueue()

	if err != nil {
		fmt.Println("error creating folder watcher", err)
	}

	defer watcher.Close()
	log.Default().Println("Watching folder:", folderPath)
	watcher.Add(folderPath)

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
