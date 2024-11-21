package enum

import validation "github.com/go-ozzo/ozzo-validation/v4"

type InvestmentTypeEnum uint8

const (
	Cdb InvestmentTypeEnum = iota + 1
	Lci
	Lca
	Cri
	Cra
	TesouroDireto
	Acao
	Fii
)

var investmentTypeNames = map[InvestmentTypeEnum]string{
	Cdb:           "CDB",
	Lci:           "LCI",
	Lca:           "LCA",
	Cri:           "CRI",
	Cra:           "CRA",
	TesouroDireto: "Tesouro Direto",
	Acao:          "Ação",
	Fii:           "FII",
}

func (e InvestmentTypeEnum) Name() string {
	return investmentTypeNames[e]
}

func (e InvestmentTypeEnum) Validate() error {
	return validation.Validate(
		e.Name(),
		validation.Required,
		validation.In(Cdb.Name(), Lci.Name(), Lca.Name(), Cri.Name(), Cra.Name(), TesouroDireto.Name(), Acao.Name(), Fii.Name()),
	)
}

type AssetClassEnum uint8

const (
	FixedIncome AssetClassEnum = iota + 1
	VariableIncome
)

var assetClassNames = map[AssetClassEnum]string{
	FixedIncome:    "Fixed Income",
	VariableIncome: "Variable Income",
}

func (e AssetClassEnum) Name() string {
	return assetClassNames[e]
}

func (e AssetClassEnum) Validate() error {
	return validation.Validate(
		e.Name(),
		validation.Required,
		validation.In(FixedIncome.Name(), VariableIncome.Name()),
	)
}
