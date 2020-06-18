package main

import (
	"log"

	"github.com/malbbako/ng_state_ward/configs"
	"github.com/malbbako/ng_state_ward/database"
	"github.com/malbbako/ng_state_ward/models"
	"github.com/malbbako/ng_state_ward/repositories"
)

func main() {
	dbUser, dbPassword, dbName := "root", "", "ng_state_ward"
	db, err := database.ConnectToDB(dbUser, dbPassword, dbName)

	//Unable to connect to database

	if err != nil {
		log.Fatalln(err)
	}

	//Ping Database
	err = db.DB().Ping()

	//error ping
	if err != nil {
		log.Fatalln(err)
	}

	//Migrations

	db.AutoMigrate(&models.State{})
	defer db.Close()

	stateRepository := repositories.NewStateRepository(db)
	route := configs.SetupRoutes(stateRepository)
	route.Run(":8000")
}
