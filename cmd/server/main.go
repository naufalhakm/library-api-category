package main

import (
	"library-api-category/internal/config"
	"library-api-category/internal/factory"
	"library-api-category/internal/routes"
	"library-api-category/pkg/database"
	"log"
)

func main() {
	config.LoadConfig()
	psqlDB, err := database.NewPqSQLClient()
	if err != nil {
		log.Fatal("Could not connect to MySQL:", err)
	}

	provider := factory.InitFactory(psqlDB)

	router := routes.RegisterRoutes(provider)
	log.Printf("Server running on :%s\n", config.ENV.ServerPort)
	log.Fatal(router.Run(":" + config.ENV.ServerPort))
}
