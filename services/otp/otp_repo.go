package otp

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type OtpRepo struct {
	collection *redis.Client
}

func NewOtpRepo(redisService *redis.Client) *OtpRepo {
	return &OtpRepo{collection: redisService}
}

var _ OtpRepoInterface = &OtpRepo{}

func (or *OtpRepo) SaveOTP(document OTP) *redis.StatusCmd {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := or.collection.Set(ctx, document.Email, document.OTP, 2*time.Minute)
	return result
}

func (or *OtpRepo) GetOTP(document GetOTP) *redis.StringCmd {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := or.collection.Get(ctx, document.Email)
	return result
}

func (or *OtpRepo) GetTTL(document GetOTP) *redis.DurationCmd {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := or.collection.TTL(ctx, document.Email)
	return result
}
