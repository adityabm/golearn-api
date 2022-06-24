package user

type RegisterUserInput struct {
	Fullname   string `json:"fullname" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type EmailCheck struct {
	Email string `json:"email" binding:"required,email"`
}