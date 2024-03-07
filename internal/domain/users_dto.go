package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type CreateUserInput struct {
	Name            string              `json:"name,omitempty"`
	InvestorProfile InvestorProfileEnum `json:"investor_type,omitempty"`
}

func (i CreateUserInput) Validate() error {
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

type GetUserByIDInput struct {
	ID string `json:"id,omitempty"`
}

func (i GetUserByIDInput) Validate() error {
	return validation.ValidateStruct(
		&i,
		validation.Field(
			&i.ID,
			validation.Required,
			is.UUIDv4,
		),
	)
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
	Id              uuid.UUID           `json:"id,omitempty"`
	Name            string              `json:"name,omitempty"`
	InvestorProfile InvestorProfileEnum `json:"investor_profile,omitempty"`
}
