package enum

import (
	"errors"

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

var investorProfileLabels = map[string]InvestorProfileEnum{
	"Conservador": Conservative,
	"Moderado":    Moderate,
	"Arrojado":    Aggressive,
}

func (e InvestorProfileEnum) String() string {
	return investorProfileNames[e]
}

func (e *InvestorProfileEnum) Scan(value interface{}) error {
	if v, ok := value.(string); ok {
		*e = investorProfileLabels[v]
		return nil
	}

	return errors.New("could not scan investment type")
}

func (e InvestorProfileEnum) Validate() error {
	return validation.Validate(
		uint8(e),
		validation.Required,
		validation.In(uint8(Conservative), uint8(Moderate), uint8(Aggressive)),
	)
}
