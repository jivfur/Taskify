package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	server "taskify/backend/server"
	validators "taskify/backend/validators"

	pb "taskify/backend/proto"
)

func CreateTaskHandler(s *server.Server, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Read the entire body content into a byte slice
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Declare a TaskRequest struct to hold the decoded data
	var taskReq pb.TaskRequest
	// Decode the JSON into our taskReq struct
	err = json.Unmarshal(body, &taskReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the task
	if err = validators.ValidateTask(taskReq.Task, false); err != nil {
		http.Error(w, fmt.Sprintf("Invalid Argument creating the task: %v", err), http.StatusBadRequest)
	}

	// Call the CreateTask method from the server struct
	res, err := s.CreateTask(r.Context(), &taskReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response to the HTTP client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.task)

}
