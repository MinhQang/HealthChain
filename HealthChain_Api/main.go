package main

import (
	"HealthChain_API/config"
	"HealthChain_API/migration"
	"HealthChain_API/routes"
	"log"
	"net/http"
)

func main() {
	config.InitDB()
	migration.Migration()
	r := routes.InitializeRoutes()

	log.Fatal(http.ListenAndServe(":8080", r))
}
