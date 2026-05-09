package main

import (
	"ascii-art-web/internal/ascii"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Result string
	Error string
}

var tmpl *template.Template

func init() {
	var err error

	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("Error loading template:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	result, err := ascii.GenerateASCII(text, banner)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Result: result,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)

	log.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
