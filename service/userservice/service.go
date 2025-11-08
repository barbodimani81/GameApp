package userservice

import (
	"Game/entity"
	"Game/pkg/phonenumber"
	"fmt"
)

type Service struct {
	repo Repository
}

type Repository interface {
	IsPhoneNumberUnique(phonenumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}

type RegisterRequest struct {
	PhoneNumber string
	Name        string
}

type RegisterResponse struct {
	User entity.User
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) Register(r RegisterRequest) (RegisterResponse, error) {
	// TODO - phone verification by sms

	// phone number validation
	if !phonenumber.IsValid(r.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("invalid phone number")
	}

	// phone number uniqueness
	if isUnique, err := s.repo.IsPhoneNumberUnique(r.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
		}
		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number already registered")
		}
	}

	// user name validation
	if len(r.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name must be more than 3 characters")
	}

	user := entity.User{
		ID:          0,
		PhoneNumber: r.PhoneNumber,
		Name:        r.Name,
	}
	// create new user in storage
	createdUser, err := s.repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	// return created user
	return RegisterResponse{User: createdUser}, nil
}
