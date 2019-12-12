package handler

import "net/http"

// HelloHanlder - for testing
func HelloHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Clean Arch."))
}

// PanicHanlder - for testing
func PanicHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	panic("ohh, panic occured")
	//w.Write([]byte("Hello Clean Arch."))
}
