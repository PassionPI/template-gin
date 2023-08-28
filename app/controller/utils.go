package controller

import (
	"app.ai_painter/pkg/qp"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) responseWithJwtToken(c *gin.Context, username string) {
	token, err := ctrl.core.Dep.Token.Generate(username)

	if err != nil {
		qp.Err(c, "Failed to sign the token")
		return
	}

	qp.Ok(c, &gin.H{"token": token})
}
