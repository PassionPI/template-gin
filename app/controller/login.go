package controller

import (
	"app-ink/app/model"
	"app-ink/pkg/rsa256"
	"app-ink/pkg/util"

	"github.com/gin-gonic/gin"
)

// Ping 测试网络通畅接口
func (ctrl *Controller) Ping(c *gin.Context) {
	util.Ok(c, "Hi!")
}

// Echo 测试接口，返回请求的 URL
func (ctrl *Controller) Echo(c *gin.Context) {
	util.Ok(c, c.Request.URL.Path)
}

// Pub 返回公钥
func (ctrl *Controller) Pub(c *gin.Context) {
	publicKey, err := rsa256.GetPublicKey()
	if err != nil {
		util.Err(c, "Failed to get public key")
		return
	}

	util.Ok(c, &gin.H{"publicKey": string(publicKey)})
}

// Sign 注册接口
func (ctrl *Controller) Sign(c *gin.Context) {
	creds, err := util.JSON[model.Credentials](c)
	ctx := c.Request.Context()

	if err != nil {
		return
	}

	username := creds.Username
	password, err := rsa256.Decrypt(creds.Password)

	if err != nil {
		util.Bad(c, "Invalid password")
		return
	}

	_, err = ctrl.core.Dep.Pg.UserFindByUsername(ctx, username)

	if err != nil {
		userSignUp := model.Credentials{
			Username: username,
			Password: password,
		}

		err := ctrl.core.Dep.Pg.UserInsert(ctx, userSignUp)

		if err != nil {
			util.Bad(c, "Failed to insert user")
			return
		}

		ctrl.responseWithJwtToken(c, username)
		return
	}
	util.Bad(c, "Username already exists")
}

// Login 登录接口
func (ctrl *Controller) Login(c *gin.Context) {
	creds, err := util.JSON[model.Credentials](c)
	ctx := c.Request.Context()

	if err != nil {
		return
	}

	username := creds.Username
	password, err := rsa256.Decrypt(creds.Password)

	message := "Invalid password"

	if err != nil {
		util.Bad(c, message)
		return
	}

	user, err := ctrl.core.Dep.Pg.UserFindByUsername(ctx, username)

	if err != nil {
		util.NotFound(c, "No account of this username found")
		return
	}

	if password != user.Password {
		util.Bad(c, message)
		return
	}

	ctrl.responseWithJwtToken(c, username)
}
