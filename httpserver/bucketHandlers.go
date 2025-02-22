package httpserver

import (
	"binvault/services"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GET /buckets
func BucketGetMany(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pagination := GetRequestPagination(r)
	buckets := services.BucketGetMany(pagination.limit, pagination.offset)
	log.Println(buckets)
	JSONResponse(w, http.StatusOK, buckets)
}

// POST /buckets
func BucketCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// GET /buckets/:bucketName
func BucketGetOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// DELETE /buckets/:bucketName
func BucketDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
