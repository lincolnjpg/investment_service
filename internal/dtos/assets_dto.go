package dtos

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type CreateAssetInput struct {
	Name         string             `json:"name,omitempty"`
	UnitPrice    float64            `json:"unit_price,omitempty"`
	Rentability  float64            `json:"rentability,omitempty"`
	DueDate      *time.Time         `json:"due_date,omitempty"`
	Ticker       *string            `json:"ticker,omitempty"`
	Type         enum.AssetTypeEnum `json:"type,omitempty"`
	AssetIndexId uuid.UUID          `json:"asset_index_id,omitempty"`
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
			&dto.Type,
			validation.Required,
			validation.Length(1, 15),
		),
		validation.Field(
			&dto.AssetIndexId,
			is.UUIDv4,
		),
	)
}

type CreateAssetOutput struct {
	Id uuid.UUID `json:"id,omitempty"`
}

type GetAssetByIdInput struct {
	Id uuid.UUID `json:"id,omitempty"`
}

func (dto GetAssetByIdInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Id,
			validation.Required,
			is.UUIDv4,
		),
	)
}

type GetAssetByIdOutput struct {
	Id           uuid.UUID          `json:"id,omitempty"`
	Name         string             `json:"name,omitempty"`
	UnitPrice    float64            `json:"unit_price,omitempty"`
	Rentability  float64            `json:"rentability,omitempty"`
	DueDate      *time.Time         `json:"due_date,omitempty"`
	Ticker       *string            `json:"ticker,omitempty"`
	Type         enum.AssetTypeEnum `json:"type,omitempty"`
	AssetIndexId *string            `json:"asset_index_id,omitempty"`
}

type UpdateAssetByIdInput struct {
	Id           uuid.UUID          `json:"id,omitempty"`
	Name         string             `json:"name,omitempty"`
	UnitPrice    float64            `json:"unit_price,omitempty"`
	Rentability  float64            `json:"rentability,omitempty"`
	DueDate      *time.Time         `json:"due_date,omitempty"`
	Ticker       *string            `json:"ticker,omitempty"`
	Type         enum.AssetTypeEnum `json:"type,omitempty"`
	AssetIndexId *string            `json:"asset_index_id,omitempty"`
}

func (dto UpdateAssetByIdInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Id,
			validation.Required,
			is.UUIDv4,
		),
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
			&dto.Type,
			validation.Required,
			validation.Length(1, 15),
		),
		validation.Field(
			&dto.AssetIndexId,
			is.UUIDv4,
		),
	)
}

type UpdateAssetByIdOutput struct {
	Id uuid.UUID `json:"id,omitempty"`
}

type DeleteAssetByIdInput struct {
	Id uuid.UUID `json:"id,omitempty"`
}

func (dto DeleteAssetByIdInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Id,
			validation.Required,
			is.UUIDv4,
		),
	)
}
