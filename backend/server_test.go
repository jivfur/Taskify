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
			expectedError: status.Error(codes.NotFound, "Exit Criteria was missing"),
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

func TestDeleteTask(t *testing.T) {
	ctx := context.Background()

	db := initializeTestingDatabase(t)
	testServer := server{
		db: db,
	}

	taskReq := &pb.TaskRequest{
		Task: &pb.Task{
			Title:        "Test Task",
			Description:  "This is the task",
			Deadline:     time.Now().Add(1 * time.Hour).Unix(),
			ExitCriteria: "Finish it",
			Complete:     false,
		},
	}
	res, err := testServer.CreateTask(ctx, taskReq)
	if err != nil {
		t.Fatalf("Task could not be stored in the database %v", err)
	}
	delRes, err := testServer.DeleteTask(ctx, &pb.TaskRequest{Task: res.Task})
	if err != nil {
		t.Errorf("The task %d: %s could not be deleted: %v", res.Task.TaskId, res.Task.Title, err)
	}
	if !delRes.Success {
		t.Errorf("Unknown Error Deleting task %d: %s", res.Task.TaskId, res.Task.Title)
	}
}

func TestDeleteTask_Task_Doesnt_Exist(t *testing.T) {
	ctx := context.Background()

	db := initializeTestingDatabase(t)
	testServer := server{
		db: db,
	}

	req := &pb.TaskRequest{
		Task: &pb.Task{
			TaskId:       1,
			Title:        "Test Task",
			Description:  "This is the task",
			Deadline:     time.Now().Add(1 * time.Hour).Unix(),
			ExitCriteria: "Finish it",
			Complete:     false,
		},
	}

	delRes, err := testServer.DeleteTask(ctx, req)
	if err != nil {
		if !strings.Contains(err.Error(), "is not found") {
			t.Errorf("The task %d: %s could not be deleted: %v", req.Task.TaskId, req.Task.Title, err)
		}
	} else {
		if delRes.Success { // It should not delete anything
			t.Errorf("It deleted task %d: %s, but it does not exist", req.Task.TaskId, req.Task.Title)
		}
	}
}

func TestDeleteTask_Delete_Deleted_Task(t *testing.T) {
	ctx := context.Background()

	db := initializeTestingDatabase(t)
	testServer := server{
		db: db,
	}

	taskReq := &pb.TaskRequest{
		Task: &pb.Task{
			Title:        "Test Task",
			Description:  "This is the task",
			Deadline:     time.Now().Add(1 * time.Hour).Unix(),
			ExitCriteria: "Finish it",
			Complete:     false,
		},
	}
	res, err := testServer.CreateTask(ctx, taskReq)

	if err != nil {
		t.Fatalf("Task could not be stored in the database %v", err)
	}
	// Delete the First Time
	delRes, err := testServer.DeleteTask(ctx, &pb.TaskRequest{Task: res.Task})
	if err != nil {
		if !strings.Contains(err.Error(), "is not found") {
			t.Errorf("The task %d: %s could not be deleted: %v", res.Task.TaskId, res.Task.Title, err)
		}
	} else {
		if !delRes.Success {
			t.Errorf("Unknown Error Deleting task %d: %s", res.Task.TaskId, res.Task.Title)
		}
	}
	// Delete the Second Time
	delRes, err = testServer.DeleteTask(ctx, &pb.TaskRequest{Task: res.Task})
	if err != nil {
		if !strings.Contains(err.Error(), "is not found") {
			t.Errorf("The task %d: %s could not be deleted: %v", res.Task.TaskId, res.Task.Title, err)
		}
	} else {
		if delRes.Success { // It should not delete anything
			t.Errorf("It deleted task %d: %s, but it does not exist", taskReq.Task.TaskId, taskReq.Task.Title)
		}
	}
}

func TestDeleteTask_Empty_Task_Id(t *testing.T) {
	ctx := context.Background()

	db := initializeTestingDatabase(t)
	testServer := server{
		db: db,
	}

	taskReq := &pb.TaskRequest{
		Task: &pb.Task{Title: "Test Task", Description: "This is the task", Deadline: time.Now().Add(1 * time.Hour).Unix(), ExitCriteria: "Finish it", Complete: false},
	}

	// Delete the First Time
	delRes, err := testServer.DeleteTask(ctx, taskReq)
	if err != nil {
		if !strings.Contains(err.Error(), "TaskId is empty") {
			t.Errorf("The task %d: %s could not be deleted: %v", taskReq.Task.TaskId, taskReq.Task.Title, err)
		}
	} else {
		if delRes.Success { // It should not delete anything
			t.Errorf("It deleted task %d: %s, but it does not exist", taskReq.Task.TaskId, taskReq.Task.Title)
		}
	}

}

