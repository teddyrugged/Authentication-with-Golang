package routes

import (
	controllers "golang/controllers"
	authMiddleware "golang/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// user has to be authenticated to use this route
	incomingRoutes.Use(authMiddleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
