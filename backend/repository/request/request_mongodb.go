package request

import (
	"awesomeProject/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.TODO()

func NewMongoDBRepo(Conn *mongo.Client) *MongoRequestRepo {
	return &MongoRequestRepo{
		collection: Conn.Database("capstoneproject").Collection("request"),
	}
}

type MongoRequestRepo struct {
	collection *mongo.Collection
}

func (m *MongoRequestRepo) CreateRequest(req *models.Request) (interface{}, error) {
	obj, err := m.collection.InsertOne(ctx, req)
	return obj.InsertedID, err
}

func (m *MongoRequestRepo) GetAllRequests() ([]*models.Request, error){
	filter := bson.D{{}}
	var requests []*models.Request

	cur, err := m.collection.Find(ctx, filter)
	if err != nil {
		return requests, err
	}

	for cur.Next(ctx){
		var req models.Request
		err := cur.Decode(&req)
		if err != nil {
			return requests, err
		}

		requests = append(requests, &req)
	}

	if err := cur.Err(); err != nil {
		return requests, err
	}

	cur.Close(ctx)

	if len(requests) == 0 {
		return requests, mongo.ErrNoDocuments
	}

	return requests, nil
}


func (m *MongoRequestRepo) DeleteRequest(req *models.Request) error {
	del, err := m.collection.DeleteOne(ctx, req)
	if err != nil{
		return err
	}

	if del.DeletedCount == 0{
		return errors.New("no record were deleted")
	}

	return nil
}

func (m *MongoRequestRepo) EditRequest(req *models.Request) error {
	return nil
}