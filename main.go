package main

import (
	"log"
	"net/http"

	"example/rest/counter"
	"example/rest/download"
	"example/rest/greeting"

	"net/http/pprof"
)

func main() {
	mux := http.NewServeMux()
	log.Println("Setting up routes...")
	greeting.GreetingHandler(mux)
	counter.CounterHandler(mux)
	download.DownloadHandler(mux)
	log.Println("Setting up PProfile...")
	setUpPProfile(mux)
	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))
}

func setUpPProfile(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
}
