package domain

import "github.com/google/uuid"

type UserRepository interface {
	Create(input CreateUserInput) (User, error)
	Update(input UpdateUserInput) (User, error)
	GetById(id uuid.UUID) (User, error)
	DeleteById(id uuid.UUID) error
}

type UserService interface {
	Create(input CreateUserInput) (CreateUserOutput, error)
	Update(input UpdateUserInput) (UpdateUserOutput, error)
	GetById(id uuid.UUID) (GetUserByIdOutput, error)
	DeleteById(id uuid.UUID) error
}
