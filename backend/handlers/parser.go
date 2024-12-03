package handlers

import (
	"fmt"
	"net/http"
	pb "taskify/backend/proto"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ParseForm(r *http.Request) (*pb.Task, error) {
	if err := r.ParseForm(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to parse form data")
	}

	// Get the deadline string from the form input (assuming it's in a format like "2024-12-03T10:30")
	deadlineStr := r.FormValue("deadline")

	// Parse the deadline string into a time.Time object (adjust format as necessary)
	deadline, err := time.Parse("2006-01-02T15:04", deadlineStr) // Adjust format if needed
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Invalid deadline format: %v", err))
	}

	complete := false
	if r.FormValue("complete") == "true" {
		complete = true
	}

	// Create a TaskRequest from the form data

	return &pb.Task{
		Title:        r.FormValue("title"),
		Description:  r.FormValue("description"),
		ExitCriteria: r.FormValue("exitCriteria"),
		Deadline:     deadline.Unix(),
		Complete:     complete,
	}, nil

}
