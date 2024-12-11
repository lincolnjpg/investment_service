package investment

import (
	"time"

	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type Investment struct {
	Id           uuid.UUID
	InvestorId   uuid.UUID
	AssetId      uuid.UUID
	Quantity     int
	Status       enum.InvestmentStatusEnum
	Message      string
	Type         enum.InvestmentTypeEnum
	PurchaseDate time.Time
}
