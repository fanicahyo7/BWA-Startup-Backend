package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(userInput RegisterUserInput) (User, error)
	LoginUser(loginInput LoginInput) (User, error)
	CekEmailUser(cekEmailInput CekEmailInput) (bool, error)
	UpdateAvatar(ID int, locationFile string) (User, error)
	GetUserByID(ID int) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(input FormUpdateUserInput) (User, error)
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

func (s *service) UpdateAvatar(ID int, locationFile string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	user.AvatarFileName = locationFile

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that ID")
	}

	return user, nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) UpdateUser(input FormUpdateUserInput) (User, error) {
	user, err := s.repository.FindByID(input.ID)
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
