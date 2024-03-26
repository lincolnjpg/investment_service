package domain

import validation "github.com/go-ozzo/ozzo-validation/v4"

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