func TestUpdate(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name          string
		task          *pb.Task
		expectedError error
	}{
		{
			name: "happy_path",
			task: &pb.Task{
				Title:        "Test Title Update",
				Description:  "This is the task Update",
				Deadline:     time.Now().Add(2 * time.Hour).Unix(),
				ExitCriteria: "Finish it Update",
				Complete:     true,
			},
		},
		{
			name: "update_title_only",
			task: &pb.Task{
				Title:        "Test Task Update",
				Description:  "This is the task",
				Deadline:     time.Now().Add(1 * time.Hour).Unix(),
				ExitCriteria: "Finish it",
				Complete:     false},
		},
		{
			name: "update_description_only",
			task: &pb.Task{
				Title:        "Test Task",
				Description:  "This is the task_update",
				Deadline:     time.Now().Add(1 * time.Hour).Unix(),
				ExitCriteria: "Finish it",
				Complete:     false},
		},
		{
			name: "update_deadline_only",
			task: &pb.Task{
				Title:        "Test Task",
				Description:  "This is the task",
				Deadline:     time.Now().Add(2 * time.Hour).Unix(),
				ExitCriteria: "Finish it",
				Complete:     false},
		},
		{
			name: "update_exit_criteria_only",
			task: &pb.Task{
				Title:        "Test Task",
				Description:  "This is the task",
				Deadline:     time.Now().Add(2 * time.Hour).Unix(),
				ExitCriteria: "Finish it Update",
				Complete:     false},
		},
		{
			name: "update_complete_only",
			task: &pb.Task{
				Title:        "Test Task",
				Description:  "This is the task",
				Deadline:     time.Now().Add(2 * time.Hour).Unix(),
				ExitCriteria: "Finish it",
				Complete:     true},
		},
		{
			name: "title_mising",
			task: &pb.Task{
				Description:  "This is the task",
				Deadline:     time.Now().Add(1 * time.Hour).Unix(),
				ExitCriteria: "Finish it",
				Complete:     false,
			},
			expectedError: status.Error(codes.NotFound, "Title was missing"),
		},
		{
			name: "description_missing",
			task: &pb.Task{
				Title:        "Test Task",
				Deadline:     time.Now().Add(1 * time.Hour).Unix(),
				ExitCriteria: "Finish it",
				Complete:     false,
			},
			expectedError: status.Error(codes.NotFound, "Description was missing"),
		},
		{
			name: "deadline_missing",
			task: &pb.Task{
				Title:        "Test Task",
				Description:  "This is the task",
				ExitCriteria: "Finish it",
				Complete:     false,
			},
			expectedError: status.Error(codes.NotFound, "Deadline was missing"),
		},
		{
			name: "exit_criteria_missing",
			task: &pb.Task{
				Title:       "Test Task",
				Description: "This is the task",
				Deadline:    time.Now().Add(1 * time.Hour).Unix(),
				Complete:    false,
			},
			expectedError: status.Error(codes.NotFound, "Exit Criteria was missing"),
		},
		{
			name: "no_changes",
			task: &pb.Task{
				Title:        "Test Task",
				Description:  "This is the task",
				Deadline:     time.Now().Add(1 * time.Hour).Unix(),
				ExitCriteria: "Finish it",
				Complete:     false,
			},
			expectedError: status.Error(codes.AlreadyExists, "no changes made"),
		},
		{
			name: "deadline_passed",
			task: &pb.Task{
				Title:        "Test Task",
				Description:  "This is the task",
				Deadline:     time.Now().Add(-1 * time.Hour).Unix(),
				ExitCriteria: "Finish it",
				Complete:     false,
			},
			expectedError: status.Error(codes.InvalidArgument, "Deadline must be in the future"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := initializeTestingDatabase(t)
			testServer := server{
				db: db,
			}
			taskReq := &pb.TaskRequest{
				Task: &pb.Task{Title: "Test Task", Description: "This is the task", Deadline: time.Now().Add(1 * time.Hour).Unix(), ExitCriteria: "Finish it", Complete: false},
			}
			res, err := testServer.CreateTask(ctx, taskReq)
			if err != nil {
				t.Fatalf("The task could not be created: %v", err)
			}
			tc.task.TaskId = res.Task.TaskId
			updateReq := &pb.TaskRequest{
				Task: tc.task,
			}
			resUp, err := testServer.UpdateTask(ctx, updateReq)
			if err != nil {
				if diff := cmp.Diff(tc.expectedError, err, cmpopts.EquateErrors()); diff != "" {
					t.Fatalf("Task %d:%s could not be updated: %v expected %v", updateReq.Task.TaskId, updateReq.Task.Title, err, tc.expectedError)
				}
			} else {
				if diff := cmp.Diff(tc.task, resUp.Task, cmpopts.IgnoreUnexported(pb.Task{})); diff != "" {
					t.Errorf("Update error (+want,-got):%v", diff)
				}
			}
		})
	}

}

func TestListTask(t *testing.T) {
	ctx := context.Background()
	testServer := server{
		db: initializeTestingDatabase(t),
	}

	task := &pb.TaskRequest{
		Task: &pb.Task{
			Title: "Title",
			// Description:  "Description",
			// Deadline:     time.Now().Unix(),
			// ExitCriteria: "Exit Criterial",
			Complete: false,
		},
	}

	_, err := testServer.ListTask(ctx, task)
	if err != nil {
		t.Errorf("ListTask had an error %v", err)
	}
}
