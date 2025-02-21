package httpserver

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Run() {
	router := httprouter.New()

	router.GET("/buckets", BucketGetMany)
	router.POST("/buckets", BucketCreate)
	router.GET("/buckets/:bucketId", BucketGetOne)
	router.DELETE("/buckets/:bucketId", BucketDelete)

	router.GET("/bucket/:bucketId/files", FileGetMany)
	router.POST("/bucket/:bucketId/files", FileCreate)
	router.GET("/bucket/:bucketId/files/:fileId", FileGetOne)
	router.DELETE("/bucket/:bucketId/files/:fileId", FileDelete)

	var err = http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
