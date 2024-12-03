package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	pb "taskify/backend/proto"
	"taskify/backend/server"
)

func ListTasksHandler(s *server.Server, w http.ResponseWriter, r *http.Request) {
	// Create a TaskRequest if necessary, or pass an empty one
	task, err := ParseForm(r, false)
	if err != nil {
		RenderErrorPage(w, fmt.Sprintf("Error parsing form: %v", err))
		return
	}

	// Call ListTask method from server
	tasks, err := s.ListTasks(r.Context(), &pb.TaskRequest{Task: task})
	if err != nil {
		RenderErrorPage(w, fmt.Sprintf("Error fetching tasks: %v", err))
		return
	}
	// Get the current working directory
	// baseDir, err := os.Getwd()
	// if err != nil {
	// 	RenderErrorPage(w, "Failed to get working directory")
	// 	return
	// }

	// Construct the path to the HTML file
	templatePath := filepath.Join("..", "frontend", "list_tasks.html")

	// Parse the HTML file
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		RenderErrorPage(w, fmt.Sprintf("Failed to load template: %s", templatePath))
		return
	}

	// Pass the list of tasks to the template
	err = tmpl.Execute(w, tasks.Tasks)
	if err != nil {
		RenderErrorPage(w, fmt.Sprintf("Failed to render template:%v", tasks.Tasks))
		return
	}
}
