package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type Investment struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	AssetId     uuid.UUID
	Quantity    int
	PuchaseDate time.Time
	Status      enum.InvestmentStatusEnum
	Message     string
	Type        enum.InvestmentTypeEnum
}
