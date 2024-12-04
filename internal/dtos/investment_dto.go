package dtos

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type CreateInvestmentInput struct {
	UserId   uuid.UUID                 `json:"user_id,omitempty"`
	AssetId  uuid.UUID                 `json:"asset_id,omitempty"`
	Quantity int                       `json:"quantity,omitempty"`
	Status   enum.InvestmentStatusEnum `json:"status,omitempty"`
	Type     enum.InvestmentTypeEnum   `json:"type,omitempty"`
}

func (dto CreateInvestmentInput) Validate() error {
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
		validation.Field(&dto.Type),
	)
}

type CreateInvestmentOutput struct {
	Id uuid.UUID `json:"id,omitempty"`
}
