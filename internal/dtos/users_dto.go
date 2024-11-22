package dtos

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type CreateUserInput struct {
	Name            string                   `json:"name,omitempty"`
	InvestorProfile enum.InvestorProfileEnum `json:"investor_profile,omitempty"`
}

func (dto CreateUserInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Name,
			validation.Required,
			validation.Length(1, 50),
		),
		validation.Field(&dto.InvestorProfile),
	)
}

type GetUserByIDInput struct {
	Id string `json:"id,omitempty"`
}

func (dto GetUserByIDInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Id,
			validation.Required,
			is.UUIDv4,
		),
	)
}

type UpdateUserInput struct {
	Id              string                   `json:"id,omitempty"`
	Name            string                   `json:"name,omitempty"`
	InvestorProfile enum.InvestorProfileEnum `json:"investor_profile,omitempty"`
}

func (dto UpdateUserInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Name,
			validation.Required,
			validation.Length(1, 100),
		),
		validation.Field(&dto.InvestorProfile),
	)
}

type DeleteUserByIDInput struct {
	Id string `json:"id,omitempty"`
}

func (dto DeleteUserByIDInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Id,
			validation.Required,
			is.UUIDv4,
		),
	)
}

type CreateUserOutput struct {
	Id string `json:"id,omitempty"`
}

type UpdateUserOutput struct {
	Id string `json:"id,omitempty"`
}

type GetUserByIdOutput struct {
	Id              string                   `json:"id,omitempty"`
	Name            string                   `json:"name,omitempty"`
	InvestorProfile enum.InvestorProfileEnum `json:"investor_profile,omitempty"`
}
