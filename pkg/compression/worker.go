package compression

import "log"

func runWorker(path string) {
	log.Default().Println("Processing file: ", path)
	// TODO: get file from database, based on mimetype, file size and file extenstion choose correct processor
}
