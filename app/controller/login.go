package controller

import (
	"context"
	"net/http"

	"app.land.x/app/model"
	"app.land.x/pkg/req"
	"app.land.x/pkg/resp"
	"app.land.x/pkg/rsa256"

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
		resp.Err(c, "Failed to get public key")
		return
	}

	resp.Ok(c, &gin.H{"publicKey": string(publicKey)})
}

// Sign 注册接口
func (ctrl *Controller) Sign(c *gin.Context) {
	creds, err := req.JSON[model.Credentials](c)

	if err != nil {
		return
	}

	username := creds.Username
	password, err := rsa256.Decrypt(creds.Password)

	if err != nil {
		resp.Err(c, "Invalid password")
		return
	}

	err = ctrl.core.Mongo.Collection.Users.FindOne(
		context.TODO(),
		bson.M{
			"username": username,
		},
	).Decode(&model.Credentials{})
	if err != nil {
		userSignUp := model.Credentials{
			Username: username,
			Password: password,
		}
		ctrl.core.Mongo.Collection.Users.InsertOne(context.TODO(), userSignUp)

		ctrl.responseWithJwtToken(c, username)
		return
	}
	resp.Err(c, "Username already exists")
}

// Login 登录接口
func (ctrl *Controller) Login(c *gin.Context) {
	creds, err := req.JSON[model.Credentials](c)

	if err != nil {
		return
	}

	username := creds.Username
	password, err := rsa256.Decrypt(creds.Password)

	message := "Invalid password"

	if err != nil {
		resp.Err(c, message)
		return
	}

	user, err := ctrl.core.Mongo.FindUserByUsername(username)

	if err != nil {
		resp.Err(c, "No account of this username found")
		return
	}

	if password != user.Password {
		resp.Err(c, message)
		return
	}

	ctrl.responseWithJwtToken(c, username)
}
