package main


import (
    "task-manager/backend/config"
    "task-manager/backend/routes"
    "github.com/gin-gonic/gin"
)


func main() {
	//initialize DB
	config.ConnectDB()
	//initialize router
	r := gin.Default()
	//register routes
	routes.RegisterRoutes(r)
	//run server
	r.Run(":8080")
}
