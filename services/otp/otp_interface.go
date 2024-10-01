package otp

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type OtpRepoInterface interface {
	SaveOTP(document OTP) *redis.StatusCmd
	GetOTP(document GetOTP) *redis.StringCmd
	GetTTL(document GetOTP) *redis.DurationCmd
}

type OtpServiceInterface interface {
	SaveOTP(document OTP) *redis.StatusCmd
	GetOTP(document GetOTP) *redis.StringCmd
	GetTTL(document GetOTP) *redis.DurationCmd
}

type OtpControllerInterface interface {
	GenerateOTP(document OTPRequest) (int, gin.H)
	VerifyOTP(document OTPVerifyRequest) (int, gin.H)
}
