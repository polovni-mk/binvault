package httpserver

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GET /buckets
func BucketGetMany(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

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
