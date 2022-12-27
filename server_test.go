package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var ViewHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello World</h1>")
	}
	server := http.Server{
		Addr:    "localhost:5000",
		Handler: ViewHandler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	fmt.Println("http://localhost:5000")
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Kifeb")
	})
	mux.HandleFunc("/riski", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query()
		fmt.Fprintf(w, "Hello %v", name)
	})
	mux.HandleFunc("/ratna", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Ratna")
	})
	mux.HandleFunc("/orang/kifeb", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello Orang Kifeb</h1>")
	})

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequestServer(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		if r.RequestURI == "/" {
			fmt.Fprint(w, "Hello Admin")
		} else {
			fmt.Fprintln(w, r.RequestURI)
		}
	}

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
