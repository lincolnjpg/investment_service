package domain

import validation "github.com/go-ozzo/ozzo-validation/v4"

const (
	CDB            = "CDB"
	LCI            = "LCI"
	LCA            = "LCA"
	CRI            = "CRI"
	CRA            = "CRA"
	TESOURO_DIRETO = "TESOURO DIRETO"
	STOCK          = "STOCK"
	FII            = "FII"
)

type InvestmentTypeEnum string

func (t InvestmentTypeEnum) Validate() error {
	return validation.Validate(
		string(t),
		validation.Required,
		validation.In(CDB, LCI, LCA, CRI, CRA, TESOURO_DIRETO, STOCK, FII),
	)
}

const (
	FIXED_INCOME    = "FIXED_INCOME"
	VARIABLE_INCOME = "VARIABLE_INCOME"
)

type AssetClassEnum string

func (t AssetClassEnum) Validate() error {
	return validation.Validate(
		string(t),
		validation.Required,
		validation.In(FIXED_INCOME, VARIABLE_INCOME),
	)
}

type AssetType struct {
	ID          string
	Name        InvestmentTypeEnum
	Description string
	Class       AssetClassEnum
}
