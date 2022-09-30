package greeting

import (
	"fmt"
	"net/http"
	"regexp"
)

const path = "/greeting/"

func GreetingHandler(h *http.ServeMux) {
	h.HandleFunc(path, greetingHandle)
}

func greetingHandle(w http.ResponseWriter, r *http.Request) {
	queryValue, prs := r.URL.Query()["name"]

	pathVariable := regexp.MustCompile("^" + path + "([-a-zA-Z0-9]+)$")
	pathValue := pathVariable.FindStringSubmatch(r.URL.RequestURI())

	switch {
	case r.Method == http.MethodGet && prs:
		fmt.Fprintf(w, "Hello %v", queryValue[0])
	case r.Method == http.MethodGet && pathValue != nil:
		fmt.Fprintf(w, "Hello %v", pathValue[1])
	case r.Method == http.MethodGet:
		fmt.Fprintf(w, "Hello World!")
	}
}
