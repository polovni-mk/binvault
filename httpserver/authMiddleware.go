package httpserver

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthContext struct {
	UserId string
}

func GetAuthContext(w http.ResponseWriter, r *http.Request, params httprouter.Params) AuthContext {
	// TODO: extract user id from request
	// return nil if is not logged in
	return AuthContext{}
}
