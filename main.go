package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"todoapi.wisnu.net/database"
	"todoapi.wisnu.net/master/activity"
	"todoapi.wisnu.net/master/todo"
	// "todoapi.wisnu.net/master/email"
	// "todoapi.wisnu.net/master/user"
)


func main(){
	r := gin.Default()

	database.ConnectDatabase()
	r.Use(CORSMiddleware())
	fmt.Println("Connection to database Establish")

	activityHandler := activity.Activity{}
	todoHandler := todo.Todo{}

	activityRoute := r.Group("/activity-groups")
	{
		activityRoute.POST("",activityHandler.CreateActivity)
		activityRoute.GET("",activityHandler.GetAllActivity)
		activityRoute.GET("/:id",activityHandler.GetOneActivity)
		activityRoute.PATCH("/:id",activityHandler.UpdateActivity)
		activityRoute.DELETE("/:id",activityHandler.DeleteActivity)
	}

	todoRoute := r.Group("/todo-items")
	{
		todoRoute.POST("",todoHandler.CreateTodo)
		todoRoute.GET("",todoHandler.GetAllTodo)
		todoRoute.GET("/:id",todoHandler.GetOneTodo)
		todoRoute.PATCH("/:id",todoHandler.UpdateTodo)
		todoRoute.DELETE("/:id",todoHandler.DeleteTodo)
	}


	
	r.Run(":3030") // listen and serve on 0.0.0.0:8080
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT,DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}