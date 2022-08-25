package routes

import (
	"cryptoguess/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	router.POST("/auth/login", controllers.AuthLogin())
	router.GET("/auth/logout", controllers.AuthLogout())
	router.GET("/auth/test", controllers.AuthTest())
}
