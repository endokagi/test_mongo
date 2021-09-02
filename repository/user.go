package repository

import (
	"context"
	"log"
	"time"

	"testmongo/driver"
	"testmongo/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllUsers() ([]models.User, error) {

	collection := driver.ConnectMongo()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodes []models.User
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	return episodes, err
}

func GetUser(id string) (models.User, error) {

	var user models.User

	collection := driver.ConnectMongo()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	return user, err
}

func CreateUser(user models.User) error {

	collection := driver.ConnectMongo()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collection.InsertOne(ctx, user)

	return err
}

func DeleteUser(id string) error {

	collection := driver.ConnectMongo()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})

	return err
}

func PutUser(user models.User) error {

	id := bson.M{"_id": user.ID}
	filter := bson.D{
		{"$set",
			user},
	}
	collection := driver.ConnectMongo()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collection.UpdateOne(ctx, id, filter)

	return err
}
