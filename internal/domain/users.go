package domain

import "github.com/google/uuid"

const (
	CONSERVATIVE = iota + 1
	MODERATE
	AGGRESSIVE
)

type InvestorTypeEnum int

type User struct {
	Id           uuid.UUID
	Name         string
	InvestorType InvestorTypeEnum
}

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
