package belajar_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(w, "SIMPLE", "Hello Guys")
}

func TestTemplate(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	rec := httptest.NewRecorder()

	SimpleHTML(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func SimpleHtmlFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.html"))
	t.ExecuteTemplate(w, "simple.html", "Hello Guys")
}
func TestTemplateFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:5000",
		Handler: http.HandlerFunc(SimpleHtmlFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TempDir(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.html"))
	t.ExecuteTemplate(w, "inde.html", "Hello Guys")
}
func TestTemplateDir(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	rec := httptest.NewRecorder()

	TempDir(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/*.html
var templates embed.FS

func TempEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.html"))
	t.ExecuteTemplate(w, "index.html", "Hello Kawan Semua")
}

func TestTemplateEmbed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	rec := httptest.NewRecorder()

	TempEmbed(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}
