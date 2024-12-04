package infra

import "github.com/google/uuid"

type Message struct {
	InvestmentId uuid.UUID `json:"user_asset_id,omitempty"`
	Ticker       string    `json:"ticker,omitempty"`
}
