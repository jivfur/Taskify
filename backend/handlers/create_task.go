package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"taskify/backend/validators"

	pb "taskify/backend/proto"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
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

	taskReq := &pb.TaskRequest{
		Task: task,
	}

	server := main.server

}
