package routes

import (
	"cryptoguess/controllers"

	"github.com/gin-gonic/gin"
)

func RootRoute(router *gin.Engine) {
	router.GET("/version", controllers.RootVersion())
	router.POST("/upload_result", controllers.RootUploadResult())
}
