package entities

import (
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type AssetIndex struct {
	Id      string
	Name    enum.AssetIndexNameEnum
	Acronym enum.AssetIndexAcronymEnum
}
