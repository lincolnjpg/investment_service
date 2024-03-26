package domain

import validation "github.com/go-ozzo/ozzo-validation/v4"

const (
	CDI_NAME  = "Certificado de Depósito Interbancário"
	IPCA_NAME = "Índice Nacional de Preços ao Consumidor Amplo"
)

type AssetIndexNameEnum string

func (t AssetIndexNameEnum) Validate() error {
	return validation.Validate(
		string(t),
		validation.Required,
		validation.In(CDI_NAME, IPCA_NAME),
	)
}

const (
	CDI_ACRONYM  = "CDI"
	IPCA_ACRONYM = "IPCA"
)

type AssetIndexAcronymEnum string

func (t AssetIndexAcronymEnum) Validate() error {
	return validation.Validate(
		string(t),
		validation.Required,
		validation.In(CDI_ACRONYM, IPCA_ACRONYM),
	)
}

type AssetIndex struct {
	Id      string
	Name    AssetIndexNameEnum
	Acronym AssetIndexAcronymEnum
}
