package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateAssetIndexInput struct {
	Name    AssetIndexNameEnum    `json:"name,omitempty"`
	Acronym AssetIndexAcronymEnum `json:"acronym,omitempty"`
}

func (i CreateAssetIndexInput) Validate() error {
	return validation.ValidateStruct(
		&i,
		validation.Field(&i.Name),
		validation.Field(&i.Acronym),
	)
}

type CreateAssetIndexOutput struct {
	Id string `json:"id,omitempty"`
}

type GetAssetIndexByIdInput struct {
	Id string `json:"id,omitempty"`
}

func (dto GetAssetIndexByIdInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Id,
			validation.Required,
			is.UUIDv4,
		),
	)
}

type GetAssetIndexByIdOutput struct {
	Id      string                `json:"id,omitempty"`
	Name    AssetIndexNameEnum    `json:"name,omitempty"`
	Acronym AssetIndexAcronymEnum `json:"acronym,omitempty"`
}

type UpdateAssetIndexByIdInput struct {
	Id      string                `json:"id,omitempty"`
	Name    AssetIndexNameEnum    `json:"name,omitempty"`
	Acronym AssetIndexAcronymEnum `json:"acronym,omitempty"`
}

func (dto UpdateAssetIndexByIdInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(&dto.Name),
		validation.Field(&dto.Acronym),
	)
}

type UpdateAssetIndexByIdOutput struct {
	Id string `json:"id,omitempty"`
}

type DeleteAssetIndexByIdInput struct {
	Id string `json:"id,omitempty"`
}

func (i DeleteAssetIndexByIdInput) Validate() error {
	return validation.ValidateStruct(
		&i,
		validation.Field(
			&i.Id,
			validation.Required,
			is.UUIDv4,
		),
	)
}
