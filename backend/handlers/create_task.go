package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	pb "taskify/backend/proto"
	server "taskify/backend/server"
	validators "taskify/backend/validators"
)

func CreateTaskHandler(s *server.Server, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	task, err := ParseForm(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Parsing the Form: %v", err), http.StatusBadRequest)
		return
	}
	// Validate the task
	if err := validators.ValidateTask(task, false); err != nil {
		http.Error(w, fmt.Sprintf("Invalid Argument creating the task: %v", err), http.StatusBadRequest)
		return
	}

	// Call the CreateTask method from the server struct
	res, err := s.CreateTask(r.Context(), &pb.TaskRequest{Task: task})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response to the HTTP client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.Task)

}
