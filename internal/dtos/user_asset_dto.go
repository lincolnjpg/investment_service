package dtos

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type CreateUserAssetInput struct {
	UserId   uuid.UUID            `json:"user_id,omitempty"`
	AssetId  uuid.UUID            `json:"asset_id,omitempty"`
	Quantity int                  `json:"quantity,omitempty"`
	Status   enum.AssetStatusEnum `json:"status,omitempty"`
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
		validation.Field(&dto.Status),
	)
}

type CreateUserAssetOutput struct {
	Id uuid.UUID `json:"id,omitempty"`
}
