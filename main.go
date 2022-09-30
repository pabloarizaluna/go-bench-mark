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
	greeting.GreetingHandler(mux)
	counter.CounterHandler(mux)
	download.DownloadHandler(mux)
	setUpPProfile(mux)
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}

func setUpPProfile(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
}
