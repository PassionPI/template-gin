package main

import (
	"time"

	"app_ink/app/controller"
	"app_ink/app/core"
	"app_ink/app/middleware"

	"github.com/gin-gonic/gin"
)

func createEngine(core *core.Core) *gin.Engine {

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
		// base := "./frontend"
		// icon := "/favicon.svg"
		// asset := "/assets"
		// index := base + "/index.html"
		// router.Static(asset, base+asset)
		// router.StaticFile("/", index)
		// router.StaticFile(icon, base+icon)
		// router.NoRoute(func(c *gin.Context) { c.File(index) })
	}

	// 上传静态资源
	{
		upload := "/upload"
		router.Static(upload, core.Dep.Env.VolumePath+upload)
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
		// open := router.Group("/open")
		// open.Use(
		// 	gin.BasicAuth(gin.Accounts{"miss": "ballad"}),
		// )

	}

	{
		// webhook := router.Group("/webhook")

		// {
		// 	github := webhook.Group("/github")
		// 	github.POST("/pr", ctrl.WebhookGithubPR)
		// }
	}

	return router
}
