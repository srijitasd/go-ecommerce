package otp

import (
	"net/http"
	"time"

	"github.com/e-commerce/lib/otpManager"
	"github.com/e-commerce/services/user"
	"github.com/gin-gonic/gin"
)

type OTPController struct {
	OTPService  OtpServiceInterface
	UserService user.UserServiceInterface
}

var _ OtpControllerInterface = &OTPController{}

func NewOTPController(OTPService OtpServiceInterface, userService user.UserServiceInterface) *OTPController {
	return &OTPController{OTPService: OTPService, UserService: userService}
}

func (oc *OTPController) GenerateOTP(document OTPRequest) (int, gin.H) {
	// Validate user exists
	userExists := oc.UserService.FindUserByEmail(document.Email)
	if userExists.Err() != nil {
		return http.StatusBadGateway, gin.H{"error": "User does not exists"}
	}

	var fetchOTP GetOTP = GetOTP{
		Email: document.Email,
	}

	ttl := oc.OTPService.GetTTL(fetchOTP)
	if ttl.Err() != nil {
		return http.StatusBadGateway, gin.H{"error": "unable to generate OTP"}
	}
	if ttl.Val() > 0 {
		return http.StatusBadGateway, gin.H{"error": "OTP already sent"}
	}

	otp := otpManager.GenerateRandomOTP(6)

	var otpDoc OTP = OTP{
		OTP:       otp,
		Email:     document.Email,
		ExpiresAt: time.Now().Add(time.Minute * 5),
	}
	res := oc.OTPService.SaveOTP(otpDoc)
	if res.Err() != nil {
		return http.StatusBadGateway, gin.H{"error": "unable to save to redis"}
	}

	return http.StatusOK, gin.H{"otp": otp}
}

func (oc *OTPController) VerifyOTP(document OTPVerifyRequest) (int, gin.H) {
	// Validate user exists
	userExists := oc.UserService.FindUserByEmail(document.Email)
	if userExists.Err() != nil {
		return http.StatusBadGateway, gin.H{"error": "User does not exists"}
	}

	var fetchOTP GetOTP = GetOTP{
		Email: document.Email,
	}

	res := oc.OTPService.GetOTP(fetchOTP)
	if res.Err() != nil {
		return http.StatusBadGateway, gin.H{"error": "unable to get data from redis"}
	}

	if res.Val() != document.OTP {
		return http.StatusBadGateway, gin.H{"error": "Invalid OTP"}
	}

	return http.StatusOK, gin.H{"status": "OTP Validated"}
}
