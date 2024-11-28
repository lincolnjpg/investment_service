package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserAsset struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	AssetId     uuid.UUID
	Quantity    int
	PuchaseDate time.Time
}
