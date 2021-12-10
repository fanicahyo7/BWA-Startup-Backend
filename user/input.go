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

type FormCreateUserInput struct {
	Name       string `form:"name" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Occupation string `form:"occupation" binding:"required"`
	Password   string `form:"password" binding:"required"`
	Error      error
}
