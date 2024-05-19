package main

import (
	"payment-system-one/cmd/server"
	"payment-system-one/internal/repository"
)

func main() {
	//Gets the environment variables
	env := server.InitDBParams()

	//Initializes the database
	db, err := repository.Initialize(env.DbUrl)
	if err != nil {
		return
	}

	//Runs the app
	server.Run(db, env.Port)
}

//comments are extremely important when writing functions

//CRUD nam

//CREATE -  creating an entry in the database e.g sign up    [POST]
//READ - get an information from the database e.g display of profile information [GET]
//UPDATE - updates information in the database e.g change of address [PUT/PATCH]
//DELETE - delete an information from the database e.g  deletion of profile [DELETE]

//downmload Postman

//Read up about http Status Codes
