package controller

import (
	"app.ai_painter/pkg/util"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) responseWithJwtToken(c *gin.Context, username string) {
	token, err := ctrl.core.Dep.Token.Generate(username)

	if err != nil {
		util.Err(c, "Failed to sign the token")
		return
	}

	util.Ok(c, &gin.H{"token": token})
}
