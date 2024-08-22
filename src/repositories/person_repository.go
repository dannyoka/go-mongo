package repositories

import (
	"context"

	"github.com/dannyoka/go-mongo/src/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPersonRepository interface {
	FindOneByName(name string) (data.Person, error)
	FindAll() ([]data.Person, error)
}

type PersonRepository struct {
	collection *mongo.Collection
}

func NewPersonRepository(db *mongo.Database) IPersonRepository {
	collection := db.Collection("test")
	return &PersonRepository{collection: collection}
}

func (r *PersonRepository) FindOneByName(name string) (data.Person, error) {
	var result data.Person
	err := r.collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&result)
	return result, err
}

func (r *PersonRepository) FindAll() ([]data.Person, error) {
	var results []data.Person
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result data.Person
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
