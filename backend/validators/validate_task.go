package validators

import (
	"strings"
	pb "taskify/backend/proto"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateTask(task *pb.Task, isUpdate bool) error {
	if len(strings.TrimSpace(task.Title)) == 0 {
		return status.Error(codes.InvalidArgument, "title is empty")
	}
	if len(strings.TrimSpace(task.Description)) == 0 {
		return status.Error(codes.InvalidArgument, "description is empty")
	}
	if len(strings.TrimSpace(task.ExitCriteria)) == 0 {
		return status.Error(codes.InvalidArgument, "exit criteria is empty")
	}

	if time.Now().Add(-5*time.Minute).Unix() > task.Deadline {
		return status.Error(codes.InvalidArgument, "deadline is set in the past")
	}
	if task.Deadline > time.Now().Add(10*365*24*time.Hour).Unix() {
		return status.Error(codes.InvalidArgument, "deadline is unreasonably far in the future")
	}
	// Validate Complete Status (only for create, skip for updates)
	if !isUpdate && task.Complete {
		return status.Error(codes.InvalidArgument, "a new task cannot be marked as complete")
	}
	return nil
}
