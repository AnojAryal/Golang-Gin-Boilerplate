package main

import (
	"github.com/anojaryal/JWT-Authentication/initializers"
	"github.com/anojaryal/JWT-Authentication/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()

}

func main() {
	r := gin.Default()

	routes.UserRoutes(r)

	r.Run()
}
