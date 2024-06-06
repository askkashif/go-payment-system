package server

import (
	"context"                                 // Importing context package for managing server shutdown
	"fmt"                                     // Importing fmt package for formatted I/O
	"github.com/joho/godotenv"                // Importing godotenv package to load environment variables from a .env file
	"gorm.io/gorm"                            // Importing gorm package for ORM (Object-Relational Mapping)
	"log"                                     // Importing log package for logging
	"net/http"                                // Importing net/http package for HTTP server
	"os"                                      // Importing os package for OS-related functionality
	"os/signal"                               // Importing os/signal package to handle OS signals
	"payment-system-four/internal/api"        // Importing the package containing API handlers
	"payment-system-four/internal/repository" // Importing the package containing repository implementations
	"time"                                    // Importing time package for handling durations
)

// Run injects all dependencies needed to run the app
func Run(db *gorm.DB, port string) {
	// Initialize a new repository using the provided database connection
	newRepo := repository.NewDB(db)

	// Create a new HTTP handler using the repository
	Handler := api.NewHTTPHandler(newRepo)

	// Set up the router with the handler and repository
	router := SetupRouter(Handler, newRepo)

	// Create an HTTP server with the specified port and router
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Printf("Listening and serving HTTP on : %v\n", port)

	// Start the server in a new goroutine
	go func() {
		// Listen and serve HTTP requests
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Create a channel to listen for OS signals
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt) // Notify for Interrupt signal (Ctrl+C)
	signal.Notify(sigChan, os.Kill)      // Notify for Kill signal

	// Block until a signal is received
	sig := <-sigChan
	log.Println("Receive terminate and shutdown gracefully", sig)

	// Create a context with a timeout for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

// Params is a data model of the data in our environment variable
type Params struct {
	Port  string // Port on which the server will run
	DbUrl string // Database URL
}

// InitDBParams gets environment variables needed to run the app
func InitDBParams() Params {
	// Load environment variables from a .env file
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println("Error loading .env file")
	}

	// Get the port from environment variables, default to "8080" if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Get the database URL from environment variables
	dbURL := os.Getenv("DATABASE_URL")

	// Return the Params struct with the loaded values
	return Params{
		Port:  port,
		DbUrl: dbURL,
	}
}
