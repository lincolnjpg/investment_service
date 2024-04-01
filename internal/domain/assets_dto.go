package domain

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateAssetInput struct {
	Name         string     `json:"name,omitempty"`
	UnitPrice    float64    `json:"unit_price,omitempty"`
	Rentability  float64    `json:"rentability,omitempty"`
	DueDate      *time.Time `json:"due_date,omitempty"`
	Ticker       *string    `json:"ticker,omitempty"`
	AssetTypeId  string     `json:"asset_type_id,omitempty"`
	AssetIndexId *string    `json:"asset_index_id,omitempty"`
}

func (dto CreateAssetInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Name,
			validation.Required,
			validation.Length(1, 100),
		),
		validation.Field(
			&dto.UnitPrice,
			validation.Required,
			validation.Min(0.0).Exclusive(),
		),
		validation.Field(
			&dto.Rentability,
			validation.Required,
			validation.Min(0.0).Exclusive(),
		),
		validation.Field(
			&dto.DueDate,
			validation.Min(time.Now()),
		),
		validation.Field(
			&dto.Ticker,
			is.UpperCase,
		),
		validation.Field(
			&dto.AssetTypeId,
			validation.Required,
			is.UUIDv4,
		),
		validation.Field(
			&dto.AssetIndexId,
			is.UUIDv4,
		),
	)
}

type CreateAssetOutput struct {
	Id string `json:"id,omitempty"`
}
