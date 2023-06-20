package main

import (
	"fmt"
	"io"
	"os"

	"app.land.x/pkg/rsa256"

	"github.com/gin-gonic/gin"
)

func initialize() {
	base := "./private"
	basePem := base + "/pem"
	baseLog := base + "/log"
	{
		rsa256.SetBasePath(basePem)
		rsa256.CreateRsaPem()
	}
	{
		err := os.MkdirAll(baseLog, os.ModePerm)
		if err != nil {
			fmt.Println("Create folder fail:", baseLog, err)
			os.Exit(1)
		}
		file, _ := os.Create(baseLog + "/gin.log")
		gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	}
}
