package registry

import (
	"github.com/e-commerce/services/otp"
	"github.com/e-commerce/services/user"
	"github.com/e-commerce/system/config"
	"github.com/redis/go-redis/v9"
)

type Container struct {
	UserRoutes *user.UserRoutes
	OtpRoutes  *otp.OTPRoutes
}

func InitContainer(mongoService *config.MongoService, redisService *redis.Client) *Container {
	// Repositories
	var userRepo user.UserRepoInterface = user.NewUserRepo(mongoService)
	var otpRepo otp.OtpRepoInterface = otp.NewOtpRepo(redisService)

	//*Services
	var userService user.UserServiceInterface = user.NewUserService(userRepo)
	var otpService otp.OtpServiceInterface = otp.NewOTPService(otpRepo)

	//* Controllers
	var userController user.UserControllerInterface = user.NewUserController(userService)
	var otpController otp.OtpControllerInterface = otp.NewOTPController(otpService, userService)

	return &Container{
		UserRoutes: user.NewUserRoutes(userController),
		OtpRoutes:  otp.NewOTPRoutes(otpController),
	}
}
