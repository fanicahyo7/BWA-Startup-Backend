package user

type UserInput struct {
	Name       string `binding:"required"`
	Occupation string `binding:"required"`
	Email      string `binding:"required,email"`
	Password   string `binding:"required"`
}
