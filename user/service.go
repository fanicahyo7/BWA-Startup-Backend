package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(userInput RegisterUserInput) (User, error)
	LoginUser(loginInput LoginInput) (User, error)
	CekEmailUser(cekEmailInput CekEmailInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(userInput RegisterUserInput) (User, error) {
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

func (s *service) LoginUser(loginInput LoginInput) (User, error) {
	email := loginInput.Email
	pswd := loginInput.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(pswd))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) CekEmailUser(cekEmailInput CekEmailInput) (bool, error) {
	input := cekEmailInput.Email

	user, err := s.repository.FindByEmail(input)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}
