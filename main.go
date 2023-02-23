package main

import "net/http"

func main() {
	mux := defaultMux()
	http.ListenAndServe(":8080", mux)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
