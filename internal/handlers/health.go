package handlers

import "net/http"

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("status: OK"))
}
