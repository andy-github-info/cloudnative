package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.Handle("/resources", http.NotFoundHandler())
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/healthz", Healthz)
	log.SetFlags(log.Ldate | log.Ltime)
	log.Fatal(http.ListenAndServe(":80", logRequest(http.DefaultServeMux)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteaddr := strings.Split(r.RemoteAddr, ":")[0]
		log.Printf("%s \n", remoteaddr)
		handler.ServeHTTP(w, r)
	})
}

func Healthz(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "200\n")
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("-----------------")
	write := SetHeader(w, req)
	io.WriteString(write, "index\n")
}

func SetHeader(w http.ResponseWriter, req *http.Request) http.ResponseWriter {
	for k, v := range req.Header {
		w.Header().Set(k, v[0])
	}
	version := os.Getenv("VERSION")
	if version != "" {
		w.Header().Set("Version", version)
	}
	return w
}
