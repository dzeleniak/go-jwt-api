package main

import (
	"github.com/dzeleniak/jwt-api/controllers"
	"github.com/dzeleniak/jwt-api/initializers"
	"github.com/dzeleniak/jwt-api/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
