package handlers

import "net/http"

// RootHandler handles requests to the root route
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Oops! The item you are looking for is not available at the moment!\n"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API Version 1\n"))
}
