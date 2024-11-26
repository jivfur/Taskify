package main

import (
	"context"
	"database/sql"
	"os"
	"strings"
	"testing"
	"time"

	pb "taskify/backend/proto"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts" // Import cmpopts for IgnoreFields
	_ "github.com/mattn/go-sqlite3"        // SQLite driver
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func initializeTestingDatabase(t *testing.T) *sql.DB {
	t.Helper() // Marks this function as a helper for better test failure output
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	// Read the schema from the init.sql file
	schema, err := os.ReadFile("../database/init.sql") // Adjust the path as needed
	if err != nil {
		t.Fatalf("Failed to read init.sql: %v", err)
	}

	// Execute the schema to set up the database
	_, err = db.Exec(string(schema))
	if err != nil {
		t.Fatalf("Failed to initialize test database schema: %v", err)
	}
	return db
}

func TestCreateTask(t *testing.T) {
	ctx := context.Background()
	db := initializeTestingDatabase(t)
	testServer := server{
		db: db,
	}

	testCases := []struct {
		name          string
		task          *pb.Task
		expectedError error
	}{
		{
			name: "happy_path",
			task: &pb.Task{
				Title:        "Test Task",
				Description:  "This is the task",
				Deadline:     time.Now().Add(1 * time.Hour).Unix(),
				ExitCriteria: "Finish it",
				Complete:     false,
			},
		},
		{
			name: "missing_description",
			task: &pb.Task{
				Title:        "Test Task",
				Description:  "",
				Deadline:     time.Now().Add(1 * time.Hour).Unix(),
				ExitCriteria: "Finish it",
				Complete:     false,
			},
			expectedError: status.Error(codes.NotFound, "Description was missing"),
		},
		{
			name: "missing_title",
			task: &pb.Task{
				Description:  "This is the task",
				Deadline:     time.Now().Add(1 * time.Hour).Unix(),
				ExitCriteria: "Finish it",
				Complete:     false,
			},
			expectedError: status.Error(codes.NotFound, "Title was missing"),
		},
		{
			name: "missing_exit_criteria",
			task: &pb.Task{
				Title:       "Test Task",
				Description: "This is the task",
				Deadline:    time.Now().Add(1 * time.Hour).Unix(),
				Complete:    false,
			},
			expectedError: status.Error(codes.NotFound, "Exist Criteria was missing"),
		},
		{
			name: "missing_deadline",
			task: &pb.Task{
				Title:        "Test Task",
				Description:  "This is the task",
				ExitCriteria: "Finish it",
				Complete:     false,
			},
			expectedError: status.Error(codes.NotFound, "Deadline was missing"),
		},
		{
			name: "deadline_passed",
			task: &pb.Task{
				Title:        "Test Task",
				Description:  "This is the task",
				Deadline:     time.Now().Add(-24 * time.Hour).Unix(),
				ExitCriteria: "Finish it",
				Complete:     false,
			},
			expectedError: status.Error(codes.InvalidArgument, "Deadline must be in the future"),
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				req := &pb.TaskRequest{
					Task: tc.task,
				}
				res, actualError := testServer.CreateTask(ctx, req)
				if actualError != nil {
					if diff := cmp.Diff(tc.expectedError, actualError, cmpopts.EquateErrors()); diff != "" {
						t.Fatalf("The task couldn't be created %v", actualError)
					}

				} else {
					if diff := cmp.Diff(req.Task, res.Task, cmpopts.IgnoreFields(pb.Task{}, "TaskId", "Deadline"), cmpopts.IgnoreUnexported(pb.Task{})); diff != "" {
						t.Errorf("Task could not be created (+want,-got) %v", diff)
					}
				}
			},
		)

	}
}

func TestCreateTask_DuplicateTask(t *testing.T) {
	ctx := context.Background()

	db := initializeTestingDatabase(t)
	testServer := server{
		db: db,
	}

	req := &pb.TaskRequest{
		Task: &pb.Task{
			Title:        "Test Task",
			Description:  "This is the task",
			Deadline:     time.Now().Add(1 * time.Hour).Unix(),
			ExitCriteria: "Finish it",
			Complete:     false,
		},
	}

	_, err := testServer.CreateTask(ctx, req) // Storing the task for the first time
	if err != nil {
		t.Fatalf("The task could not be created %v", err)
	}

	_, err = testServer.CreateTask(ctx, req) // Trying to store the task for a second time
	if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
		t.Fatalf("There's a bigger issue trying to create a task in the DB %v", err)
	}

}
