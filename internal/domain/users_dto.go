package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
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
