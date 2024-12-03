package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"taskify/backend/handlers"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	pb "taskify/backend/proto"
	server "taskify/backend/server"
)

func main() {
	// Initialize the database
	db, err := server.InitializeDatabase()
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer db.Close()
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//
	srv := &server.Server{Db: db}
	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the service
	pb.RegisterTaskServiceServer(grpcServer, srv)

	r := mux.NewRouter()

	r.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTaskHandler(srv, w, r) // Pass server instance to the handler
	}).Methods("POST")
	r.HandleFunc("/createTask", handlers.CreateTaskPageHandler).Methods("GET")

	r.HandleFunc("/listTasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.ListTasksHandler(srv, w, r)
	}).Methods("GET")

	r.HandleFunc("/deleteTask/{taskId}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTaskHandler(srv, w, r)
	}).Methods("GET", "POST")

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.RenderErrorPage(w, "Page not found.")
	})

	go func() {
		fmt.Println("HTTP Server running on port 8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("failed to start HTTP server: %v", err)
		}
	}()

	// Start the server
	fmt.Println("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
