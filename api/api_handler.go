package api

import (
	"net/http"
	"strings"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")

	if strings.HasPrefix(r.URL.Path, "create_user") {
		CreateUserHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "scan_login") {
		ScanLoginHandler(w, r)
		return
	}
}
