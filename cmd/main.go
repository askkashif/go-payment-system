package main

import (
	"payment-system-four/cmd/server"
	"payment-system-four/internal/repository"
)

func main() {
	// Gets the environment variables required for database connection and server configuration
	env := server.InitDBParams()

	// Initializes the database connection using the provided URL
	db, err := repository.Initialize(env.DbUrl)
	if err != nil {
		return // Exit the program if there's an error initializing the database
	}

	// Runs the server application, passing the initialized database connection and server port
	server.Run(db, env.Port)
}

// Comments are extremely important when writing functions to explain their purpose and usage.

// CRUD Operations:
// CREATE - Creating an entry in the database, e.g., sign up [POST]
// READ - Retrieving information from the database, e.g., display of profile information [GET]
// UPDATE - Updating information in the database, e.g., change of address [PUT/PATCH]
// DELETE - Deleting information from the database, e.g., deletion of a profile [DELETE]

// Download Postman (https://www.postman.com/downloads/) to interact with HTTP APIs.

// Read up about HTTP Status Codes to understand the meaning of different status codes returned by HTTP responses.
