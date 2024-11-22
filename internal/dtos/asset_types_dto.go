package dtos

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/lincolnjpg/investment_service/internal/enum"
)

type CreateAssetTypeInput struct {
	Name        enum.InvestmentTypeEnum `json:"name,omitempty"`
	Description string                  `json:"description,omitempty"`
	Class       enum.AssetClassEnum     `json:"class,omitempty"`
}

func (dto CreateAssetTypeInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(&dto.Name),
		validation.Field(
			&dto.Description,
			validation.Required,
			validation.Length(1, 100),
		),
		validation.Field(&dto.Class),
	)
}

type CreateAssetTypeOutput struct {
	Id string `json:"id,omitempty"`
}

type GetAssetTypeByIDInput struct {
	Id string `json:"id,omitempty"`
}

type GetAssetTypeByIDOutput struct {
	Id          string                  `json:"id,omitempty"`
	Name        enum.InvestmentTypeEnum `json:"name,omitempty"`
	Description string                  `json:"description,omitempty"`
	Class       enum.AssetClassEnum     `json:"class,omitempty"`
}

func (dto GetAssetTypeByIDInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Id,
			validation.Required,
			is.UUIDv4,
		),
	)
}

type UpdateAssetTypeByIdInput struct {
	Id          string                  `json:"id,omitempty"`
	Name        enum.InvestmentTypeEnum `json:"name,omitempty"`
	Description string                  `json:"description,omitempty"`
	Class       enum.AssetClassEnum     `json:"class,omitempty"`
}

func (dto UpdateAssetTypeByIdInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Id,
			validation.Required,
			is.UUIDv4,
		),
		validation.Field(&dto.Name),
		validation.Field(
			&dto.Description,
			validation.Required,
			validation.Length(1, 100),
		),
		validation.Field(&dto.Class),
	)
}

type UpdateAssetTypeByIdOutput struct {
	Id string `json:"id,omitempty"`
}

type DeleteAssetTypeByIdInput struct {
	Id string `json:"id,omitempty"`
}

func (dto DeleteAssetTypeByIdInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Id,
			validation.Required,
			is.UUIDv4,
		),
	)
}
