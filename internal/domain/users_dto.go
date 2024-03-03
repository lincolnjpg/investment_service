package domain

import "github.com/google/uuid"

type CreateUserInput struct {
	Name         string           `json:"name,omitempty"`
	InvestorType InvestorTypeEnum `json:"investor_type,omitempty"`
}

type UpdateUserInput struct {
	Name         string
	InvestorType InvestorTypeEnum
}

type CreateUserOutput struct {
	Id uuid.UUID
}

type UpdateUserOutput struct {
	Id uuid.UUID
}

type GetUserByIdOutput struct {
	Id           uuid.UUID
	Name         string
	InvestorType InvestorTypeEnum
}
