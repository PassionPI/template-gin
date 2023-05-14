package mgo

import (
	"context"

	"app_land_x/app/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (mongo *Mongo) FindUserByUsername(username string) (credentials model.Credentials, err error) {
	err = mongo.Collection.Users.FindOne(
		context.TODO(),
		bson.M{
			"username": username,
		},
	).Decode(&credentials)
	return credentials, err
}
