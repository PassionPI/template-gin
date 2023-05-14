package controller

import (
	"context"
	"net/http"

	"app_land_x/app/model"
	"app_land_x/pkg/req"
	"app_land_x/pkg/resp"
	"app_land_x/pkg/rsa256"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Ping(c *gin.Context) {
	c.String(http.StatusOK, "Hello!")
}

func (ctrl *Controller) Pub(c *gin.Context) {
	publicKey, err := rsa256.GetPublicKey()
	if err != nil {
		resp.Err(c, "Failed to get public key")
		return
	}

	resp.Ok(c, &gin.H{"publicKey": string(publicKey)})
}

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

	_, err = ctrl.Mongo.FindUserByUsername(username)
	if err != nil {
		userSignUp := model.Credentials{
			Username: username,
			Password: password,
		}
		ctrl.Mongo.Collection.Users.InsertOne(context.TODO(), userSignUp)

		ctrl.responseWithJwtToken(c, username)
		return
	}
	resp.Err(c, "Username already exists")
}

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

	user, err := ctrl.Mongo.FindUserByUsername(username)

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
