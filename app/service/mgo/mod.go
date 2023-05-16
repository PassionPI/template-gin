package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *MongoCollection
}

type MongoCollection struct {
	Users *mongo.Collection
}

func New(uri string) *Mongo {
	Client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	MongoDatabase := Client.Database("sample_demo")
	MongoCollectionUser := MongoDatabase.Collection("users")

	MongoCollection := &MongoCollection{
		Users: MongoCollectionUser,
	}

	return &Mongo{
		Client,
		MongoDatabase,
		MongoCollection,
	}
}

// userSignUp := model.Credentials{
// 	Username: "land.pan",
// 	Password: "123456",
// }
// MongoCollectionUser.InsertOne(context.TODO(), userSignUp)
// var user model.Credentials
// id, _ := primitive.ObjectIDFromHex("645a8d097074707223093bfc")
// err = MongoCollectionUser.FindOne(
// 	context.TODO(),
// 	bson.D{
// 		{"_id", id},
// 	},
// ).Decode(&user)
// if err != nil {
// 	fmt.Println("Mongo FindOne", err.Error())
// }
// fmt.Println("username", user.Username, "password", user.Password)
