//
// Inspired by https://golang.org/doc/articles/wiki/#tmp_3
//
// Generated certs running
//
// go run /usr/lib/go/src/crypto/tls/generate_cert.go --host localhost
//
// https://localhost:10443/
//

package main

import (
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path[1:]
	if len(p) == 0 {
		p = "root"
	}
	fmt.Fprintf(w, "Hi there, loading %s!", p)
}

func monkey(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Monkey!")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/monkey", monkey)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}
