package domain

import "github.com/google/uuid"

const (
	CONSERVATIVE = "conservative"
	MODERATE     = "moderate"
	AGGRESSIVE   = "aggressive"
)

type InvestorProfileEnum string

type User struct {
	Id              uuid.UUID
	Name            string
	InvestorProfile InvestorProfileEnum
}
