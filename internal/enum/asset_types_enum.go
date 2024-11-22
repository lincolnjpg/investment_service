package enum

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type InvestmentTypeEnum uint8

const (
	Cdb InvestmentTypeEnum = iota + 1
	Lci
	Lca
	Cri
	Cra
	Debenture
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
	Debenture:     "Debênture",
	TesouroDireto: "Tesouro Direto",
	Acao:          "Ação",
	Fii:           "FII",
}

var investmentTypeLabels = map[string]InvestmentTypeEnum{
	"CDB":            Cdb,
	"LCI":            Lci,
	"LCA":            Lca,
	"CRI":            Cri,
	"CRA":            Cra,
	"Debênture":      Debenture,
	"Tesouro Direto": TesouroDireto,
	"Ação":           Acao,
	"FII":            Fii,
}

func (e InvestmentTypeEnum) String() string {
	return investmentTypeNames[e]
}

func (e *InvestmentTypeEnum) Scan(value interface{}) error {
	if v, ok := value.(string); ok {
		*e = investmentTypeLabels[v]
		return nil
	}

	return errors.New("could not scan investment type")
}

func (e InvestmentTypeEnum) Validate() error {
	return validation.Validate(
		e.String(),
		validation.Required,
		validation.In(
			Cdb.String(),
			Lci.String(),
			Lca.String(),
			Cri.String(),
			Cra.String(),
			TesouroDireto.String(),
			Acao.String(),
			Fii.String(),
		),
	)
}
