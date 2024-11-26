package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc" // gRPC package
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "taskify/backend/proto" // Import your proto package (path should match where task.pb.go is located)

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type server struct {
	pb.UnimplementedTaskServiceServer         // Embedding the Unimplemented service for forward compatibility
	db                                *sql.DB // Database
}

func initializeDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "database/taskify.db") // Update the path if needed
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	schema, err := os.ReadFile("database/init.sql") // Read schema from init.sql
	if err != nil {
		return nil, fmt.Errorf("failed to read schema file: %w", err)
	}

	_, err = db.Exec(string(schema)) // Execute schema
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database schema: %w", err)
	}

	return db, nil

}

// Validate task is complete

func (s *server) validateTask(ctx context.Context, task *pb.Task) error {
	if len(strings.TrimSpace(task.Title)) == 0 {
		return status.Error(codes.NotFound, "Title was missing")
	}

	if len(strings.TrimSpace(task.Description)) == 0 {
		return status.Error(codes.NotFound, "Description was missing")
	}

	if len(strings.TrimSpace(task.Title)) == 0 {
		return status.Error(codes.NotFound, "Title was missing")
	}

	if len(strings.TrimSpace(task.ExitCriteria)) == 0 {
		return status.Error(codes.NotFound, "Exist Criteria was missing")
	}

	if task.Deadline == 0 {
		return status.Error(codes.NotFound, "Deadline was missing")
	}

	if task.Deadline <= time.Now().Unix() {
		return status.Error(codes.InvalidArgument, "Deadline must be in the future")
	}
	return nil

}

func (s *server) GetTask(ctx context.Context, id int64) (*pb.Task, error) {
	var taskId, deadline, complete int
	var title, description, exitCriteria string
	err := s.db.QueryRow("SELECT * FROM tasks WHERE taskId = ?", id).Scan(&taskId, &title, &description, &deadline, &exitCriteria, &complete)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "task %d not found %v", id, err)
		} else {
			return nil, status.Errorf(codes.Aborted, "fatal error: %v", err)
		}
	}
	return &pb.Task{
		TaskId:       int64(taskId),
		Title:        title,
		Description:  description,
		Deadline:     int64(deadline),
		ExitCriteria: exitCriteria,
		Complete:     complete == 1,
	}, nil
}

// CreateTask will store the TaskRequest in the Database
func (s *server) CreateTask(ctx context.Context, in *pb.TaskRequest) (*pb.TaskResponse, error) {
	if in == nil {
		return nil, status.Error(codes.NotFound, "task is nil")
	}

	err := s.validateTask(ctx, in.Task)
	if err != nil {
		return nil, err
	}

	// Prepare the INSERT statement
	query := `INSERT INTO tasks (title, description, deadline, exitCriteria, complete) 
		VALUES (?, ?, ?, ?, ?)`

	// Execute the insert query
	task := in.Task
	res, err := s.db.Exec(query, task.Title, task.Description, task.Deadline, task.ExitCriteria, task.Complete)
	if err != nil {
		return nil, err
	}

	taskId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	task, err = s.GetTask(ctx, taskId)
	if err != nil {
		return nil, err
	}

	return &pb.TaskResponse{
		Task: task,
	}, nil
}

// Updatetask will store update the TaskRequest in the Database
func (s *server) UpdateTask(ctx context.Context, in *pb.TaskRequest) (*pb.TaskResponse, error) {
	return nil, status.Error(codes.Unimplemented, "UpdateTask method is not implemented yet")
}

// DeleteTask will delete the task from the database
func (s *server) DeleteTask(ctx context.Context, in *pb.TaskRequest) (*pb.DeleteTaskResponse, error) {
	return nil, status.Error(codes.Unimplemented, "DeleteTask method is not implemented yet")
}

// ListTask retrieves all the tasks, filtered by dates, status, etc.
func (s *server) ListTask(ctx context.Context, in *pb.ListTaskRequest) (*pb.ListTaskResponse, error) {
	return nil, status.Error(codes.Unimplemented, "ListTask method is not implemented yet")
}

func main() {
	// Initialize the database
	db, err := initializeDatabase()
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer db.Close()
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the service
	pb.RegisterTaskServiceServer(s, &server{db: db})

	// Start the server
	fmt.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
