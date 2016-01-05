// https://github.com/astaxie/build-web-application-with-golang/blob/master/en/03.4.md

package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hellow route!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}