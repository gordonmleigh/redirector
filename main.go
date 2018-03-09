// import "github.com/gordonmleigh/redirector"
package main

import (
	"net/http"
	"strings"
)

func main() {
	server := &http.Server{
		Addr:    ":4000",
		Handler: http.HandlerFunc(handler),
	}
	server.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := *r.URL
	parts := strings.SplitN(strings.TrimLeft(url.Path, "/"), "/", 3)
	url.Scheme = parts[0]
	url.Host = parts[1]
	url.Path = parts[2]
	http.Redirect(w, r, url.String(), 301)
}
