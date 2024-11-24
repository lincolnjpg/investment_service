package entities

import (
	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type User struct {
	Id              uuid.UUID
	Name            string
	InvestorProfile enum.InvestorProfileEnum
}
