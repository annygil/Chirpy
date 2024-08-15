package main

import (
	"net/http"
)

func main() {
	//Create a new http.ServeMux
	mux := http.NewServeMux()
	//Create a new http.Server struct and use the new "ServeMux" as the server's handler
	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.Handle("/assets", http.FileServer(http.Dir("./assets/")))
	//Use the server's ListenAndServe method to start the server
	srv.ListenAndServe()

}
