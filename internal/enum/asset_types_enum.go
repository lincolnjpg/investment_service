package enum

import validation "github.com/go-ozzo/ozzo-validation/v4"

const (
	CDB            = "CDB"
	LCI            = "LCI"
	LCA            = "LCA"
	CRI            = "CRI"
	CRA            = "CRA"
	TESOURO_DIRETO = "TESOURO DIRETO"
	ACAO           = "AÇÃO"
	FII            = "FII"
)

type InvestmentTypeEnum string

const (
	FIXED_INCOME    = "FIXED_INCOME"
	VARIABLE_INCOME = "VARIABLE_INCOME"
)

type AssetClassEnum string

func (t InvestmentTypeEnum) Validate() error {
	return validation.Validate(
		string(t),
		validation.Required,
		validation.In(CDB, LCI, LCA, CRI, CRA, TESOURO_DIRETO, ACAO, FII),
	)
}

func (t AssetClassEnum) Validate() error {
	return validation.Validate(
		string(t),
		validation.Required,
		validation.In(FIXED_INCOME, VARIABLE_INCOME),
	)
}
