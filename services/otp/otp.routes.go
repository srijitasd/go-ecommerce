package otp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OTPRoutes struct {
	OTPController OtpControllerInterface
}

func NewOTPRoutes(OtpController OtpControllerInterface) *OTPRoutes {
	return &OTPRoutes{OTPController: OtpController}
}

func (ur *OTPRoutes) RegisterRoutes(router *gin.Engine) {
	OtpRoutes := router.Group("/api/otp")

	OtpRoutes.POST("/generate", ur.GenerateOTP)
	OtpRoutes.POST("/validate", ur.VerifyOTP)
}

// InsertUser API to insert a new user
func (h *OTPRoutes) GenerateOTP(c *gin.Context) {
	var otp OTPRequest
	if err := c.ShouldBindJSON(&otp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, response := h.OTPController.GenerateOTP(otp)
	c.JSON(status, response)
}

// InsertUser API to insert a new user
func (h *OTPRoutes) VerifyOTP(c *gin.Context) {
	var otp OTPVerifyRequest
	if err := c.ShouldBindJSON(&otp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, response := h.OTPController.VerifyOTP(otp)
	c.JSON(status, response)
}
