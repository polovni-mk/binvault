package processor

import (
	"binvault/filesystem"
	"log"
	"sync"
)

func RunWorkers(workers int) {
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for filePath := range filesystem.FileQueue {
				runWorker(filePath)
			}
		}()
	}
}

func runWorker(path string) {
	log.Default().Println("Processing file: ", path)
	// TODO: get file from database, based on mimetype, file size and file extenstion choose correct processor
}
