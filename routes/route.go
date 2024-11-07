package routes

import (
	"github.com/AllergySnipe/go-auth-api/controllers"
	"github.com/AllergySnipe/go-auth-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)

	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/protectedroute", controllers.ProtectedRoute)
	auth.POST("/revoke", controllers.RevokeToken)
	auth.POST("/refresh", controllers.RefreshToken)

	return r
}
