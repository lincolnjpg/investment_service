package entities

import (
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type AssetType struct {
	Id          string                  `json:"id,omitempty"`
	Name        enum.InvestmentTypeEnum `json:"name,omitempty"`
	Description string                  `json:"description,omitempty"`
	Class       enum.AssetClassEnum     `json:"class,omitempty"`
}
