package user

import (
	"context"
	"time"

	"github.com/e-commerce/system/config"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepo struct {
	collection *mongo.Collection
}

var _ UserRepoInterface = &UserRepo{}

func NewUserRepo(mongoService *config.MongoService) *UserRepo {
	return &UserRepo{collection: mongoService.DB.Collection("users")}
}

func (ur *UserRepo) InsertUser(document CreateUser) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := ur.collection.InsertOne(ctx, document)
	return result, err
}
func (ur *UserRepo) FindUserByEmail(email string) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := ur.collection.FindOne(ctx, bson.M{"email": email})
	return result
}
