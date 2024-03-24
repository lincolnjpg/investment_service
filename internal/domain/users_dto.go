package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateUserInput struct {
	Name            string              `json:"name,omitempty"`
	InvestorProfile InvestorProfileEnum `json:"investor_profile,omitempty"`
}

func (i CreateUserInput) Validate() error {
	return validation.ValidateStruct(
		&i,
		validation.Field(
			&i.Name,
			validation.Required,
			validation.Length(1, 50),
		),
		validation.Field(&i.InvestorProfile),
	)
}

type GetUserByIDInput struct {
	Id string `json:"id,omitempty"`
}

func (i GetUserByIDInput) Validate() error {
	return validation.ValidateStruct(
		&i,
		validation.Field(
			&i.Id,
			validation.Required,
			is.UUIDv4,
		),
	)
}

type UpdateUserInput struct {
	Id              string              `json:"id,omitempty"`
	Name            string              `json:"name,omitempty"`
	InvestorProfile InvestorProfileEnum `json:"investor_profile,omitempty"`
}

func (i UpdateUserInput) Validate() error {
	return validation.ValidateStruct(
		&i,
		validation.Field(
			&i.Name,
			validation.Required,
			validation.Length(1, 100),
		),
		validation.Field(&i.InvestorProfile),
	)
}

type DeleteUserByIDInput struct {
	Id string `json:"id,omitempty"`
}

func (i DeleteUserByIDInput) Validate() error {
	return validation.ValidateStruct(
		&i,
		validation.Field(
			&i.Id,
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
	Id              string              `json:"id,omitempty"`
	Name            string              `json:"name,omitempty"`
	InvestorProfile InvestorProfileEnum `json:"investor_profile,omitempty"`
}
