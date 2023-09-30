package controller

import (
	"net/http"

	"app.ai_painter/app/model"
	"app.ai_painter/pkg/rsa256"
	"app.ai_painter/pkg/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Ping 测试网络通畅接口
func (ctrl *Controller) Ping(c *gin.Context) {
	c.String(http.StatusOK, "Hello!")
}

// Echo 测试接口，返回请求的 URL
func (ctrl *Controller) Echo(c *gin.Context) {
	c.String(http.StatusOK, c.Request.URL.Path)
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
		util.Err(c, "Invalid password")
		return
	}

	err = ctrl.core.Dep.Mongo.Collection.Users.FindOne(
		ctx,
		bson.M{
			"username": username,
		},
	).Decode(&model.Credentials{})
	if err != nil {
		userSignUp := model.Credentials{
			Username: username,
			Password: password,
		}
		_, err := ctrl.core.Dep.Mongo.Collection.Users.InsertOne(ctx, userSignUp)

		if err != nil {
			util.Err(c, "Failed to insert user")
			return
		}

		ctrl.responseWithJwtToken(c, username)
		return
	}
	util.Err(c, "Username already exists")
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
		util.Err(c, message)
		return
	}

	user, err := ctrl.core.Dep.Mongo.FindUserByUsername(ctx, username)

	if err != nil {
		util.Err(c, "No account of this username found")
		return
	}

	if password != user.Password {
		util.Err(c, message)
		return
	}

	ctrl.responseWithJwtToken(c, username)
}
