package user

// User represents the structure of a user object in the request body
type CreateUser struct {
	Name     string `json:"name" binding:"required" bson:"name"`
	Email    string `json:"email" binding:"required,email" bson:"email"`
	Password string `json:"password" binding:"required" bson:"password"`
}

type GetOTP struct {
	Email string `json:"email" binding:"required,email" bson:"email"`
}
