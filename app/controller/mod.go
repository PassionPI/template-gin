package controller

import (
	"app.land.x/app/service/mgo"
	"app.land.x/pkg/resp"
	"app.land.x/pkg/token"

	"github.com/gin-gonic/gin"
)

// Controller 结构体
// 包含了所有的控制器
// 以及整个应用所有依赖
type Controller struct {
	Mongo *mgo.Mongo
	Token *token.Token
}

func (ctrl *Controller) responseWithJwtToken(c *gin.Context, username string) {
	token, err := ctrl.Token.Generate(username)

	if err != nil {
		resp.Err(c, "Failed to sign the token")
		return
	}

	resp.Ok(c, &gin.H{"token": token})
}
