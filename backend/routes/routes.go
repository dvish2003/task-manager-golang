package routes

import (
	"context"
	"fmt"
	"net/http"
	"task-manager/backend/config"
	"task-manager/backend/models"

	"github.com/gin-gonic/gin"
)



func RegisterRoutes(r *gin.Engine){
	//register 
	r.POST("/api/loguser/register", func(c *gin.Context) {
	  var logUser models.LogUser
	  c.BindJSON(&logUser)
	  fmt.Println("Registering user:", logUser.Email)
	  config.DB.Collection("loguser").InsertOne(
		context.TODO(),logUser)
		//send message
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	})
	//Login logUser
	r.POST("/api/loguser/login", func(c *gin.Context) {
		var logUser models.LogUser
		c.BindJSON(&logUser)
		var result models.LogUser
		errorFetch := config.DB.Collection("loguser").FindOne(
			context.TODO(),
			gin.H{"email": logUser.Email, "password": logUser.Password},
		).Decode(&result)

		if errorFetch != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})
	
	//create User
	//getUserByEmail
	//getAll Users


	//create Task
	//getTaskByID
	//getAll Tasks
}