package main

import (
	"net/http"
	"strconv"
)

type apiConfig struct {
	fileserverHits int
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) middlewareMetricsReset(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits = 0
		next.ServeHTTP(w, r)
	})
}
func (cfg *apiConfig) middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits := "Hits: " + strconv.Itoa(cfg.fileserverHits)
		w.Write([]byte(hits))
		next.ServeHTTP(w, r)
	})
}
func customHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)

}
func main() {
	//Create a new http.ServeMux
	mux := http.NewServeMux()
	//Create a new http.Server struct and use the new "ServeMux" as the server's handler
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	apiCfg := &apiConfig{}
	//mux.HandleFunc("/users", customHandler)
	mux.HandleFunc("/healthz", customHandler)

	mux.Handle("/metrics", apiCfg.middlewareLog(http.StripPrefix("/metrics", http.FileServer(http.Dir(".")))))
	mux.Handle("/reset", apiCfg.middlewareMetricsReset(http.StripPrefix("/reset", http.FileServer(http.Dir(".")))))
	mux.Handle("/app/*", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	mux.Handle("/assets", http.FileServer(http.Dir("./assets/")))

	//Use the server's ListenAndServe method to start the server
	srv.ListenAndServe()

}
