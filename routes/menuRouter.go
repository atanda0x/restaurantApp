package routes

import (
	controller "github.com/atanda0x/restaurantApp/controllers"
	"github.com/gin-gonic/gin"
)

func MemuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.GetMenu())
	incomingRoutes.GET("/memus/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.CreateMenu)
	incomingRoutes.PATCH("/menus/:menu_id", controller.UpdateMenu())
}
