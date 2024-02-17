package main

import "net/http"

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	http.ListenAndServe(":3000", router)
}
