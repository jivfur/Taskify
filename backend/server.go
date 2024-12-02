package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc" // gRPC package
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"taskify/backend/handlers"
	pb "taskify/backend/proto" // Import your proto package (path should match where task.pb.go is located)

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type Server struct {
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

func (s *Server) validateTask(ctx context.Context, task *pb.Task) error {
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
		return status.Error(codes.NotFound, "Exit Criteria was missing")
	}

	if task.Deadline == 0 {
		return status.Error(codes.NotFound, "Deadline was missing")
	}

	if task.Deadline <= time.Now().Unix() {
		return status.Error(codes.InvalidArgument, "Deadline must be in the future")
	}
	return nil

}

func (s *Server) GetTask(ctx context.Context, id int64) (*pb.Task, error) {
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
func (s *Server) CreateTask(ctx context.Context, in *pb.TaskRequest) (*pb.TaskResponse, error) {
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
func (s *Server) UpdateTask(ctx context.Context, in *pb.TaskRequest) (*pb.TaskResponse, error) {
	err := s.validateTask(ctx, in.Task)
	if err != nil {
		return nil, err
	}

	task, err := s.GetTask(ctx, in.Task.TaskId)
	if err != nil {
		return nil, fmt.Errorf("retrieving the task: %v", err)
	}

	if diff := cmp.Diff(in.Task, task, cmpopts.IgnoreUnexported(pb.Task{})); diff == "" {
		return nil, status.Error(codes.AlreadyExists, "no changes made")
	}

	query := "UPDATE tasks SET title = ?, description = ?, deadline = ?, exitCriteria = ?, complete = ? WHERE taskId = ?;"

	res, err := s.db.Exec(query, in.Task.Title, in.Task.Description, in.Task.Deadline, in.Task.ExitCriteria, in.Task.Complete, in.Task.TaskId)
	if err != nil {
		return nil, err
	}
	lastAffectedRow, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	task, err = s.GetTask(ctx, lastAffectedRow)
	if err != nil {
		return nil, fmt.Errorf("retrieving the task: %v", err)
	}

	return &pb.TaskResponse{Task: task}, nil
}

// DeleteTask will delete the task from the database
func (s *Server) DeleteTask(ctx context.Context, in *pb.TaskRequest) (*pb.DeleteTaskResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "Task is nil")
	}
	if in.Task.TaskId == 0 {
		return nil, status.Error(codes.InvalidArgument, "TaskId is empty")
	}

	// Delete Query

	query := `DELETE FROM tasks WHERE  taskId = ?`
	res, err := s.db.Exec(query, in.Task.TaskId)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &pb.DeleteTaskResponse{Success: rowsAffected == 1}, nil
}

// ListTask retrieves all the tasks, filtered by dates, status, etc.
func (s *Server) ListTask(ctx context.Context, in *pb.TaskRequest) (*pb.ListTaskResponse, error) {
	var whereClause []string
	if len(strings.TrimSpace(in.Task.Title)) != 0 {
		whereClause = append(whereClause, "title LIKE %"+strings.TrimSpace(in.Task.Title)+"%")
	}
	log
	if len(strings.TrimSpace(in.Task.Description)) != 0 {
		whereClause = append(whereClause, "description LIKE %"+strings.TrimSpace(in.Task.Description)+"%")
	}

	if len(strings.TrimSpace(in.Task.ExitCriteria)) != 0 {
		whereClause = append(whereClause, "exitCriteria LIKE %"+strings.TrimSpace(in.Task.ExitCriteria)+"%")
	}

	if in.Task.Deadline != 0 {
		whereClause = append(whereClause, "deadline = "+fmt.Sprintf("%d", in.Task.Deadline))
	}

	if in.Task.Complete {
		whereClause = append(whereClause, "complete = 1")
	}

	query := "SELECT * FROM tasks "
	if len(whereClause) > 0 {
		query += "WHERE "
	}
	query = query + strings.Join(whereClause, " AND ") + " ORDER BY taskId ASC"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Failed to query tasks: %v", err))
	}
	defer rows.Close() // Ensure the rows are properly closed when done.

	var tasks []*pb.Task
	for rows.Next() {
		var taskId int
		var title, description, exitCriteria string
		var deadline int64
		var complete bool

		// Adjust the scan parameters based on your database schema
		if err := rows.Scan(&taskId, &title, &description, &deadline, &exitCriteria, &complete); err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("Failed to scan row: %v", err))
		}

		tasks = append(tasks, &pb.Task{
			TaskId:       int64(taskId),
			Title:        title,
			Description:  description,
			ExitCriteria: exitCriteria,
			Deadline:     deadline,
			Complete:     complete,
		})
	}

	return &pb.ListTaskResponse{Tasks: tasks}, nil
}

// Get Completed Tasks.
func (s *Server) CompletedTasks(ctx context.Context, in *pb.TaskRequest) (*pb.ListTaskResponse, error) {
	rows, err := s.db.Query("SELECT * FROM tasks WHERE complete = 1")
	if err != nil {
		log.Fatalf("Failed to query tasks: %v", err)
	}
	defer rows.Close() // Ensure the rows are properly closed when done.

	var completedTasks []*pb.Task
	for rows.Next() {
		var taskId int
		var title, description, exitCriteria string
		var deadline int64
		var complete bool

		// Adjust the scan parameters based on your database schema
		if err := rows.Scan(&taskId, &title, &description, &deadline, &exitCriteria, &complete); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}

		completedTasks = append(completedTasks, &pb.Task{
			TaskId:       int64(taskId),
			Title:        title,
			Description:  description,
			ExitCriteria: exitCriteria,
			Deadline:     deadline,
			Complete:     complete,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error encountered during iteration: %v", err))
	}
	return &pb.ListTaskResponse{Tasks: completedTasks}, nil
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
	pb.RegisterTaskServiceServer(s, &Server{db: db})

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTaskHandler(s, w, r) // Pass server instance to the handler
	})

	// Start the server
	fmt.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
