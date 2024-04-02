package domain

import "time"

const MAX_CDI_RENTABILITY float64 = 150.0

type Asset struct {
	Id           string     `json:"id,omitempty"`
	Name         string     `json:"name,omitempty"`
	UnitPrice    float64    `json:"unit_price,omitempty"`
	Rentability  float64    `json:"rentability,omitempty"`
	DueDate      *time.Time `json:"due_date,omitempty"`
	Ticker       *string    `json:"ticker,omitempty"`
	AssetTypeId  string     `json:"asset_type_id,omitempty"`
	AssetIndexId *string    `json:"asset_index_id,omitempty"`
}
