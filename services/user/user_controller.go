package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService UserServiceInterface
}

var _ UserControllerInterface = &UserController{}

func NewUserController(userService UserServiceInterface) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) InsertUser(document CreateUser) (int, gin.H) {
	existingUser := uc.userService.FindUserByEmail(document.Email)
	if existingUser.Err() == nil {
		return http.StatusBadRequest, gin.H{"error": "Email already exists"}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(document.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusBadGateway, gin.H{"error": err}
	}

	document.Password = string(hashedPassword)

	res, err := uc.userService.InsertOne(document)
	if err != nil {
		return http.StatusBadGateway, gin.H{"error": err}
	}

	return http.StatusOK, gin.H{"data": res}
}
