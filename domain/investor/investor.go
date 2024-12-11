package investor

import (
	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type Investor struct {
	Id              uuid.UUID
	Name            string
	InvestorProfile enum.InvestorProfileEnum
}
