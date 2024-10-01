package user

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserService struct {
	userRepo UserRepoInterface
}

var _ UserServiceInterface = &UserService{}

func NewUserService(userRepo UserRepoInterface) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) InsertOne(document CreateUser) (*mongo.InsertOneResult, error) {
	return us.userRepo.InsertUser(document)
}

func (us *UserService) FindUserByEmail(email string) *mongo.SingleResult {
	fmt.Println("email", email)
	return us.userRepo.FindUserByEmail(email)
}
