package repository

import "go.mongodb.org/mongo-driver/mongo"

type personRepository struct {
	collection *mongo.Collection
	db *mongo.Client
}

type PersonRepository interface {

}

func NewPersonRepository(db *mongo.Client) PersonRepository {
	return &personRepository {
		db : db,
		collection: db.Database("user_db").Collection("persons"),
	}
}
