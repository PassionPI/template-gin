package main

import (
	"time"

	"app.land.x/app/controller"
	"app.land.x/app/core"
	"app.land.x/app/middleware"
	"app.land.x/pkg/rsa256"

	"github.com/gin-gonic/gin"
)

func initialize() {
	rsa256.CreateRsaPem()
}

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
		// index := "./frontend/index.html"
		// router.Static("/assets", "./frontend/assets")
		// router.StaticFile("/", index)
		// router.StaticFile("/favicon.svg", "./frontend/favicon.svg")
		// router.NoRoute(func(c *gin.Context) { c.File(index) })
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
			mids.Empty(),
			mids.AuthValidator(),
		)

		api.POST("/ping", ctrl.Ping)

		{
			SSH := api.Group("/ssh")
			SSH.POST("/new", ctrl.Echo)
			SSH.POST("/del", ctrl.Echo)
			SSH.POST("/put", ctrl.Echo)
			SSH.POST("/get", ctrl.Echo)
			SSH.POST("/list", ctrl.Echo)

			{
				command := SSH.Group("/command")
				command.POST("/run/:id", ctrl.Echo)
			}
		}

		{
			user := api.Group("/user")
			user.POST("/privilege/put", ctrl.Echo)
			user.POST("/privilege/get", ctrl.Echo)
		}

		{
			project := api.Group("/project")
			project.POST("/new", ctrl.Echo)
			project.POST("/del", ctrl.Echo)
			project.POST("/put", ctrl.Echo)
			project.POST("/list", ctrl.Echo)
			project.GET("/:project", func(ctx *gin.Context) {
				project := ctx.Param("project")
				ctx.JSON(200, gin.H{
					"message": project,
				})
			})

			{
				ENV := project.Group("/:project")
				ENV.POST("/new", ctrl.Echo)
				ENV.POST("/del", ctrl.Echo)
				ENV.POST("/put", ctrl.Echo)
				ENV.POST("/list", ctrl.Echo)
				ENV.GET("/:env", func(ctx *gin.Context) {
					project := ctx.Param("project")
					env := ctx.Param("env")
					ctx.JSON(200, gin.H{
						"message": project + "+" + env,
					})
				})

				{
					App := ENV.Group("/:env")
					App.POST("/new", ctrl.Echo)
					App.POST("/del", ctrl.Echo)
					App.POST("/put", ctrl.Echo)
					App.POST("/list", ctrl.Echo)
					App.POST("/:app/deploy", ctrl.Echo)
					App.POST("/:app/release", ctrl.Echo)
					App.GET("/:app", func(ctx *gin.Context) {
						project := ctx.Param("project")
						env := ctx.Param("env")
						app := ctx.Param("app")
						ctx.JSON(200, gin.H{
							"message": project + "+" + env + "+" + app,
						})
					})
				}
			}
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
