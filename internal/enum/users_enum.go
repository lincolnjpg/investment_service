package enum

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type InvestorProfileEnum uint8

const (
	Conservative InvestorProfileEnum = iota + 1
	Moderate
	Aggressive
)

var investorProfileNames = map[InvestorProfileEnum]string{
	Conservative: "Conservador",
	Moderate:     "Moderado",
	Aggressive:   "Arrojado",
}

func (e InvestorProfileEnum) String() string {
	return investorProfileNames[e]
}

func (e InvestorProfileEnum) Validate() error {
	return validation.Validate(
		e.String(),
		validation.Required,
		validation.In(Conservative.String(), Moderate.String(), Aggressive.String()),
	)
}
