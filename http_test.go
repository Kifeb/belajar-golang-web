package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	// name := r.URL.Query().Get("name")
	name := r.URL.Query().Get("name")
	addr := r.URL.Query().Get("addr")
	if name == "" {
		fmt.Fprint(w, "Hello Admin")
	} else {
		fmt.Fprintf(w, "Hello My Name is %s From %s", name, addr)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000?name=Riski&addr=Bekasi", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func MultipleParam(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprint(w, names[0], names[1], names[2])
}

func TestMulitQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000?name=Riski&name=Ratna&name=Kifeb", nil)
	recorder := httptest.NewRecorder()

	MultipleParam(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
