package main

import "net/http"

func (a *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("status: OK"))
}
