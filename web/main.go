package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//go:embed templates/*
var content embed.FS

func RunDashboard(port int) {
	tmpl := template.Must(template.ParseFS(content, "templates/*.html"))
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Region": "Makurdi",
			"Constants": map[string]float64{
				"Gumbel_C": 404.11,
				"Gumbel_M": 0.1733,
			},
		}
		tmpl.ExecuteTemplate(w, "index.html", data)
	})
	
	addr := ":" + strconv.Itoa(port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
