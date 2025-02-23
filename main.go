package main

import (
	"binvault/core"
	"binvault/database"
	"binvault/filesystem"
	"binvault/httpserver"
	"binvault/processor"
	"runtime"
)

var workers = runtime.NumCPU()

func main() {
	go filesystem.WatchFolder(core.GetPath("TEMP_PATH"))
	go processor.RunWorkers(workers)

	database.Init()
	httpserver.Run()
}
