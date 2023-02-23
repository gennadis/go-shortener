package main

import "net/http"

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/net":   "https://pkg.go.dev/net",
		"/os":    "https://pkg.go.dev/os",
		"/bufio": "https://pkg.go.dev/bufio",
		"/flag":  "https://pkg.go.dev/flag",
		"/json":  "https://pkg.go.dev/encoding/json",
		"/html":  "https://pkg.go.dev/html",
	}
	mapHandler := MapHandler(pathsToUrls, mux)
	http.ListenAndServe(":8080", mapHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
