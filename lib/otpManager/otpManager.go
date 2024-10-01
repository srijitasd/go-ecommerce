package otpManager

import "crypto/rand"

// Helper function to generate a random OTP
func GenerateRandomOTP(length int) string {
	const charset = "0123456789"
	otp := make([]byte, length)
	_, _ = rand.Read(otp)

	for i := range otp {
		otp[i] = charset[int(otp[i])%len(charset)]
	}

	return string(otp)
}
