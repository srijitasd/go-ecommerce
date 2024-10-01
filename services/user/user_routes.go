package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userController UserControllerInterface
}

func NewUserRoutes(userController UserControllerInterface) *UserRoutes {
	return &UserRoutes{userController: userController}
}

func (ur *UserRoutes) RegisterRoutes(router *gin.Engine) {
	userRoutes := router.Group("/api/user")

	userRoutes.POST("/", ur.InsertUser)
}

// InsertUser API to insert a new user
func (h *UserRoutes) InsertUser(c *gin.Context) {
	var user CreateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, response := h.userController.InsertUser(user)
	c.JSON(status, response)
}
