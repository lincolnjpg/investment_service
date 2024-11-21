package enum

import validation "github.com/go-ozzo/ozzo-validation/v4"

const (
	CdiName AssetIndexNameEnum = iota + 1
	IpcaName
)

var assetIndexNames = map[AssetIndexNameEnum]string{
	CdiName:  "Certificado de Depósito Interbancário",
	IpcaName: "Índice Nacional de Preços ao Consumidor Amplo",
}

type AssetIndexNameEnum uint8

func (e AssetIndexNameEnum) Name() string {
	return assetIndexNames[e]
}

func (e AssetIndexNameEnum) Validate() error {
	return validation.Validate(
		e.Name(),
		validation.Required,
		validation.In(CdiName.Name(), IpcaName.Name()),
	)
}

const (
	CdiAcronym AssetIndexAcronymEnum = iota + 1
	IpcaAcronym
)

var asetIndexAcronymNames = map[AssetIndexAcronymEnum]string{
	CdiAcronym:  "CDI",
	IpcaAcronym: "IPCA",
}

type AssetIndexAcronymEnum uint8

func (e AssetIndexAcronymEnum) Name() string {
	return asetIndexAcronymNames[e]
}

func (e AssetIndexAcronymEnum) Validate() error {
	return validation.Validate(
		e.Name(),
		validation.Required,
		validation.In(CdiAcronym.Name(), IpcaAcronym.Name()),
	)
}
