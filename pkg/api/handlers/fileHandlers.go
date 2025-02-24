package handlers

import (
	"binvault/pkg/api/helpers"
	"binvault/pkg/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GET /bucket/:bucketName/files
func FileGetMany(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pagination := helpers.GetRequestPagination(r)
	bucketName := params.ByName("bucketName")
	files := services.FileGetMany(bucketName, pagination.Limit, pagination.Offset)
	helpers.JSONResponse(w, http.StatusOK, files)
}

// POST /bucket/:bucketName/files
func FileCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// GET /bucket/:bucketName/files/:fileId
func FileGetOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bucketName := params.ByName("bucketName")
	fileId := params.ByName("fileId")
	file := services.FileGetOne(bucketName, fileId)
	helpers.JSONResponse(w, http.StatusOK, file)
}

// GET /bucket/:bucketName/files/:fileId/content
func FileGetContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// DELETE /bucket/:bucketName/files/:fileId
func FileDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
