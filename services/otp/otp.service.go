package otp

import (
	"github.com/redis/go-redis/v9"
)

type OTPService struct {
	otpRepo OtpRepoInterface
}

var _ OtpServiceInterface = &OTPService{}

func NewOTPService(OTPRepo OtpRepoInterface) *OTPService {
	return &OTPService{otpRepo: OTPRepo}
}

func (us *OTPService) SaveOTP(document OTP) *redis.StatusCmd {
	return us.otpRepo.SaveOTP(document)
}

func (us *OTPService) GetOTP(document GetOTP) *redis.StringCmd {
	return us.otpRepo.GetOTP(document)
}

func (us *OTPService) GetTTL(document GetOTP) *redis.DurationCmd {
	return us.otpRepo.GetTTL(document)
}
