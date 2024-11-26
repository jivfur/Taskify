package main

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"testing"
	"time"

	pb "taskify/backend/proto"

	"github.com/google/go-cmp/cmp"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
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
	testServer := &pb.server{
		db: db,
	}

	testCases := []struct {
		name          string
		task          *pb.Task
		expectedError error
	}{
		{
			Name: "happy_path",
			Task: &pb.Task{
				Title:        "Test Task",
				Description:  "This is the task",
				Deadline:     time.Now().Unix(),
				ExitCriteria: "Finish it",
				Complete:     false,
			},
		},
		{
			Name: "missing_description",
			Task: &pb.Task{
				Title:        "Test Task",
				Deadline:     time.Now().Unix(),
				ExitCriteria: "Finish it",
				Complete:     false,
			},
			expectedError: errors.New("Description was missing"),
		},
		{
			Name: "missing_title",
			Task: &pb.Task{
				Description:  "This is the task",
				Deadline:     time.Now().Unix(),
				ExitCriteria: "Finish it",
				Complete:     false,
			},
			expectedError: errors.New("Title was missing"),
		},
		{
			Name: "missing_exit_criteria",
			Task: &pb.Task{
				Title:       "Test Task",
				Description: "This is the task",
				Deadline:    time.Now().Unix(),
				Complete:    false,
			},
			expectedError: errors.New("Exit Criteria was missing"),
		},
		{
			Name: "deadline_passed",
			Task: &pb.Task{
				Title:        "Test Task",
				Description:  "This is the task",
				Deadline:     time.Now().Add(-24 * time.Hour).Unix(),
				ExitCriteria: "Finish it",
				Complete:     false,
			},
			expectedError: errors.New("Deadline already passed"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := &pb.TaskRequest{
				Task: tc.task,
			}
			res, actualError := testServer.CreateTask(ctx, tc.task)
			if actualError != nil {
				if diff := cmp.Diff(tc.wantError, actualError); diff != "" {
					t.Fatalf("The task couldn't be created %v", actualError)
				}

			}
			if diff := cmp.Diff(req, res, cmp.IgnoreFields(pb.TaskRequest{}, "Deadline")); diff != "" {
				t.Errorf("Task could not be created (+want,-got) %v", diff)
			}
		})

	}
}
