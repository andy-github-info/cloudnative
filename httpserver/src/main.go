package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.Handle("/resources", http.NotFoundHandler())
	http.Handle("/", SetHeader(http.HandlerFunc(IndexHandler)))
	http.Handle("/healthz", SetHeader(http.HandlerFunc(HealthzHandler)))
	log.SetFlags(log.Ldate | log.Ltime)
	log.Fatal(http.ListenAndServe(":80", logRequest(http.DefaultServeMux)))
}

func HealthzHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "200\n")
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "index\n")
}

func SetHeader(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			w.Header().Set(k, v[0])
		}
		version := os.Getenv("VERSION")
		if version != "" {
			w.Header().Set("VERSION", version)
		}
		handler.ServeHTTP(w, r)
	})
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteaddr := strings.Split(r.RemoteAddr, ":")[0]
		log.Printf("%s %d\n", remoteaddr, http.StatusOK)
		handler.ServeHTTP(w, r)
	})
}