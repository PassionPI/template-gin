package main

import (
	"fmt"
	"io"
	"os"

	"app_ink/pkg/rsa256"

	"github.com/gin-gonic/gin"
)

func initialize(VolumePath string) {
	basePem := VolumePath + "/pem"
	baseLog := VolumePath + "/log"
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
