package asset

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type FixedIncomeAsset struct {
	Id          uuid.UUID          `json:"id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Type        enum.AssetTypeEnum `json:"type,omitempty"` // value object
	UnitPrice   money.Money        `json:"unit_price,omitempty"`
	Rentability float64            `json:"rentability,omitempty"`
	DueDate     time.Time          `json:"due_date,omitempty"`
	Index       string             `json:"index,omitempty"` // value object
	IssuedBy    string             `json:"issued_by,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty"`
}
