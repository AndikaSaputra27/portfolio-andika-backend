package main

import (
	"final-project/config"
	"final-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	routes.SetupRoutes(r)
	r.Run(":8080") // http://localhost:8080
}
