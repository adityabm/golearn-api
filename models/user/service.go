package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Register(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	EmailCheck(input EmailCheck) (bool, error)
	SaveProfilePicture(ID int, FileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(input RegisterUserInput) (User, error) {
	checkUser, err := s.repository.SearchByEmail(input.Email)
	if err != nil {
		return checkUser, err
	}

	if(checkUser.ID != 0) {
		return checkUser, errors.New("Email already exists")
	}

	user := User{}
	user.Fullname = input.Fullname
	user.Occupation = input.Occupation
	user.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Create(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.SearchByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) EmailCheck(input EmailCheck) (bool, error) {
	email := input.Email

	user, err := s.repository.SearchByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true, nil
		}

		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveProfilePicture(ID int, FileLocation string) (User, error) {
	user, err := s.repository.SearchByID(ID)
	if err != nil {
		return user, err
	}

	user.ProfilePicture = FileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.SearchByID(ID)
	if err != nil {
		return user, err
	}

	return user, nil
}