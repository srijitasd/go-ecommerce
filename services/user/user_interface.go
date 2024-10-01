package user

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepoInterface interface {
	InsertUser(document CreateUser) (*mongo.InsertOneResult, error)
	FindUserByEmail(email string) *mongo.SingleResult
}

type UserServiceInterface interface {
	InsertOne(document CreateUser) (*mongo.InsertOneResult, error)
	FindUserByEmail(email string) *mongo.SingleResult
}

type UserControllerInterface interface {
	InsertUser(document CreateUser) (int, gin.H)
}
