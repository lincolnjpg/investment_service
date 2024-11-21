package enum

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type InvestorProfileEnum uint8

const (
	Conservative InvestorProfileEnum = iota + 1
	Moderate
	Aggressive
)

var investorProfileNames = map[InvestorProfileEnum]string{
	Conservative: "conservative",
	Moderate:     "moderate",
	Aggressive:   "aggressive",
}

func (e InvestorProfileEnum) Name() string {
	return investorProfileNames[e]
}

func (e InvestorProfileEnum) Validate() error {
	fmt.Println(e)
	return validation.Validate(
		e.Name(),
		validation.Required,
		validation.In(Conservative.Name(), Moderate.Name(), Aggressive.Name()),
	)
}
