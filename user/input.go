package user

type RegisterUserInput struct {
	Name       string `binding:"required"`
	Occupation string `binding:"required"`
	Email      string `binding:"required,email"`
	Password   string `binding:"required"`
}

type LoginInput struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

type CekEmailInput struct {
	Email string `binding:"required,email"`
}
