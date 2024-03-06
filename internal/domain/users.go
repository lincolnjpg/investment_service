package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

const (
	CONSERVATIVE = "conservative"
	MODERATE     = "moderate"
	AGGRESSIVE   = "aggressive"
)

type InvestorProfileEnum string

func (t InvestorProfileEnum) Validate() error {
	return validation.Validate(
		string(t),
		validation.Required,
		validation.In(CONSERVATIVE, MODERATE, AGGRESSIVE),
	)
}

type User struct {
	Id              uuid.UUID
	Name            string
	InvestorProfile InvestorProfileEnum
}
