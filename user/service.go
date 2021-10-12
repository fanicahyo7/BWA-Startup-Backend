package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(userInput UserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(userInput UserInput) (User, error) {
	user := User{}
	user.Name = userInput.Name
	user.Email = userInput.Email
	user.Occupation = userInput.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.MinCost)
	if err != nil {
		return user, nil
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "USER"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, nil
	}

	return newUser, nil
}
