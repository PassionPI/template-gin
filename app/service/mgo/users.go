package mgo

import (
	"context"

	"app_ink/app/model"

	"go.mongodb.org/mongo-driver/bson"
)

// FindUserByUsername 根据用户名查找用户
func (mongo *Mongo) FindUserByUsername(ctx context.Context, username string) (credentials model.Credentials, err error) {
	err = mongo.Collection.Users.FindOne(
		ctx,
		bson.M{
			"username": username,
		},
	).Decode(&credentials)
	return credentials, err
}
