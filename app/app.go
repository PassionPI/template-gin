package main

import (
	"time"

	"app.ai_painter/app/controller"
	"app.ai_painter/app/core"
	"app.ai_painter/app/middleware"

	"github.com/gin-gonic/gin"
)

func createEngine() *gin.Engine {
	initialize()

	core := core.New()
	ctrl := controller.New(core)
	mids := middleware.New(core)

	router := gin.New()
	router.Use(
		gin.Logger(),
		gin.Recovery(),
		mids.RateLimiter(30, time.Minute),
	)

	// 前端静态资源
	{
		base := "./frontend"
		icon := "/favicon.svg"
		asset := "/assets"
		index := base + "/index.html"
		router.StaticFile("/", index)
		router.StaticFile(icon, base+icon)
		router.Static(asset, base+asset)
		router.NoRoute(func(c *gin.Context) { c.File(index) })
	}

	{
		pub := router.Group("/api")
		pub.POST("/pub", ctrl.Pub)
		pub.POST("/sign", ctrl.Sign)
		pub.POST("/login", ctrl.Login) // need throttle, lock
	}

	{
		api := router.Group("/api")
		api.Use(
			mids.AuthValidator(),
		)

		api.POST("/ping", ctrl.Ping)

		{
			user := api.Group("/user")
			user.POST("/privilege/put", ctrl.Echo)
			user.POST("/privilege/get", ctrl.Echo)
		}
	}

	{
		open := router.Group("/open")
		open.Use(
			gin.BasicAuth(gin.Accounts{"miss": "ballad"}),
		)

	}

	{
		webhook := router.Group("/webhook")

		{
			github := webhook.Group("/github")
			github.POST("/pr", ctrl.WebhookGithubPR)
		}
	}

	return router
}
