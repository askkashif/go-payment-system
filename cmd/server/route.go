package server

import (
	"github.com/gin-contrib/cors"            // Importing CORS middleware for handling Cross-Origin Resource Sharing
	"github.com/gin-gonic/gin"               // Importing the Gin framework for creating the HTTP server
	"payment-system-one/internal/api"        // Importing the package containing API handlers
	"payment-system-one/internal/middleware" // Importing the package containing middleware
	"payment-system-one/internal/ports"      // Importing the package containing port interfaces for dependency injection
	"time"                                   // Importing the time package for handling durations
)

// SetupRouter is where router endpoints are configured
func SetupRouter(handler *api.HTTPHandler, repository ports.Repository) *gin.Engine {
	// Create a default Gin router
	router := gin.Default()

	// Configure CORS settings to allow cross-origin requests
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                     // Allow all origins
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"}, // Allow specified HTTP methods
		AllowHeaders:     []string{"*"},                                     // Allow all headers
		ExposeHeaders:    []string{"Content-Length"},                        // Expose Content-Length header
		AllowCredentials: true,                                              // Allow credentials
		MaxAge:           12 * time.Hour,                                    // Cache preflight request for 12 hours
	}))

	// Create a route group for the root path
	r := router.Group("/")
	{
		// Define a GET endpoint for readiness check
		r.GET("/", handler.Readiness)
		// Define a POST endpoint to create a new user
		r.POST("/create", handler.CreateUser)
		// Define a POST endpoint for user login
		r.POST("/login", handler.LoginUer)
	}

	// Create a route group for admin endpoints
	authorizeAdmin := r.Group("/admin")
	// Apply middleware to the admin group to authorize admin users
	authorizeAdmin.Use(middleware.AuthorizeAdmin(repository.FindUserByEmail, repository.TokenInBlacklist))
	{
		// Define a GET endpoint to retrieve user information by email
		authorizeAdmin.GET("/user", handler.GetUserByEmail)
	}

	// Return the configured router
	return router
}
