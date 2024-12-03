package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func RenderErrorPage(w http.ResponseWriter, errorMessage string) {
	// Parse the error.html template
	tmpl, err := template.ParseFiles("../frontend/error.html")
	if err != nil {
		log.Printf("Failed to load error template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Render the template with the error message
	data := struct {
		Message string
	}{
		Message: errorMessage,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Failed to render error template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
