package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-Kifeb-Ganteng"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success Set Cookie")
}
func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-Kifeb-Ganteng")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(w, "Hello %s", name)
	}

}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
