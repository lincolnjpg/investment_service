package enum

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type InvestmentStatusEnum uint8

const (
	Pending InvestmentStatusEnum = iota + 1
	Done
	Canceled
)

var investmentStatusNames = map[InvestmentStatusEnum]string{
	Pending:  "Pending",
	Done:     "Done",
	Canceled: "Canceled",
}

var investmentStatusLabels = map[string]InvestmentStatusEnum{
	"Pending":  Pending,
	"Done":     Done,
	"Canceled": Canceled,
}

func (e InvestmentStatusEnum) String() string {
	return investmentStatusNames[e]
}

func (e *InvestmentStatusEnum) Scan(value interface{}) error {
	if v, ok := value.(string); ok {
		*e = investmentStatusLabels[v]
		return nil
	}

	return errors.New("could not scan investment status")
}

func (e InvestmentStatusEnum) Validate() error {
	return validation.Validate(
		uint8(e),
		validation.In(uint8(Pending), uint8(Done), uint8(Canceled)),
	)
}

type InvestmentTypeEnum uint8

const (
	Buy InvestmentTypeEnum = iota + 1
	Sell
)

var investmentTypesNames = map[InvestmentTypeEnum]string{
	Buy:  "Buy",
	Sell: "Sell",
}

var investmentTypesLabels = map[string]InvestmentTypeEnum{
	"Buy":  Buy,
	"Sell": Sell,
}

func (e InvestmentTypeEnum) String() string {
	return investmentTypesNames[e]
}

func (e *InvestmentTypeEnum) Scan(value interface{}) error {
	if v, ok := value.(string); ok {
		*e = investmentTypesLabels[v]
		return nil
	}

	return errors.New("could not scan investment type")
}

func (e InvestmentTypeEnum) Validate() error {
	return validation.Validate(
		uint8(e),
		validation.Required,
		validation.In(uint8(Buy), uint8(Sell)),
	)
}
