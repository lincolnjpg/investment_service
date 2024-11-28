package dtos

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type CreateUserAssetInput struct {
	UserId   uuid.UUID `json:"user_id"`
	AssetId  uuid.UUID `json:"asset_id"`
	Quantity int       `json:"quantity"`
}

func (dto CreateUserAssetInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.UserId,
			validation.Required,
			is.UUIDv4,
		),
		validation.Field(
			&dto.AssetId,
			validation.Required,
			is.UUIDv4,
		),
		validation.Field(
			&dto.Quantity,
			validation.Required,
			validation.Min(1),
		),
	)
}

type CreateUserAssetOutput struct {
	Id uuid.UUID
}
