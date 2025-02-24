package helpers

import (
	"net/http"
	"strconv"
)

type RequestPagination struct {
	Limit  int
	Offset int
}

func GetRequestPagination(r *http.Request) *RequestPagination {
	query := r.URL.Query()
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(query.Get("offset"))
	if err != nil {
		offset = 0
	}

	if (limit < 1) || (limit > 100) {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return &RequestPagination{limit, offset}
}
