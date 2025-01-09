package routes

import (
	"github.com/anojaryal/JWT-Authentication/controllers"
	"github.com/anojaryal/JWT-Authentication/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
}
