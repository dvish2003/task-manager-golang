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
	r.POST("/api/loguser/register", func(c *gin.Context) {
	  var logUser models.LogUser
	  c.BindJSON(&logUser)
	  fmt.Println("Registering user:", logUser.Email)
	  config.DB.Collection("loguser").InsertOne(
		context.TODO(),logUser)
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	})

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
	
	r.POST("/api/user/saveUser",func(c *gin.Context) {
		var user models.User
		c.BindJSON(&user)
		fmt.Println("Saving user:", user.Name)
		config.DB.Collection("users").InsertOne(
			context.TODO(),user)
			c.JSON(http.StatusOK, gin.H{"message": "User saved successfully"})
	})

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

	r.POST("/api/task/saveTask",func(c *gin.Context) {
		var task models.Task
		c.BindJSON(&task)
		fmt.Println("Saving task:", task)
		if task == (models.Task{}) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid task data"})
			return
		}
		config.DB.Collection("tasks").InsertOne(
			context.TODO(),task)
			c.JSON(http.StatusOK, gin.H{"message": "Task saved successfully"})	

	})

	r.GET("/api/task/getTaskByID/:id", func(c *gin.Context) {
		id := c.Param("id")
		var result models.Task
		errorFetch := config.DB.Collection("tasks").FindOne(
			context.TODO(),
			gin.H{"id": id},
		).Decode(&result)

		if errorFetch != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
			return
		}
		c.JSON(http.StatusOK, result)
	})

	r.GET("/api/task/getAllTasks", func(c *gin.Context) {
		cursor, err := config.DB.Collection("tasks").Find(context.TODO(), gin.H{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching tasks"})
			return
		}
		var tasks []models.Task
		if err = cursor.All(context.TODO(), &tasks); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error decoding tasks"})
			return
		}
		c.JSON(http.StatusOK, tasks)
	})
}