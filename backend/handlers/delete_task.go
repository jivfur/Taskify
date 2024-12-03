package handlers

import (
	"net/http"
	"strconv"

	pb "taskify/backend/proto"
	server "taskify/backend/server"
)

func DeleteTaskHandler(s *server.Server, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	taskId := vars["taskId"] // This is the TaskId part from the URL

	// Convert the taskId to an integer (assuming taskId is an int64)
	id, err := strconv.ParseInt(taskId, 10, 64)
	if err != nil {
		// Handle error if the TaskId is invalid
		RenderErrorPage(w, "Invalid Task ID")
		return
	}

	task := &pb.Task{
		TaskId: id,
	}
	// Call the CreateTask method from the server struct
	_, err = s.DeleteTask(r.Context(), &pb.TaskRequest{Task: task})
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
