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
