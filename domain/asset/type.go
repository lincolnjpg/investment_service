package asset

import "github.com/google/uuid"

type AssetType struct {
	Id      uuid.UUID
	Name    string
	Acronym string
}
