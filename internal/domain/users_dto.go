package domain

import "github.com/google/uuid"

type CreateUserInput struct {
	Name             string
	InvestorTypeEnum InvestorTypeEnum
}

type UpdateUserInput struct {
	Name             string
	InvestorTypeEnum InvestorTypeEnum
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
