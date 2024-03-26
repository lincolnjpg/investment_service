package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateAssetTypeInput struct {
	Name        InvestmentTypeEnum `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	IndexId     *string            `json:"index_id,omitempty"`
	Class       AssetClassEnum     `json:"class,omitempty"`
}

func (i CreateAssetTypeInput) Validate() error {
	return validation.ValidateStruct(
		&i,
		validation.Field(&i.Name),
		validation.Field(
			&i.Description,
			validation.Required,
			validation.Length(1, 100),
		),
		validation.Field(
			&i.IndexId,
			is.UUIDv4,
		),
		validation.Field(&i.Class),
	)
}

type CreateAssetTypeOutput struct {
	Id string `json:"id,omitempty"`
}

type GetAssetTypeByIDInput struct {
	Id string `json:"id,omitempty"`
}

type GetAssetTypeByIDOutput struct {
	Id          string             `json:"id,omitempty"`
	Name        InvestmentTypeEnum `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	IndexId     *string            `json:"index_id,omitempty"`
	Class       AssetClassEnum     `json:"class,omitempty"`
}

func (i GetAssetTypeByIDInput) Validate() error {
	return validation.ValidateStruct(
		&i,
		validation.Field(
			&i.Id,
			validation.Required,
			is.UUIDv4,
		),
	)
}

type UpdateAssetTypeByIdInput struct {
	Id          string             `json:"id,omitempty"`
	Name        InvestmentTypeEnum `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	IndexId     *string            `json:"index_id,omitempty"`
	Class       AssetClassEnum     `json:"class,omitempty"`
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
		validation.Field(
			&dto.IndexId,
			is.UUIDv4,
		),
		validation.Field(&dto.Class),
	)
}

type UpdateAssetTypeByIdOutput struct {
	Id string `json:"id,omitempty"`
}
