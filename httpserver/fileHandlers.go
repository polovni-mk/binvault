package httpserver

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GET /bucket/:bucketName/files
func FileGetMany(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// POST /bucket/:bucketName/files
func FileCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// GET /bucket/:bucketName/files/:fileId
func FileGetOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// DELETE /bucket/:bucketName/files/:fileId
func FileDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
