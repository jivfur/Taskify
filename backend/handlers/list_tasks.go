package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"taskify/backend/server"
)

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Create a TaskRequest if necessary, or pass an empty one
	taskReq := &PB.TaskRequest{}
	// Call ListTask method from server
	tasks, err := server.ListTasks(taskReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching tasks: %v", err), http.StatusInternalServerError)
		return
	}

	// Render the listTasks.html template with the tasks data
	tmpl, err := template.ParseFiles("frontend/listTasks.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error loading template: %v", err), http.StatusInternalServerError)
		return
	}

	// Pass the list of tasks to the template
	tmpl.Execute(w, tasks)
}
