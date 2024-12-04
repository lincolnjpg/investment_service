package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type Investment struct {
	Id           uuid.UUID
	UserId       uuid.UUID
	AssetId      uuid.UUID
	Quantity     int
	Status       enum.InvestmentStatusEnum
	Type         enum.InvestmentTypeEnum
	PurchaseDate time.Time
	Message      *string
}
