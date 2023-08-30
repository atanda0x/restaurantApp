package routes

import (
	"github.com/atanda0x/restaurantApp/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUser())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.Signup())
	incomingRoutes.POST("/users/login", controller.Login())

}
