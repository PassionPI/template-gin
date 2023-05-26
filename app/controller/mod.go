package controller

import (
	"app_land_x/app/service/mgo"
	"app_land_x/app/service/rds"
	"app_land_x/pkg/resp"
	"app_land_x/pkg/token"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Rds   *rds.Rds
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
