package main

import (
	"os"

	"app_land_x/app/controller"
	"app_land_x/app/middleware"
	"app_land_x/app/service/mgo"
	"app_land_x/pkg/rsa256"
	"app_land_x/pkg/token"

	"github.com/gin-gonic/gin"
)

func createController() *controller.Controller {
	mongoURI := os.Getenv("MONGODB_URI") // "mongodb://localhost:27017"
	jwtSecret := os.Getenv("JWT_SECRET") // "Wia3d3zRH84SuLo5n6WCfR5YNU09qLLZHlBnWeGnFZ"

	return &controller.Controller{
		Mongo: mgo.New(mongoURI),
		Token: token.New(jwtSecret),
	}
}

func createEngine() *gin.Engine {
	rsa256.CreateRsaPem()

	ctrl := createController()
	mids := middleware.New(ctrl)

	router := gin.New()
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	// 前端静态资源
	ctrl.Frontend(router)

	router.POST("/pub", ctrl.Pub)
	router.POST("/sign", ctrl.Sign)
	router.POST("/login", ctrl.Login) // need throttle, lock

	{
		authorized := router.Group("/api")
		authorized.Use(mids.AuthValidator())

		authorized.POST("/", ctrl.Protected)
		authorized.POST("/ping", ctrl.Ping)
	}

	return router
}
