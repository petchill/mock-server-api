package response

import "go.mongodb.org/mongo-driver/mongo"

type responseRepository struct {
	mongoDB *mongo.Database
}

func NewResponseRepository(mongoDB *mongo.Database) IResponseRepository {
	return &responseRepository{
		mongoDB,
	}
}
