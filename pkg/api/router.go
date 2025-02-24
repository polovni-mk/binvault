package api

import (
	"binvault/pkg/api/handlers"

	"github.com/julienschmidt/httprouter"
)

func initRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/buckets", handlers.BucketGetMany)
	router.POST("/api/buckets", handlers.BucketCreate)
	router.GET("/api/buckets/:bucketName", handlers.BucketGetOne)
	router.DELETE("/api/buckets/:bucketName", handlers.BucketDelete)

	router.GET("/api/bucket/:bucketName/files", handlers.FileGetMany)
	router.POST("/api/bucket/:bucketName/files", handlers.FileCreate)
	router.GET("/api/bucket/:bucketName/files/:fileId", handlers.FileGetOne)
	router.DELETE("/api/bucket/:bucketName/files/:fileId", handlers.FileDelete)

	return router
}
