package main

import (
	"github.com/anojaryal/JWT-Authentication/controllers"
	"github.com/anojaryal/JWT-Authentication/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
	
}

func main() {
	r := gin.Default()
	r.POST("/signup",controllers.SignUp)
	r.POST("/login",controllers.Login)
	
	r.Run() 
}