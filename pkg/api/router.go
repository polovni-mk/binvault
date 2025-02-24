package api

import (
	"binvault/pkg/api/handlers"

	"github.com/julienschmidt/httprouter"
)

func initRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/buckets", handlers.BucketGetMany)
	router.POST("/buckets", handlers.BucketCreate)
	router.GET("/buckets/:bucketId", handlers.BucketGetOne)
	router.DELETE("/buckets/:bucketId", handlers.BucketDelete)

	router.GET("/bucket/:bucketId/files", handlers.FileGetMany)
	router.POST("/bucket/:bucketId/files", handlers.FileCreate)
	router.GET("/bucket/:bucketId/files/:fileId", handlers.FileGetOne)
	router.DELETE("/bucket/:bucketId/files/:fileId", handlers.FileDelete)

	return router
}
