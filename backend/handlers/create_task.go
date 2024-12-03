package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	pb "taskify/backend/proto"
	server "taskify/backend/server"
	validators "taskify/backend/validators"
)

func CreateTaskHandler(s *server.Server, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	task, err := ParseForm(r, true)
	if err != nil {
		RenderErrorPage(w, fmt.Sprintf("Error Parsing the Form: %v", err))
		return
	}
	// Validate the task
	if err := validators.ValidateTask(task, false); err != nil {
		RenderErrorPage(w, fmt.Sprintf("Invalid Argument creating the task: %v", err))
		return
	}

	// Call the CreateTask method from the server struct
	_, err = s.CreateTask(r.Context(), &pb.TaskRequest{Task: task})
	if err != nil {
		RenderErrorPage(w, err.Error())
		return
	}

	http.Redirect(w, r, "/listTasks", http.StatusSeeOther) // Redirect to the listTasks route

	// // Return the response to the HTTP client
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(res.Task)

}

func CreateTaskPageHandler(w http.ResponseWriter, r *http.Request) {
	// Define the path to your HTML file
	tmpl, err := template.ParseFiles("../frontend/create_task.html")
	if err != nil {
		RenderErrorPage(w, "Failed to load template: "+err.Error())
		return
	}

	// Render the template
	err = tmpl.Execute(w, nil) // Pass any data if needed
	if err != nil {
		RenderErrorPage(w, "Error rendering template: "+err.Error())
	}
}
