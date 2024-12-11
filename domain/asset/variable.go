package asset

import (
	"time"

	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type VariableIncomeAsset struct {
	Id        uuid.UUID          `json:"id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Type      enum.AssetTypeEnum `json:"type,omitempty"`
	Ticker    string             `json:"ticker,omitempty"`
	Quote     float64            `json:"quote,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
}
