package domain

import "github.com/google/uuid"

type CreateUserInput struct {
	Name            string              `json:"name,omitempty"`
	InvestorProfile InvestorProfileEnum `json:"investor_type,omitempty"`
}

type UpdateUserInput struct {
	Name            string
	InvestorProfile InvestorProfileEnum
}

type CreateUserOutput struct {
	Id uuid.UUID `json:"id,omitempty"`
}

type UpdateUserOutput struct {
	Id uuid.UUID
}

type GetUserByIdOutput struct {
	Id              uuid.UUID
	Name            string
	InvestorProfile InvestorProfileEnum
}
