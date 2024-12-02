package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	server "taskify/backend/server"
	validators "taskify/backend/validators"

	pb "taskify/backend/proto"
)

func CreateTaskHandler(s *server.Server, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Get the deadline string from the form input (assuming it's in a format like "2024-12-03T10:30")
	deadlineStr := r.FormValue("deadline")

	// Parse the deadline string into a time.Time object (adjust format as necessary)
	deadline, err := time.Parse("2006-01-02T15:04", deadlineStr) // Adjust format if needed
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid deadline format: %v", err), http.StatusBadRequest)
		return
	}

	// Create a TaskRequest from the form data
	taskReq := &pb.TaskRequest{
		Task: &pb.Task{
			Title:        r.FormValue("title"),
			Description:  r.FormValue("description"),
			ExitCriteria: r.FormValue("exitCriteria"),
			Deadline:     deadline.Unix(),
			Complete:     false,
		},
	}

	// Validate the task
	if err := validators.ValidateTask(taskReq.Task, false); err != nil {
		http.Error(w, fmt.Sprintf("Invalid Argument creating the task: %v", err), http.StatusBadRequest)
		return
	}

	// Call the CreateTask method from the server struct
	res, err := s.CreateTask(r.Context(), taskReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response to the HTTP client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.Task)

}
