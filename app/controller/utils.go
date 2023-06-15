package controller

import (
	"app.land.x/pkg/resp"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) responseWithJwtToken(c *gin.Context, username string) {
	token, err := ctrl.core.Token.Generate(username)

	if err != nil {
		resp.Err(c, "Failed to sign the token")
		return
	}

	resp.Ok(c, &gin.H{"token": token})
}
