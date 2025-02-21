package main

import (
	"binvault/database"
	"binvault/httpserver"
)

func main() {
	database.Init()
	httpserver.Run()
}
