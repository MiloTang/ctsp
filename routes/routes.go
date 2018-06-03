package routes

import (
	"net/http"
	"webnet/action"
)

var (
	Routes = map[string]func(w http.ResponseWriter, r *http.Request){
		"/":      action.Index,
		"/admin": action.Admin,
	}
)
