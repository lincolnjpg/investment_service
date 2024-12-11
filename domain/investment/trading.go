package investment

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type Trading struct {
	Id           uuid.UUID
	InvestorId   uuid.UUID
	AssetId      uuid.UUID
	Quantity     int
	TradingPrice money.Money
	Status       enum.InvestmentStatusEnum
	Message      string
	Type         enum.InvestmentTypeEnum
	PurchaseDate time.Time
}
