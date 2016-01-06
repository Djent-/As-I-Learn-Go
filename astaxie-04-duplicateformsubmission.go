// https://github.com/astaxie/build-web-application-with-golang/blob/master/en/04.1.md

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"crypto/md5"
	"time"
	"io"
	"strconv"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse url parameters passed, then parse the
	// response packet for the POST body
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // get request method
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		// you know this isn't cryptographically secure, right?
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login2.gtpl")
		t.Execute(w, token)
	} else {
		log.Printf("Handling: ", r.Method)
		//login request
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// check token validity
		} else {
			// give error if no token
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func main() {
	http.HandleFunc("/", sayhelloName) // set router rule
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // set listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
