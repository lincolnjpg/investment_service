package entities

import (
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	Id           uuid.UUID  `json:"id,omitempty"`
	Name         string     `json:"name,omitempty"`
	UnitPrice    float64    `json:"unit_price,omitempty"`
	Rentability  float64    `json:"rentability,omitempty"`
	DueDate      *time.Time `json:"due_date,omitempty"`
	Ticker       *string    `json:"ticker,omitempty"`
	Type         string     `json:"type,omitempty"`
	AssetIndexId *string    `json:"asset_index_id,omitempty"`
}
