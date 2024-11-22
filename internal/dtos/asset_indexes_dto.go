package dtos

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateAssetIndexInput struct {
	Name    string `json:"name,omitempty"`
	Acronym string `json:"acronym,omitempty"`
}

func (dto CreateAssetIndexInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Name,
			validation.Required,
		),
		validation.Field(
			&dto.Acronym,
			validation.Required,
		),
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
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Acronym string `json:"acronym,omitempty"`
}

type UpdateAssetIndexByIdInput struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Acronym string `json:"acronym,omitempty"`
}

func (dto UpdateAssetIndexByIdInput) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Name,
			validation.Required,
		),
		validation.Field(
			&dto.Acronym,
			validation.Required,
		),
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
