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
	r.POST("/api/user/saveUser",func(c *gin.Context) {
		var user models.User
		c.BindJSON(&user)
		fmt.Println("Saving user:", user.Name)
		config.DB.Collection("users").InsertOne(
			context.TODO(),user)
			//send message
			c.JSON(http.StatusOK, gin.H{"message": "User saved successfully"})
	})
	//getUserByEmail
	r.GET("/api/user/getUserByEmail/:email", func(c *gin.Context) {
		email := c.Param("email")
		var result models.User
		errorFetch := config.DB.Collection("users").FindOne(
			context.TODO(),
			gin.H{"email": email},
		).Decode(&result)

		if errorFetch != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		c.JSON(http.StatusOK, result)
	})
	//getAll Users
	r.GET("/api/user/getAllUsers", func(c *gin.Context) {
		cursor, err := config.DB.Collection("users").Find(context.TODO(), gin.H{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching users"})
			return
		}
		var users []models.User
		if err = cursor.All(context.TODO(), &users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error decoding users"})
			return
		}
		c.JSON(http.StatusOK, users)
	})


	//create Task
	//getTaskByID
	//getAll Tasks
}