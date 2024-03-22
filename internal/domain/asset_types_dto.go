package domain

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CreateAssetTypeInput struct {
	Name        InvestmentTypeEnum `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
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
		validation.Field(&i.Class),
	)
}

type CreateAssetTypeOutput struct {
	ID string `json:"id,omitempty"`
}
