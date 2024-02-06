package main

import (
	"github.com/ezaz-ahmed/EventHub/db"
	"github.com/ezaz-ahmed/EventHub/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InintDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
