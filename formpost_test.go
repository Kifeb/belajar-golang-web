package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	name := r.PostForm.Get("name")
	age := r.PostForm.Get("age")
	intAge, _ := strconv.Atoi(age)
	if name == "" || intAge == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Form Ada Yang Harus Diisi")
	} else {
		fmt.Fprintf(w, "Hello %s My Age %v", name, intAge)
	}
}

func TestPostForm(t *testing.T) {
	reqBody := strings.NewReader("name=Riski&age=0")
	req := httptest.NewRequest(http.MethodPost, "http://localhost:5000", reqBody)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	FormPost(rec, req)
	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
	fmt.Println(res.StatusCode)
}
