package routes

import (
	"cryptoguess/controllers"

	"github.com/gin-gonic/gin"
)

func ResourcesRoute(router *gin.Engine) {
	router.GET("/resources/coins_today", controllers.ResourceGetCoinsToday())
	router.GET("/resources/archive/:file", controllers.ResourceGetArchiveFile())
}
