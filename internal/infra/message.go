package infra

import (
	"github.com/lincolnjpg/investment_service/internal/dtos"
)

type Message struct {
	Asset      dtos.UpdateAssetByIdInput      `json:"asset,omitempty"`
	Investment dtos.UpdateInvestmentByIdInput `json:"investment,omitempty"`
	Ticker     string                         `json:"ticker,omitempty"`
}
