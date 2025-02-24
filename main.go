package main

import (
	"binvault/pkg/api"
	"binvault/pkg/cfg"
	"binvault/pkg/clients/filesystem"
	"binvault/pkg/compression"
	"binvault/pkg/database"
	"runtime"
)

var workers = runtime.NumCPU()

func main() {
	go filesystem.WatchFolder(cfg.GetPath("TEMP_PATH"))
	go compression.Init(workers)

	database.Init()
	api.StartServer()
}
