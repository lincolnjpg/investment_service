package enum

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AssetStatusEnum uint8

const (
	Pending AssetStatusEnum = iota + 1
	Done
	Canceled
)

var assetStatusNames = map[AssetStatusEnum]string{
	Pending:  "Pending",
	Done:     "Done",
	Canceled: "Canceled",
}

var assetStatusLabels = map[string]AssetStatusEnum{
	"Pending":  Pending,
	"Done":     Done,
	"Canceled": Canceled,
}

func (e AssetStatusEnum) String() string {
	return assetStatusNames[e]
}

func (e *AssetStatusEnum) Scan(value interface{}) error {
	if v, ok := value.(string); ok {
		*e = assetStatusLabels[v]
		return nil
	}

	return errors.New("could not scan asset status")
}

func (e AssetStatusEnum) Validate() error {
	return validation.Validate(
		uint8(e),
		validation.In(uint8(Pending), uint8(Done), uint8(Canceled)),
	)
}
