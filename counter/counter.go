package counter

import (
	"fmt"
	"net/http"
)

func CounterHandler(mux *http.ServeMux) {
	ref := make(chan int, 1)
	initRef(ref)
	mux.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
		getHandle(w, r, ref)
	})
	mux.HandleFunc("/up/", func(w http.ResponseWriter, r *http.Request) {
		upHandle(w, r, ref)
	})
	mux.HandleFunc("/down/", func(w http.ResponseWriter, r *http.Request) {
		downHandle(w, r, ref)
	})
}

func initRef(ref chan int) {
	ref <- 0
}

func upHandle(w http.ResponseWriter, r *http.Request, ref chan int) {
	if r.Method != http.MethodGet {
		return
	}

	temp := (<-ref) + 1
	fmt.Fprintf(w, "Counter: %v", temp)
	ref <- temp
}

func downHandle(w http.ResponseWriter, r *http.Request, ref chan int) {
	if r.Method != http.MethodGet {
		return
	}

	temp := (<-ref) - 1
	fmt.Fprintf(w, "Counter: %v", temp)
	ref <- temp
}

func getHandle(w http.ResponseWriter, r *http.Request, ref chan int) {
	if r.Method != http.MethodGet {
		return
	}

	temp := <-ref
	fmt.Fprintf(w, "Counter: %v", temp)
	ref <- temp
}
