package main

import (
	"cryptoguess/configs"
	"cryptoguess/routes"
	"cryptoguess/scripts"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

func main() {
	//* Redis
	configs.ConnectRedis()

	//* Archive configuration
	configs.SetupArchive()
	scripts.UpdateToday()

	//* Scheduled tasks
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(15).Minutes().Do(func() { scripts.UpdateCoinPrices() })
	scheduler.Every(1).Day().At("00:00").Do(func() { scripts.UpdateToday() })
	scheduler.StartAsync()

	//* Configure Gin
	router := gin.Default()
	router.Use(sessions.Sessions("cryptoguess", memstore.NewStore([]byte(configs.EnvCookieSecret()))))

	//* Configure CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//* Configure Routes
	routes.RootRoute(router)
	routes.ResourcesRoute(router)
	routes.AuthRoute(router)

	router.Run(":8000")
}
