package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	ct := r.Header.Get("Content-Type")
	fmt.Fprint(w, ct)
}

func TestContentType(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	RequestHeader(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "34")
	fmt.Fprint(w, "Oke")
}

func TestResponseHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	ResponseHeader(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)
	ct := res.Header.Get("x-powered-by")
	intteger, _ := strconv.Atoi(ct)

	fmt.Println(string(body))
	fmt.Println(intteger)
}
