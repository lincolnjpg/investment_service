package entities

import "github.com/google/uuid"

type AssetIndex struct {
	Id      uuid.UUID
	Name    string
	Acronym string
}
