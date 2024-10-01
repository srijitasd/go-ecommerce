package otp

import "time"

type OTP struct {
	OTP       string    `json:"otp"`
	Email     string    `json:"email"`
	UserID    string    `bson:"user_id"`
	ExpiresAt time.Time `bson:"expires_at"`
}

type OTPRequest struct {
	Email string `json:"email"`
}

type OTPVerifyRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type GetOTP struct {
	Email string `json:"email"`
}
