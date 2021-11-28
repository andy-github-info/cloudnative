package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	var srv http.Server
	flag.Set("v", "4")
	http.Handle("/resources", http.NotFoundHandler())
	http.Handle("/", SetHeader(http.HandlerFunc(IndexHandler)))
	http.Handle("/healthz", SetHeader(http.HandlerFunc(HealthzHandler)))
	log.SetFlags(log.Ldate | log.Ltime)
	log.Printf("Starting http server...")
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := http.ListenAndServe(":80", logRequest(http.DefaultServeMux)); err != nil && err != http.ErrServerClosed{
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server Started")
	<-done
	log.Printf("Server Stoped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Printf("Server Exited Properly")
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