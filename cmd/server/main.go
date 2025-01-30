package main

import (
	"library-api-category/internal/config"
	"library-api-category/internal/factory"
	"library-api-category/internal/grpc/client"
	"library-api-category/internal/routes"
	"library-api-category/pkg/database"
	"log"
	"sync"
)

func main() {
	config.LoadConfig()
	psqlDB, err := database.NewPqSQLClient()
	if err != nil {
		log.Fatal("Could not connect to PqSQL:", err)
	}

	provider := factory.InitFactory(psqlDB)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		runHTTPServer(provider)
	}()

	wg.Wait()
}

func runHTTPServer(provider *factory.Provider) {
	authClient, err := client.NewAuthClient(config.ENV.UserGRPC)
	if err != nil {
		log.Fatalf("Failed to initialize auth client: %v", err)
	}
	defer authClient.Close()

	router := routes.RegisterRoutes(provider, authClient)
	log.Printf("REST API server running on port %s\n", config.ENV.ServerPort)
	log.Fatal(router.Run(":" + config.ENV.ServerPort))
}
