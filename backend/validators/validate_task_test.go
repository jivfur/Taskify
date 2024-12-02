package validators

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "taskify/backend/proto"
)

func TestValidateTask(t *testing.T) {
	testCases := []struct {
		name          string
		task          *pb.Task
		expectedError error
		isUpdate      bool
	}{
		{
			name: "Happy Path",
			task: &pb.Task{
				Title:        "Title",
				Description:  "Description",
				Deadline:     time.Now().Add(24 * time.Hour).Unix(),
				ExitCriteria: "Exit Criteria",
				Complete:     false,
			},
			expectedError: nil,
		},
		{
			name: "missing title",
			task: &pb.Task{
				Description:  "Description",
				Deadline:     time.Now().Add(24 * time.Hour).Unix(),
				ExitCriteria: "Exit Criteria",
				Complete:     false,
			},
			expectedError: status.Error(codes.InvalidArgument, "title is empty"),
		},
		{
			name: "missing description",
			task: &pb.Task{
				Title:        "Title",
				Deadline:     time.Now().Add(24 * time.Hour).Unix(),
				ExitCriteria: "Exit Criteria",
				Complete:     false,
			},
			expectedError: status.Error(codes.InvalidArgument, "description is empty"),
		},
		{
			name: "missing exit criteria",
			task: &pb.Task{
				Title:       "Title",
				Description: "Description",
				Deadline:    time.Now().Add(24 * time.Hour).Unix(),
				Complete:    false,
			},
			expectedError: status.Error(codes.InvalidArgument, "exit criteria is empty"),
		},
		{
			name: "deadline in the past",
			task: &pb.Task{
				Title:        "Title",
				Description:  "Description",
				Deadline:     time.Now().Add(-24 * time.Hour).Unix(),
				ExitCriteria: "Exit Criteria",
				Complete:     false,
			},
			expectedError: status.Error(codes.InvalidArgument, "deadline is set in the past"),
		},
		{
			name: "deadline in the future",
			task: &pb.Task{
				Title:        "Title",
				Description:  "Description",
				Deadline:     time.Now().Add(24 * 50 * 365 * time.Hour).Unix(),
				ExitCriteria: "Exit Criteria",
				Complete:     false,
			},
			expectedError: status.Error(codes.InvalidArgument, "deadline is unreasonably far in the future"),
		},
		{
			name: "complete set to true when creating",
			task: &pb.Task{
				Title:        "Title",
				Description:  "Description",
				Deadline:     time.Now().Add(24 * time.Hour).Unix(),
				ExitCriteria: "Exit Criteria",
				Complete:     true,
			},
			expectedError: status.Error(codes.InvalidArgument, "a new task cannot be marked as complete"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateTask(tc.task, tc.isUpdate)
			if err != nil {
				if diff := cmp.Diff(tc.expectedError, err, cmpopts.EquateErrors()); diff != "" {
					t.Fatalf("ValidateTask(%v):%v, (-want,+got):%v", tc.task, err, diff)
				}
			}
		})

	}

}
