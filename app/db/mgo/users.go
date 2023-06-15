package mgo

import (
	"context"

	"app.land.x/app/model"

	"go.mongodb.org/mongo-driver/bson"
)

// FindUserByUsername 根据用户名查找用户
func (mongo *Mongo) FindUserByUsername(username string) (credentials model.Credentials, err error) {
	err = mongo.Collection.Users.FindOne(
		context.TODO(),
		bson.M{
			"username": username,
		},
	).Decode(&credentials)
	return credentials, err
}
