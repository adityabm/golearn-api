package user

type RegisterUserInput struct {
	Fullname   string `json:"fullname" binding:"required" example:"John Doe"`
	Occupation string `json:"occupation" binding:"required" example:"Software Engineer"`
	Email      string `json:"email" binding:"required,email" example:"mail@mail.com"`
	Password   string `json:"password" binding:"required,min=8" example:"12345678"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type EmailCheck struct {
	Email string `json:"email" binding:"required,email"`
}