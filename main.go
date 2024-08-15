package main

import (
	"net/http"
)

func customHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

}
func main() {
	//Create a new http.ServeMux
	mux := http.NewServeMux()
	//Create a new http.Server struct and use the new "ServeMux" as the server's handler
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	//mux.HandleFunc("/users", customHandler)
	mux.HandleFunc("/healthz", customHandler)
	mux.Handle("/app/*", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	mux.Handle("/assets", http.FileServer(http.Dir("./assets/")))

	//Use the server's ListenAndServe method to start the server
	srv.ListenAndServe()

}
